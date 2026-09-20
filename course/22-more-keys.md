# Step 22 · More keys

Roguelike players expect the vi keys (`hjkl` plus `yubn` for diagonals), the numpad, and diagonals on Home/End/PgUp/PgDn. A `switch` with twenty cases would be tedious; two lookup tables do it, and adding a binding becomes one line.

{{diff input.go}}

- `map[tcell.Key][2]int` is a map from a key to a two-element array. `moveKeys[ev.Key()]` looks one up; the two-value form `d, ok := moveKeys[...]` tells us whether the key was in the map at all.
- Printable characters arrive as `tcell.KeyRune` with the character in `ev.Rune()`, so `moveRunes` is keyed by `rune`.
- `.` and `5` wait a turn, which needs a new action:

{{diff actions.go}}

!!! Run it: move with `hjkl`, diagonally with `y u b n`, or the numpad. `.` passes a turn (nothing visible happens yet).
