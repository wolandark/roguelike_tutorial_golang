# Step 41 · A message line

### The problem

Attacking should say what happened, and so will much else later. We cannot `fmt.Println` while tcell owns the screen: it would print over the game. The game needs its own place for messages. A proper message log comes later; for now the engine keeps the last four messages and draws them under the map, where the five spare rows are.

>>> 1. In `engine.go`, add a field `Messages []string` to `Engine`, and a method `func (e *Engine) Log(msg string)` that appends the message and drops the oldest one when there are more than four.
>>> 2. In `Engine.Render`, after the map, draw each message on its own row starting at row `e.GameMap.Height`.
>>> 3. In `main.go`, after `engine.UpdateFOV()`, log a welcome: `engine.Log("Hello and welcome, adventurer, to yet another dungeon!")`.

!!! The welcome message under the map.

--- reveal

{{diff engine.go}}

- `e.Messages[1:]` is a **slice expression**: everything from index 1 on. Assigning it back drops the first, oldest message.
- The map is 45 rows high and the screen 50, so rows 45 to 49 are free for messages.

{{diff main.go}}

--- end

%%% Remove the `if` that drops old messages and log something on every key press (for example in `HandleEvent`). The list grows past row 49 and the extra messages are simply not drawn: tcell ignores cells outside the screen. The trimming is what keeps the newest messages visible.
