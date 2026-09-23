# Step 30 · Lit tiles

### The problem

Visible cells are drawn with the same colours as before. When the next step adds remembered cells, the two have to look different: the cells you can see now should be bright, the rest dim. A tile therefore needs a second look, a **lit** one, and `Render` should use it for visible cells.

>>> 1. In `tiles.go`, add a field `Light Glyph` to `Tile`, and give `floor` and `wall` a lit glyph: a space on RGB 200, 180, 50 for floor and on RGB 130, 110, 50 for wall.
>>> 2. In `Render`, draw the tile's `Light` glyph for visible cells instead of its `Dark` glyph.

!!! The square around you is now yellow-brown, floor lighter than wall. The rest is black.

--- reveal

{{diff tiles.go}}

- `floor` and `wall` are struct literals with named fields; adding a field to `Tile` only means adding a line to each.

{{diff gamemap.go}}

--- end

%%% Leave out the `Light:` line of `wall`. The field gets its zero value, a `Glyph` with rune 0 and default colours, and lit walls turn into black holes in the lit square. Every tile kind has to define every look.
