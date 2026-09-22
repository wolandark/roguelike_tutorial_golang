# Step 52 · Experience

### The problem

Kills should be worth something. Monsters carry a value, the player accumulates it, and past a threshold that grows with each level the player is owed an upgrade. One `Level` component covers both sides: monsters only set `XPGiven`; the player has the thresholds. The threshold formula, `base + level*factor`, is the entire balancing lever, so it lives in one method.

>>> Create `level.go` with `Level{CurrentLevel, CurrentXP, LevelUpBase, LevelUpFactor, XPGiven}`, `ExperienceToNextLevel`, `RequiresLevelUp`, `AddXP(engine, xp)` (no-op for entities without thresholds), `increaseLevel`, and three `Increase...` methods (max HP, power, defense). Add the component to `Entity` (copied in `Spawn`), award the corpse's XP to the player in `Die`, and set the templates: orcs 35, trolls 100, the player level 1 with base 200 and factor 150.

!!! Kill an orc: "You gain 35 experience points." After 350: "You advance to level 2!", but nothing happens yet.

--- reveal

{{file level.go}}

{{diff entity.go}}

{{diff fighter.go}}

{{diff entity_factories.go}}

--- end

%%% Set the player's `LevelUpFactor` to 0. Every level costs 200 XP, forever; after a few floors you are unkillable. Then set it to 500 and see how far you get. The formula is where the game's pacing lives.
