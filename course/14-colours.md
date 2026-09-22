# Step 14 · Colours

### The problem

Dots for floor are a fine tradition, but this game draws its floor as **colour**: a space with a coloured background, so the cell's character stays free for whatever stands on it. A look is a character plus two colours, and it will be needed for every kind of tile, so it gets a small type of its own, `Glyph`, with a method that turns it into the `tcell.Style` that `SetContent` wants.

>>> Create `tiles.go` with a `Glyph` (rune, foreground, background) and a `Style()` method, and one package-level `floorGlyph`: a space on a blue background. Draw the rectangle with it.

!!! The rectangle is solid blue now, and the `@` is drawn on it with the default black background, which the next steps fix.

--- reveal

{{file tiles.go}}

- `Style()` has a *value* receiver, `(g Glyph)`: it only reads `g`, so no pointer is needed. Compare `Move` in step 11, which had to change its receiver.
- `tcell.NewRGBColor(50, 50, 150)` is a true-colour value; tcell approximates it on terminals with fewer colours.
- `var floorGlyph = ...` is a package-level variable, visible in every file of the package.

{{diff main.go}}

--- end

%%% Change the glyph to `Glyph{'.', tcell.ColorWhite, tcell.NewRGBColor(50, 50, 150)}`: dots on blue. Both conventions exist in real roguelikes; the choice is one line.
