# Step 40 · Sharing the direction

### The problem

Walking into a monster should attack it. That means a second action, an attack, which needs exactly what movement has: a direction, and the cell that direction points at. A third action will soon decide between the two. Copying `DX, DY` and the destination arithmetic into every one of them is the obvious way and the fragile one.

Go has no inheritance, but it has **embedding**: a struct can contain another struct without naming the field, and then the outer struct gets the inner one's fields and methods as if they were its own. So the direction becomes a small struct of its own, and every directional action embeds it. This step does only that, for movement.

>>> 1. In `actions.go`, declare a struct `ActionWithDirection` with fields `DX, DY int`, and a method `func (a ActionWithDirection) Dest(entity *Entity) (int, int)` that returns the entity's position plus the direction.
>>> 2. Change `MovementAction` to `type MovementAction struct{ ActionWithDirection }`, and in its `Perform` compute the destination with `a.Dest(entity)`.
>>> 3. In `input.go`, create movements as `MovementAction{ActionWithDirection{0, -1}}` and so on.

!!! Exactly the same game as before.

--- reveal

{{diff actions.go}}

- `type MovementAction struct{ ActionWithDirection }` **embeds** the struct: the field has no name, only a type. `a.DX`, `a.DY` and `a.Dest(...)` all work on a `MovementAction` directly.
- `entity.Move(a.DX, a.DY)` did not change: the promoted fields `a.DX` and `a.DY` read exactly like the old ones.

{{diff input.go}}

- The composite literal needs the inner struct spelled out: `MovementAction{ActionWithDirection{0, -1}}`.

--- end

%%% Write `MovementAction{DX: 0, DY: -1}` in `input.go`. The compiler says `unknown field DX in struct literal`. Promoted fields can be *read* through the outer struct, but a literal has to build the embedded struct itself.
