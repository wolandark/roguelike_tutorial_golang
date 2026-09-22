# Step 8 · Actions

### The problem

The loop now looks at *which key* was pressed and changes x and y right there. Think ahead a few chapters: a step left can come from the left arrow, from `h`, from numpad 4, from a mouse click on a tile, or from an orc's brain deciding to approach you. If every one of those has to know how to change x and y, the movement rules (walls, monsters in the way) get copied five times.

The fix is to separate two questions. Input answers only "what does the player want to do?", with a small value we call an **action**. The loop performs actions, without caring where they came from. Later, monsters produce the same action values. There is nothing Go-specific about this idea; what is Go-specific is how to say "one of several kinds of value": an interface and a type switch.

This is also the first time the program is more than one file. All `.go` files in the folder belong to `package main` and see each other's names; there is nothing to import.

>>> Create `actions.go` with an `Action` type that any value can be (an empty interface), plus `EscapeAction` (no data) and `MovementAction` with `DX, DY int`. Create `input.go` with `handleKey(ev *tcell.EventKey) Action` that moves the tagless switch there and returns those values (arrows and vi keys to movements, Escape/Ctrl-C to escape, `nil` otherwise). In the loop, switch on the action's type instead of the key.

!!! Behaves exactly like step 7. The structure is what changed.

--- reveal

{{file actions.go}}

- `type Action interface{}`: an **interface** lists methods a value must have; an empty one lists none, so *any* value is an `Action`. For now it is only a common return type; in step 11 it gets a real method.
- `type EscapeAction struct{}` is a struct with no fields: a value that carries no data, only meaning. `MovementAction` carries the step.

{{file input.go}}

- `MovementAction{DX: 0, DY: -1}` is a *struct literal*. `return nil` for keys we do not handle; `nil` is a valid value of an interface type.

{{diff main.go}}

- `switch action := handleKey(ev).(type)` is a **type switch**: it branches on the concrete type inside the interface value, and in each `case` the variable `action` *has* that type, so `action.DX` is available in the `MovementAction` branch. `case nil:` catches the ignored keys.

--- end

%%% Add a case to `handleKey` that returns `MovementAction{DX: 1, DY: 1}` for `ev.Rune() == 'n'`. The loop did not change and `n` moves you diagonally: input and consequences are now independent. (The chapter exercise adds all four diagonals.)
