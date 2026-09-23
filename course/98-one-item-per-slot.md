# Step 98 · One item per slot

### The problem

You can wear only one weapon and one piece of armour. Equipping a sword while holding a dagger should take the dagger off first. For that, the wearer must answer "what is in my weapon slot?". With the flag on the item, the answer is a scan of the inventory, which holds at most 26 items.

Dropping a worn item should take it off first, or a sword lying on the floor would still count as worn. And with slots in place, the total bonus of everything worn is a sum over the slots: that is what the next step adds to attack and defense.

>>> 1. In `equipment.go`, add a method `func (e *Entity) EquippedItem(slot EquipmentType) *Entity` that returns the worn item of that type in the inventory, or `nil`.
>>> 2. In `ToggleEquip`, before equipping, unequip the item that `EquippedItem` returns for the new item's slot, if any.
>>> 3. Add two methods, `PowerBonus() int` and `DefenseBonus() int`, that add up the bonuses of the items in the `Weapon` and `Armor` slots.
>>> 4. In `inventory.go`, in `Drop`, call `owner.ToggleEquip(engine, item)` first when the item is worn.

!!! With the test table from step 97, pick up two swords. Equip one, then the other: "You remove the Sword." and "You equip the Sword." in the same turn, and only one shows `(E)`. Drop the sword: "You remove the Sword." then "You dropped the Sword.".

--- reveal

{{diff equipment.go}}

- `if current := e.EquippedItem(...); current != nil` declares `current` for the `if` only, the shape from step 2.
- `[]EquipmentType{Weapon, Armor}` is a slice literal made on the spot, just to range over it.

{{diff inventory.go}}

--- end

%%% In `EquippedItem`, remove the `item.Equippable.Type == slot` test. Equip a sword, then chain mail: the chain mail replaces the sword, because every worn item now seems to sit in every slot.
