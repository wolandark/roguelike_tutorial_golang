package main

import (
	"errors"
	"os"

	"github.com/gdamore/tcell/v2"
)

func NewGame() *Engine {
	engine := &Engine{MessageLog: &MessageLog{}}
	engine.Player = playerTemplate.Spawn(nil, 0, 0)
	engine.GameWorld = &GameWorld{
		MaxRooms:    maxRooms,
		RoomMinSize: roomMinSize,
		RoomMaxSize: roomMaxSize,
	}
	engine.GameWorld.GenerateFloor(engine)
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
		if errors.Is(err, os.ErrNotExist) {
			return &PopupMessage{Parent: m, Text: "No saved game to load."}
		}
		if err != nil {
			return &PopupMessage{Parent: m, Text: "Failed to load save: " + err.Error()}
		}
		return &MainGameEventHandler{Engine: engine}
	case 'n', 'N':
		return &MainGameEventHandler{Engine: NewGame()}
	}
	return m
}

type PopupMessage struct {
	Parent EventHandler
	Text   string
}

func (p *PopupMessage) OnRender(screen tcell.Screen) {
	p.Parent.OnRender(screen)
	dimScreen(screen)
	style := tcell.StyleDefault.Foreground(colorWhite).Background(colorBlack)
	drawText(screen, (screenWidth-len([]rune(p.Text)))/2, screenHeight/2, p.Text, style)
}

func (p *PopupMessage) HandleEvent(ev tcell.Event) EventHandler {
	if _, ok := ev.(*tcell.EventKey); ok {
		return p.Parent
	}
	return p
}

func dimScreen(screen tcell.Screen) {
	w, h := screen.Size()
	for y := 0; y < h; y++ {
		for x := 0; x < w; x++ {
			ch, _, style, _ := screen.GetContent(x, y)
			screen.SetContent(x, y, ch, nil, style.Dim(true))
		}
	}
}
