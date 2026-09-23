# Step 72 · Consumables

### The problem

A potion should do something when used. What it does is the potion's business, not the inventory's: a health potion heals, a scroll later will strike or confuse. So "can be used up" is one more component, and because every item behaves differently, the component is an interface, `Consumable`, like `AI` in step 51.

It answers two questions. First, what should happen when the player selects the item? For a potion, "perform an item action now". A scroll that needs aiming will answer "ask for a target first", so the answer is an `Action` *or* an `EventHandler`. Second, what is the effect when the action is performed? That is `Activate`, and it can refuse with `Impossible` from step 69.

The action itself is an `ItemAction`: which item, and which cell it targets. Its `Perform` just asks the item's consumable to activate. Nothing can use an item yet, so this step only builds the pieces.

>>> 1. In `fighter.go`, add a method `func (f *Fighter) Heal(amount int) int` that raises `HP` by `amount` but not above `MaxHP`, and returns how much was actually recovered.
>>> 2. In `actions.go`, declare a struct `ItemAction` with the fields `Item *Entity` and `TargetX, TargetY int`. Give it a method `TargetActor(engine *Engine) *Entity` that returns the actor at the target cell, and a `Perform` that returns `a.Item.Consumable.Activate(engine, a, entity)`.
>>> 3. Create `consumable.go` with an interface `Consumable` with two methods, `GetAction(engine *Engine, consumer, item *Entity) (Action, EventHandler)` and `Activate(engine *Engine, action ItemAction, consumer *Entity) error`. Add a function `func consume(consumer, item *Entity)` that removes the item from the consumer's inventory, and a struct `HealingConsumable` with a field `Amount int` and both methods.
>>> 4. In `entity.go`, add a field `Consumable Consumable` to `Entity`. In `entity_factories.go`, give the potion `Consumable: HealingConsumable{Amount: 4}`. In `colors.go`, add `colorHealthRecovered` (green).

!!! Nothing changes on screen yet. The program compiles.

--- reveal

{{diff fighter.go}}

{{diff actions.go}}

{{file consumable.go}}

- `HealingConsumable.GetAction` returns an `ItemAction` and a `nil` handler: use it right away, no aiming.
- `Activate` returns `Impossible` at full health. The potion is not consumed and, by the rule from step 68, no turn passes. `consume` runs only on success.
- `Consumable` and `ItemAction` refer to each other. Go allows that inside a package: the order of declarations does not matter.

{{diff entity.go}}

- `Consumable` holds a value, not a pointer, like `AI`. `HealingConsumable` has no state that changes, so `Spawn` can copy it as it is.

{{diff entity_factories.go}}

{{diff colors.go}}

--- end

%%% In `consumable.go`, delete the `GetAction` method of `HealingConsumable`. The build fails in `entity_factories.go`: `HealingConsumable does not implement Consumable (missing method GetAction)`. A type satisfies an interface only with every method, and the compiler checks it where the value is stored. Put the method back.
