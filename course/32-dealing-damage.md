# Step 32 · Dealing damage

### The problem

Melee should hurt. The formula the tutorial uses is the simplest that works: damage is attacker power minus defender defense, only applied if positive. Two things need deciding around it. You can only attack *living fighters*, not corpses or the potions of chapter 8, so the map needs a lookup that returns only those. And the player needs to see their HP, which for now is a line of text under the map.

>>> 1. In `gamemap.go`, add a method `func (m *GameMap) GetActorAt(x, y int) *Entity` that returns the living entity with a `Fighter` on that cell.
>>> 2. In `actions.go`, add a method `func (a ActionWithDirection) TargetActor(engine *Engine, entity *Entity) *Entity`.
>>> 3. Rewrite `MeleeAction.Perform` to compute `Power - Defense`, log the result and apply it with `SetHP`.
>>> 4. Change `BumpAction.Perform` to decide with `TargetActor` instead of `BlockingEntity`.
>>> 5. In `engine.go`, draw `HP: x/y` on row 46 and move the messages to column 21.

!!! You have 30 HP, power 5, defense 2. Orcs (10 HP) die in two hits and leave a red `%`. Trolls take four. The monsters do not fight back yet.

--- reveal

{{diff gamemap.go}}

{{diff actions.go}}

- `%d` formats an integer. `BumpAction` now asks for a living actor rather than any blocker, so bumping a corpse walks over it.

{{diff engine.go}}

- The messages move to column 21 to leave room for the health bar of chapter 7.

--- end

%%% Give the troll template `Defense: 5`. Your power is 5, so every hit "does no damage" and trolls are invulnerable. Now you know why the level-up menu in chapter 11 offers attack as an option.
