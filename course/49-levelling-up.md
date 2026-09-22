# Step 49 · Levelling up

### The problem

When the player has enough XP, the game should ask which attribute to improve, and unlike every other menu the question cannot be dismissed. Where does that interruption belong? After the enemies' turn in `runAction`, next to the game-over check: both are "the state changed, switch modes". A character sheet on `c` is the same window pattern as before.

>>> In `runAction`, after the game-over check, return a `LevelUpHandler` when the player `RequiresLevelUp`. The handler draws three choices (`a` +20 max HP, `b` +1 power, `c` +1 defense), logs "Invalid entry." and stays on anything else, and returns to the main game after a choice. Add a `CharacterScreenHandler` (parent, level, XP, next threshold, attack, defense) bound to `c`.

!!! Kill a few orcs. At 350 XP the level-up window appears and will not go away until you pick `a`, `b` or `c`. `c` in the game shows your sheet.

--- reveal

{{diff input.go}}

- `LevelUpHandler` has no `Parent`, on purpose: there is nothing to return to until a choice is made.

--- end

%%% Move the level-up check *before* the game-over check in `runAction`. Kill an orc with your last hit while an adjacent troll kills you in the same turn: you get to level up as a corpse, then the game-over screen never appears because the level-up handler returns to the main game. Order of checks again.
