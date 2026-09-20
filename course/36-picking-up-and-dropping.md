# Step 36 · Picking up and dropping

Two actions: `g` picks up whatever item is on your tile; dropping needs a menu (next step) but the action exists now.

{{diff actions.go}}

- `PickupAction` loops over the map's entities looking for an item (an entity with a consumable) on the player's tile. Capacity full is `Impossible`; otherwise the item leaves the map (`RemoveEntity`) and joins the inventory. Nothing found is also `Impossible`, so a wasted `g` costs no turn.
- `DropItem` hands over to `Inventory.Drop`.

{{diff input.go}}

- The single `if` for `.` and `5` becomes a `switch` on the rune, with `g` added.

!!! Run it: stand on a `!` and press `g`: "You picked up the Health Potion!". Press `g` again: "There is nothing here to pick up." in grey, no turn spent.
