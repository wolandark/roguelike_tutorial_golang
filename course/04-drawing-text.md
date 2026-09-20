# Step 4 · Drawing text

### The problem

`SetContent` draws one character. We want to write a sentence. A string in Go is a sequence of *bytes*, and a character like `é` or `┤` is more than one byte, so "one byte per cell" would put later characters in the wrong columns. We need to walk the string character by character.

>>> Write `Hello, roguelike! Press any key to quit.` on row 1, starting at column 1, one character per cell. Hint: `range` over `[]rune(msg)`, not over `msg`.

!!! The greeting on row 1, the `@` still on row 5.

--- reveal

{{diff main.go}}

- `msg := "..."` declares a variable with `:=`; Go infers the type (`string`).
- `[]rune(msg)` converts the string to a slice of runes, one per character. `for i, r := range` gives the index and the rune. Ranging over the string itself would give **byte** offsets.
- `1+i` puts the first character in column 1 and each next one one cell further right.

--- end

%%% Change the message to `Héllo` and range over `msg` instead of `[]rune(msg)`. The `é` is followed by a gap: it is two bytes, so `i` jumps by two. Put `[]rune` back and the gap disappears.
