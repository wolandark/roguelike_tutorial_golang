# Step 7 · Reading the arrow keys

Now the update phase: arrow keys change the position, Escape quits.

{{diff main.go}}

- `ev, ok := screen.PollEvent().(*tcell.EventKey)` combines polling and the type assertion. If it is not a key, `continue` jumps straight to the next loop iteration (which redraws).
- `switch ev.Key()` branches on which key it was. tcell names special keys: `tcell.KeyUp`, `KeyDown`, `KeyLeft`, `KeyRight`, `KeyEscape`, `KeyCtrlC`. Go's `switch` does not fall through; each `case` is its own branch, and `case tcell.KeyEscape, tcell.KeyCtrlC:` lists two values for one branch.
- `playerY--` for *up*: row 0 is at the top, so up means a smaller y.
- Why handle Ctrl-C ourselves? In raw mode the terminal no longer turns it into a signal; it arrives as an ordinary key. Treating it as quit is friendlier than trapping the player.

!!! Run it: the `@` moves with the arrows. You can walk off the screen; there is no map yet. Escape or Ctrl-C quits.
