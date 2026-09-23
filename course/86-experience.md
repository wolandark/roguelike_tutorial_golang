# Step 86 · Experience

### The problem

Kills should be worth something. Monsters carry a value, the player accumulates it, and past a threshold that grows with each level the player is owed an upgrade. One `Level` component covers both sides: monsters only set `XPGiven`; the player has the thresholds. The threshold formula, `base + level*factor`, is the entire balancing lever, so it lives in one method.

>>> 1. Create `level.go` with a struct `Level` (`CurrentLevel`, `CurrentXP`, `LevelUpBase`, `LevelUpFactor`, `XPGiven`, all `int`).
>>> 2. Add the methods `ExperienceToNextLevel() int`, `RequiresLevelUp() bool`, `AddXP(engine *Engine, xp int)` and a helper `increaseLevel()`.
>>> 3. Add three methods that apply a level: `IncreaseMaxHP`, `IncreasePower` and `IncreaseDefense`, each taking `(engine *Engine, entity *Entity, amount int)`.
>>> 4. In `entity.go`, add a field `Level *Level` and copy it in `Spawn`. In `fighter.go`, award the dead entity's `XPGiven` to the player in `Die`.
>>> 5. In `entity_factories.go`, give the player `Level{CurrentLevel: 1, LevelUpBase: 200, LevelUpFactor: 150}`, orcs `XPGiven: 35`, trolls `XPGiven: 100`.

!!! Kill an orc: "You gain 35 experience points." After 350: "You advance to level 2!", but nothing happens yet.

--- reveal

{{file level.go}}

{{diff entity.go}}

{{diff fighter.go}}

{{diff entity_factories.go}}

--- end

%%% Set the player's `LevelUpFactor` to 0. Every level costs 200 XP, forever; after a few floors you are unkillable. Then set it to 500 and see how far you get. The formula is where the game's pacing lives.
