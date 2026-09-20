package main

import "github.com/gdamore/tcell/v2"

var (
	playerTemplate = Entity{Char: '@', Color: tcell.ColorWhite, Name: "Player", BlocksMovement: true}
	orc            = Entity{Char: 'o', Color: tcell.NewRGBColor(63, 127, 63), Name: "Orc", BlocksMovement: true}
	troll          = Entity{Char: 'T', Color: tcell.NewRGBColor(0, 127, 0), Name: "Troll", BlocksMovement: true}
)
