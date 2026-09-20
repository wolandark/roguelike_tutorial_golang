# Step 18 · Templates and spawning
## Chapter: Placing enemies

A dungeon needs many orcs. Rather than writing the orc's glyph, colour and name every time, we keep one **template** per kind of creature and copy it when we place one. Go copies a struct on assignment, and a method with a *value receiver* receives a copy for free; that is all `Spawn` needs.

{{diff entity.go}}

- Two new fields: a `Name` for messages, and `BlocksMovement`, because soon you will not be able to walk through an orc.
- `func (e Entity) Spawn(x, y int) *Entity` has a **value receiver** (`e Entity`, not `*Entity`), so `e` is already a copy of the template it was called on. `clone := e` copies once more into a variable whose address we can return. The template is never touched; every spawned orc is its own struct. This is what `copy.deepcopy` does in the Python tutorial.

Create `entity_factories.go` with the templates:

{{file entity_factories.go}}

- Three package-level values. `playerTemplate` rather than `player`, because `player` is the name of the *spawned* one in `main`.

And spawn from them in `main.go`. The NPC is replaced by an orc and a troll next to you, just to look at:

{{diff main.go}}

!!! Run it: a green `o` to your right, a `T` to your left. You can still walk through them; that changes next step.
