package main

import "github.com/gdamore/tcell/v2"

func NewGame() *Engine {
	engine := &Engine{MessageLog: &MessageLog{}}
	engine.GameMap = NewGameMap(mapWidth, mapHeight)
	engine.Player = playerTemplate.Spawn(engine.GameMap, 0, 0)
	GenerateDungeon(engine.GameMap, maxRooms, roomMinSize, roomMaxSize, maxMonstersPerRoom, maxItemsPerRoom, engine.Player)
	engine.UpdateFOV()
	engine.Log("Hello and welcome, adventurer, to yet another dungeon!", colorWelcomeText)
	return engine
}

type MainMenu struct{}

var menuTitle = []string{
	"▄▄▄▄▄ ▄▄▄▄▄ ▄   ▄ ▄▄▄▄  ▄▄▄▄▄",
	"  █   █   █ ██ ██ █   █ █    ",
	"  █   █   █ █ █ █ █▄▄▄  ▀▀▀▀▄",
	"  █   ▀▄▄▄▀ █   █ █   █ ▄▄▄▄▀",
	"                              ",
	"   of the Ancient Kings       ",
}

func (MainMenu) OnRender(screen tcell.Screen) {
	title := tcell.StyleDefault.Foreground(colorMenuTitle)
	for i, line := range menuTitle {
		drawText(screen, (screenWidth-len([]rune(line)))/2, screenHeight/2-10+i, line, title)
	}
	drawText(screen, (screenWidth-20)/2, screenHeight-2, "By (Your name here)", title)

	options := []string{"[N] Play a new game", "[C] Continue last game", "[Q] Quit"}
	menu := tcell.StyleDefault.Foreground(colorMenuText).Background(colorBlack)
	for i, opt := range options {
		drawText(screen, screenWidth/2-12, screenHeight/2-2+i, "  "+opt+"                ", menu)
	}
}

func (m MainMenu) HandleEvent(ev tcell.Event) EventHandler {
	key, ok := ev.(*tcell.EventKey)
	if !ok {
		return m
	}
	if key.Key() == tcell.KeyEscape || key.Key() == tcell.KeyCtrlC {
		return nil
	}
	if key.Key() != tcell.KeyRune {
		return m
	}
	switch key.Rune() {
	case 'q', 'Q':
		return nil
	case 'c', 'C':
		engine, err := LoadGame(saveFile)
		if err != nil {
			return m
		}
		return &MainGameEventHandler{Engine: engine}
	case 'n', 'N':
		return &MainGameEventHandler{Engine: NewGame()}
	}
	return m
}
