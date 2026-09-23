# Step 48 · Fireball

### The problem

An area effect: everything within a radius of the target takes damage, you included if you stand too close. The cursor handler needs one more thing so the player can see what they are about to hit: a hook to draw extra decoration, here a red frame around the cursor. That is an optional function field, `nil` when unused, and the reason step 40 kept `drawFrame` border-only.

>>> 1. In `input.go`, add a field `OnRenderExtra func(screen tcell.Screen)` to `SelectIndexHandler` and call it in `OnRender` when it is not nil.
>>> 2. Add a function `func NewAreaRangedAttackHandler(engine *Engine, radius int, callback func(x, y int) Action) *SelectIndexHandler` that sets `OnRenderExtra` to draw a red frame around the cursor.
>>> 3. In `consumable.go`, declare a struct `FireballDamageConsumable` with `Damage int` and `Radius int`. Its `GetAction` opens the area handler; its `Activate` damages every actor within `Radius` of the target, or returns `Impossible`.
>>> 4. In `entity_factories.go`, add a `fireballScroll` template; in `procgen.go` give it a 10% chance.

!!! A red `~`. Use it and a red square follows the cursor; confirm on a group of monsters. Stand inside the square yourself and learn why the message says "The Player is engulfed in a fiery explosion".

--- reveal

{{diff consumable.go}}

{{diff input.go}}

{{diff entity_factories.go}}

{{diff procgen.go}}

--- end

%%% Change the blast check to Euclidean distance (`dx*dx+dy*dy <= r*r`). The frame still shows a square but the blast is now a circle; the corners of the square are safe. Whatever distance you choose, the reticle must draw the same shape.
