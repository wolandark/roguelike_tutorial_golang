# Step 42 · Bumping into monsters

### The problem

Now for the attack itself. Walking into a monster should attack it; walking into an empty cell should move. Both begin by looking at the cell in the chosen direction, so there are three directional actions: move, attack, and a **bump** that decides between them. The keys will produce bumps.

>>> 1. In `actions.go`, add a method `func (a ActionWithDirection) BlockingEntity(engine *Engine, entity *Entity) *Entity` that returns the blocking entity on the destination cell, or `nil`.
>>> 2. Declare `type MeleeAction struct{ ActionWithDirection }` with a `Perform` that logs "You kick the <name>, much to its annoyance!" when there is a blocking entity at the destination. Import `fmt` for `fmt.Sprintf`.
>>> 3. Declare `type BumpAction struct{ ActionWithDirection }` with a `Perform` that performs a `MeleeAction` when something blocks the destination and a `MovementAction` otherwise.
>>> 4. In `input.go`, create `BumpAction{ActionWithDirection{0, -1}}` and so on instead of `MovementAction`s.

!!! Walk into an orc: "You kick the Orc, much to its annoyance!" appears under the map.

--- reveal

{{diff actions.go}}

- `MeleeAction` and `BumpAction` embed the same struct as `MovementAction`, so all three get `Dest` and `BlockingEntity` from one definition.
- `BumpAction` builds the other action from its own embedded struct: `MeleeAction{a.ActionWithDirection}`, then calls `Perform` on it.
- `fmt.Sprintf` formats a string the way `fmt.Printf` prints one; `%s` is replaced by the name.

{{diff input.go}}

--- end

%%% Change `BumpAction` to always perform `MovementAction`. Bumping a monster now silently fails (the entity check in movement refuses it). The bump is the only place that decides *what a collision means*; everything else just reports one.
