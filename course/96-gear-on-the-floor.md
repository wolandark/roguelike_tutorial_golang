# Step 96 · Gear on the floor
## Chapter: Equipment

### In this chapter

Daggers, swords, leather armour and chain mail: items that are worn rather than used up, adding to attack and defense. First gear exists and can be picked up, then it can be worn, one item per slot, and finally it changes the numbers. One design decision here departs from the Python tutorial, because of how saving works in Go.

### The problem

A wearable item needs a slot (weapon or armour), its bonuses, and whether it is worn right now. That is one more component, `Equippable`.

Where does "worn" live? The Python tutorial stores it on the *wearer*, as pointers to the worn items. Those items are also in the inventory, and after a `gob` round trip (step 83) the wearer's pointers would point at *copies*, the same problem as the player in step 83. A flag `Equipped` **on the item** avoids that: no shared pointers, nothing to re-link after loading.

Picking up must accept gear too. Until now "an item" meant "has a `Consumable`". A small method, `IsItem`, will say "has a `Consumable` or an `Equippable`".

>>> 1. Create `equipment.go` with `type EquipmentType int` and the constants `Weapon` and `Armor` using `iota`. Declare a struct `Equippable` with the fields `Type EquipmentType`, `PowerBonus int`, `DefenseBonus int` and `Equipped bool`.
>>> 2. In `entity.go`, add a field `Equippable *Equippable` and copy it in `Spawn`. Add a method `func (e *Entity) IsItem() bool` that returns true when the entity has a `Consumable` or an `Equippable`.
>>> 3. In `actions.go`, in `PickupAction.Perform`, skip entities for which `IsItem()` is false instead of checking `Consumable`.
>>> 4. In `entity_factories.go`, add the templates `dagger`, `sword`, `leatherArmor` and `chainMail`. In `procgen.go`, add `sword` with weight 5 on floor 4 and `chainMail` with weight 15 on floor 6 to `itemChances`.

!!! From floor 4 on you can find a sword, `/`, and from floor 6 chain mail, `[`, and pick them up with `g`. Don't select them in the inventory yet: `useItem` expects a `Consumable`, and gear has none, so the game panics.

--- reveal

{{file equipment.go}}

- `iota` numbers the constants 0, 1, ..., as for `RenderOrder` in step 49. Giving them their own type, `EquipmentType`, keeps a slot from being mixed up with any other `int`.

{{diff entity.go}}

{{diff actions.go}}

{{diff entity_factories.go}}

{{diff procgen.go}}

--- end

%%% Change `sword` in `itemChances` to floor 0 with weight 500, pick up a sword on floor 1, press `i` and select it. The game crashes with a nil pointer dereference: `useItem` calls `GetAction` on a `nil` `Consumable`. The next step handles gear there. Put the table back.
