package main

import "github.com/gdamore/tcell/v2"

func main() {
	screen, _ := tcell.NewScreen()
	screen.Init()
	screen.SetContent(10, 5, '@', nil, tcell.StyleDefault)
	screen.Show()
	screen.PollEvent()
	screen.PollEvent()
	screen.Fini()
}
