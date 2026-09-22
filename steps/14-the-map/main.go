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
	entities := []*Entity{npc, player}

	gameMap := NewGameMap(mapWidth, mapHeight)

	for {
		screen.Clear()
		gameMap.Render(screen)
		for _, e := range entities {
			screen.SetContent(e.X, e.Y, e.Char, nil, tcell.StyleDefault.Foreground(e.Color))
		}
		screen.Show()

		ev, ok := screen.PollEvent().(*tcell.EventKey)
		if !ok {
			continue
		}
		switch action := handleKey(ev).(type) {
		case nil:
		case EscapeAction:
			return nil
		case MovementAction:
			player.Move(action.DX, action.DY)
		}
	}
}

func main() {
	if err := run(); err != nil {
		fmt.Fprintln(os.Stderr, "error:", err)
		os.Exit(1)
	}
}
