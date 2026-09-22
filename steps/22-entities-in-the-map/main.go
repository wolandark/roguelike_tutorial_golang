package main

import (
	"fmt"
	"os"

	"github.com/gdamore/tcell/v2"
)

const (
	screenWidth  = 80
	screenHeight = 50
	mapWidth     = 80
	mapHeight    = 45

	roomMaxSize = 10
	roomMinSize = 6
	maxRooms    = 30
)

func run() error {
	screen, err := tcell.NewScreen()
	if err != nil {
		return err
	}
	if err := screen.Init(); err != nil {
		return err
	}
	defer screen.Fini()

	gameMap := NewGameMap(mapWidth, mapHeight)
	player := playerTemplate.Spawn(gameMap, 0, 0)
	GenerateDungeon(gameMap, maxRooms, roomMinSize, roomMaxSize, player)
	orc.Spawn(gameMap, player.X+3, player.Y)
	troll.Spawn(gameMap, player.X-3, player.Y)

	engine := &Engine{Player: player, GameMap: gameMap}
	engine.UpdateFOV()

	for {
		engine.Render(screen)
		if quit := engine.HandleEvent(screen.PollEvent()); quit {
			return nil
		}
	}
}

func main() {
	if err := run(); err != nil {
		fmt.Fprintln(os.Stderr, "error:", err)
		os.Exit(1)
	}
}
