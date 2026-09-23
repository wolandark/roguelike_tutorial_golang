package main

import "github.com/gdamore/tcell/v2"

var (
	colorWhite = tcell.NewRGBColor(0xFF, 0xFF, 0xFF)
	colorBlack = tcell.NewRGBColor(0, 0, 0)
	colorRed   = tcell.NewRGBColor(0xFF, 0, 0)

	colorPlayerAtk = tcell.NewRGBColor(0xE0, 0xE0, 0xE0)
	colorEnemyAtk  = tcell.NewRGBColor(0xFF, 0xC0, 0xC0)

	colorPlayerDie = tcell.NewRGBColor(0xFF, 0x30, 0x30)
	colorEnemyDie  = tcell.NewRGBColor(0xFF, 0xA0, 0x30)

	colorWelcomeText = tcell.NewRGBColor(0x20, 0xA0, 0xFF)
)
