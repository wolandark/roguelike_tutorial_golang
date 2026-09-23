# Step 47 · Dealing damage

### The problem

Bumping into a monster only kicks it. Now it should hurt. The formula is the simplest one that works: damage is the attacker's power minus the defender's defense, applied only when it is positive.

There is a second question: what can be attacked? Only *living fighters*. A corpse is not a target, and neither will be a potion lying on the floor. So the map gets a lookup that returns only those, and the bump decides between attacking and moving with it instead of with "anything blocking".

>>> 1. In `gamemap.go`, add a method `func (m *GameMap) GetActorAt(x, y int) *Entity` that returns the entity on that cell that has a `Fighter` and is alive, or `nil`.
>>> 2. In `actions.go`, add a method `func (a ActionWithDirection) TargetActor(engine *Engine, entity *Entity) *Entity` that calls `GetActorAt` on the destination.
>>> 3. Rewrite `MeleeAction.Perform`: find the target with `TargetActor`, compute `entity.Fighter.Power - target.Fighter.Defense`, log "<attacker> attacks <target> for <n> hit points." and apply it with `SetHP` when it is positive, or log "... but does no damage." otherwise.
>>> 4. Change `BumpAction.Perform` to decide with `TargetActor` instead of `BlockingEntity`.

!!! Orcs die in two hits and leave a red `%`; trolls take four. The monsters still do not fight back.

--- reveal

{{diff gamemap.go}}

{{diff actions.go}}

- `%d` in a format string is replaced by an integer.
- Because `BumpAction` now asks for a living actor, bumping a corpse walks onto it.

--- end

%%% Give the troll template `Defense: 5`. Your power is 5, so every hit "does no damage" and trolls cannot be killed. That is why raising attack is one of the level-up choices later.
