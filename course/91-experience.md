# Step 91 · Experience

### The problem

Kills should be worth something. Each monster is worth some experience points (XP), the player collects them, and past a threshold the player is owed an upgrade. The threshold grows with each level, so that later levels take longer.

One `Level` component covers both sides. Monsters only set `XPGiven`, the XP they are worth. The player sets the threshold fields. The formula for the threshold, `LevelUpBase + CurrentLevel*LevelUpFactor`, is the game's main balancing lever, so it lives in one method.

>>> 1. Create `level.go` with a struct `Level` with the `int` fields `CurrentLevel`, `CurrentXP`, `LevelUpBase`, `LevelUpFactor` and `XPGiven`.
>>> 2. Add the methods `ExperienceToNextLevel() int` (the formula), `RequiresLevelUp() bool` (enough XP for the next level) and `AddXP(engine *Engine, xp int)`, which adds the XP, logs "You gain N experience points." and, when a level is due, "You advance to level N!".
>>> 3. In `entity.go`, add a field `Level *Level` and copy it in `Spawn`, like the `Fighter`. In `fighter.go`, at the end of `Die`, give the dead entity's `XPGiven` to the player when both have a `Level`.
>>> 4. In `entity_factories.go`, give the player `Level: &Level{CurrentLevel: 1, LevelUpBase: 200, LevelUpFactor: 150}`, orcs `Level: &Level{XPGiven: 35}` and trolls `Level: &Level{XPGiven: 100}`.

!!! Start a new game and kill an orc: "You gain 35 experience points." At 350 XP: "You advance to level 2!". Nothing else happens yet.

--- reveal

{{file level.go}}

- The comments after the fields say what each one means. A comment starting with `//` runs to the end of the line.
- `AddXP` returns early for entities without thresholds (`LevelUpBase == 0`), so a monster that somehow gains XP does not level.

{{diff entity.go}}

{{diff fighter.go}}

- `Die` also runs when the player dies. The player has a `Level` but `XPGiven` is 0, and `AddXP` ignores 0.

{{diff entity_factories.go}}

--- end

%%% Give the orc template `Level: &Level{XPGiven: 400}` and kill one orc. Both messages appear at once, because 400 is past the first threshold of 350. Put it back to 35.
