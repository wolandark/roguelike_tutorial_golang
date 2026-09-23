# Step 59 · Wielding and wearing

### The problem

Gear exists but does nothing. Attack and defense must become base values plus bonuses, and because the bonuses need the inventory, the sums belong on `Entity`, not on `Fighter`. Selecting gear in the inventory should equip or remove it, the menu should mark worn items, dropping something worn should take it off first, and the player should start with a dagger and leather armour so the numbers stay what they were.

>>> 1. In `fighter.go`, rename the fields to `BaseDefense` and `BasePower`, and add two methods on `Entity`: `Power() int` and `Defense() int`, each adding the equipment bonus.
>>> 2. In `entity_factories.go`, set the player to `BasePower: 2, BaseDefense: 1` and rename the fields in the other templates.
>>> 3. In `actions.go`, use `entity.Power() - target.Defense()` in melee, and declare a struct `EquipAction` with a field `Item *Entity` whose `Perform` calls `ToggleEquip`.
>>> 4. In `inventory.go`, unequip an item in `Drop` before removing it.
>>> 5. In `input.go`, return an `EquipAction` from `useItem` for equippable items, add ` (E)` to worn items in the menu, and use the base values in the level-up menu and `Power()`/`Defense()` on the character sheet. In `level.go`, raise the base fields.
>>> 6. In `setup_game.go`, give the new player a dagger and leather armour and equip them silently.

!!! `i` shows `(a) Dagger (E)` and `(b) Leather Armor (E)`; `c` shows attack 4, defense 2. Find a sword on floor 4 or deeper and equip it: "You remove the Dagger." then "You equip the Sword." in one turn.

--- reveal

{{diff fighter.go}}

{{diff entity_factories.go}}

{{diff actions.go}}

{{diff inventory.go}}

{{diff input.go}}

{{diff level.go}}

{{diff setup_game.go}}

- `Spawn(nil, ...)` copies a template without placing it; the two items go straight into the inventory.

--- end

%%% Give the dagger `DefenseBonus: 1` too. `Defense()` sums both slots, so the dagger now also protects, without any change to the sums: the slot loop already covers every worn item.

## You built a roguelike

About 2,300 lines of Go, and you argued your way to every one of them before seeing it. Some directions from here: doors and traps (new `Tile` values and a little `procgen`), an archer AI that keeps its distance (reuse `FindPath`), a ranged weapon for the player (reuse `SelectIndexHandler`), a version number in `saveData`, or a graphical front end: the engine only touches tcell through `tcell.Screen` and `tcell.Event`, so an [Ebiten](https://ebitengine.org/) renderer with tiles is within reach. Now go make it *yours*. 🗡️🌸
