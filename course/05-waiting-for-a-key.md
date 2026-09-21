# Step 5 · Waiting for a key

### The problem

"Press any key" is a lie in step 4: move the mouse or resize the window and the program ends, because `PollEvent` returns for *every* event the terminal reports. The double `PollEvent` from step 1 only papered over the first of those, the resize tcell sends after `Init`. We only want keys, and the way to get that is to look at what arrived and keep waiting otherwise.

tcell hands us events as values of the interface type `tcell.Event`; the concrete type says what happened: `*tcell.EventKey`, `*tcell.EventMouse`, `*tcell.EventResize`. So the question is how to ask "is this a key?" and keep waiting otherwise.

>>> Replace the two `PollEvent` calls with a loop that polls until the event is a `*tcell.EventKey`, then returns. Look up the two-result form of a type assertion.

!!! Resizing the window or moving the mouse no longer quits. A key does.

--- reveal

{{diff main.go}}

- `for { ... }` with no condition is Go's infinite loop; it only ends through `return`.
- `ev.(*tcell.EventKey)` is a **type assertion**: "is this event a key event?". With two results, `_, ok :=`, it never panics: `ok` is `true` when the assertion holds. The key itself is ignored (`_`) for now.
- Where did `screen.Fini()` go? Nowhere: it is still the `defer` from step 3, and it runs when `run` **returns**. The `return nil` inside the loop is that return, so the moment a key arrives the loop ends, `run` ends, and the deferred `Fini()` restores the terminal before `main` continues. This is the case `defer` was chosen for: the function now has its exit in the middle of a loop, and there is still exactly one `Fini`, in the place it was declared.

--- end

%%% Use the one-result form, `key := ev.(*tcell.EventKey)`, and resize the window: the program panics with `interface conversion`. The two-result form exists so that a wrong guess is a `false`, not a crash.
