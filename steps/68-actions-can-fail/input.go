package main

import "github.com/gdamore/tcell/v2"

type EventHandler interface {
	HandleEvent(ev tcell.Event) EventHandler
	OnRender(screen tcell.Screen)
}

var moveKeys = map[tcell.Key][2]int{
	tcell.KeyUp:    {0, -1},
	tcell.KeyDown:  {0, 1},
	tcell.KeyLeft:  {-1, 0},
	tcell.KeyRight: {1, 0},
	tcell.KeyHome:  {-1, -1},
	tcell.KeyEnd:   {-1, 1},
	tcell.KeyPgUp:  {1, -1},
	tcell.KeyPgDn:  {1, 1},
}

var moveRunes = map[rune][2]int{
	'h': {-1, 0}, 'j': {0, 1}, 'k': {0, -1}, 'l': {1, 0},
	'y': {-1, -1}, 'u': {1, -1}, 'b': {-1, 1}, 'n': {1, 1},
	'1': {-1, 1}, '2': {0, 1}, '3': {1, 1}, '4': {-1, 0},
	'6': {1, 0}, '7': {-1, -1}, '8': {0, -1}, '9': {1, -1},
}

type MainGameEventHandler struct {
	Engine *Engine
}

func (h *MainGameEventHandler) OnRender(screen tcell.Screen) { h.Engine.Render(screen) }

func (h *MainGameEventHandler) HandleEvent(ev tcell.Event) EventHandler {
	key, ok := ev.(*tcell.EventKey)
	if !ok {
		trackMouse(h.Engine, ev)
		return h
	}
	if key.Key() == tcell.KeyRune && key.Rune() == 'v' {
		return NewHistoryViewer(h.Engine, h)
	}
	action := handleKey(key)
	if action == nil {
		return h
	}
	if _, ok := action.(EscapeAction); ok {
		return nil
	}

	if err := action.Perform(h.Engine, h.Engine.Player); err != nil {
		h.Engine.Log(err.Error(), colorImpossible)
		return h
	}
	h.Engine.HandleEnemyTurns()
	h.Engine.UpdateFOV()

	if !h.Engine.Player.Alive {
		return &GameOverEventHandler{Engine: h.Engine}
	}
	return h
}

func handleKey(ev *tcell.EventKey) Action {
	if d, ok := moveKeys[ev.Key()]; ok {
		return BumpAction{ActionWithDirection{d[0], d[1]}}
	}
	switch ev.Key() {
	case tcell.KeyEscape, tcell.KeyCtrlC:
		return EscapeAction{}
	case tcell.KeyRune:
		if d, ok := moveRunes[ev.Rune()]; ok {
			return BumpAction{ActionWithDirection{d[0], d[1]}}
		}
		if ev.Rune() == '.' || ev.Rune() == '5' {
			return WaitAction{}
		}
	}
	return nil
}

func trackMouse(engine *Engine, ev tcell.Event) {
	if mouse, ok := ev.(*tcell.EventMouse); ok {
		x, y := mouse.Position()
		if engine.GameMap.InBounds(x, y) {
			engine.MouseX, engine.MouseY = x, y
		}
	}
}

type GameOverEventHandler struct {
	Engine *Engine
}

func (h *GameOverEventHandler) OnRender(screen tcell.Screen) { h.Engine.Render(screen) }

func (h *GameOverEventHandler) HandleEvent(ev tcell.Event) EventHandler {
	if key, ok := ev.(*tcell.EventKey); ok {
		if key.Key() == tcell.KeyEscape || key.Key() == tcell.KeyCtrlC {
			return nil
		}
	}
	return h
}

type HistoryViewer struct {
	Engine    *Engine
	Parent    EventHandler
	LogLength int
	Cursor    int
}

func NewHistoryViewer(engine *Engine, parent EventHandler) *HistoryViewer {
	n := len(engine.MessageLog.Messages)
	return &HistoryViewer{Engine: engine, Parent: parent, LogLength: n, Cursor: n - 1}
}

func (h *HistoryViewer) OnRender(screen tcell.Screen) {
	h.Parent.OnRender(screen)
	w, ht := screenWidth-6, screenHeight-6
	clearRect(screen, 3, 3, w, ht, tcell.StyleDefault)
	drawFrame(screen, 3, 3, w, ht, "┤Message history├", tcell.StyleDefault)
	renderMessages(screen, 4, 4, w-2, ht-2, h.Engine.MessageLog.Messages[:h.Cursor+1])
}

func (h *HistoryViewer) HandleEvent(ev tcell.Event) EventHandler {
	key, ok := ev.(*tcell.EventKey)
	if !ok {
		return h
	}
	step := 0
	switch key.Key() {
	case tcell.KeyUp:
		step = -1
	case tcell.KeyDown:
		step = 1
	case tcell.KeyPgUp:
		step = -10
	case tcell.KeyPgDn:
		step = 10
	case tcell.KeyHome:
		h.Cursor = 0
		return h
	case tcell.KeyEnd:
		h.Cursor = h.LogLength - 1
		return h
	default:
		return h.Parent
	}
	if step < 0 && h.Cursor == 0 {
		h.Cursor = h.LogLength - 1
	} else if step > 0 && h.Cursor == h.LogLength-1 {
		h.Cursor = 0
	} else {
		h.Cursor = max(0, min(h.Cursor+step, h.LogLength-1))
	}
	return h
}
