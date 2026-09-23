# Step 61 · Confusion

### The problem

A confusion scroll makes one enemy stumble around at random for ten turns, then behave normally again. Two ideas from earlier steps meet. Because behaviour is an interface (step 48), confusing a monster is **swapping its AI** for one that remembers the old one and hands control back when its timer runs out. And because `GetAction` may return a handler (step 56), the scroll can open the cursor from step 60 and build its item action from the chosen cell, through a closure that remembers which item is being used.

>>> 1. In `ai.go`, declare a struct `ConfusedEnemy` with fields `PreviousAI AI` and `TurnsRemaining int`, and a method with a pointer receiver, `func (c *ConfusedEnemy) Perform(engine *Engine, entity *Entity) error`, that bumps in a random direction and restores `PreviousAI` when the turns run out.
>>> 2. In `input.go`, add a function `func NewSingleRangedAttackHandler(engine *Engine, callback func(x, y int) Action) *SelectIndexHandler` that runs the callback's action through `runAction`.
>>> 3. In `consumable.go`, declare a struct `ConfusionConsumable` with `NumberOfTurns int`. Its `GetAction` returns `nil` and a `NewSingleRangedAttackHandler`; its `Activate` checks the target and replaces its AI with a `*ConfusedEnemy`.
>>> 4. In `colors.go`, add `colorNeedsTarget` and `colorStatusEffectApplied`; in `entity_factories.go` add a `confusionScroll` template; in `procgen.go` give it a 15% chance.

!!! Find a purple `~`. Use it: "Select a target location." and the cursor appears; put it on an orc and press Enter. Watch the orc wander for ten turns, then "The Orc is no longer confused." Confuse an orc next to another orc and enjoy the brawl.

--- reveal

{{diff ai.go}}

- `ConfusedEnemy` has state, so its `Perform` has a pointer receiver and it is used as `*ConfusedEnemy`; `HostileEnemy` stays a value. Once it restores `PreviousAI`, nothing references the wrapper and the garbage collector frees it.

{{diff consumable.go}}

- `GetAction` returns `nil, handler`: no action yet, switch to aiming. The callback is a **closure**: it captures `item` and builds the action once the player confirms.

{{diff input.go}}

- The small dance with `var h` first lets the closure refer to the handler it is part of, so an `Impossible` target keeps you aiming with the scroll intact.

{{diff colors.go}}

{{diff entity_factories.go}}

{{diff procgen.go}}

- A tagless `switch` with an init statement: `roll` is drawn once and compared against thresholds.

--- end

%%% Confuse the same orc twice. The second scroll wraps the *confused* AI, so when the inner timer ends the orc is still confused by the outer one. Decide whether that is a bug; the chapter exercise in the full course asks you to make a second scroll extend the timer instead.
