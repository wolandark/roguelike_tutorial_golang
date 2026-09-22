# Step 31 · The message log
## Chapter: The interface

### In this chapter

Four lines of plain text under the map is not an interface. This chapter adds the real message log with colours and stacking, a health bar, mouse-over names, and a scrollable message history in a window. Nothing here is hard; it is where the game starts to feel like one, and where the windows and frames used by every menu later are built.

### The problem

The four-line stand-in loses messages, cannot colour them, and shows "Orc attacks Player for 1 hit points." three times in a row. A log should keep everything, in colour, stack repeats as `(x3)`, wrap long lines, and draw its last lines into a box bottom-up so the newest is always at the bottom. That is a small type with a render function; the history window in step 34 will reuse the same render function on a slice of the messages, which is why rendering takes the message list as a parameter.

>>> Create `colors.go` with named colours for attacks, deaths, the welcome text and the bar. Create `messagelog.go` with `Message` (text, colour, count), `MessageLog.AddMessage(text, color, stack)` that bumps the count on a repeat, `wrap(text, width)` that breaks at spaces, and a `renderMessages` that fills a box bottom-up. Replace the engine's string slice with a `*MessageLog`, make `Log` take a colour, move `drawText` into `engine.go` as a shared helper, and colour the combat and death messages.

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
