# Step 3 · Restoring the terminal, always

There is a bug in step 2: if something goes wrong *after* `Init()`, we return without calling `Fini()`, and the terminal is left broken. Rather than remembering to call `Fini()` on every path, we ask Go to do it for us with `defer`, and we move the work into its own function so that errors travel out cleanly.

{{diff main.go}}

- `func run() error` is a function that returns an error. Every `return err` hands the problem to the caller; `return nil` at the end means success.
- `defer screen.Fini()` schedules `Fini()` to run **when `run` returns**, however it returns: normally, through any `return err`, even through a panic. It is placed right after `Init()` succeeded, so we never try to restore a screen that was never set up.
- `main` calls `run()`; if it got an error it prints it to **standard error** with `fmt.Fprintln(os.Stderr, ...)` and exits with status 1 via `os.Exit(1)`. This ordering matters: the deferred `Fini()` has already restored the terminal by the time `main` prints, so the message is not drawn over the game screen. Never call `os.Exit` or `log.Fatal` *inside* `run`: they skip deferred calls.
- `os.Exit` takes the exit code as an argument. Status 1 tells the shell something failed.

!!! Run it: unchanged on screen. Every later step keeps this shape: `run` does the work, `main` reports.
