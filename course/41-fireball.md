# Step 41 · Fireball

### The problem

An area effect: everything within a radius of the target takes damage, you included if you stand too close. The cursor handler needs one more thing so the player can see what they are about to hit: a hook to draw extra decoration, here a red frame around the cursor. That is an optional function field, `nil` when unused, and the reason step 33 kept `drawFrame` border-only.

>>> Add `OnRenderExtra func(screen)` to `SelectIndexHandler`, called after the game is drawn if set. Add `NewAreaRangedAttackHandler(engine, radius, callback)` that builds a single-target handler and sets a closure drawing a red frame of `2*radius+1` cells around the cursor. Add `FireballDamageConsumable{Damage, Radius}`: visible cell required, damage every actor within `Radius` (king's move), `Impossible` if nobody was hit. Template, 10% spawn chance.

!!! A red `~`. Use it and a red square follows the cursor; confirm on a group of monsters. Stand inside the square yourself and learn why the message says "The Player is engulfed in a fiery explosion".

--- reveal

{{diff consumable.go}}

{{diff input.go}}

{{diff entity_factories.go}}

{{diff procgen.go}}

--- end

%%% Change the blast check to Euclidean distance (`dx*dx+dy*dy <= r*r`). The frame still shows a square but the blast is now a circle; the corners of the square are safe. Whatever distance you choose, the reticle must draw the same shape.
