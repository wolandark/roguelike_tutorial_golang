# Step 60 · The message log

### The problem

The engine throws old messages away after four. A player wants to scroll back through the whole fight later, so every message should be kept, and only the *display* should be limited to the last few. Keeping messages and drawing them is a job of its own, so it gets its own type: a `MessageLog` with a list of messages, a method to add one, and a method to draw the newest ones into a box, bottom-up, so the newest message is always on the bottom row.

Drawing takes the list of messages as a parameter rather than always drawing the whole log, because the history window later draws a *part* of it with the same code.

>>> 1. In `messagelog.go`, declare a struct `MessageLog` with a field `Messages []Message`, and a method `func (l *MessageLog) AddMessage(text string, color tcell.Color)` that appends a message.
>>> 2. Add a function `func renderMessages(screen tcell.Screen, x, y, width, height int, messages []Message)` that draws messages from the newest backwards, one per row, starting on the bottom row of the box and moving up until the box is full; and a method `func (l *MessageLog) Render(screen tcell.Screen, x, y, width, height int)` that calls it with all messages.
>>> 3. In `engine.go`, replace the `Messages` field with `MessageLog *MessageLog`, make `Log` call `e.MessageLog.AddMessage(msg, color)`, and in `Render` replace the message loop with `e.MessageLog.Render(screen, 21, 45, 40, 5)`.
>>> 4. In `main.go`, create the engine with `MessageLog: &MessageLog{}`.

!!! The same as before, except that the newest message is now always on the bottom row.

--- reveal

{{diff messagelog.go}}

- `yOffset` starts at the bottom row of the box (`height - 1`) and goes up by one per message. The loop stops when the box is full, so a log of a thousand messages draws only five.
- `width` is not used yet; the next steps need it.

{{diff engine.go}}

{{diff main.go}}

- `&MessageLog{}` creates an empty log. Without it `e.MessageLog` is `nil`, and the first `Log` call would crash.

--- end

%%% Leave out `MessageLog: &MessageLog{}` in `main.go`. The game crashes at start-up with `invalid memory address or nil pointer dereference`, in `AddMessage`, called from the welcome `Log`. A pointer field that nobody set is `nil`.
