# Step 21 · Bumping into monsters

Walking *into* a monster should attack it. Three actions now share the idea of "the tile next to me in some direction", so that part is factored out and **embedded**, Go's way of sharing fields and methods without inheritance. `actions.go` is rewritten:

{{diff actions.go}}

- `ActionWithDirection` holds the direction and two helpers: `Dest` (the target tile) and `BlockingEntity` (whatever solid thing stands there).
- `type MeleeAction struct{ ActionWithDirection }` **embeds** it: `MeleeAction` has no fields of its own but gets `DX`, `DY`, `Dest` and `BlockingEntity` as if they were its own. Same for `MovementAction` and `BumpAction`.
- `MeleeAction.Perform` only logs a message for now. `fmt.Sprintf` formats a string; `%s` is replaced by the name.
- `BumpAction` is what the keys produce: it looks at the destination and turns itself into a melee or a movement by building the other action from its own embedded struct: `MeleeAction{a.ActionWithDirection}`.

The keys return bump actions now:

{{diff input.go}}

There is no message log yet, and we cannot `fmt.Println` while tcell owns the screen: it would draw over the game. So the engine keeps the last four messages and draws them under the map, a stand-in until chapter 7:

{{diff engine.go}}

- `e.Messages[1:]` re-slices from the second element: the oldest message drops off.

!!! Run it: walk into an orc: "You kick the Orc, much to its annoyance!" appears under the map.
