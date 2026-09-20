# Step 2 · Checking errors

Both `NewScreen` and `Init` can fail: no terminal attached, an unknown `TERM`, a broken pipe. Go reports failure with a returned `error` value, and the convention is to check it right away.

{{diff main.go}}

- `screen, err := tcell.NewScreen()` now keeps the error. `if err != nil` is the Go idiom: an error value that is `nil` means "nothing went wrong".
- `if err := screen.Init(); err != nil {` does two things in one line: the part before the semicolon runs first and declares `err` only for this `if`. It is the standard shape for "call, then check".
- `fmt.Println("error:", err)` prints the error (an `error` knows how to print itself) and `return` leaves `main`, which ends the program. `fmt` comes from the standard library, so the import block now lists two packages, and `gofmt` keeps standard-library imports in a group above third-party ones.

!!! Run it: it looks exactly like step 1. To see the error path, run it with no terminal: `go run . < /dev/null | cat` prints `error: ...`.

??? Forgetting the `/v2` in the import path is the most common way to see an error here: the compiler complains that `tcell.NewScreen` has a different signature. Check the import line.
