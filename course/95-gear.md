# Step 95 · Gear
## Chapter: Equipment

### In this chapter

Daggers, swords, leather armour and chain mail: items that are worn rather than used up, adding to attack and defense. Two steps: make gear exist and get picked up, then make it count. The design decision here is one place where this course knowingly departs from the Python tutorial, because of how saving works in Go.

### The problem

An item that is worn needs a slot (weapon or armour), bonuses, and a way to know whether it is currently worn. The Python tutorial stores that last fact on the *wearer*, as pointers to the worn items. Those items also live in the inventory, so after a gob round trip (step 84) the wearer's pointers would point at *copies*. Keeping one `Equipped` flag **on the item** means there are no shared pointers and nothing to re-link after loading; "what is in my weapon slot" becomes a scan of the inventory, which is at most 26 items.

>>> 1. Create `equipment.go` with `type EquipmentType int` and the constants `Weapon` and `Armor` using `iota`.
>>> 2. Declare a struct `Equippable` with fields `Type EquipmentType`, `PowerBonus int`, `DefenseBonus int` and `Equipped bool`.
>>> 3. Add methods on `Entity`: `EquippedItem(slot EquipmentType) *Entity`, `ItemIsEquipped(item *Entity) bool`, `PowerBonus() int`, `DefenseBonus() int`, and `ToggleEquip(engine *Engine, item *Entity, addMessage bool)` with a helper `unequip`.
>>> 4. In `entity.go`, add a field `Equippable *Equippable`, copy it in `Spawn`, and add a method `func (e *Entity) IsItem() bool`. Use `IsItem` in `PickupAction`.
>>> 5. In `entity_factories.go`, add the templates `dagger`, `sword`, `leatherArmor` and `chainMail`; add `sword` on floor 4 and `chainMail` on floor 6 in `itemChances`.

!!! On deeper floors you can find and pick up `/` and `[`. Selecting one in the inventory does nothing sensible yet.

--- reveal

{{file equipment.go}}

{{diff entity.go}}

{{diff actions.go}}

{{diff entity_factories.go}}

{{diff procgen.go}}

--- end

%%% Equip an item, save, load, and check `ItemIsEquipped` still says yes (it does: the flag travelled with the item). Then imagine the Python design: the wearer's pointer would now point at a copy that is *not* in the inventory. That is the whole reason for the flag.
