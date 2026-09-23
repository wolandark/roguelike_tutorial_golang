# Step 97 · Equipping

### The problem

Selecting gear in the inventory should put it on, and selecting it again should take it off. That is an action like any other, `EquipAction`, and it costs a turn. The inventory menu should also show which items are worn, with ` (E)` after the name.

The work of toggling belongs to the wearer: a method `ToggleEquip` on `Entity` flips the item's flag and logs what happened.

>>> 1. In `equipment.go`, add a method `func (e *Entity) ItemIsEquipped(item *Entity) bool`, and a method `ToggleEquip(engine *Engine, item *Entity)` that calls a helper `unequip(engine, item)` when the item is worn, and otherwise sets `Equipped` and logs "You equip the ...".
>>> 2. The helper `func (e *Entity) unequip(engine *Engine, item *Entity)` clears `Equipped` and logs "You remove the ...".
>>> 3. In `actions.go`, declare a struct `EquipAction` with a field `Item *Entity` and a `Perform` that calls `entity.ToggleEquip(engine, a.Item)` and returns `nil`.
>>> 4. In `input.go`, make `useItem` return `EquipAction{Item: item}, nil` for items with an `Equippable`. In `InventoryHandler.OnRender`, build each line in a variable `label` and add ` (E)` when `ItemIsEquipped` is true.

!!! Gear is rare on the first floors. To test, temporarily add `{&sword, 50}, {&chainMail, 50}` to floor 0 of `itemChances`. Pick up a sword, press `i` and select it: "You equip the Sword." and the menu shows `(a) Sword (E)`. Select it again: "You remove the Sword.".

--- reveal

{{diff equipment.go}}

{{diff actions.go}}

{{diff input.go}}

- `label += " (E)"` appends to a string. Strings in Go cannot be changed in place; `+=` builds a new string and assigns it.

--- end

%%% Pick up two swords and equip both. The menu marks both with `(E)`: nothing stops you from wielding two weapons. The next step adds the one-item-per-slot rule.
