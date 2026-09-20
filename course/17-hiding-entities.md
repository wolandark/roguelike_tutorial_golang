# Step 17 · Hiding entities

The map hides tiles, but the engine still draws every entity. Only visible ones should show, so monsters can lurk in the dark. To have something to test with, `main.go` puts a yellow `@` twelve cells to your right.

{{diff engine.go}}

{{diff main.go}}

!!! Run it: the yellow `@` is invisible until your field of view reaches it (walk towards it; it is often through a wall), then it appears and disappears as you move.
