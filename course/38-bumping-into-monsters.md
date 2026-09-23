# Step 38 · Bumping into monsters

### The problem

Walking *into* a monster should attack it; walking into an empty cell should move. Both actions look at the same thing first, "what is on the cell next to me in this direction?", so there are now three actions (move, attack, and the bump that decides between them) sharing a direction and the helpers that go with it. Copying `DX, DY` and the destination arithmetic into each is the obvious way; Go's way to share fields and methods without inheritance is **embedding**, and this is the step where it earns its place.

There is also a small practical problem: there is no message log yet and we cannot `fmt.Println` while tcell owns the screen. The engine will keep the last few messages and draw them under the map, a stand-in until chapter 7.

>>> 1. In `actions.go`, declare a struct `ActionWithDirection` with fields `DX, DY int`.
>>> 2. Add a method `func (a ActionWithDirection) Dest(entity *Entity) (int, int)` and a method `func (a ActionWithDirection) BlockingEntity(engine *Engine, entity *Entity) *Entity`.
>>> 3. Redeclare `MovementAction` as `struct{ ActionWithDirection }` (embedding) and use `a.Dest(entity)` in its `Perform`.
>>> 4. Declare `type MeleeAction struct{ ActionWithDirection }` with a `Perform` that logs "You kick the <name>, much to its annoyance!" when something blocks the destination.
>>> 5. Declare `type BumpAction struct{ ActionWithDirection }` with a `Perform` that runs `MeleeAction` if something blocks the destination and `MovementAction` otherwise.
>>> 6. In `input.go`, make the four movement cases return `BumpAction{ActionWithDirection{dx, dy}}`.
>>> 7. In `engine.go`, add a field `Messages []string`, a method `func (e *Engine) Log(msg string)` that keeps the last four, and draw them under the map in `Render`.

!!! Walk into an orc: "You kick the Orc, much to its annoyance!" appears under the map.

--- reveal

{{diff actions.go}}

- `type MeleeAction struct{ ActionWithDirection }` **embeds** the struct: `MeleeAction` has no fields of its own but gets `DX`, `DY`, `Dest` and `BlockingEntity` as if they were its own.
- `BumpAction` builds the other action from its own embedded struct: `MeleeAction{a.ActionWithDirection}`.

{{diff input.go}}

{{diff engine.go}}

- `e.Messages[1:]` re-slices from the second element, dropping the oldest.

--- end

%%% Change `BumpAction` to always perform `MovementAction`. Bumping a monster now silently fails (the entity check in movement refuses it). The bump is the only place that decides *what a collision means*; everything else just reports one.
