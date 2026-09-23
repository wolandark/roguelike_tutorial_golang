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

	roomMaxSize        = 10
	roomMinSize        = 6
	maxRooms           = 30
	maxMonstersPerRoom = 2
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
	GenerateDungeon(gameMap, maxRooms, roomMinSize, roomMaxSize, maxMonstersPerRoom, player)

	engine := &Engine{Player: player, GameMap: gameMap}
	engine.UpdateFOV()
	engine.Log("Hello and welcome, adventurer, to yet another dungeon!")

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
