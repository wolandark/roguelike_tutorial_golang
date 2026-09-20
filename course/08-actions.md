# Step 8 · Actions

The loop currently looks at *which key* was pressed and changes x and y right there. Soon a step can come from an arrow key, a vi key, a numpad digit, a mouse click or a monster's brain, so we separate two questions: input answers "what does the player want to do?" with an **action** value, and the loop performs the action without caring where it came from. Two new files; this is also the first time the program is more than one file. All files in the folder belong to `package main` and see each other's names.

Create `actions.go`:

{{file actions.go}}

- `type Action interface{}` declares a type. An **interface** lists methods a value must have; an empty interface lists none, so *any* value is an `Action`. For now it is just a common return type.
- `type EscapeAction struct{}` is a struct with no fields: a value that carries no data, only meaning.
- `MovementAction` carries the step as two `int` fields. `DX, DY int` declares both with one type.

Create `input.go`:

{{file input.go}}

- `handleKey` takes a key event and returns an `Action`. `MovementAction{DX: 0, DY: -1}` is a *struct literal*: a new value with those fields. `return nil` for keys we do not handle; `nil` is a valid value for an interface type.

And `main.go` uses it:

{{diff main.go}}

- `switch action := handleKey(ev).(type)` is a **type switch**: it branches on the concrete type inside the interface value, and in each `case` the variable `action` *has* that type, so `action.DX` is available in the `MovementAction` branch. `case nil:` catches the ignored keys.

!!! Run it: behaves exactly like step 7. The structure is what changed, and everything from here on builds on it.
