# Step 99 · Wielding and wearing

### The problem

Gear can be worn but changes nothing. Attack and defense become a base value plus the bonuses of worn gear. The bonuses need the inventory, so the sums are methods on `Entity`, next to `PowerBonus` and `DefenseBonus`, not on `Fighter`. `Fighter` keeps the base values, and levelling up raises those.

The player should start with a dagger (+2 attack) and leather armour (+1 defense), and their base values drop to 2 and 1. That makes attack 4 and defense 2 at the start, one point of attack less than before. Equipping the starting gear should not fill the log with messages, so `ToggleEquip` gets a flag that turns them off.

>>> 1. In `fighter.go`, rename the fields `Power` and `Defense` to `BasePower` and `BaseDefense`, and add two methods on `Entity`, `Power() int` and `Defense() int`, that add the equipment bonus to the base value.
>>> 2. Follow the compiler errors: in `actions.go`, compute melee damage as `entity.Power() - target.Defense()`. In `level.go`, raise the base fields. In `input.go`, show the base values in the level-up menu and `Power()` and `Defense()` on the character sheet. In `entity_factories.go`, rename the fields and set the player to `BasePower: 2, BaseDefense: 1`.
>>> 3. In `equipment.go`, add a parameter `addMessage bool` to `ToggleEquip` and `unequip`, and log only when it is true. Pass `true` in `EquipAction` and `Drop`.
>>> 4. In `setup_game.go`, in `NewGame`, spawn a dagger and leather armour with `Spawn(nil, 0, 0)`, append them to the player's inventory and equip them with `addMessage` set to `false`.

!!! Delete `savegame.sav`: a save from before this step stores `Power` and `Defense`, which no longer exist, so its base values would load as 0. Start a new game. `i` shows `(a) Dagger (E)` and `(b) Leather Armor (E)`, and `c` shows attack 4, defense 2. Equip a sword: attack 6.

--- reveal

{{diff fighter.go}}

- `Power` and `Defense` are now methods on `Entity`, while `BasePower` and `BaseDefense` are fields on `Fighter`. A field and a method cannot share a name on the same type, which is one reason for the rename.

{{diff actions.go}}

{{diff level.go}}

{{diff input.go}}

{{diff entity_factories.go}}

{{diff equipment.go}}

{{diff inventory.go}}

{{diff setup_game.go}}

- `[]Entity{dagger, leatherArmor}` is a slice of template values. `tpl.Spawn(nil, 0, 0)` copies each one without placing it on a map, as in step 89.

--- end

%%% Give the dagger `DefenseBonus: 1` too. `Defense()` sums both slots, so the dagger now protects as well, with no change to the sums.

## You built a roguelike

About 2,000 lines of Go, and you argued your way to every one of them before seeing it. Some directions from here: doors and traps (new `Tile` values and a little `procgen`), an archer AI that keeps its distance (reuse `FindPath`), a ranged weapon for the player (reuse `SelectIndexHandler`), a version number in `saveData`, or a graphical front end: the engine only touches tcell through `tcell.Screen` and `tcell.Event`, so an [Ebiten](https://ebitengine.org/) renderer with tiles is within reach. Now go make it *yours*. 🗡️
