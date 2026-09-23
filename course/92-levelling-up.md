# Step 92 · Levelling up

### The problem

When the player has enough XP, the game should ask which attribute to improve: more hit points, more attack or more defense. Unlike every other window, this question cannot be dismissed. Each choice raises one stat, logs a message, and moves the player to the next level, keeping any XP left over.

Where does the interruption belong? In `runAction`, after the enemies' turn and next to the game-over check. Both mean "the state changed, switch modes".

>>> 1. In `level.go`, add a method `increaseLevel()` that subtracts `ExperienceToNextLevel()` from `CurrentXP` and then adds 1 to `CurrentLevel`.
>>> 2. Add three methods, `IncreaseMaxHP`, `IncreasePower` and `IncreaseDefense`, each with the parameters `(engine *Engine, entity *Entity, amount int)`. Each raises one `Fighter` stat by `amount` (`IncreaseMaxHP` raises `HP` too), logs a message and calls `increaseLevel`.
>>> 3. In `input.go`, declare a struct `LevelUpHandler` with a field `Engine *Engine`. Its `OnRender` draws the game and a framed window with the three choices `a`, `b` and `c`. Its `HandleEvent` applies the chosen increase and returns the main game handler, or logs "Invalid entry." for any other letter and stays open.
>>> 4. In `runAction`, after the game-over check, return `&LevelUpHandler{Engine: engine}` when `engine.Player.Level.RequiresLevelUp()`.

!!! Start a new game and kill orcs until 350 XP. The level-up window appears, and only `a`, `b` or `c` closes it. Choose `a`: "Your health improves!" and the bar shows the new maximum.

--- reveal

{{diff level.go}}

- `increaseLevel` starts with a lower-case letter, so by Go's convention it is meant for use inside the package only. The three exported methods are what the menu calls.
- The order in `increaseLevel` matters: the XP is subtracted with the threshold of the level just finished, then the level goes up.

{{diff input.go}}

- `LevelUpHandler` has no `Parent`, on purpose: there is nothing to return to until a choice is made.
- The check after `!engine.Player.Alive` means a dead player never sees the level-up window.

--- end

%%% In `LevelUpHandler.HandleEvent`, make the `default` case return `&MainGameEventHandler{Engine: h.Engine}`. Any key now closes the window, but after your next move it opens again: `RequiresLevelUp` is still true, and `runAction` checks it after every turn. The XP stays owed until it is spent.
