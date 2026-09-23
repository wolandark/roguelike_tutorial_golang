# Step 58 · A text helper
## Chapter: The interface

### In this chapter

Four plain lines of text under the map is not an interface. This chapter builds the real one in small steps: a helper for writing text, coloured messages, a message log, stacked repeats, wrapped lines, a health bar, names under the mouse, and a message history window with a frame and scrolling. Nothing here is hard; it is where the game starts to look like one, and where the windows used by every menu later are built.

### The problem

Writing a string to the screen appears twice in `Engine.Render` already, each time as a loop over runes calling `SetContent`, and the interface will need it many more times. A small function that writes a string at a position, in a style, removes the repetition.

>>> 1. In `engine.go`, add a function `func drawText(screen tcell.Screen, x, y int, text string, style tcell.Style)` that writes each rune of `text` in its own cell, starting at `x, y`.
>>> 2. In `Engine.Render`, use `drawText` for the HP line and for each message instead of the two loops.

!!! Exactly the same screen as before.

--- reveal

{{diff engine.go}}

- `drawText` ranges over `[]rune(text)`, so a character that is several bytes long still takes one cell, the same reasoning as in step 4.
- It is a plain function, not a method: it needs only a screen, not the engine.

--- end

%%% Make `drawText` range over `text` instead of `[]rune(text)`, and log a message containing `é`. A gap appears after it, as in step 4: the index jumps by the character's byte length.
