package main

import (
	"fmt"
	"os"

	"github.com/gdamore/tcell/v2"
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

	screen.SetContent(10, 5, '@', nil, tcell.StyleDefault)
	msg := "Hello, roguelike! Press any key to quit."
	for i, r := range []rune(msg) {
		screen.SetContent(1+i, 1, r, nil, tcell.StyleDefault)
	}
	screen.Show()
	screen.PollEvent()
	screen.PollEvent()
	return nil
}

func main() {
	if err := run(); err != nil {
		fmt.Fprintln(os.Stderr, "error:", err)
		os.Exit(1)
	}
}
