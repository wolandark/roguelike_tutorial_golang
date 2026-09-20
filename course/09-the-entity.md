# Step 9 · The entity
## Chapter: Entities and the map

### In this chapter

Two ideas carry the rest of the game. An **entity** is anything that sits somewhere and has a glyph: the player, a monster, a potion, a staircase. The **map** is a grid of tiles the entities live on. This chapter builds both, then lets actions consult the map so that walls stop you. Along the way the loop moves into an **engine**, and `main.go` shrinks to a few lines it will keep until the end.

### The problem

Two integers describe the player. A second thing on screen would need two more, a third two more again, and every one of them wants a glyph and a colour too. That is a struct. The design question is what goes *in* it: position, glyph, colour, and, soon, a name and abilities. And the practical question is how several parts of the program can refer to the *same* player: by pointer.

>>> Create `entity.go` with an `Entity` struct (position, `Char rune`, `Color tcell.Color`) and a `Move(dx, dy int)` method that changes the entity's position. In `main.go`, make the player an `*Entity`, add a yellow `@` bystander, keep both in a slice and draw the slice each frame. Movement should call `player.Move`.

!!! Your white `@` and a yellow `@` five cells to its left. You can walk through the yellow one; nothing checks for other entities yet.

--- reveal

{{file entity.go}}

- `type Entity struct { ... }` groups fields into one type. `Color tcell.Color` is tcell's colour type.
- `func (e *Entity) Move(dx, dy int)` is a **method**: a function attached to a type through the *receiver* `(e *Entity)`. Because the receiver is a **pointer**, changes to `e.X` change the caller's entity. With a value receiver `(e Entity)` the method would work on a copy and the caller would see nothing.

{{diff main.go}}

- `&Entity{X: ..., Char: '@', Color: tcell.ColorWhite}` builds a struct value and `&` takes its address, so `player` is a `*Entity`. Fields left out get their zero value.
- `entities := []*Entity{npc, player}` is a *slice* of pointers. The slice and the `player` variable point at the **same** entity: move it through one, the other sees it.
- `tcell.StyleDefault.Foreground(e.Color)` gives each entity its colour; `Foreground` returns a *new* style, styles are values.

--- end

%%% Change the receiver of `Move` to `(e Entity)` (no star). It compiles, and the `@` never moves: the method moved a copy. Put the star back. This is the single most common Go mistake in a game.

%%% Make `entities` a `[]Entity` of values instead of pointers and store `*player` in it. Moving the player no longer moves what is drawn: the slice holds a copy taken at that moment.
