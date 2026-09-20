# Step 7 · Reading the arrow keys

### The problem

We have a key event and ignore what key it was. tcell tells us through `ev.Key()`, which returns named constants for special keys (`tcell.KeyUp`, `tcell.KeyEscape`, ...). One thing to know about the terminal first: in raw mode Ctrl-C is no longer a signal that kills the program; it arrives as an ordinary key. If we do not handle it, the player is trapped. So we need arrows to move and two keys to quit.

>>> Branch on `ev.Key()`: the four arrows change `playerX`/`playerY` (remember row 0 is at the top), Escape and Ctrl-C return. Other keys do nothing.

!!! The `@` moves with the arrows. It can walk off the screen, since there is no map yet. Escape or Ctrl-C quits.

--- reveal

{{diff main.go}}

- `ev, ok := screen.PollEvent().(*tcell.EventKey)` combines polling and the assertion; `continue` on a non-key jumps to the next iteration, which redraws.
- `switch ev.Key()` branches on the key. Go's `switch` does not fall through: each `case` is its own branch, and `case tcell.KeyEscape, tcell.KeyCtrlC:` lists two values for one branch.
- `playerY--` for *up*, because up means a smaller row number.

--- end

%%% Remove the `tcell.KeyCtrlC` case and press Ctrl-C while the game runs. Nothing happens: that is raw mode. Escape still works. Put it back.
