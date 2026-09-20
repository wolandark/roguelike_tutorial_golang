package main

import (
	"errors"
	"fmt"

	"github.com/gdamore/tcell/v2"
)

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
	if key.Key() == tcell.KeyRune {
		switch key.Rune() {
		case 'v':
			return NewHistoryViewer(h.Engine, h)
		case 'i':
			return NewInventoryHandler(h.Engine, h, "Select an item to use", useItem)
		case 'd':
			return NewInventoryHandler(h.Engine, h, "Select an item to drop", dropItem)
		case '/':
			return NewLookHandler(h.Engine)
		}
	}
	action := handleKey(key)
	if _, ok := action.(EscapeAction); ok {
		return nil
	}
	return runAction(h.Engine, h, action)
}

func runAction(engine *Engine, self EventHandler, action Action) EventHandler {
	if action == nil {
		return self
	}
	if err := action.Perform(engine, engine.Player); err != nil {
		var imp Impossible
		if errors.As(err, &imp) {
			engine.Log(imp.Msg, colorImpossible)
			return self
		}
		engine.Log(err.Error(), colorError)
		return self
	}
	engine.HandleEnemyTurns()
	engine.UpdateFOV()
	if !engine.Player.Alive {
		return &GameOverEventHandler{Engine: engine}
	}
	return &MainGameEventHandler{Engine: engine}
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
		switch ev.Rune() {
		case '.', '5':
			return WaitAction{}
		case 'g':
			return PickupAction{}
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

type InventoryHandler struct {
	Engine   *Engine
	Parent   EventHandler
	Title    string
	OnSelect func(engine *Engine, item *Entity) (Action, EventHandler)
}

func NewInventoryHandler(engine *Engine, parent EventHandler, title string,
	onSelect func(*Engine, *Entity) (Action, EventHandler)) *InventoryHandler {
	return &InventoryHandler{Engine: engine, Parent: parent, Title: title, OnSelect: onSelect}
}

func useItem(engine *Engine, item *Entity) (Action, EventHandler) {
	return item.Consumable.GetAction(engine, engine.Player, item)
}

func dropItem(engine *Engine, item *Entity) (Action, EventHandler) {
	return DropItem{Item: item}, nil
}

func (h *InventoryHandler) OnRender(screen tcell.Screen) {
	h.Parent.OnRender(screen)
	items := h.Engine.Player.Inventory.Items
	height := max(len(items)+2, 3)
	x := 0
	if h.Engine.Player.X <= 30 {
		x = 40
	}
	width := len([]rune(h.Title)) + 8
	style := tcell.StyleDefault.Foreground(colorWhite).Background(colorBlack)
	clearRect(screen, x, 0, width, height, style)
	drawFrame(screen, x, 0, width, height, h.Title, style)
	if len(items) == 0 {
		drawText(screen, x+1, 1, "(Empty)", tcell.StyleDefault)
		return
	}
	for i, item := range items {
		drawText(screen, x+1, 1+i, fmt.Sprintf("(%c) %s", 'a'+i, item.Name), tcell.StyleDefault)
	}
}

func (h *InventoryHandler) HandleEvent(ev tcell.Event) EventHandler {
	key, ok := ev.(*tcell.EventKey)
	if !ok {
		return h
	}
	if key.Key() == tcell.KeyRune && key.Rune() >= 'a' && key.Rune() <= 'z' {
		index := int(key.Rune() - 'a')
		items := h.Engine.Player.Inventory.Items
		if index >= len(items) {
			h.Engine.Log("Invalid entry.", colorInvalid)
			return h
		}
		action, next := h.OnSelect(h.Engine, items[index])
		if next != nil {
			return next
		}
		return runAction(h.Engine, h, action)
	}
	return h.Parent
}

type SelectIndexHandler struct {
	Engine        *Engine
	OnSelect      func(x, y int) EventHandler
	OnRenderExtra func(screen tcell.Screen)
}

func NewSelectIndexHandler(engine *Engine, onSelect func(x, y int) EventHandler) *SelectIndexHandler {
	engine.MouseX, engine.MouseY = engine.Player.X, engine.Player.Y
	return &SelectIndexHandler{Engine: engine, OnSelect: onSelect}
}

func (h *SelectIndexHandler) OnRender(screen tcell.Screen) {
	h.Engine.Render(screen)
	if h.OnRenderExtra != nil {
		h.OnRenderExtra(screen)
	}
	x, y := h.Engine.MouseX, h.Engine.MouseY
	ch, _, style, _ := screen.GetContent(x, y)
	screen.SetContent(x, y, ch, nil, style.Reverse(true))
}

func (h *SelectIndexHandler) HandleEvent(ev tcell.Event) EventHandler {
	switch e := ev.(type) {
	case *tcell.EventMouse:
		x, y := e.Position()
		if h.Engine.GameMap.InBounds(x, y) {
			h.Engine.MouseX, h.Engine.MouseY = x, y
			if e.Buttons()&tcell.Button1 != 0 {
				return h.OnSelect(x, y)
			}
		}
		return h
	case *tcell.EventKey:
		d, ok := moveKeys[e.Key()]
		if !ok && e.Key() == tcell.KeyRune {
			d, ok = moveRunes[e.Rune()]
		}
		if ok {
			step := 1
			if e.Modifiers()&tcell.ModShift != 0 {
				step *= 5
			}
			if e.Modifiers()&tcell.ModCtrl != 0 {
				step *= 10
			}
			if e.Modifiers()&tcell.ModAlt != 0 {
				step *= 20
			}
			x := max(0, min(h.Engine.MouseX+d[0]*step, h.Engine.GameMap.Width-1))
			y := max(0, min(h.Engine.MouseY+d[1]*step, h.Engine.GameMap.Height-1))
			h.Engine.MouseX, h.Engine.MouseY = x, y
			return h
		}
		if e.Key() == tcell.KeyEnter {
			return h.OnSelect(h.Engine.MouseX, h.Engine.MouseY)
		}
		return &MainGameEventHandler{Engine: h.Engine}
	}
	return h
}

func NewLookHandler(engine *Engine) *SelectIndexHandler {
	return NewSelectIndexHandler(engine, func(x, y int) EventHandler {
		return &MainGameEventHandler{Engine: engine}
	})
}

func NewSingleRangedAttackHandler(engine *Engine, callback func(x, y int) Action) *SelectIndexHandler {
	var h *SelectIndexHandler
	h = NewSelectIndexHandler(engine, func(x, y int) EventHandler {
		return runAction(engine, h, callback(x, y))
	})
	return h
}

func NewAreaRangedAttackHandler(engine *Engine, radius int, callback func(x, y int) Action) *SelectIndexHandler {
	h := NewSingleRangedAttackHandler(engine, callback)
	h.OnRenderExtra = func(screen tcell.Screen) {
		size := radius*2 + 1
		style := tcell.StyleDefault.Foreground(colorRed)
		drawFrame(screen, engine.MouseX-radius, engine.MouseY-radius, size, size, "", style)
	}
	return h
}
