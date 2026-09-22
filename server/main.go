// Cutesy Rogue tutorial server.
//
// Serves the static site and offers four things to the browser:
//
//	POST /api/run     build and run a stdout lesson, return its output
//	POST /api/build   vet a lesson or part (with the learner's edits overlaid)
//	POST /api/check   run an exercise's hidden go test against the learner's file
//	GET  /api/pty     websocket: build the part and run it in a real PTY,
//	                  streamed to the in-page terminal
//	POST /api/open    build the part and launch it in a real terminal
//	                  emulator window on this machine
//
// Everything builds in isolated temp directories with timeouts.
package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"strings"
	"time"

	"github.com/coder/websocket"
	"github.com/creack/pty"
)

type Exercise struct {
	Prompt  string `json:"prompt"`
	Base    string `json:"base"` // lessons/NN-slug or parts/NN-slug
	File    string `json:"file"`
	Starter string `json:"starter"`
}

var (
	exercises map[string]Exercise
	sem       = make(chan struct{}, 2) // max 2 concurrent go builds
	safeName  = regexp.MustCompile(`^[a-z0-9_]+\.go$`)
	safeBase  = regexp.MustCompile(`^steps/[0-9]{2}-[a-z0-9-]+$`)
)

// stageModule copies a lesson's or part's sources into a fresh temp module
// and overlays the learner's files on top. Only file names that already
// exist there are accepted, so the browser cannot write anywhere else.
func stageModule(base string, overlay map[string]string) (string, error) {
	if !safeBase.MatchString(base) {
		return "", fmt.Errorf("bad lesson name")
	}
	src := filepath.FromSlash(base)
	entries, err := os.ReadDir(src)
	if err != nil {
		return "", fmt.Errorf("unknown lesson %q", base)
	}
	dir, err := os.MkdirTemp("", "rogue-run-*")
	if err != nil {
		return "", err
	}
	known := map[string]bool{}
	for _, e := range entries {
		if !strings.HasSuffix(e.Name(), ".go") || strings.HasSuffix(e.Name(), "_test.go") {
			continue
		}
		known[e.Name()] = true
		data, err := os.ReadFile(filepath.Join(src, e.Name()))
		if err != nil {
			return dir, err
		}
		if err := os.WriteFile(filepath.Join(dir, e.Name()), data, 0644); err != nil {
			return dir, err
		}
	}
	for name, code := range overlay {
		if !safeName.MatchString(name) || !known[name] {
			return dir, fmt.Errorf("unknown file %q", name)
		}
		if err := os.WriteFile(filepath.Join(dir, name), []byte(code), 0644); err != nil {
			return dir, err
		}
	}
	for _, f := range []string{"go.mod", "go.sum"} {
		data, err := os.ReadFile(f)
		if err != nil {
			return dir, err
		}
		if err := os.WriteFile(filepath.Join(dir, f), data, 0644); err != nil {
			return dir, err
		}
	}
	return dir, nil
}

// runGo runs a go subcommand in dir with a timeout and returns its output.
func runGo(dir string, timeout time.Duration, args ...string) (string, bool) {
	sem <- struct{}{}
	defer func() { <-sem }()
	cmd := exec.Command("go", args...)
	cmd.Dir = dir
	cmd.Env = append(os.Environ(), "GOFLAGS=-mod=mod -buildvcs=false")
	var out strings.Builder
	cmd.Stdout, cmd.Stderr = &out, &out
	if err := cmd.Start(); err != nil {
		return err.Error(), false
	}
	done := make(chan error, 1)
	go func() { done <- cmd.Wait() }()
	select {
	case err := <-done:
		return out.String(), err == nil
	case <-time.After(timeout):
		cmd.Process.Kill()
		return out.String() + "\ntimeout: took longer than " + timeout.String() + " 🐌", false
	}
}

// tidy makes go tool output nicer to read in the browser: temp paths
// become plain file names.
func tidy(dir, s string) string {
	s = strings.ReplaceAll(s, dir+string(filepath.Separator), "")
	s = strings.ReplaceAll(s, "# cutesy-rogue\n", "")
	s = strings.ReplaceAll(s, "# [cutesy-rogue]\n", "")
	return strings.TrimSpace(s)
}

// buildPart stages the lesson with the learner's edits and compiles it to
// dir/game. On failure the temp dir is removed and the compiler output
// returned.
func buildPart(base string, overlay map[string]string) (dir, output string, err error) {
	dir, err = stageModule(base, overlay)
	if err != nil {
		if dir != "" {
			os.RemoveAll(dir)
		}
		return "", err.Error(), err
	}
	out, ok := runGo(dir, 120*time.Second, "build", "-o", "game", ".")
	if !ok {
		os.RemoveAll(dir)
		return "", tidy(dir, out), errors.New("build failed")
	}
	return dir, "", nil
}

// handleRun builds a stdout lesson with the learner's edits and runs it,
// returning what it printed. This is the cutesy "run it" button.
func handleRun(w http.ResponseWriter, r *http.Request) {
	var req struct {
		Base  string            `json:"base"`
		Files map[string]string `json:"files"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "bad request", 400)
		return
	}
	res := runProgram(req.Base, req.Files)
	json.NewEncoder(w).Encode(res)
}

type runResult struct {
	Output   string `json:"output"`
	Stderr   string `json:"stderr"`
	ExitCode int    `json:"exitCode"`
	BuildErr string `json:"buildErr,omitempty"`
}

// runProgram compiles and runs a lesson with a 15 second limit, capturing
// stdout and stderr separately.
func runProgram(base string, overlay map[string]string) runResult {
	dir, out, err := buildPart(base, overlay)
	if err != nil {
		return runResult{BuildErr: out}
	}
	defer os.RemoveAll(dir)
	cmd := exec.Command(filepath.Join(dir, "game"))
	cmd.Dir = dir
	var stdout, stderr strings.Builder
	cmd.Stdout, cmd.Stderr = &stdout, &stderr
	if err := cmd.Start(); err != nil {
		return runResult{BuildErr: err.Error()}
	}
	done := make(chan error, 1)
	go func() { done <- cmd.Wait() }()
	select {
	case <-done:
	case <-time.After(15 * time.Second):
		cmd.Process.Kill()
		return runResult{Stderr: "timeout: program ran longer than 15s (infinite loop?) 🐌", ExitCode: -1}
	}
	exit := 0
	if cmd.ProcessState != nil {
		exit = cmd.ProcessState.ExitCode()
	}
	return runResult{Output: stdout.String(), Stderr: stderr.String(), ExitCode: exit}
}

func handleBuild(w http.ResponseWriter, r *http.Request) {
	var req struct {
		Base  string            `json:"base"`
		Files map[string]string `json:"files"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "bad request", 400)
		return
	}
	dir, err := stageModule(req.Base, req.Files)
	if dir != "" {
		defer os.RemoveAll(dir)
	}
	if err != nil {
		json.NewEncoder(w).Encode(map[string]any{"ok": false, "output": err.Error()})
		return
	}
	out, ok := runGo(dir, 120*time.Second, "vet", ".")
	if ok {
		out = "it compiles and vets clean! 🎉"
	}
	json.NewEncoder(w).Encode(map[string]any{"ok": ok, "output": tidy(dir, out)})
}

func handleCheck(w http.ResponseWriter, r *http.Request) {
	var req struct {
		Slug string `json:"slug"`
		Code string `json:"code"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "bad request", 400)
		return
	}
	ex, ok := exercises[req.Slug]
	if !ok {
		json.NewEncoder(w).Encode(map[string]any{"pass": false, "output": "unknown exercise"})
		return
	}
	pass, out := checkExercise(ex, req.Slug, req.Code)
	json.NewEncoder(w).Encode(map[string]any{"pass": pass, "output": out})
}

// checkExercise overlays the learner's file on the lesson and either runs
// the hidden check_test.go, or runs the program and compares its output
// with expected.txt.
func checkExercise(ex Exercise, slug, code string) (bool, string) {
	if expected, err := os.ReadFile(filepath.Join("exercises", slug, "expected.txt")); err == nil {
		r := runProgram(ex.Base, map[string]string{ex.File: code})
		if r.BuildErr != "" {
			return false, "build error 🙀:\n" + r.BuildErr
		}
		got := norm(r.Output)
		if got == norm(string(expected)) {
			return true, r.Output
		}
		msg := "your output:\n" + r.Output
		if strings.TrimSpace(r.Stderr) != "" {
			msg += "\n(stderr) " + r.Stderr
		}
		return false, msg + "\n\nexpected output:\n" + string(expected)
	}
	dir, err := stageModule(ex.Base, map[string]string{ex.File: code})
	if dir != "" {
		defer os.RemoveAll(dir)
	}
	if err != nil {
		return false, err.Error()
	}
	test, err := os.ReadFile(filepath.Join("exercises", slug, "check_test.go"))
	if err != nil {
		return false, "exercise has no test 😿"
	}
	if err := os.WriteFile(filepath.Join(dir, "check_test.go"), test, 0644); err != nil {
		return false, err.Error()
	}
	out, ok := runGo(dir, 120*time.Second, "test", "-count=1", "-run", ".", ".")
	return ok, tidy(dir, out)
}

// norm trims trailing whitespace on every line and around the whole text,
// so an extra newline never fails an exercise.
func norm(s string) string {
	lines := strings.Split(strings.TrimSpace(s), "\n")
	for i, l := range lines {
		lines[i] = strings.TrimRight(l, " \t")
	}
	return strings.Join(lines, "\n")
}

// gameCmd prepares the compiled game to run from the lesson's directory
// (so savegame.sav lands next to it) with a colourful terminal.
func gameCmd(dir, base string) *exec.Cmd {
	cmd := exec.Command(filepath.Join(dir, "game"))
	cmd.Dir = filepath.FromSlash(base)
	cmd.Env = append(os.Environ(), "TERM=xterm-256color", "COLORTERM=truecolor")
	return cmd
}

// handlePTY is the in-page terminal: build, run under a PTY, and pipe
// bytes both ways over a websocket. The browser sends one JSON message
// with its edited files first, then raw keystrokes.
func handlePTY(w http.ResponseWriter, r *http.Request) {
	part := r.URL.Query().Get("base")
	c, err := websocket.Accept(w, r, &websocket.AcceptOptions{OriginPatterns: []string{"*"}})
	if err != nil {
		return
	}
	defer c.CloseNow()
	ctx := r.Context()
	status := func(msg string) { c.Write(ctx, websocket.MessageText, []byte(msg)) }

	var first struct {
		Files map[string]string `json:"files"`
	}
	if _, data, err := c.Read(ctx); err != nil || json.Unmarshal(data, &first) != nil {
		return
	}
	status("building " + part + "…")
	dir, out, err := buildPart(part, first.Files)
	if err != nil {
		status("build failed 🙀\r\n" + strings.ReplaceAll(out, "\n", "\r\n"))
		c.Close(websocket.StatusNormalClosure, "build failed")
		return
	}
	defer os.RemoveAll(dir)

	cmd := gameCmd(dir, part)
	tty, err := pty.StartWithSize(cmd, &pty.Winsize{Rows: 50, Cols: 80})
	if err != nil {
		status("could not start a pty: " + err.Error())
		return
	}
	defer func() { cmd.Process.Kill(); cmd.Wait(); tty.Close() }()

	go func() {
		buf := make([]byte, 8192)
		for {
			n, err := tty.Read(buf)
			if n > 0 {
				if c.Write(ctx, websocket.MessageBinary, buf[:n]) != nil {
					return
				}
			}
			if err != nil { // EIO when the game exits and closes its side
				result := "exit status 0"
				if werr := cmd.Wait(); werr != nil {
					result = werr.Error()
				}
				status("[game exited: " + result + "] press ▶ to play again")
				c.Close(websocket.StatusNormalClosure, "exited")
				return
			}
		}
	}()
	// ROGUE_PTY_LOG=path appends every byte the browser sends to the PTY;
	// handy when a terminal emulator answers a query the game did not expect.
	var inLog *os.File
	if path := os.Getenv("ROGUE_PTY_LOG"); path != "" {
		inLog, _ = os.OpenFile(path, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		defer inLog.Close()
	}
	for {
		_, data, err := c.Read(ctx)
		if err != nil {
			return
		}
		if inLog != nil {
			fmt.Fprintf(inLog, "%q\n", data)
		}
		tty.Write(data)
	}
}

// handleOpen builds the part and launches it in a real terminal window on
// this machine. The binary is kept for a few hours because some terminals
// (wezterm, kitty) hand the command to an already running instance and
// return immediately, so we cannot wait for it.
func handleOpen(w http.ResponseWriter, r *http.Request) {
	var req struct {
		Base  string            `json:"base"`
		Files map[string]string `json:"files"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "bad request", 400)
		return
	}
	dir, out, err := buildPart(req.Base, req.Files)
	if err != nil {
		json.NewEncoder(w).Encode(map[string]any{"ok": false, "output": out})
		return
	}
	time.AfterFunc(4*time.Hour, func() { os.RemoveAll(dir) })
	cwd, _ := filepath.Abs(filepath.FromSlash(req.Base))
	name, err := launchTerminal(filepath.Join(dir, "game"), cwd)
	if err != nil {
		json.NewEncoder(w).Encode(map[string]any{"ok": false, "output": err.Error()})
		return
	}
	json.NewEncoder(w).Encode(map[string]any{"ok": true, "output": "opened in " + name + " 🗡️  (look for a new window)"})
}

// launchTerminal tries $TERMINAL first, then the usual suspects, asking
// each for an 80x50 window where it has a flag for it.
func launchTerminal(bin, cwd string) (string, error) {
	// The window manager usually resizes a brand-new window a few tens of
	// milliseconds after it appears, which reaches the program as an extra
	// resize event. Waiting for the window to settle keeps the early
	// steps, which exit on the first event after tcell's own resize, alive.
	script := fmt.Sprintf(`sleep 0.5; cd %q && TERM=xterm-256color COLORTERM=truecolor %q; echo; echo "[game exited] press Enter to close this window"; read _`, cwd, bin)
	sh := []string{"sh", "-c", script}
	var candidates [][]string
	if t := os.Getenv("TERMINAL"); t != "" {
		candidates = append(candidates, append([]string{t, "-e"}, sh...))
	}
	candidates = append(candidates,
		append([]string{"wezterm", "--config", "initial_cols=80", "--config", "initial_rows=50", "start", "--"}, sh...),
		append([]string{"kitty", "-o", "initial_window_width=80c", "-o", "initial_window_height=50c"}, sh...),
		append([]string{"alacritty", "-o", "window.dimensions.columns=80", "-o", "window.dimensions.lines=50", "-e"}, sh...),
		append([]string{"foot", "-W", "80x50"}, sh...),
		append([]string{"ghostty", "-e"}, sh...),
		append([]string{"gnome-terminal", "--geometry=80x50", "--"}, sh...),
		append([]string{"konsole", "-e"}, sh...),
		append([]string{"xfce4-terminal", "--geometry=80x50", "-x"}, sh...),
		append([]string{"xterm", "-geometry", "80x50", "-e"}, sh...),
		append([]string{"x-terminal-emulator", "-e"}, sh...),
	)
	if os.Getenv("DISPLAY") == "" && os.Getenv("WAYLAND_DISPLAY") == "" {
		return "", errors.New("no display: this only works when the server runs on your desktop (go run ./server), not in Docker. Use ▶ play here instead.")
	}
	for _, c := range candidates {
		if _, err := exec.LookPath(c[0]); err != nil {
			continue
		}
		cmd := exec.Command(c[0], c[1:]...)
		if err := cmd.Start(); err != nil {
			continue
		}
		go cmd.Wait()
		return c[0], nil
	}
	return "", errors.New("no terminal emulator found; set $TERMINAL to yours and restart the server")
}

// pruneOldBuilds removes leftover build dirs from earlier runs.
func pruneOldBuilds() {
	old, _ := filepath.Glob(filepath.Join(os.TempDir(), "rogue-run-*"))
	for _, d := range old {
		if info, err := os.Stat(d); err == nil && time.Since(info.ModTime()) > 4*time.Hour {
			os.RemoveAll(d)
		}
	}
}

func jsonHandler(fn http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS")
		if r.Method == "OPTIONS" {
			return
		}
		w.Header().Set("Content-Type", "application/json")
		fn(w, r)
	}
}

func main() {
	data, err := os.ReadFile("exercises.json")
	if err != nil {
		panic(err)
	}
	if err := json.Unmarshal(data, &exercises); err != nil {
		panic(err)
	}
	pruneOldBuilds()

	fs := http.FileServer(http.Dir("."))
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Cache-Control", "no-cache")
		fs.ServeHTTP(w, r)
	})
	http.HandleFunc("/api/run", jsonHandler(handleRun))
	http.HandleFunc("/api/build", jsonHandler(handleBuild))
	http.HandleFunc("/api/check", jsonHandler(handleCheck))
	http.HandleFunc("/api/open", jsonHandler(handleOpen))
	http.HandleFunc("/api/pty", handlePTY)

	// ROGUE_ADDR overrides the listen address (default :8378), e.g. for a
	// second instance next to a running one.
	addr := os.Getenv("ROGUE_ADDR")
	if addr == "" {
		addr = ":8378"
	}
	fmt.Println("🗡️  cutesy rogue tutorial server on http://127.0.0.1" + addr)
	if err := http.ListenAndServe(addr, nil); err != nil {
		panic(err)
	}
}
