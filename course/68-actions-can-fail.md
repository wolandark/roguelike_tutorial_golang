# Step 68 · Actions can fail
## Chapter: Items and inventory

### In this chapter

Potions on the floor, an inventory to carry them, a menu to use or drop them. Before any of that, a rule that every roguelike needs and that shapes all later actions: an action that *cannot* be done should say why and must not cost a turn.

### The problem

Walk into a wall in step 67: nothing moves, but the monsters still get their turn. The handler cannot tell a real move from a refused one, because `Perform` returns nothing. It will matter more with items: drinking a potion at full health should not waste the turn.

`Perform` could return a `bool`, "did something happen?", but then it cannot say *why* not. Go's usual answer is to return an `error`. `error` is a built-in interface with one method, `Error() string`, and `nil` means "no error". `errors.New("some text")` makes an error that carries that text.

>>> 1. In `actions.go`, change the `Action` interface method to `Perform(engine *Engine, entity *Entity) error`. Add `error` to every `Perform` method and make each one end with `return nil`.
>>> 2. In `MovementAction.Perform`, return `errors.New("That way is blocked.")` instead of the plain `return`s. In `MeleeAction.Perform`, return `errors.New("Nothing to attack.")` when there is no target. In `BumpAction.Perform`, return the result of the action it calls.
>>> 3. In `ai.go`, change the `AI` interface and `HostileEnemy.Perform` the same way: return `error`, and return what the called action returns.
>>> 4. In `colors.go`, add `colorImpossible` (grey). In `input.go`, when `Perform` returns an error, log `err.Error()` in that colour and return `h` before the enemy turns.

!!! Walk into a wall: a grey "That way is blocked." and the monsters do **not** move. Take a real step and they do.

--- reveal

{{diff actions.go}}

- `errors.New` lives in the `errors` package, so `actions.go` now imports two packages.
- `return MeleeAction{...}.Perform(...)` passes the inner action's result straight through.
- `MovementAction` now tests bounds and walkability in one `if` with `||`. `||` stops at the first true operand, so `TileAt` is still never called outside the map.

{{diff ai.go}}

{{diff colors.go}}

{{diff input.go}}

- `if err := ...; err != nil` declares `err` for the `if` only, the shape from step 2.
- `HandleEnemyTurns` in `engine.go` still calls `ent.AI.Perform(e, ent)` and ignores the result. Go lets you drop a return value. A monster that bumps into something simply loses its turn.

--- end

%%% In `input.go`, move the `if err := ...` block below `h.Engine.HandleEnemyTurns()`. The message still appears, but walking into a wall costs a turn again. Where the check sits decides the rule.
