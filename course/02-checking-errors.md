# Step 2 · Checking errors

### The problem

Run step 1 with no terminal attached: `go run . < /dev/null | cat`. It crashes with a nil-pointer panic, a stack trace, and no hint of what went wrong. `NewScreen` returned an error and a `nil` screen; we ignored the error and called `Init` on nothing.

Go has no exceptions. A function that can fail returns an `error` as its last result, and the caller decides what to do. The convention is blunt: check it immediately, on the next line. That is the whole change in this step, but it is the pattern you will write more than any other in Go, so it gets a step of its own.

>>> Keep the error from `NewScreen` instead of discarding it, and check the one `Init` returns too. When either is not `nil`, print it and stop, before touching the screen.

!!! Same picture as step 1. With `go run . < /dev/null | cat` you now get one line, `error: ...`, instead of a stack trace.

--- reveal

{{diff main.go}}

- `screen, err := tcell.NewScreen()` keeps the error. `if err != nil` is the idiom: a `nil` error means "nothing went wrong".
- `if err := screen.Init(); err != nil {` does two things in one line: the statement before the semicolon runs first and declares `err` only for this `if`. It is the standard shape for "call, then check".
- `fmt.Println("error:", err)` prints the error (an `error` knows how to print itself) and `return` leaves `main`, which ends the program. `fmt` comes from the standard library, so the import block now lists two packages; `gofmt` keeps standard-library imports in a group above third-party ones.

--- end

%%% Make `Init` fail on purpose: run with `TERM=nonsense go run .`. tcell cannot find that terminal in its database and `Init` returns an error, which now shows as one clean line.

??? If the compiler says `tcell.NewScreen` has a different signature than expected, the import path is missing the `/v2`. Check the import line.
