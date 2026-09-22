# Step 56 · Wielding and wearing

### The problem

Gear exists but does nothing. Attack and defense must become base values plus bonuses, and because the bonuses need the inventory, the sums belong on `Entity`, not on `Fighter`. Selecting gear in the inventory should equip or remove it, the menu should mark worn items, dropping something worn should take it off first, and the player should start with a dagger and leather armour so the numbers stay what they were.

>>> Rename the Fighter fields to `BaseDefense`/`BasePower`; add `Power()` and `Defense()` methods on `Entity` that add the bonuses, and use them in melee, the character sheet and the level-up menu (base values there). Add `EquipAction`; make `useItem` return it for equippables; mark worn items `(E)` in the menu; unequip in `Drop`. Lower the player to power 2 / defense 1 and give them a dagger and leather armour, equipped silently, in `NewGame`.

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
