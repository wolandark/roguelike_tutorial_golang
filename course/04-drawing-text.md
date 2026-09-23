# Step 4 · Drawing text

### The problem

`SetContent` draws one character. We want to write a sentence. A string in Go is a sequence of *bytes*, and `range` over a string gives byte offsets. For plain ASCII that is harmless: every character is one byte, so the offsets are the character positions and `for i, r := range msg` draws the sentence correctly. It stops being harmless the moment a message contains `é`, `┤` or an emoji, which are two to four bytes each: the offset jumps and a gap appears. Since the game will draw box-drawing characters in chapter 7, we walk the string character by character from the start, by converting it to runes.

>>> 1. In `run`, after the `SetContent` for the `@`, declare `msg := "Hello, roguelike! Press any key to quit."`.
>>> 2. Add a loop `for i, r := range []rune(msg)` that calls `screen.SetContent(1+i, 1, r, nil, tcell.StyleDefault)`.

!!! The greeting on row 1, the `@` still on row 5.

--- reveal

{{diff main.go}}

- `msg := "..."` declares a variable with `:=`; Go infers the type (`string`).
- `[]rune(msg)` converts the string to a slice of runes, one per character. `for i, r := range` gives the index and the rune. Ranging over the string itself gives **byte** offsets, which are the same numbers for ASCII and different for anything else.
- `1+i` is the column. `SetContent` draws exactly one cell, so the column has to move with the character: `i` is 0 for the first rune, 1 for the second, and so on, and the `1+` only shifts the whole line one cell in from the left edge. Write `SetContent(1, 1, r, ...)` with a fixed column and every rune lands on the same cell, each overwriting the last, leaving just the final `.` on screen. The same applies to a classic `for i := 0; i < len(runes); i++` loop: the index is what makes the text advance.

--- end

%%% Range over `msg` instead of `[]rune(msg)` and run: with the ASCII message nothing changes. Now put an accent in it, `Héllo, roguelike!`, and run again: the `é` is followed by a gap, because it is two bytes and `i` jumps by two. Put `[]rune` back and the gap disappears. Both halves matter: the bug only exists for non-ASCII text, which is exactly why it is easy to ship.
