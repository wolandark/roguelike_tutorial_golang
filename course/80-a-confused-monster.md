# Step 80 · A confused monster

### The problem

A confusion scroll should make one monster stumble around at random for a few turns, then behave normally again. Because behaviour is an interface (step 51), confusing a monster means **swapping its AI** for another one. The new AI must remember the old one, count down its turns, and put the old one back when the count reaches zero.

Unlike `HostileEnemy`, this AI has state that changes every turn: the counter. So its method has a pointer receiver, and the monster holds a `*ConfusedEnemy`. With a value, each turn would count down a copy and the monster would stay confused forever.

>>> 1. In `ai.go`, declare a struct `ConfusedEnemy` with the fields `PreviousAI AI` and `TurnsRemaining int`, and a package-level variable `directions` listing the eight neighbour offsets as `[][2]int`.
>>> 2. Add a method `func (c *ConfusedEnemy) Perform(engine *Engine, entity *Entity) error`. When `TurnsRemaining` is 0 or less, log "The ... is no longer confused.", set `entity.AI = c.PreviousAI` and return `nil`.
>>> 3. Otherwise decrease `TurnsRemaining` and return `BumpAction{...}.Perform(engine, entity)` for a random entry of `directions`.

!!! Nothing changes yet: no monster gets this AI until the scroll in the next step. The program compiles.

--- reveal

{{diff ai.go}}

- `ai.go` now imports `fmt` and `math/rand/v2`, grouped in parentheses.
- A confused monster uses `BumpAction`, so it attacks whatever it stumbles into, other monsters included.
- Once `entity.AI` points back at `PreviousAI`, nothing refers to the `*ConfusedEnemy` any more, and the garbage collector frees it.

--- end

%%% In `entity_factories.go`, give the orc template `AI: &ConfusedEnemy{PreviousAI: HostileEnemy{}, TurnsRemaining: 10}` and take three steps. The log says "The Orc is no longer confused. (x6)" or similar: every orc spawned from the template shares one `*ConfusedEnemy`, and ten turns in total run out almost at once. `Spawn` copies the pointer, not the struct behind it, the trap from step 45. Put `HostileEnemy{}` back.
