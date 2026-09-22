# Step 13 · Tiles

### The problem

The world is empty black. Floors and walls are the next thing, and before there is a map to hold them we need to decide what one **tile** is. A tile has to answer three questions later on: can you stand on it, can you see through it (chapter 5), and how does it look. The look is a character plus colours, and roguelike maps are mostly colour: a floor tile and a wall tile are both a space, with different backgrounds. This step defines tiles and draws a few by hand; the map comes next.

>>> Create `tiles.go` with a `Glyph` (rune, foreground, background) and a `Style()` method turning it into a `tcell.Style`; a `Tile` with `Walkable`, `Transparent` and a `Dark` glyph; and two package-level tiles, `floor` and `wall`, both a space with a blue background, the wall darker. In `main.go`, draw a strip of twenty floor tiles on row 30 with one wall tile in the middle.

!!! A blue strip below the `@` with a darker cell in its middle.

--- reveal

{{file tiles.go}}

- `Style()` has a *value* receiver, `(g Glyph)`: it only reads `g`, so no pointer is needed. Compare `Move` in step 11, which had to change its receiver.
- `tcell.NewRGBColor(50, 50, 150)` is a true-colour value; tcell approximates it on terminals with fewer colours.
- `var ( ... )` declares package-level variables. `floor` and `wall` are *values*; anything that copies them gets its own copy, so a map cell can never change these.

{{diff main.go}}

--- end

%%% Give the wall a glyph of `'#'` instead of a space and a white foreground. Now walls read as walls without any colour; that is how most terminal roguelikes draw them, and the choice is one line here.
