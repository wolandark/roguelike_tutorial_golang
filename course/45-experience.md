# Step 45 · Experience

Kills award experience, and enough experience earns a level. One `Level` component serves both sides: monsters only set `XPGiven`; the player has the thresholds. Choosing what to improve is the next step; this one does the bookkeeping.

Create `level.go`:

{{file level.go}}

- `ExperienceToNextLevel` is `base + level*factor`: 350 XP for level 2, 500 more for level 3, and so on. The formula is the whole balancing lever.
- `AddXP` does nothing for entities without thresholds (monsters), otherwise adds, logs, and announces a pending level up.
- `increaseLevel` spends the XP of one level (leftover carries over). The three `Increase...` methods each apply a bonus, log, and call it.

{{diff entity.go}}

{{diff fighter.go}}

- When something dies, its `XPGiven` goes to the player. `engine.Player.Level != nil` guards against a player without the component.

{{diff entity_factories.go}}

- Orcs are worth 35, trolls 100. The player starts at level 1 with base 200 and factor 150.

!!! Run it: kill an orc: "You gain 35 experience points." After 350: "You advance to level 2!", but nothing happens yet. `c` for the character sheet comes next step.
