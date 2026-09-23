# Step 54 · Walking back

### The problem

`FindPath` knows, for every cell it reached, which cell it came from. Starting at the goal and following those links leads back to the start, one cell at a time. That gives the path in reverse order: goal first. Reversing the list gives the cells to walk, in order.

>>> 1. In `FindPath`, replace `return nil` with: return `nil` if the goal's cost is still -1 (unreachable); otherwise start at the goal and follow `prev` until the start, appending each cell to `path`.
>>> 2. Reverse `path` in place and return it.

!!! Nothing changes on screen yet; the next step makes monsters follow the path.

--- reveal

{{diff pathfinding.go}}

- `p != [2]int{x1, y1}` compares two arrays; Go compares arrays element by element.
- `path[i], path[j] = path[j], path[i]` swaps two elements in one statement. Moving `i` up from the front and `j` down from the back until they meet reverses the slice.
- The start cell itself is not in the path: the first element is the first step to take.

--- end

%%% Leave out the reversing loop. The path now starts at the player's cell, so in the next step a monster tries to jump straight to you and fails every turn, because that cell is not next to it.
