# Step 38 · Lightning
## Chapter: Scrolls and targeting

The first scroll needs no aiming: lightning strikes the **closest visible enemy** within range. That is a scan over all living fighters keeping the nearest one, a pattern worth knowing on its own. If nobody qualifies the scroll is not wasted.

{{diff gamemap.go}}

- `Actors` returns every living fighter, so consumables do not have to filter corpses and items themselves.

{{diff consumable.go}}

- Another type implementing `Consumable`. `GetAction` is the same as the potion's: an `ItemAction` right away.
- `Activate` starts `closest` at "farther than allowed", so the first candidate in range always wins, then keeps the closest. `if d := ...; d < closest` declares `d` for the `if` only. The consumer is skipped, and so is anything out of sight. 20 damage kills anything on the early floors.

{{diff entity_factories.go}}

{{diff procgen.go}}

- 70% potions, 30% lightning scrolls for now.

!!! Run it: pick up a yellow `~`, wait for an orc to come into view, use the scroll from `i`. With nobody near: "No enemy is close enough to strike." in grey, scroll kept.
