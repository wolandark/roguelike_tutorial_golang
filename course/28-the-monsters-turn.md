# Step 28 · The monsters' turn

### The problem

Monsters never act. In a turn-based game the rule is: the player acts, then every other entity acts, then the screen is redrawn. The place for that is right after the player's action in `HandleEvent`. Real behaviour needs hit points and a brain, which is the next chapter; for now each monster just announces that it would like a turn, which proves the loop works and shows the order of events.

>>> Add `HandleEnemyTurns` to the engine that logs a complaint for every entity except the player, and call it after the player's action and before the field of view is recomputed.

!!! Every move fills the message area with grumbling monsters.

--- reveal

{{diff engine.go}}

- `ent != e.Player` compares pointers: the player is skipped. `engine.go` imports `fmt` now; the import block gains a standard-library group.

--- end

%%% Move `HandleEnemyTurns` *before* the player's action. The messages look the same, but from chapter 6 on the difference is real: monsters would move before you do, and a monster next to you would hit you before your attack lands. Order of turns is a design decision; roguelikes let the player go first.
