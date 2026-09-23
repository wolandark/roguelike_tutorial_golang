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

	player := &Entity{X: screenWidth / 2, Y: screenHeight / 2, Char: '@', Color: tcell.ColorWhite}
	npc := &Entity{X: screenWidth/2 - 5, Y: screenHeight / 2, Char: '@', Color: tcell.ColorYellow}
	gameMap := NewGameMap(mapWidth, mapHeight)
	for x := 30; x < 33; x++ {
		gameMap.SetTile(x, 22, wall)
	}

	engine := &Engine{
		Entities: []*Entity{npc, player},
		Player:   player,
		GameMap:  gameMap,
	}

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
