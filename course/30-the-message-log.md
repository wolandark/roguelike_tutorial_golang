# Step 30 · The message log
## Chapter: The interface

The four-line stand-in becomes a real log: messages have colours, repeats stack as `(x3)`, long lines wrap, and the log fills its box bottom-up. Two new files.

`colors.go` names every colour in one place:

{{file colors.go}}

`messagelog.go` is the log:

{{file messagelog.go}}

- A `Message` is text, colour and a repeat counter; `FullText` appends `(xN)` when the counter is above 1.
- `AddMessage` with `stack` set compares against the last message and bumps its counter instead of appending. `l.Messages[len(l.Messages)-1]` is the last element.
- `renderMessages` draws **bottom-up**: it walks the messages from newest to oldest, wraps each one, and writes lines from the bottom row upwards until `yOffset` goes negative. `Render` on the log is the same for the whole log; the history viewer in step 33 reuses `renderMessages` with a slice of the messages.
- `wrap` splits at spaces with `strings.Fields` and packs words into lines no longer than `width`. It is a small greedy word wrapper.
- `tcell.StyleDefault.Foreground(msg.Color)` draws each line in its colour.

The engine swaps the slice for the log, `Log` takes a colour, and `drawText` becomes a shared helper:

{{diff engine.go}}

The fighting messages get colours, and the welcome is the first log entry:

{{diff fighter.go}}

{{diff actions.go}}

{{diff main.go}}

!!! Run it: the blue welcome message bottom right. Fight an orc and watch repeated hits stack with `(x2)`. Long messages wrap inside the 40-column box.
