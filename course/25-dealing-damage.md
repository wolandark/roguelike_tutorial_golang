# Step 25 · Dealing damage

Melee becomes arithmetic: attacker power minus defender defense. You can only attack *living fighters*, so the map gets a lookup for those, and your HP is shown under the map.

{{diff gamemap.go}}

- `GetActorAt` returns only entities that have a `Fighter` and are alive. Corpses are neither attackable nor in the way.

{{diff actions.go}}

- `TargetActor` is a third helper on `ActionWithDirection`. `MeleeAction` uses it, computes the damage, logs one of two messages (`%d` formats an integer) and applies it through `SetHP`, which handles death.
- `BumpAction` now asks for a living actor rather than any blocker, so bumping a corpse walks over it.

{{diff engine.go}}

- The HP text goes on row 46, and the messages move to column 21 to leave room for the health bar of chapter 7.

!!! Run it: you have 30 HP, power 5, defense 2. Orcs (10 HP) die in two hits and leave a red `%`. Trolls take four. The monsters still do not fight back.
