# Step 2 · Checking errors

### The problem

Run step 1 with a terminal type tcell has never heard of: `TERM=nonsense go run .`. It crashes with a nil-pointer panic, a stack trace, and no hint of what went wrong. `NewScreen` could not find a terminfo entry for that name, returned an error and a `nil` screen; we ignored the error and called `Init` on nothing. (Redirecting stdin or stdout does *not* trigger this, by the way: tcell talks to `/dev/tty` directly, so `go run . < /dev/null | cat` works fine.)

Go has no exceptions. A function that can fail returns an `error` as its last result, and the caller decides what to do. The convention is blunt: check it immediately, on the next line. That is the whole change in this step, but it is the pattern you will write more than any other in Go, so it gets a step of its own.

>>> Keep the error from `NewScreen` instead of discarding it, and check the one `Init` returns too. When either is not `nil`, print it and stop, before touching the screen.

!!! Same picture as step 1. With `TERM=nonsense go run .` you now get one line, `error: couldn't open terminfo ($TERM) file for nonsense`, instead of a stack trace.

--- reveal

{{diff main.go}}

- `screen, err := tcell.NewScreen()` keeps the error. `if err != nil` is the idiom: a `nil` error means "nothing went wrong".
- `if err := screen.Init(); err != nil {` does two things in one line: the statement before the semicolon runs first and declares `err` only for this `if`. It is the standard shape for "call, then check".
- `fmt.Println("error:", err)` prints the error (an `error` knows how to print itself) and `return` leaves `main`, which ends the program. `fmt` comes from the standard library, so the import block now lists two packages; `gofmt` keeps standard-library imports in a group above third-party ones.

--- end

%%% `TERM=nonsense` makes `NewScreen` fail. To make `Init` fail instead, take the terminal away: `setsid -w go run .` runs the program in a new session with no controlling terminal, so `/dev/tty` cannot be opened. Step 2 reports `error: open /dev/tty: no such device or address`; step 1 would panic. Two different failures, two checks, both needed.

??? If the compiler says `tcell.NewScreen` has a different signature than expected, the import path is missing the `/v2`. Check the import line.
