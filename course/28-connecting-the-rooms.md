# Step 28 · Connecting the rooms

### The problem

The rooms are separate islands. Every room needs a way to every other room, and there is a cheap way to guarantee that without any graph algorithm: connect each new room to the room kept just before it. Room 2 connects to room 1, room 3 to room 2, and so on, so every room is reachable from the first, where the player starts.

>>> 1. In `GenerateDungeon`, after carving a room that is not the first, get the centre of the previous room (`rooms[len(rooms)-1]`) and of the new room, and call `tunnelBetween` between them.

!!! A different dungeon every time. You start in the middle of a room and every room is reachable.

--- reveal

{{diff procgen.go}}

- `rooms[len(rooms)-1]` is the last element of the slice: the room kept before this one. It is read *before* the new room is appended.
- The `if` for the first room grows an `else`: the first room gets the player, every other room gets a tunnel.

--- end

%%% Set `maxRooms` to 300. The level fills up with rooms until nothing more fits; the loop still ends, because it counts attempts, not kept rooms.
