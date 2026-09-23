# Step 76 · Dropping

### The problem

`d` should show the same menu, and the chosen item should land on the floor where you stand. With `OnSelect` from step 75 that needs no new handler, only a new action and a new function to pass in.

Dropping is the reverse of picking up: remove the item from the inventory, put it at the owner's position and append it to the map's entities. The inventory owns that logic, so it becomes a method on `Inventory`, and the action only calls it.

>>> 1. In `inventory.go`, add a method `func (inv *Inventory) Drop(engine *Engine, owner, item *Entity)` that removes `item`, sets its position to the owner's, appends it to `engine.GameMap.Entities` and logs "You dropped the ...".
>>> 2. In `actions.go`, declare a struct `DropItem` with a field `Item *Entity` and a `Perform` that calls `entity.Inventory.Drop(engine, entity, a.Item)` and returns `nil`.
>>> 3. In `input.go`, add a function `dropItem` with the same signature as `useItem` that returns `DropItem{Item: item}, nil`.
>>> 4. In `MainGameEventHandler.HandleEvent`, add a `case 'd'` that opens the menu with the title "Select an item to drop" and `dropItem`.

!!! Pick up a potion, walk a few steps, press `d` then `a`: "You dropped the Health Potion." and a `!` appears under the `@` when you step off.

--- reveal

{{diff inventory.go}}

{{diff actions.go}}

{{diff input.go}}

--- end

%%% Add a third menu bound to `x` with the title "Select an item to examine" and an `OnSelect` that logs the item's name and returns `nil, nil`. It takes a few lines and no new type. A `nil` action goes through `runAction`, which returns the menu itself, so the menu stays open.
