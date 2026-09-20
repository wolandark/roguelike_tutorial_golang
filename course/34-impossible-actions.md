# Step 34 · Impossible actions
## Chapter: Items and inventory

Walking into a wall should tell you so, and it should **not** cost a turn: monsters must not get a free hit because you bumped a wall. Go's tool for "this could not be done" is an error value, so actions now return `error`, and one special error type carries the message for the player.

Create `exceptions.go`:

{{file exceptions.go}}

- Any type with an `Error() string` method satisfies Go's built-in `error` interface. `Impossible` is such a type: a struct holding the message to show.

Every action returns an error now:

{{diff actions.go}}

- `EscapeAction` and `WaitAction` return `nil` (no problem). Melee with no target, and movement into a wall or another entity, return an `Impossible` with a message. `return MeleeAction{...}.Perform(...)` passes the inner action's result straight through.

The AI interface follows suit:

{{diff ai.go}}

The main handler's tail becomes a shared function:

{{diff input.go}}

- `runAction` is the one place where the rule lives. It performs the action; if the error **is an `Impossible`**, the message is logged in grey and the same handler is returned: no enemy turn. `errors.As(err, &imp)` checks the error's type and copies it into `imp` if it matches; it also sees through errors that were wrapped. Any other error is logged in red. On success the enemies act, the field of view updates, and the next handler is game-over or a fresh main-game handler.
- `import "errors"` is standard library, so it goes in the first group.

Two colours for the new messages:

{{diff colors.go}}

- The engine ignores the errors monsters' actions return: a monster that walks into a wall simply loses its turn.

!!! Run it: walk into a wall: a grey "That way is blocked." and the monsters do **not** move. Compare with a real step, after which they do.
