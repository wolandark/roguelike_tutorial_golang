package main

import "github.com/gdamore/tcell/v2"

var (
	playerTemplate = Entity{
		Char: '@', Color: tcell.ColorWhite, Name: "Player", BlocksMovement: true, Alive: true,
		RenderOrder: RenderActor,
		Fighter:     &Fighter{HP: 30, MaxHP: 30, Defense: 2, Power: 5},
	}
	orc = Entity{
		Char: 'o', Color: tcell.NewRGBColor(63, 127, 63), Name: "Orc", BlocksMovement: true, Alive: true,
		RenderOrder: RenderActor,
		Fighter:     &Fighter{HP: 10, MaxHP: 10, Defense: 0, Power: 3},
	}
	troll = Entity{
		Char: 'T', Color: tcell.NewRGBColor(0, 127, 0), Name: "Troll", BlocksMovement: true, Alive: true,
		RenderOrder: RenderActor,
		Fighter:     &Fighter{HP: 16, MaxHP: 16, Defense: 1, Power: 4},
	}
)
