# Step 71 · An inventory

### The problem

The player should be able to carry things. That is a new ability, and abilities in this game are components: a `Fighter` makes an entity fight, an `AI` makes it act. An `Inventory` makes it carry. It holds a capacity and a slice of entities. Only the player gets one, so orcs and trolls keep a `nil` inventory.

`Spawn` has the same trap as with the `Fighter` in step 45. The template's `Inventory` is a pointer, and a copy of the entity copies the pointer, not the inventory. Worse, the inventory holds a slice, and copying a slice copies only its header, so the backing array would still be shared. `Spawn` must copy both.

>>> 1. Create `inventory.go` with a struct `Inventory` with the fields `Capacity int` and `Items []*Entity`.
>>> 2. In the same file, add a method `func (inv *Inventory) Remove(item *Entity)` that deletes `item` from `inv.Items`, keeping the order of the others.
>>> 3. In `entity.go`, add a field `Inventory *Inventory` to `Entity`. In `Spawn`, when the template has an inventory, copy the struct and give the copy its own copy of `Items`.
>>> 4. In `entity_factories.go`, give the player `Inventory: &Inventory{Capacity: 26}`.

!!! Nothing changes on screen. The program compiles and runs as before.

--- reveal

{{file inventory.go}}

- `append(inv.Items[:i], inv.Items[i+1:]...)` deletes element `i`: it appends everything after `i` to everything before it. The `...` passes the elements of a slice as separate arguments.
- `Capacity` is 26 because the menu will label items with the letters `a` to `z`.

{{diff entity.go}}

- `append([]*Entity(nil), e.Inventory.Items...)` appends all items to an empty slice, which allocates a new backing array. That is the idiom for copying a slice.

{{diff entity_factories.go}}

--- end

%%% In `Spawn`, remove the `*` from `inv := *e.Inventory`. The build fails with `cannot use &inv (value of type **Inventory) as *Inventory value in assignment`. Without the `*`, `inv` is a copy of the pointer, not of the inventory, and `&inv` is a pointer to that pointer. Put the `*` back.
