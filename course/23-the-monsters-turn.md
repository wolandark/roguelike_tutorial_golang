# Step 23 · The monsters' turn

After the player acts, every other entity should get a turn. The real behaviour comes in the next chapter; for now they just complain, which proves the loop works.

{{diff engine.go}}

- `HandleEnemyTurns` runs after the player's action and before the field of view is recomputed. `ent != e.Player` compares pointers: the player is skipped.
- The message is built with `fmt.Sprintf`, so `engine.go` imports `fmt` now; the import block gains a standard-library group above the tcell line.

!!! Run it: every move fills the message area with grumbling monsters. Satisfying enough; in the next chapter they get teeth.
