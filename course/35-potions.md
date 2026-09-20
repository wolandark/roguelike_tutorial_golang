# Step 35 · Potions

### The problem

An item is an entity lying on the floor until someone carries it. Two new abilities, so two new components: an **Inventory** for entities that can carry, and a **Consumable** for entities that can be used up. The consumable is the interesting design: what a potion *does* is the potion's business, not the inventory's, so `Consumable` is an interface. It answers two questions: what should happen when the player selects me (usually "perform an item action now", but chapter 9 will need "ask for a target first", so the answer can be an action *or* a handler), and what my effect is when activated.

This step puts potions on the floor and defines the components; picking them up is the next step, so nothing changes in play yet.

>>> Create `inventory.go` (capacity, items, `Remove`, `Drop`) and `consumable.go` (the `Consumable` interface with `GetAction` returning `(Action, EventHandler)` and `Activate`; a `consume` helper; `HealingConsumable`). Add `Inventory` and `Consumable` fields to `Entity` and copy the inventory in `Spawn`. Add `Fighter.Heal(amount) int` returning what was actually recovered. Add `ItemAction` (item, target position; `Perform` calls `Activate`). Add a `healthPotion` template, `EntityAt` and `RemoveEntity` to the map, and place 0 to 2 potions per room.

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

%%% Remove the inventory copy from `Spawn` and (after step 37) pick up a potion: the template's inventory grows too, and every later player would start with it. Same trap as the Fighter, one component later.
