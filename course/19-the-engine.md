# Step 19 · The engine

### The problem

`run` now holds everything: the entities, the player, the map, the drawing and the input rules. It will keep growing with every feature. The world belongs in one place that the rest of the code can be handed around with, and the natural first thing to move there is drawing: an **Engine** struct holding the entities, the player and the map, with a `Render` method.

>>> 1. Create `engine.go` and declare a struct `Engine` with fields `Entities []*Entity`, `Player *Entity` and `GameMap *GameMap`.
>>> 2. Add a method `func (e *Engine) Render(screen tcell.Screen)` that clears the screen, renders the map, draws every entity and shows the result.
>>> 3. In `main.go`, delete the `entities` slice, build `engine := &Engine{Entities: []*Entity{npc, player}, Player: player, GameMap: gameMap}`, and replace the drawing code at the top of the loop with `engine.Render(screen)`.

!!! Plays exactly like step 18. The drawing now lives in the engine.

--- reveal

{{file engine.go}}

- `Engine` holds pointers, so `engine.Player` and the `player` variable in `run` are the same entity; moving one moves the other.

{{diff main.go}}

- The struct literal is written one field per line. Go requires a trailing comma after the last field when the closing brace is on its own line.

--- end

%%% Delete the comma after `GameMap:  gameMap` in the literal. The compiler reports `syntax error: unexpected newline in composite literal; possibly missing comma or }`. When the closing brace is on its own line, every field line ends with a comma, the last one included.
