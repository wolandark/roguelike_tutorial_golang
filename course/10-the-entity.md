# Step 10 · The entity
## Chapter: Entities and the map

### In this chapter

Two ideas carry the rest of the game. An **entity** is anything that sits somewhere and has a glyph: the player, a monster, a potion, a staircase. The **map** is a grid of tiles the entities live on. This chapter builds both in small steps: a struct for the entity, a method on it, a second entity, then tiles, then the map, and finally actions that consult the map so that walls stop you. Along the way the loop moves into an **engine**, and `main.go` shrinks to a few lines it keeps until the end.

### The problem

Two integers describe the player. A second thing on screen would need two more, a third two more again, and every one of them wants a glyph and a colour too. Four related values that always travel together are a **struct**. This step only introduces it and moves the player into it; nothing else changes.

>>> Create `entity.go` with an `Entity` struct: `X, Y int`, `Char rune`, `Color tcell.Color`. In `main.go`, replace the two integers with `player := &Entity{...}`, draw `player.Char` at `player.X, player.Y` in `player.Color`, and move by changing `player.X` and `player.Y`.

!!! Looks and plays exactly like step 9. The `@` is now a struct.

--- reveal

{{file entity.go}}

- `type Entity struct { ... }` groups fields into one type. `X, Y int` declares two fields of the same type on one line. `Color tcell.Color` is tcell's colour type.

{{diff main.go}}

- `&Entity{X: ..., Char: '@', Color: tcell.ColorWhite}` is a *struct literal* with named fields, and `&` takes its address, so `player` is a `*Entity`, a pointer to the struct. Fields you leave out get their zero value.
- Reading a field through a pointer needs no special syntax: `player.X` works on a `*Entity` as on an `Entity`.
- `tcell.StyleDefault.Foreground(player.Color)` gives the glyph its colour; `Foreground` returns a *new* style, styles are values.

--- end

%%% Write `player := Entity{...}` without the `&`. Everything still works: `player.X += 1` changes the local struct directly. The `&` earns its keep in the next two steps, when a method and a slice need to share the same player.
