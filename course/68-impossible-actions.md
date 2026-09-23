# Step 68 · Impossible actions
## Chapter: Items and inventory

### In this chapter

Potions on the floor, an inventory to carry them, a menu to use or drop them. Before any of that, a rule that every roguelike needs and that shapes all later actions: an action that *cannot* be done should say why and must not cost a turn.

### The problem

Walk into a wall in step 67 and the monsters get a free turn. That is unfair, and it gets worse with items: drinking a potion at full health should not waste the potion or the turn. Two designs are possible. `Perform` could return a `bool`, "did something happen?", but then it cannot say *why* not. Or it could return an `error`, which carries a message, and a special error type can mark the "not possible, no turn" case so that the loop can tell it apart from a real failure. Go's `errors.As` is built for exactly that distinction.

>>> 1. Create `exceptions.go` with a struct `type Impossible struct { Msg string }` and a method `func (i Impossible) Error() string` returning `i.Msg`.
>>> 2. In `actions.go`, change the interface method to `Perform(engine *Engine, entity *Entity) error`, and make every `Perform` return an `error`: `nil` on success, `Impossible{"That way is blocked."}` for walls and blockers, `Impossible{"Nothing to attack."}` for melee without a target.
>>> 3. In `ai.go`, change the `AI` interface and `HostileEnemy.Perform` to return `error` the same way.
>>> 4. In `input.go`, add a function `func runAction(engine *Engine, self EventHandler, action Action) EventHandler`: on an `Impossible` (check with `errors.As`) log it in grey and return `self`; on another error log it in red; otherwise run the enemy turns and FOV and return the next handler. End `MainGameEventHandler.HandleEvent` with it.
>>> 5. In `colors.go`, add `colorInvalid`, `colorImpossible` and `colorError`.

!!! Walk into a wall: a grey "That way is blocked." and the monsters do **not** move. Compare with a real step, after which they do.

--- reveal

{{file exceptions.go}}

- Any type with an `Error() string` method satisfies Go's built-in `error` interface.

{{diff actions.go}}

- `return MeleeAction{...}.Perform(...)` passes the inner action's result straight through.

{{diff ai.go}}

{{diff input.go}}

- `errors.As(err, &imp)` checks the error's type and copies it into `imp` if it matches, seeing through wrapped errors too. `runAction` is the one place where the rule lives, and every handler that can produce a player action will end with it.

{{diff colors.go}}

--- end

%%% In `runAction`, replace `errors.As` with `err == (Impossible{"That way is blocked."})`. Only walls are free now; attacking nothing still costs a turn. Comparing errors by type, not by value or message, is what makes the rule general.
