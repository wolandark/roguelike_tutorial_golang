# Step 44 · The message log
## Chapter: The interface

### In this chapter

Four lines of plain text under the map is not an interface. This chapter adds the real message log with colours and stacking, a health bar, mouse-over names, and a scrollable message history in a window. Nothing here is hard; it is where the game starts to feel like one, and where the windows and frames used by every menu later are built.

### The problem

The four-line stand-in loses messages, cannot colour them, and shows "Orc attacks Player for 1 hit points." three times in a row. A log should keep everything, in colour, stack repeats as `(x3)`, wrap long lines, and draw its last lines into a box bottom-up so the newest is always at the bottom. That is a small type with a render function; the history window in step 47 will reuse the same render function on a slice of the messages, which is why rendering takes the message list as a parameter.

>>> 1. Create `colors.go` with a `var ( )` block of named `tcell.Color` values: attack colours, death colours, the welcome text and the bar colours.
>>> 2. Create `messagelog.go` with a struct `Message` (`Text string`, `Color tcell.Color`, `Count int`) and a method `func (m Message) FullText() string` that appends `(xN)`.
>>> 3. In the same file, declare `type MessageLog struct { Messages []Message }` with a method `AddMessage(text string, color tcell.Color, stack bool)` that bumps the count of a repeated last message.
>>> 4. Add a function `func wrap(text string, width int) []string`, a function `renderMessages(screen, x, y, width, height int, messages []Message)` that draws bottom-up, and a method `Render` on `MessageLog` that calls it.
>>> 5. In `engine.go`, replace `Messages []string` with `MessageLog *MessageLog`, change `Log` to `func (e *Engine) Log(msg string, color tcell.Color)`, draw the log in a 40 by 5 box at column 21, row 45, and add a helper `func drawText(screen tcell.Screen, x, y int, text string, style tcell.Style)`.
>>> 6. In `fighter.go` and `actions.go`, pass the matching colour to every `Log` call.
>>> 7. In `main.go`, create the engine with `MessageLog: &MessageLog{}` and log a welcome message in `colorWelcomeText`.

!!! The blue welcome message bottom right. Fight an orc and watch repeated hits stack with `(x2)`. Long messages wrap inside the 40-column box.

--- reveal

{{file colors.go}}

{{file messagelog.go}}

- `l.Messages[len(l.Messages)-1]` is the last element; `AddMessage` bumps its counter instead of appending when `stack` is set and the text repeats.
- `renderMessages` walks messages newest to oldest, wraps each, and writes lines from the bottom row upwards until `yOffset` goes negative.
- `wrap` splits at spaces with `strings.Fields` and packs words greedily.

{{diff engine.go}}

{{diff fighter.go}}

{{diff actions.go}}

{{diff main.go}}

--- end

%%% Pass `false` as the `stack` argument in `Engine.Log`. Every hit gets its own line again and the welcome message scrolls away after five hits. Stacking is a small thing that makes logs readable.
