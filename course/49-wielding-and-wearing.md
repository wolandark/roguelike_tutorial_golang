# Step 49 · Wielding and wearing

The last step: attack and defense become base values plus equipment bonuses, selecting gear in the inventory equips it, the menu marks worn items with `(E)`, and you start with a dagger and leather armour.

{{diff fighter.go}}

- `Defense` and `Power` are renamed `BaseDefense` and `BasePower`, and `Power()` and `Defense()` become **methods on `Entity`**, because they need the inventory to add the bonuses. Level-ups raise the base values.

{{diff entity_factories.go}}

- The player's own stats drop to power 2 / defense 1; the starting dagger (+2) and leather armour (+1) bring them back to 4 / 2.

{{diff actions.go}}

- Melee uses the methods now. `EquipAction` is one call to `ToggleEquip`.

{{diff inventory.go}}

- Dropping something you wear takes it off first.

{{diff input.go}}

- `useItem` returns an `EquipAction` for gear, and asks the consumable otherwise. The inventory shows `(E)`. The level-up menu shows base values (that is what it changes); the character sheet shows totals.

{{diff level.go}}

{{diff setup_game.go}}

- The starting kit: `Spawn(nil, ...)` copies a template without placing it; the two items go straight into the inventory and are equipped silently.

!!! Run it: `i` shows `(a) Dagger (E)` and `(b) Leather Armor (E)`; `c` shows attack 4, defense 2. Find a sword on floor 4 or deeper and equip it: "You remove the Dagger." then "You equip the Sword." in one turn.

## You built a roguelike

About 2,300 lines of Go, every one of them added for a reason you have read. Some directions from here: doors and traps (new `Tile` values plus a little `procgen`), an archer AI that keeps its distance (reuse `FindPath`), a ranged weapon for the player (reuse `SelectIndexHandler`), a version number in `saveData`, or a graphical front end: the engine only touches tcell through `tcell.Screen` and `tcell.Event`, so an [Ebiten](https://ebitengine.org/) renderer with tiles is within reach. Now go make it *yours*. 🗡️🌸
