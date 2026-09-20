# Step 41 · Fireball

An area effect: everything within a radius of the target takes damage, you included if you stand too close. The aiming handler gains a hook to draw the blast area so you can see what you are about to hit.

{{diff consumable.go}}

- `FireballDamageConsumable.Activate` loops over all actors and damages those within `Radius` (king's-move distance, so the blast is a square). It does not skip the consumer. If nobody was hit it is `Impossible` and the scroll is kept.

{{diff input.go}}

- `OnRenderExtra` is an optional function field: `nil` means nothing extra. `NewAreaRangedAttackHandler` builds a single-target handler and then sets a closure that draws a red frame of `2*radius+1` cells around the cursor. This is why `drawFrame` in step 33 only paints the border.

{{diff entity_factories.go}}

{{diff procgen.go}}

!!! Run it: a red `~`. Use it and a red square follows the cursor; confirm on a group of monsters. Stand inside the square yourself and learn why the message says "The Player is engulfed in a fiery explosion".
