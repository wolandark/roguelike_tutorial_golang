# Step 63 · Levelling up

### The problem

When the player has enough XP, the game should ask which attribute to improve, and unlike every other menu the question cannot be dismissed. Where does that interruption belong? After the enemies' turn in `runAction`, next to the game-over check: both are "the state changed, switch modes". A character sheet on `c` is the same window pattern as before.

>>> 1. In `input.go`, in `runAction`, after the game-over check, return a `LevelUpHandler` when `engine.Player.Level.RequiresLevelUp()`.
>>> 2. Declare a struct `LevelUpHandler` with a field `Engine *Engine`, an `OnRender` showing three choices, and a `HandleEvent` that applies `a`, `b` or `c` and logs "Invalid entry." for anything else, staying open.
>>> 3. Declare a struct `CharacterScreenHandler` with fields `Engine *Engine` and `Parent EventHandler`, an `OnRender` showing level, XP and stats, and a `HandleEvent` that returns `Parent` on any key. Open it for `c`.

!!! Kill a few orcs. At 350 XP the level-up window appears and will not go away until you pick `a`, `b` or `c`. `c` in the game shows your sheet.

--- reveal

{{diff input.go}}

- `LevelUpHandler` has no `Parent`, on purpose: there is nothing to return to until a choice is made.

--- end

%%% Move the level-up check *before* the game-over check in `runAction`. Kill an orc with your last hit while an adjacent troll kills you in the same turn: you get to level up as a corpse, then the game-over screen never appears because the level-up handler returns to the main game. Order of checks again.
