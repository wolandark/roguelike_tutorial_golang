# Step 1 · Hello, tcell
## Chapter: Hello, tcell

### In this chapter

A roguelike is a program that owns the whole terminal: no scrolling text, no prompt, just a grid of cells it draws into and keys it reacts to one at a time. Before any dungeon, we need to be able to do exactly that, and to hand the terminal back in one piece afterwards. This chapter gets there in five small steps, and by the end of it you will have met the shape that every later step keeps: a `run` function that does the work and a `main` that reports.

### The problem

Print `@` with `fmt.Println` and you get an `@` at the bottom of a scrolling log, with a prompt after it. That is not a game screen. We need three things a normal program never asks for: to draw at a *position*, to be told about a key *the moment it is pressed* (no Enter), and to leave no trace when we quit. Terminals can do all of this through escape sequences, but the details differ per terminal and platform, so we use a library that knows them: [tcell](https://github.com/gdamore/tcell).

The question for this step is only: what is the smallest program that takes the terminal over, draws one character, and gives it back?

>>> Make a `main.go` that opens a tcell screen, draws an `@` somewhere, waits for something to happen, and restores the terminal. You will need `go get github.com/gdamore/tcell/v2` first (note the `/v2`). The tcell functions you need are `NewScreen`, and on the screen: `Init`, `SetContent`, `Show`, `PollEvent`, `Fini`. Ignore errors for now. If your program exits on its own before you can see the `@`, you have found the surprise of this step; read on.

!!! A blank screen with a single `@`; the next key, click or resize ends the program and your shell is back exactly as it was.

--- reveal

{{file main.go}}

- `import "github.com/gdamore/tcell/v2"`: the **`/v2`** is part of the path; without it you get the old, incompatible v1.
- `tcell.NewScreen()` returns two values, a screen and an error. `screen, _ :=` keeps the first and throws the second away with the blank identifier `_`. That is a bad habit, and the next step is about why.
- `screen.Init()` takes over the terminal: raw mode (keys arrive one at a time, nothing is echoed) and the *alternate screen*, the one `vim` and `less` use, so your shell history is untouched.
- `screen.SetContent(10, 5, '@', nil, tcell.StyleDefault)` puts the rune `'@'` at column 10, row 5. Row 0 is the **top**. The `nil` is for combining characters (accents), never needed here. `tcell.StyleDefault` is the terminal's plain colours.
- Nothing is visible until `screen.Show()` copies tcell's buffer to the real terminal. tcell keeps a buffer so that it can send only what changed.
- `screen.PollEvent()` blocks until the terminal reports *anything*. So why call it **twice**? Because tcell always delivers one event right after `Init`: a resize, telling you the screen size. A single `PollEvent` returns immediately with that resize and the program ends before you can blink. Calling it twice swallows the resize and then waits for a real event. It is a cute hack, and it is honest about what it does not do: a second resize would end the program too. Step 5 replaces it with a loop that asks *what kind* of event arrived, which is the real fix.
- `screen.Fini()` undoes `Init()`.

--- end

%%% Delete one of the two `screen.PollEvent()` lines and run it. The `@` flashes for a frame and the program is gone: the first event was the initial resize. Put the line back, then resize the window while the program waits: it ends, because the second resize is an event too.

%%% Delete the `screen.Fini()` line and run it. When the program ends your shell has no echo and no cursor. Type `reset` and press Enter to fix it. Now you know what `Fini` is for, and why step 3 will make it impossible to forget.

%%% Change `'@'` to `"@"` (double quotes). The compiler refuses: `"@"` is a `string`, `'@'` is a `rune`, one character, and `SetContent` wants a rune.
