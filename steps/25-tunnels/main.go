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

	gameMap := NewGameMap(mapWidth, mapHeight)
	room1 := NewRectangularRoom(20, 15, 10, 15)
	room2 := NewRectangularRoom(40, 28, 12, 8)
	room1.carve(gameMap)
	room2.carve(gameMap)
	x1, y1 := room1.Center()
	x2, y2 := room2.Center()
	tunnelBetween(gameMap, x1, y1, x2, y2)

	px, py := room1.Center()
	player := &Entity{X: px, Y: py, Char: '@', Color: tcell.ColorWhite}
	nx, ny := room2.Center()
	npc := &Entity{X: nx, Y: ny, Char: '@', Color: tcell.ColorYellow}

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
