# Step 46 · Levelling up

When the player has enough XP, a window asks which attribute to improve, and unlike every other menu it **cannot be dismissed**. `c` shows a character sheet.

{{diff input.go}}

- `runAction` gains one more check after the enemies' turn: if the player `RequiresLevelUp`, the next handler is the `LevelUpHandler`. A level up therefore appears right after the killing blow's turn resolves.
- `LevelUpHandler` draws three choices (`a` +20 max HP, `b` +1 power, `c` +1 defense). Any other key logs "Invalid entry." and **stays**; there is no `Parent` to return to. Each choice calls the matching `Increase...` method and returns to the main game.
- `CharacterScreenHandler` (`c` key) is a plain parent-then-window screen with level, XP, XP to next level, attack and defense; any key closes it.

!!! Run it: kill a few orcs. At 350 XP the level-up window appears and will not go away until you pick `a`, `b` or `c`. Press `c` in the game to see your sheet.
