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

	playerX, playerY := screenWidth/2, screenHeight/2

	for {
		screen.Clear()
		screen.SetContent(playerX, playerY, '@', nil, tcell.StyleDefault)
		screen.Show()

		ev := screen.PollEvent()
		if _, ok := ev.(*tcell.EventKey); ok {
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
