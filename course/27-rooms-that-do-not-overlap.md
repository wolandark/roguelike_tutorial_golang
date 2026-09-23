# Step 27 · Rooms that do not overlap

### The problem

Overlapping rooms melt into shapeless blobs. The fix is to throw away any new room that overlaps one we already kept, and simply try again with the next random room. That needs a test for two rectangles overlapping.

It is easiest to think about when two rooms do **not** overlap: one is entirely to the left of the other, or entirely to the right, above, or below. If none of those is true, they overlap. Written positively, that is four comparisons joined with `&&`.

>>> 1. In `procgen.go`, add a method `func (r RectangularRoom) Intersects(other RectangularRoom) bool` that returns true when the two rooms overlap: `r` starts before `other` ends and ends after `other` starts, both horizontally and vertically.
>>> 2. In `GenerateDungeon`, after creating `newRoom`, loop over the rooms kept so far; if `newRoom` intersects any of them, skip it with `continue`.

!!! Separate, clean rectangles on every run, fewer than 30 of them because many attempts are thrown away. They are still not connected.

--- reveal

{{diff procgen.go}}

- `Intersects` compares edges: `r.X1 <= other.X2 && r.X2 >= other.X1` means the two overlap horizontally, and the same for `Y` means they overlap vertically. Both must hold.
- The `overlaps` flag is needed because `continue` inside the inner `for` would only skip to the next *other* room, not to the next attempt. `break` leaves the inner loop, and the `continue` after it skips the attempt.

--- end

%%% Move the `continue` into the inner loop in place of `break`, and delete the flag. Overlapping rooms come back: the `continue` now only moves on to the next room in the comparison, and the new room is carved anyway.
