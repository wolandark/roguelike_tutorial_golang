# Step 30 · Dealing damage

### The problem

Melee should hurt. The formula the tutorial uses is the simplest that works: damage is attacker power minus defender defense, only applied if positive. Two things need deciding around it. You can only attack *living fighters*, not corpses or the potions of chapter 8, so the map needs a lookup that returns only those. And the player needs to see their HP, which for now is a line of text under the map.

>>> Add `GetActorAt(x, y)` to the map (an entity with a Fighter that is alive) and a `TargetActor` helper on `ActionWithDirection`. Make `MeleeAction` compute the damage, log one of two messages, and apply it through `SetHP`. Make `BumpAction` decide based on `TargetActor`. Draw `HP: x/y` on row 46.

!!! You have 30 HP, power 5, defense 2. Orcs (10 HP) die in two hits and leave a red `%`. Trolls take four. The monsters do not fight back yet.

--- reveal

{{diff gamemap.go}}

{{diff actions.go}}

- `%d` formats an integer. `BumpAction` now asks for a living actor rather than any blocker, so bumping a corpse walks over it.

{{diff engine.go}}

- The messages move to column 21 to leave room for the health bar of chapter 7.

--- end

%%% Give the troll template `Defense: 5`. Your power is 5, so every hit "does no damage" and trolls are invulnerable. Now you know why the level-up menu in chapter 11 offers attack as an option.
