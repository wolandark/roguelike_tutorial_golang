package main

import (
	"fmt"
	"os"

	"github.com/gdamore/tcell/v2"
)

const (
	screenWidth  = 80
	screenHeight = 50
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

	for {
		screen.Clear()
		screen.SetContent(player.X, player.Y, player.Char, nil, tcell.StyleDefault.Foreground(player.Color))
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
