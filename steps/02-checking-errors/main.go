package main

import (
	"fmt"

	"github.com/gdamore/tcell/v2"
)

func main() {
	screen, err := tcell.NewScreen()
	if err != nil {
		fmt.Println("error:", err)
		return
	}
	if err := screen.Init(); err != nil {
		fmt.Println("error:", err)
		return
	}
	screen.SetContent(10, 5, '@', nil, tcell.StyleDefault)
	screen.Show()
	screen.PollEvent()
	screen.Fini()
}
