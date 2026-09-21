# Step 4 · Drawing text

### The problem

`SetContent` draws one character. We want to write a sentence. A string in Go is a sequence of *bytes*, and a character like `é` or `┤` is more than one byte, so "one byte per cell" would put later characters in the wrong columns. We need to walk the string character by character.

>>> Write `Hello, roguelike! Press any key to quit.` on row 1, starting at column 1, one character per cell. Hint: `range` over `[]rune(msg)`, not over `msg`.

!!! The greeting on row 1, the `@` still on row 5.

--- reveal

{{diff main.go}}

- `msg := "..."` declares a variable with `:=`; Go infers the type (`string`).
- `[]rune(msg)` converts the string to a slice of runes, one per character. `for i, r := range` gives the index and the rune. Ranging over the string itself would give **byte** offsets.
- `1+i` is the column. `SetContent` draws exactly one cell, so the column has to move with the character: `i` is 0 for the first rune, 1 for the second, and so on, and the `1+` only shifts the whole line one cell in from the left edge. Write `SetContent(1, 1, r, ...)` with a fixed column and every rune lands on the same cell, each overwriting the last, leaving just the final `.` on screen. The same applies to a classic `for i := 0; i < len(runes); i++` loop: the index is what makes the text advance.

--- end

%%% Change the message to `Héllo` and range over `msg` instead of `[]rune(msg)`. The `é` is followed by a gap: it is two bytes, so `i` jumps by two. Put `[]rune` back and the gap disappears.
