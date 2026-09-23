# Step 45 · Potions

### The problem

An item is an entity lying on the floor until someone carries it. Two new abilities, so two new components: an **Inventory** for entities that can carry, and a **Consumable** for entities that can be used up. The consumable is the interesting design: what a potion *does* is the potion's business, not the inventory's, so `Consumable` is an interface. It answers two questions: what should happen when the player selects me (usually "perform an item action now", but chapter 9 will need "ask for a target first", so the answer can be an action *or* a handler), and what my effect is when activated.

This step puts potions on the floor and defines the components; picking them up is the next step, so nothing changes in play yet.

>>> 1. Create `inventory.go` with a struct `Inventory` (`Capacity int`, `Items []*Entity`), a method `Remove(item *Entity)` and a method `Drop(engine *Engine, owner, item *Entity)`.
>>> 2. Create `consumable.go` with an interface `Consumable` with two methods: `GetAction(engine *Engine, consumer, item *Entity) (Action, EventHandler)` and `Activate(engine *Engine, action ItemAction, consumer *Entity) error`.
>>> 3. In the same file, add a helper `func consume(consumer, item *Entity)`, a struct `HealingConsumable` with `Amount int`, and its two methods (heal, or `Impossible` at full health).
>>> 4. In `entity.go`, add the fields `Inventory *Inventory` and `Consumable Consumable`, and make `Spawn` copy the inventory and its slice.
>>> 5. In `fighter.go`, add a method `func (f *Fighter) Heal(amount int) int` that returns how much was recovered.
>>> 6. In `actions.go`, declare a struct `ItemAction` (`Item *Entity`, `TargetX, TargetY int`) with `TargetActor` and a `Perform` that calls `Item.Consumable.Activate`.
>>> 7. In `gamemap.go`, add the methods `EntityAt(x, y int) *Entity` and `RemoveEntity(entity *Entity)`.
>>> 8. In `entity_factories.go`, give the player `Inventory: &Inventory{Capacity: 26}` and add a `healthPotion` template.
>>> 9. In `procgen.go`, add a method `func (r RectangularRoom) randomTile() (int, int)`, add a `maxItems` parameter to `placeEntities` and `maxItemsPerRoom` to `GenerateDungeon`, and place potions. In `main.go`, add the constant and pass it.
>>> 10. In `colors.go`, add `colorHealthRecovered`.

!!! Purple `!` in some rooms. You can walk over them; nothing happens yet.

--- reveal

{{file inventory.go}}

- `Remove` deletes from a slice while keeping order: `append(items[:i], items[i+1:]...)`.

{{file consumable.go}}

- `HealingConsumable.Activate` refuses with `Impossible` at full health, so a wasted potion costs neither the potion nor the turn, and calls `consume` only on success.

{{diff entity.go}}

{{diff fighter.go}}

{{diff actions.go}}

{{diff entity_factories.go}}

{{diff gamemap.go}}

{{diff procgen.go}}

{{diff main.go}}

{{diff colors.go}}

--- end

%%% Remove the inventory copy from `Spawn` and (after step 47) pick up a potion: the template's inventory grows too, and every later player would start with it. Same trap as the Fighter, one component later.
