# Step 73 · Picking up

### The problem

Standing on a potion, `g` should put it in your pack. Picking up is an action like moving, so it follows the rules from step 69. If there is nothing to pick up, or the inventory is full, it returns `Impossible` and no turn passes. Otherwise it moves the entity from the map's list into the inventory's list, logs a message, and costs a turn.

"Something to pick up" means an entity on the player's cell that has a `Consumable`. That excludes corpses, which lie on cells too. Taking it off the map needs a way to delete one entity from `m.Entities`.

>>> 1. In `gamemap.go`, add a method `func (m *GameMap) RemoveEntity(entity *Entity)` that deletes `entity` from `m.Entities`, the same way `Inventory.Remove` does.
>>> 2. In `actions.go`, declare `type PickupAction struct{}` with a `Perform` that loops over the map's entities and skips every entity that has no `Consumable` or is not on the player's cell.
>>> 3. For the first entity it does not skip: return `Impossible{"Your inventory is full."}` if `len(Items)` has reached `Capacity`. Otherwise remove it from the map, append it to the inventory, log "You picked up the ...!" and return `nil`. After the loop, return `Impossible{"There is nothing here to pick up."}`.
>>> 4. In `input.go`, in `handleKey`, turn the `if` for `.` and `5` into a `switch` on `ev.Rune()`, and add a `case 'g'` that returns `PickupAction{}`.

!!! Stand on a `!` and press `g`: "You picked up the Health Potion!" and the `!` is gone. Press `g` again: "There is nothing here to pick up." in grey, and no monster moves.

--- reveal

{{diff gamemap.go}}

{{diff actions.go}}

- The loop returns as soon as it has handled one item, so removing from `m.Entities` while ranging over it is safe here: the loop never continues after the removal.

{{diff input.go}}

- `case '.', '5':` matches either rune. A `switch` case can list several values.

--- end

%%% Set the player's inventory capacity to 1 in `entity_factories.go`. Pick up one potion, then stand on another and press `g`: "Your inventory is full." in grey, and no turn passes. Put the capacity back to 26.
