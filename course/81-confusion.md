# Step 81 · Confusion

### The problem

The confusion scroll makes the player aim, then swaps the target's AI. Aiming is the cursor from step 78. Because `GetAction` may return a handler instead of an action (step 72), the scroll can return a cursor handler. When the player confirms a cell, the cursor's callback builds the `ItemAction` with that cell as its target and runs it through `runAction` from step 75.

The callback needs to know which item is being used, but the cursor handler only passes `x, y`. A **closure** solves that: the function literal written inside `GetAction` can use `item` from around it, and keeps it for as long as the function exists.

>>> 1. In `input.go`, add a function `func NewSingleRangedAttackHandler(engine *Engine, callback func(x, y int) Action) *SelectIndexHandler` that creates a select handler whose `OnSelect` returns `runAction(engine, h, callback(x, y))`, where `h` is the handler itself.
>>> 2. In `consumable.go`, declare a struct `ConfusionConsumable` with a field `NumberOfTurns int`. Its `GetAction` logs "Select a target location." in `colorNeedsTarget` and returns `nil` and a `NewSingleRangedAttackHandler` whose callback builds `ItemAction{Item: item, TargetX: x, TargetY: y}`.
>>> 3. Its `Activate` returns `Impossible` when the target cell is not visible, has no actor, or is the consumer. Otherwise it logs a message in `colorStatusEffectApplied`, sets `target.AI = &ConfusedEnemy{PreviousAI: target.AI, TurnsRemaining: c.NumberOfTurns}` and consumes the item.
>>> 4. In `colors.go`, add `colorNeedsTarget` and `colorStatusEffectApplied`. In `entity_factories.go`, add a `confusionScroll` template. In `procgen.go`, replace the `if` with a `switch` so a potion spawns 70% of the time, a confusion scroll 15% and a lightning scroll 15%.

!!! Find a purple `~` and use it: "Select a target location." and the cursor appears. Put it on an orc and press Enter. The orc wanders for ten turns, then "The Orc is no longer confused." Confuse an orc that stands next to another orc and watch them fight.

--- reveal

{{diff input.go}}

- `var h *SelectIndexHandler` declares `h` before the handler exists, so that the function literal can refer to it. By the time the player confirms, `h` holds the handler. With `h` as `self`, an `Impossible` target keeps you aiming, with the scroll still in your pack.

{{diff consumable.go}}

- `GetAction` returns `nil, handler`: no action yet, switch to aiming.

{{diff colors.go}}

{{diff entity_factories.go}}

{{diff procgen.go}}

- `switch roll := rand.Float64(); {` is a `switch` with an init statement and no value after it. `roll` is drawn once, and each `case` is a condition, tested from top to bottom.

--- end

%%% Confuse the same orc twice. The second scroll wraps the confused AI, so when the second timer ends the orc goes back to the first confusion, not to hunting. Decide whether that is a bug, and how you would make a second scroll extend the timer instead.
