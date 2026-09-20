# Step 15 · Visible and explored
## Chapter: Field of view

Seeing the whole dungeon at once spoils exploration. A roguelike tracks two things per tile: is it **visible** right now, and has it ever been **explored**. Visible tiles are drawn lit, explored ones dark, the rest hidden. This step adds the bookkeeping with a deliberately naive rule, "everything within 8 cells is visible, walls or not"; the next step replaces the rule with real line of sight.

Tiles get a lit look, and there is a glyph for the unknown:

{{diff tiles.go}}

The map tracks visibility in two slices indexed exactly like `Tiles`:

{{diff gamemap.go}}

- `Visible` and `Explored` are `[]bool`, one flag per tile, made alongside `Tiles`.
- `setVisible` marks a tile visible **and** explored. It is the only place that writes `Explored`, so memory can never get out of sync with sight. Its name starts lower-case: a helper.
- `ComputeFOV` clears the old visibility and then lights a square around the viewer. `for i := range m.Visible` with only an index is the idiom for "every position of the slice".
- `Render` now chooses the glyph with a **tagless `switch`**: the first `case` whose condition is true wins, so visible beats explored beats the shroud.

The engine recomputes the field of view after every action, and `main.go` does it once before the first frame (otherwise the first screen would be black until you pressed a key):

{{diff engine.go}}

{{diff main.go}}

!!! Run it: a lit square around you, dark blue where you have been, black elsewhere. Walls do not block your view yet: you can see into neighbouring rooms through the rock.
