# Step 52 · A priority queue

### The problem

A monster walking straight at you stops at the first wall. It needs a route around obstacles. The classic tool is **Dijkstra's algorithm**, which explores outwards from the monster and always continues from the cheapest cell reached so far. "Always take the cheapest one next" is a **priority queue**, and Go's standard library has one in `container/heap`. The package does not provide a queue type; it works on any slice type of yours that has five methods.

This step builds only the queue's type. The next steps use it.

>>> 1. Create `pathfinding.go` with a struct `type pathNode struct{ x, y, cost int }` and a slice type `type pathQueue []pathNode`.
>>> 2. Give `pathQueue` the three methods of `sort.Interface`, with value receivers: `Len() int`, `Less(i, j int) bool` comparing `cost`, and `Swap(i, j int)`.
>>> 3. Give it `Push(x any)` and `Pop() any`, with **pointer** receivers, that append a `pathNode` and remove the last element.

!!! Nothing changes on screen.

--- reveal

{{file pathfinding.go}}

- `type pathQueue []pathNode` defines a new named slice type, so it can have methods.
- `Less` decides the order: the node with the smaller `cost` comes first. `container/heap` uses it to keep the cheapest node at the front.
- `Push` and `Pop` need pointer receivers because they change the slice's length, and a value receiver would change a copy of the slice header.
- `any` is Go's name for the empty interface: a value of any type. `x.(pathNode)` turns it back into a `pathNode`.

--- end

%%% Give `Push` a value receiver, `(q pathQueue)`, and assign `q = append(q, x.(pathNode))`. It compiles, but once the next steps use it, nothing is ever added to the queue: `append` builds a new slice header and the assignment only changes the method's copy.
