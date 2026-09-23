# Step 8 · The vi keys

### The problem

Roguelike players expect `h j k l` next to the arrows. The character keys arrive differently from the arrows: tcell reports a printable character with `ev.Key()` equal to `tcell.KeyRune` and the character itself in `ev.Rune()`. So each direction now has two triggers of different kinds, and the obvious move is to bolt the second one onto the existing case:

```go
case tcell.KeyUp || ev.Rune() == 'k':
```

Try it. The compiler refuses:

```
invalid operation: tcell.KeyUp || ev.Rune() == 'k' (mismatched types tcell.Key and untyped bool)
```

In `switch ev.Key()` a `case` is a **value** that gets compared with `ev.Key()`, not a condition. `tcell.KeyUp || ...` tries to `||` a `tcell.Key` with a boolean, which is meaningless, and even if it were allowed the result would be a `bool`, not a key to compare against. `case tcell.KeyUp, tcell.KeyRune:` lists two *values*, which is the only thing a valued case can do, and it cannot look at the rune.

Go's `switch` has a second form for exactly this: written without a value, each `case` is a full boolean expression, and the first true one runs. The comparison with `ev.Key()` moves into the cases, and the rune check can sit next to it with `||`.

>>> 1. In the up case, try `case tcell.KeyUp || ev.Rune() == 'k':` and read the compiler error.
>>> 2. Change `switch ev.Key() {` to a tagless `switch {`.
>>> 3. Rewrite each case as a condition: `case ev.Key() == tcell.KeyUp || ev.Rune() == 'k':`, and the same with `j` for down, `h` for left, `l` for right.
>>> 4. Rewrite the quit case as `case ev.Key() == tcell.KeyEscape || ev.Key() == tcell.KeyCtrlC:`.

!!! The `@` moves with the arrows and with `h j k l`. Escape or Ctrl-C quits.

--- reveal

{{diff main.go}}

- `switch {` with no value is a chain of conditions; nothing falls through, so the first matching case is the only one that runs.
- `ev.Rune()` is only meaningful for printable characters. For special keys such as the arrows it returns 0, so `== 'k'` is simply false there; that is what makes mixing the two kinds in one condition safe.

--- end

%%% Put `case tcell.KeyUp || ev.Rune() == 'k':` back into a `switch ev.Key()` and read the compiler message once more, slowly. `mismatched types tcell.Key and untyped bool` is Go telling you which two things you tried to combine. Learning to read that line is worth more than this step.

%%% Now the mirror image: inside the tagless switch, write the quit case the old way, `case tcell.KeyEscape, tcell.KeyCtrlC:`. It does not compile either: `tcell.KeyEscape` is a `tcell.Key`, and every case of a tagless switch must be a `bool`. The two forms of `switch` do not mix, in either direction.
