# Step 5 · Waiting for a key

"Press any key to quit" is a lie in step 4: moving the mouse or resizing the window also ends the program, because `PollEvent` returns for *every* event. We loop until the event is a key.

{{diff main.go}}

- `for { ... }` with no condition is Go's infinite loop. It only ends through `return`.
- `ev := screen.PollEvent()` gives us an event of interface type `tcell.Event`. The concrete type depends on what happened: `*tcell.EventKey`, `*tcell.EventMouse`, `*tcell.EventResize`.
- `ev.(*tcell.EventKey)` is a **type assertion**: "is this event a key event?". With two results, `_, ok :=`, it never panics: `ok` is `true` when the assertion holds. We ignore the key itself (`_`) and return; any other event loops again.

!!! Run it: resizing the window or moving the mouse no longer quits. A key does.
