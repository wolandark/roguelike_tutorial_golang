# Step 45 · Hit points
## Chapter: Doing (and taking) damage

### In this chapter

Combat, in small steps: hit points, death, dealing damage, showing your own health, drawing corpses under the living, a distance measure, a hostile brain, pathfinding in four parts, and finally a game-over mode. The Go theme is optional behaviour: not every entity fights, not every entity thinks, and Go has no class hierarchy to say "a fighting entity". Components, pointer fields that may be `nil`, are the answer, and they come with one trap around copying.

### The problem

Monsters and the player need hit points, an attack value and a defense value. A potion, which comes later, needs none of them. In a language with inheritance you would make a subclass for things that fight. Go has no inheritance, and its alternative fits better anyway: `Entity` gets an optional **component**, a pointer to a `Fighter` struct that is `nil` when the entity cannot fight. "Can it fight?" becomes `e.Fighter != nil`.

A pointer field brings one trap. `Spawn` copies the template struct, and copying a struct copies a pointer inside it, not what it points to. Without care, every orc would share the template's one `Fighter`.

>>> 1. Create `fighter.go` and declare a struct `Fighter` with fields `HP, MaxHP int`, `Defense int` and `Power int`.
>>> 2. In `entity.go`, add a field `Fighter *Fighter` to `Entity`.
>>> 3. In `Spawn`, when the template has a `Fighter`, copy the struct it points to and point the clone at the copy.
>>> 4. In `entity_factories.go`, give the templates a `Fighter`: the player 30 HP, defense 2, power 5; the orc 10 HP, defense 0, power 3; the troll 16 HP, defense 1, power 4.

!!! Nothing changes on screen; nothing uses the numbers yet.

--- reveal

{{file fighter.go}}

{{diff entity.go}}

- `f := *e.Fighter` reads the struct *behind* the pointer into a new variable, a copy. `&f` is a pointer to that copy, so the clone gets its own `Fighter`.

{{diff entity_factories.go}}

- `Fighter: &Fighter{...}` inside a struct literal is a pointer to a new `Fighter` value.
- The literals no longer fit on one line, so each template is written over several.

--- end

%%% Remove the lines that copy the `Fighter` in `Spawn`. Nothing looks different yet, but in step 47, when you hit one orc, every orc on the level loses the same hit points, and so does the template. Remember this one; it is the classic trap of copying structs that contain pointers.
