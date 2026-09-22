# Step 3 · Restoring the terminal, always

### The problem

Step 2 has a bug you cannot see yet. Suppose something fails *after* `Init()` succeeded, in code we will add later. We would `return`, and `Fini()` would never run: broken terminal, exactly like the experiment in step 1. Sprinkling `screen.Fini()` before every `return` is fragile; someone will forget one.

There are two design options. Keep everything in `main` and remember `Fini` everywhere, or move the work into its own function where Go can guarantee that `Fini` runs whenever that function ends, for any reason. Go has `defer` for exactly that guarantee, and it works per function, which is why the second option needs a second function. Once the work has its own function, it should *report* failures rather than print them, because printing while tcell owns the screen would draw over the game; and `main` becomes the only place that prints.

>>> Move the work into `func run() error`. Right after `Init` succeeds, `defer screen.Fini()`. Make `run` return errors instead of printing them, and make `main` print whatever `run` returns to standard error and exit with status 1.

!!! Unchanged on screen. But now any error after `Init`, in this step or in step 50, restores the terminal before it is reported.

--- reveal

{{diff main.go}}

- `func run() error` returns an error. Each `return err` hands the problem to the caller; `return nil` at the end means success.
- `defer screen.Fini()` schedules `Fini()` for **when `run` returns**, however it returns: normally, through any `return err`, even through a panic. It sits right after `Init()` succeeded, so we never try to restore a screen that was never set up.
- `main` calls `run()`; on error it prints to **standard error** with `fmt.Fprintln(os.Stderr, ...)` and exits with `os.Exit(1)`. The order matters: by the time `main` prints, the deferred `Fini()` has already restored the terminal, so the message lands in your shell, not on the game screen.
- Never call `os.Exit` or `log.Fatal` *inside* `run`: both end the process immediately and skip deferred calls.

--- end

%%% Add `return fmt.Errorf("boom")` right after the `defer` line and run it. The terminal is restored and `error: boom` appears in your shell. Now replace that line with `os.Exit(1)` and run again: no `Fini`, broken terminal, `reset` needed. That difference is the reason for this whole step.

??? `os.Exit` takes the exit code as an argument: `os.Exit(1)`. Without it the compiler says `not enough arguments in call to os.Exit`. And a function that contains `return err` must declare `error` as its result: `func run() error`.
