# Step 57 · Picking up and dropping

### The problem

Standing on a potion should let you take it with `g`, and there should be a way to put it down. Both are actions, so both follow the rules from step 55: nothing to pick up is `Impossible` (no turn), a full inventory is `Impossible`, a successful pickup moves the entity from the map's list into the inventory's list and costs a turn. Dropping is the reverse; it needs a menu to choose what, which is the next step, but the action can exist now.

>>> 1. In `actions.go`, declare `type PickupAction struct{}` with a `Perform` that moves an item on the player's cell from the map into the inventory, or returns `Impossible`.
>>> 2. Declare a struct `DropItem` with a field `Item *Entity` and a `Perform` that calls `entity.Inventory.Drop`.
>>> 3. In `input.go`, return `PickupAction{}` for the `g` key.

!!! Stand on a `!` and press `g`: "You picked up the Health Potion!". Press `g` again: "There is nothing here to pick up." in grey, no turn spent.

--- reveal

{{diff actions.go}}

{{diff input.go}}

- The single `if` for `.` and `5` becomes a `switch` on the rune.

--- end

%%% Set the player's inventory capacity to 1 in the template. The second pickup says "Your inventory is full." and, being `Impossible`, does not cost a turn. Rules compose without new code.
