# Step 69 · Impossible actions

### The problem

Every error from `Perform` is shown in grey and costs no turn. Today that is right, because every error means "you can't do that". But `error` is a general type: a later action could fail for a real reason, a bug or a failed file write, and that should not look like a harmless grey note.

So the "not possible, no turn" case needs its own type. Any type with an `Error() string` method satisfies the `error` interface, so a small struct can be returned wherever an `error` is expected. The handler then asks "is this error an `Impossible`?" with `errors.As`, and treats everything else as a real error.

>>> 1. Create `exceptions.go` with a struct `type Impossible struct { Msg string }` and a method `func (i Impossible) Error() string` that returns `i.Msg`.
>>> 2. In `actions.go`, replace each `errors.New("...")` with `Impossible{"..."}` and remove the `errors` import.
>>> 3. In `colors.go`, add `colorError` (red).
>>> 4. In `input.go`, when `Perform` returns an error, declare `var imp Impossible`. If `errors.As(err, &imp)` is true, log `imp.Msg` in `colorImpossible`. Otherwise log `err.Error()` in `colorError`. Return `h` in both cases.

!!! The same as step 68: a grey "That way is blocked." and no monster turn. The difference only shows for errors that are not `Impossible`.

--- reveal

{{file exceptions.go}}

- `Impossible{"That way is blocked."}` is a struct literal with the fields in order, here the only one, `Msg`.

{{diff actions.go}}

{{diff colors.go}}

{{diff input.go}}

- `errors.As(err, &imp)` checks whether `err` is an `Impossible`. If it is, it copies it into `imp` and returns `true`. It takes a pointer so that it can write into `imp`.

--- end

%%% In `MovementAction.Perform`, change the first `Impossible{"That way is blocked."}`, the one for walls, back to `errors.New("That way is blocked.")` and add the `errors` import again. Walk into a wall: the same text, now in red. The handler decides by the type of the error, not by its text.
