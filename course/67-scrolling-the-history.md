# Step 67 · Scrolling the history

### The problem

The history window shows only as many messages as fit. To read older ones the player has to scroll. The viewer keeps a **cursor**, the index of the message shown on the bottom row, which starts at the newest message. Up and Down move it by one, PgUp and PgDn by ten, Home and End jump to the ends, and moving past either end wraps around to the other. Any other key still closes the window.

>>> 1. In `input.go`, add the fields `LogLength int` and `Cursor int` to `HistoryViewer`, and set both in `NewHistoryViewer`: the number of messages, and the index of the last one.
>>> 2. In `OnRender`, draw only the messages up to and including the cursor.
>>> 3. In `HandleEvent`, move the cursor with Up, Down, PgUp and PgDn, jump with Home and End, wrap around at both ends, and return `h.Parent` for any other key.

!!! Fight a bit, press `v`, scroll with the arrows and PgUp/PgDn; any other key closes the window.

--- reveal

{{diff input.go}}

- `Messages[:h.Cursor+1]` is a **slice expression**: every message up to and including the cursor. Drawing that slice bottom-up puts the cursor's message on the bottom row, which is what scrolling back looks like.
- `max(0, min(h.Cursor+step, h.LogLength-1))` keeps the cursor inside the log; the two `if`s before it wrap around at the ends.

--- end

%%% Remove the two `if`s that wrap around and keep only the `max(0, min(...))` line. Scrolling now stops at the oldest and newest message instead of jumping to the other end.
