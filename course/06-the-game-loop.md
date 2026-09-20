# Step 6 · The game loop
## Chapter: Moving the @

Every turn-based game is the same loop: **draw**, **wait for input**, **update**. Step 5 already draws once and waits; now the drawing moves *inside* the loop, and the `@` gets a position we can change. The greeting goes.

{{diff main.go}}

- `const ( ... )` declares named constants. The game is designed for an 80 by 50 cell console, the classic size; naming the numbers keeps them out of the code.
- `playerX, playerY := screenWidth/2, screenHeight/2` declares two variables at once. Integer division: the `@` starts at (40, 25).
- `screen.Clear()` at the top of every iteration wipes the buffer, so whatever we drew last time does not linger. Then draw, then `Show()`. Because the loop redraws on every event, a resize now repaints the screen automatically.
- The update part is still missing: any key quits. Next step.

!!! Run it: the `@` sits in the middle of the screen.
