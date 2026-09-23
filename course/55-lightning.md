# Step 55 · Lightning
## Chapter: Scrolls and targeting

### In this chapter

Three scrolls: lightning picks its own target, confusion and fireball need you to aim. Aiming needs a cursor mode, which also gives a free "look" command. The Go theme is functions as values: callbacks and closures let one cursor handler serve looking, single-target and area-target scrolls.

### The problem

A lightning scroll strikes the nearest visible enemy within range. The consumable interface from step 52 fits without changes: `GetAction` returns an item action, `Activate` finds the target. Finding "the closest of these" is a loop that keeps the best candidate so far; starting the best distance at "just out of range" means the first in-range candidate always wins and out-of-range ones never do. The map needs one helper, "all living fighters", so scrolls do not filter corpses themselves.

>>> 1. In `gamemap.go`, add a method `func (m *GameMap) Actors() []*Entity` returning every living fighter.
>>> 2. In `consumable.go`, declare a struct `LightningDamageConsumable` with `Damage int` and `MaximumRange int`, and its `GetAction` (an `ItemAction`) and `Activate` (strike the closest visible actor in range, or `Impossible`).
>>> 3. In `entity_factories.go`, add a `lightningScroll` template (`~`, yellow).
>>> 4. In `procgen.go`, spawn a potion 70% of the time and a lightning scroll otherwise.

!!! Pick up a yellow `~`, wait for an orc to come into view, use it from `i`. With nobody near: "No enemy is close enough to strike." in grey, scroll kept.

--- reveal

{{diff gamemap.go}}

{{diff consumable.go}}

- `if d := ...; d < closest` declares `d` for the `if` only.

{{diff entity_factories.go}}

{{diff procgen.go}}

--- end

%%% Remove the `actor == consumer` check. With no enemy in sight, the closest visible actor is you: the scroll strikes the caster. Self-targeting bugs are the classic scroll bug; every targeted effect below has the same guard.
