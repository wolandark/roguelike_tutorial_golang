# Step 40 · Confusion

A confusion scroll makes one enemy stumble around at random. Two ideas meet here. Because behaviour is an interface, confusing a monster is a matter of **swapping its AI** for one that remembers the old one. And because `GetAction` may return a handler, the scroll can ask for a target with the cursor from the last step and build its action from the answer.

{{diff ai.go}}

- `ConfusedEnemy` wraps the previous AI and counts down. It has state, so its `Perform` has a pointer receiver and it is used as `*ConfusedEnemy`; `HostileEnemy` stays a value. Each turn it performs a `BumpAction` in a random direction, so a confused orc happily attacks another orc. At zero it puts `PreviousAI` back; from then on nothing references the wrapper and Go's garbage collector frees it.
- `directions` is a slice of the eight neighbours; `rand.IntN(len(directions))` picks one.

{{diff consumable.go}}

- `ConfusionConsumable.GetAction` returns `nil, handler`: no action yet, switch to the targeting handler instead. The **callback** it passes is a closure: it captures `item` and builds the `ItemAction` with the chosen coordinates once the player confirms. The handler knows nothing about scrolls.
- `Activate` validates: the tile must be visible, hold a living actor, and not be the player. Then one assignment swaps the AI.

{{diff input.go}}

- `NewSingleRangedAttackHandler` wires the cursor to `runAction`, so an `Impossible` target keeps you in aiming mode with the scroll intact. The small dance with `var h` first lets the closure refer to the handler it is part of.

{{diff colors.go}}

{{diff entity_factories.go}}

{{diff procgen.go}}

- A tagless `switch` with an init statement: `roll` is drawn once and compared against thresholds.

!!! Run it: find a purple `~`. Use it: "Select a target location." and the cursor appears; put it on an orc and press Enter. Watch the orc wander for ten turns, then "The Orc is no longer confused." Confuse an orc standing next to another orc and enjoy the brawl.
