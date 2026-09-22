# Step 23 · More keys

### The problem

Arrows and `hjkl` cannot move diagonally, and roguelike players expect `yubn` for that, the numpad, and Home/End/PgUp/PgDn. That is twenty bindings; the tagless switch from step 8 would grow to twenty conditions, and adding one means finding the right spot. A map from key to direction is shorter, and adding a binding is one line. Two maps, because tcell reports special keys through `Key()` and printable characters through `Rune()`, which is the same split the `||` in each case has been bridging. Waiting a turn also needs a key, and therefore an action.

>>> Replace the switch in `handleKey` with two lookup tables: `moveKeys map[tcell.Key][2]int` for arrows and the four diagonal special keys, `moveRunes map[rune][2]int` for `hjklyubn` and the numpad digits. Add a `WaitAction` and return it for `.` and `5`.

!!! Move with `hjkl`, diagonally with `y u b n`, or the numpad. `.` passes a turn (nothing visible happens yet).

--- reveal

{{diff input.go}}

- `map[tcell.Key][2]int` maps a key to a two-element array. The two-value lookup `d, ok := moveKeys[...]` tells us whether the key was in the map at all.

{{diff actions.go}}

--- end

%%% Add `'w': {0, -1}, 'a': {-1, 0}, 's': {0, 1}, 'd': {1, 0}` to `moveRunes` for WASD. One line each, nothing else changes. (You will want to remove `d` again in chapter 8, where it drops items.)
