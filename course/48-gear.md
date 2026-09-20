# Step 48 · Gear
## Chapter: Equipment

Daggers, swords, leather armour and chain mail: items that are worn rather than used up. This step makes them exist, lie around and get picked up; the next one makes them do something.

Create `equipment.go`:

{{file equipment.go}}

- `EquipmentType` is another `iota` enumeration: the slot an item goes into.
- `Equippable` is the component for gear: a slot, two bonuses, and `Equipped`. A deliberate departure from the Python tutorial, which keeps pointers to the worn items on the *wearer*. Those same items also live in the inventory, so after a gob round trip (step 42) the wearer's pointers would point at *copies*. Keeping one `Equipped` flag **on the item** means there are no shared pointers and nothing to relink.
- The wearer-side helpers are methods on `Entity`: `EquippedItem` scans the inventory for the worn item of a slot; `PowerBonus` and `DefenseBonus` add up both slots; `ToggleEquip` equips or removes, and equipping into an occupied slot removes the old item first. `addMessage` lets the starting gear be equipped silently.

{{diff entity.go}}

- `IsItem` now means "consumable **or** equippable", and pickup uses it:

{{diff actions.go}}

{{diff entity_factories.go}}

{{diff procgen.go}}

- Swords appear from floor 4 with a low weight, chain mail from floor 6.

!!! Run it: on deeper floors you can find and pick up `/` and `[`. Selecting one in the inventory does nothing sensible yet (it has no consumable), which the next step fixes.
