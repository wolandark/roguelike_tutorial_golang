# Golang Roguelike Tutorial — build a roguelike in Go, one step at a time

An interactive, cute course that grows **one program** from ten lines into a complete,
classic terminal roguelike in **Go** with [tcell](https://github.com/gdamore/tcell), in the
style of the `cutesy` Go-by-example site. The game design follows the classic
[python-tcod roguelike tutorial](https://rogueliketutorials.com/tutorials/tcod/v2/).

## How the course works

- **49 steps in 14 chapters.** Step 1 opens the screen and draws an `@`. Every later step
  adds a handful of lines: a new file shown in full, or the **exact diff** against the
  previous step, generated from the real code so it can never drift.
- **Every new piece of Go is explained the first time it appears**: error values, `defer`,
  methods and pointer receivers, slices and maps, interfaces, type switches, embedding,
  closures, `iota`, `container/heap`, `encoding/gob`... and the roguelike ideas alongside:
  the game loop, actions, a flat tile grid, rooms and tunnels, shadowcasting, components,
  Dijkstra, event handlers, spawn tables, equipment.
- Each step ends with **what you should see** when you run it; the whole program after the
  step is one tab away, editable.
- ▶ **play here** builds the step (with your edits) and runs it in a real PTY streamed to a
  terminal in the page; 🖥️ **play in a terminal window** launches it in your desktop
  terminal (wezterm, kitty, alacritty, foot, ghostty, gnome-terminal, konsole,
  xfce4-terminal, xterm or `$TERMINAL`; needs the server running on your desktop).
- 🏋️ an **exercise at the end of every chapter**, checked by a hidden Go test against the
  step's code.

## Run the site

```sh
docker compose up          # everything except "play in a terminal window"
go run ./server            # natively; enables the terminal-window button
```

Then open <http://127.0.0.1:8378>. Any Go ≥ 1.22 works.

## Run the steps yourself

```sh
go run ./steps/01-hello-tcell
go run ./steps/49-wielding-and-wearing   # the finished game (needs an 80x50 terminal)
```

Keys: arrows / `hjkl` / numpad move, `.` wait, `g` pick up, `i` inventory, `d` drop,
`v` message history, `/` look, `c` character sheet, `>` descend, `Esc` save & quit.

## Layout and rebuilding

- `steps/NN-slug/*.go` is the program as it stands after step NN.
- `course/NN-slug.md` explains the step. It starts with `# Step N · Title`, may open a
  chapter with `## Chapter: Name`, and uses `{{file x.go}}` (whole file), `{{diff x.go}}`
  (exact changes vs the previous step), `!!! what you should see` and `??? common mistake`.
- `exercises/<slug>/`: `meta.json` (`base`, `file`, `prompt`), `starter.go`, `solution.go`,
  `check_test.go`. `exercises/go.mod` fences these files off the root module on purpose.

```sh
python3 build.py     # steps + course + exercises -> data.js, exercises.json
python3 validate.py  # vet everything; every solution must pass, every starter must fail
```

The server exposes `/api/build`, `/api/check`, `/api/pty` (websocket PTY) and `/api/open`
(terminal window). Everything builds in isolated temp directories with timeouts.
