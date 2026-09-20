# Step 4 · Drawing text

`SetContent` draws one character. A string is a loop of characters, and a Go string is a sequence of *bytes*, so we convert it to runes first.

{{diff main.go}}

- `msg := "..."` declares a variable with `:=`, letting Go infer the type (`string`).
- `for i, r := range []rune(msg)` loops over the characters. `[]rune(msg)` converts the string to a slice of runes; `range` then gives the index `i` and the rune `r`. Ranging over the string directly would give **byte** offsets, and a character like `┤` is three bytes, so later characters would land in the wrong cells. Ranging over `[]rune` keeps one character per cell.
- `1+i` puts the first character in column 1 and each next one one cell to the right.

!!! Run it: the greeting on row 1 and the `@` still on row 5.
