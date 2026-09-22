# Step 6 · The game loop
## Chapter: Moving the @

### In this chapter

Every turn-based game is one loop: **draw** the world, **wait** for input, **update** the world, repeat. This chapter builds it over four steps. Step 6 (this one) turns the program into that loop and gives the `@` a position, but nothing moves yet. Step 7 reads the arrow keys and makes it move. Step 8 adds the vi keys, and runs into the difference between a value and a condition. Step 9 adds a small piece of structure, *actions*, that separates "which key" from "what happens"; that looks like overkill for a few keys, and it is what lets monsters, mice and menus produce the same movements later.

### The problem

Step 5 draws once and then waits. To move something, the drawing has to happen again after every input, and the `@` needs a position that can change. Where should the position live? Two integers in `run` are enough for now; they become a struct in the next chapter when there is more than one thing on screen.

>>> Give the `@` a position, start it in the middle of an 80 by 50 screen (name those two numbers as constants), and restructure `run` so that every loop iteration clears the buffer, draws the `@` at its position and shows it, then waits for an event. Drop the greeting.

!!! The `@` sits in the middle of the screen. Any key still quits; that is the next step.

--- reveal

{{diff main.go}}

- `const ( ... )` declares named constants. The game is designed for an 80 by 50 console, the classic size.
- `playerX, playerY := screenWidth/2, screenHeight/2` declares two variables at once; integer division puts the `@` at (40, 25).
- `screen.Clear()` at the top of every iteration wipes the buffer so nothing lingers, then draw, then `Show()`. Because the loop redraws on every event, a resize now repaints the screen by itself.

--- end

%%% Remove `screen.Clear()`, and in the next step, move around: the `@` leaves a trail, because the buffer still holds the old ones. Clearing per frame is cheap; tcell only sends the cells that actually changed.
