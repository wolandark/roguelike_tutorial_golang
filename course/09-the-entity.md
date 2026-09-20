# Step 9 · The entity
## Chapter: Entities and the map

Two integers for the player will not scale to a dungeon full of orcs. Anything with a position and a glyph becomes an **Entity**: the player, a yellow bystander for now, monsters and items later.

Create `entity.go`:

{{file entity.go}}

- `type Entity struct { ... }` groups fields into one type. `Color tcell.Color` is tcell's colour type; named colours like `tcell.ColorWhite` exist, and later we will mix our own.
- `func (e *Entity) Move(dx, dy int)` is a **method**: a function attached to a type through the *receiver* `(e *Entity)`. Inside, `e` is the entity it was called on, and because the receiver is a **pointer** (`*Entity`), changes to `e.X` change the caller's entity. With a value receiver `(e Entity)` the method would work on a copy and the caller would see nothing.

Then use it in `main.go`:

{{diff main.go}}

- `&Entity{X: ..., Char: '@', Color: tcell.ColorWhite}` builds a struct value and `&` takes its address, so `player` is a `*Entity`, a pointer. Fields left out get their zero value.
- `entities := []*Entity{npc, player}` is a *slice* of pointers. The slice and the `player` variable point at the **same** entity: move it through one, and the other sees it. That is why the world will always keep pointers.
- `tcell.StyleDefault.Foreground(e.Color)` gives each entity its colour when drawing. `Foreground` returns a *new* style; styles are values.
- `player.Move(action.DX, action.DY)` replaces the two `+=` lines.

!!! Run it: your white `@` and a yellow `@` five cells to the left. You can walk through it; nothing checks for other entities yet.
