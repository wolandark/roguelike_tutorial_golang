# Step 35 · Potions

Two new components: an **Inventory** for things that can carry, and a **Consumable** for things that can be used up. A health potion is an entity with a consumable and no fighter. This step puts potions on the floor; picking them up is the next step.

Create `inventory.go`:

{{file inventory.go}}

- `Remove` deletes an item from the slice while keeping order: everything before it, then everything after it, appended together. `inv.Items[:i]` and `inv.Items[i+1:]` are slice expressions; `...` spreads the second into `append`'s arguments.
- `Drop` puts an item back on the map at the owner's feet.

Create `consumable.go`:

{{file consumable.go}}

- A consumable answers two questions. `GetAction`: what should happen when the player selects me in the inventory? For a potion, "perform an `ItemAction` now". The second return value exists because some items (chapter 9) must ask for a target first and return an event handler instead. `Activate` is the effect itself.
- `HealingConsumable` heals through the new `Fighter.Heal`, refuses with `Impossible` at full health, and finally `consume`s itself out of the inventory.

The entity gets both components, and `Spawn` copies the inventory (and its slice) too:

{{diff entity.go}}

`Heal` returns how much was **actually** recovered, so the message can be honest:

{{diff fighter.go}}

`ItemAction` carries the item to use, plus a target position for later:

{{diff actions.go}}

Templates: the player can carry 26 items (one per letter), and here is the potion:

{{diff entity_factories.go}}

The map gets two helpers, `EntityAt` (anything on a tile) and `RemoveEntity` (for pickup, next step):

{{diff gamemap.go}}

Room generation scatters 0 to 2 potions per room:

{{diff procgen.go}}

- `randomTile` factors out the "random floor cell inside this room" arithmetic that was inline before.

{{diff main.go}}

{{diff colors.go}}

!!! Run it: purple `!` in some rooms. You can walk over them; nothing happens yet.
