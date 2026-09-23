package main

func NewGame() *Engine {
	engine := &Engine{MessageLog: &MessageLog{}}
	engine.GameMap = NewGameMap(mapWidth, mapHeight)
	engine.Player = playerTemplate.Spawn(engine.GameMap, 0, 0)
	GenerateDungeon(engine.GameMap, maxRooms, roomMinSize, roomMaxSize, maxMonstersPerRoom, maxItemsPerRoom, engine.Player)
	engine.UpdateFOV()
	engine.Log("Hello and welcome, adventurer, to yet another dungeon!", colorWelcomeText)
	return engine
}
