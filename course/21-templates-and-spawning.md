# Step 21 · Templates and spawning
## Chapter: Placing enemies

### In this chapter

A dungeon needs inhabitants. This chapter fills rooms with orcs and trolls, makes them solid, lets you attack by walking into them, adds the rest of the traditional keys, and gives monsters a turn of their own (a placeholder turn; teeth come in chapter 6). The Go theme is copying: when is a struct a copy, when is it shared, and how to make many orcs from one description.

### The problem

Writing `&Entity{Char: 'o', Color: ..., Name: "Orc"}` every time an orc appears is repetitive and error-prone. We want one description per kind of creature, a **template**, and a way to stamp copies out of it. In Python the tutorial uses `copy.deepcopy`. In Go a struct assignment already copies, and a method with a *value receiver* receives a copy for free. The catch, which bites in step 27, is that copying a struct copies pointers inside it, not what they point to; for now `Entity` has no pointers, so a plain copy is a full copy.

>>> Add `Name string` and `BlocksMovement bool` to `Entity`, and a method `Spawn(x, y int) *Entity` with a **value** receiver that returns a pointer to a copy placed at (x, y). Create `entity_factories.go` with three templates: `playerTemplate`, `orc` (`o`, green), `troll` (`T`, darker green). Spawn the player and two monsters in `main.go`.

!!! A green `o` to your right, a `T` to your left. You can still walk through them.

--- reveal

{{diff entity.go}}

- `func (e Entity) Spawn(...)` has a **value receiver**, so `e` is already a copy of the template. `clone := e` copies once more into a variable whose address we can return. The template is never touched.

{{file entity_factories.go}}

- `playerTemplate` rather than `player`, because `player` names the *spawned* one in `main`.

{{diff main.go}}

--- end

%%% Change `Spawn` to a pointer receiver and return `e` itself instead of a clone. Both monsters are now the same struct as the template: move one (later) and the template moves. The value receiver is not a style choice here; it is the copy.
