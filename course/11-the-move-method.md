# Step 11 · The Move method

### The problem

`player.X += action.DX` and `player.Y += action.DY` are the entity moving itself, written out in the loop. Monsters will move too, later, from a different place in the code. Behaviour that belongs to a type is attached to it as a **method**, so it is written once and called from anywhere: `player.Move(dx, dy)`.

There is one decision to make, and it is the one Go beginners get wrong most often: does the method work on the caller's entity, or on a copy?

>>> 1. In `entity.go`, add a method `func (e *Entity) Move(dx, dy int)` that adds `dx` to `e.X` and `dy` to `e.Y`.
>>> 2. In `main.go`, replace the two `+=` lines in the `MovementAction` case with `player.Move(action.DX, action.DY)`.

!!! Still moves exactly as before.

--- reveal

{{diff entity.go}}

- `func (e *Entity) Move(dx, dy int)` is a **method**: a function attached to a type through the *receiver* `(e *Entity)` written before the name. Inside, `e` is the entity the method was called on.
- The receiver is a **pointer** (`*Entity`), so `e.X += dx` changes the caller's entity. With a value receiver, `(e Entity)`, the method would receive a *copy*, change the copy, and the caller would see nothing.

{{diff main.go}}

--- end

%%% Change the receiver to `(e Entity)`, no star. It compiles, and the `@` never moves: the method moved a copy that was thrown away. Put the star back. This is the single most common Go mistake in a game, and now you have seen it once.
