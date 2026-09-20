# Step 1 · Hello, tcell
## Chapter: Hello, tcell

The whole course grows one program, and this is the seed: ten lines that open the terminal in "game mode", draw an `@`, wait for something to happen and hand the terminal back. Type it into `main.go`.

{{file main.go}}

- `import "github.com/gdamore/tcell/v2"` brings in tcell, the library that talks to the terminal for us. The **`/v2`** at the end is part of the path; without it you get the old, incompatible v1. You need it in your module first: `go get github.com/gdamore/tcell/v2`.
- `tcell.NewScreen()` returns two values: a screen and an error. `screen, _ :=` keeps the first and throws the second away with the blank identifier `_`. Ignoring errors is a bad habit that we fix in the very next step; for now it keeps the program short.
- `screen.Init()` is the call that takes over the terminal: raw mode (keys arrive one at a time, nothing is echoed) and the *alternate screen*, the one `vim` and `less` use, so your shell history is untouched.
- `screen.SetContent(10, 5, '@', nil, tcell.StyleDefault)` puts the character `'@'` at column 10, row 5. Row 0 is the **top** of the screen. `'@'` with single quotes is a `rune`, one character. The `nil` is for combining characters (accents); we never need it. `tcell.StyleDefault` is the terminal's plain colours.
- Nothing is visible until `screen.Show()` copies tcell's buffer to the real terminal.
- `screen.PollEvent()` blocks until the terminal reports *anything*: a key, a mouse move, a resize. We ignore what it was.
- `screen.Fini()` undoes `Init()`. Skip it and your shell is left without echo and with a hidden cursor (if that ever happens: type `reset`, Enter).

!!! Run it: `go run .` (or ▶ play here). A blank screen with an `@`; any key, click or resize ends the program and your shell is back as it was.
