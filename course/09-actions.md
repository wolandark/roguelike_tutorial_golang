# Step 9 · Actions

### The problem

The loop now looks at *which key* was pressed and changes x and y right there. Think ahead a few chapters: a step left can come from the left arrow, from `h`, from numpad 4, from a mouse click on a tile, or from an orc's brain deciding to approach you. If every one of those has to know how to change x and y, the movement rules (walls, monsters in the way) get copied five times.

The fix is to separate two questions. Input answers only "what does the player want to do?", with a small value we call an **action**. The loop performs actions, without caring where they came from. Later, monsters produce the same action values. There is nothing Go-specific about this idea; what is Go-specific is how to say "one of several kinds of value": an interface and a type switch.

This is also the first time the program is more than one file. All `.go` files in the folder belong to `package main` and see each other's names; there is nothing to import.

>>> 1. Create `actions.go`. Declare an interface `type Action interface{}`.
>>> 2. In the same file, declare two structs: `type EscapeAction struct{}` and `type MovementAction struct { DX, DY int }`.
>>> 3. Create `input.go` with a function `func handleKey(ev *tcell.EventKey) Action`. Move the tagless switch there and make each case *return* a value instead of changing the position: `MovementAction{DX: 0, DY: -1}` for up and so on, `EscapeAction{}` for quit, and `return nil` at the end.
>>> 4. In `main.go`, replace the switch in the loop with a type switch: `switch action := handleKey(ev).(type)` with `case nil:` (do nothing), `case EscapeAction:` (return nil) and `case MovementAction:` (add `action.DX` and `action.DY` to the position).

!!! Behaves exactly like step 7. The structure is what changed.

--- reveal

{{file actions.go}}

- `type Action interface{}`: an **interface** lists methods a value must have; an empty one lists none, so *any* value is an `Action`. For now it is only a common return type; in step 18 it gets a real method.
- `type EscapeAction struct{}` is a struct with no fields: a value that carries no data, only meaning. `MovementAction` carries the step.

{{file input.go}}

- `MovementAction{DX: 0, DY: -1}` is a *struct literal*. `return nil` for keys we do not handle; `nil` is a valid value of an interface type.

{{diff main.go}}

- `switch action := handleKey(ev).(type)` is a **type switch**: it branches on the concrete type inside the interface value, and in each `case` the variable `action` *has* that type, so `action.DX` is available in the `MovementAction` branch. `case nil:` catches the ignored keys.

--- end

%%% Add a case to `handleKey` that returns `MovementAction{DX: 1, DY: 1}` for `ev.Rune() == 'n'`. The loop did not change and `n` moves you diagonally: input and consequences are now independent. (The chapter exercise adds all four diagonals.)
