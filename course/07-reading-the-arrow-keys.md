# Step 7 · Reading the arrow keys

### The problem

We have a key event and ignore what key it was. tcell reports keys in two ways: special keys through `ev.Key()`, which returns named constants (`tcell.KeyUp`, `tcell.KeyEscape`, ...), and printable characters through `ev.Rune()`, with `ev.Key()` set to `tcell.KeyRune` for those. Roguelike players expect both the arrows and the vi keys `h j k l`, so each direction has two triggers of different kinds, and a plain `switch ev.Key()` cannot express "this key *or* that character". Go's `switch` without a value can: each `case` is then a full condition.

One thing to know about the terminal first: in raw mode Ctrl-C is no longer a signal that kills the program; it arrives as an ordinary key. If we do not handle it, the player is trapped.

>>> Write a tagless `switch` with one case per direction, each matching an arrow *or* a vi key (`ev.Key() == tcell.KeyUp || ev.Rune() == 'k'`), changing `playerX`/`playerY` (row 0 is at the top). Escape and Ctrl-C return. Other keys do nothing.

!!! The `@` moves with the arrows and with `h j k l`. It can walk off the screen, since there is no map yet. Escape or Ctrl-C quits.

--- reveal

{{diff main.go}}

- `ev, ok := screen.PollEvent().(*tcell.EventKey)` combines polling and the assertion; `continue` on a non-key jumps to the next iteration, which redraws.
- `switch {` with no value is a chain of conditions: the first `case` whose expression is true runs, and nothing falls through. `case ev.Key() == tcell.KeyUp || ev.Rune() == 'k':` therefore reads "up arrow or the k key".
- `ev.Rune()` is only meaningful for printable characters; for special keys it returns 0, so `== 'k'` is simply false there. That is what makes mixing the two kinds safe.
- `playerY--` for *up*, because up means a smaller row number.

--- end

%%% Remove `|| ev.Key() == tcell.KeyCtrlC` from the last case and press Ctrl-C while the game runs. Nothing happens: that is raw mode. Escape still works. Put it back.

%%% Try to write the up case as `case tcell.KeyUp || ev.Rune() == 'k':` inside a `switch ev.Key()`. The compiler refuses: a `case` of a valued switch is a value to compare, not a condition, and `tcell.KeyUp || ...` tries to `||` a key with a boolean. The tagless switch exists for exactly this.
