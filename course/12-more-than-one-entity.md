# Step 12 · More than one entity

### The problem

A game has many entities, and the draw loop should not know how many. That means a **slice** of entities, drawn in a loop. The question is what the slice holds: entities, or pointers to entities? The player is referenced from two places, the `player` variable and the slice, and both must see the same position when it moves. Two copies of a struct are two different entities; two pointers to one struct are one entity seen twice.

>>> Add a second entity, a yellow `@` five cells to the left of the player, keep both in a slice `[]*Entity`, and draw the slice with a `for range` loop each frame.

!!! Your white `@` and a yellow `@` five cells to its left. You can walk through the yellow one; nothing checks for other entities yet.

--- reveal

{{diff main.go}}

- `entities := []*Entity{npc, player}` is a slice of pointers. The slice and the `player` variable point at the **same** entity: `player.Move` through one, and the loop draws the new position through the other.
- `for _, e := range entities` gives each element in turn; `_` discards the index we do not need.

--- end

%%% Make it a `[]Entity` of values and store `*player` in it (the struct behind the pointer). Moving the player no longer moves what is drawn: the slice holds a copy taken at that moment. Pointers are how the world will always keep its entities.
