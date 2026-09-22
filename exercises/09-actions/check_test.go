package main

import (
	"testing"

	"github.com/gdamore/tcell/v2"
)

func TestDiagonalKeys(t *testing.T) {
	cases := map[rune]MovementAction{'y': {DX: -1, DY: -1}, 'u': {DX: 1, DY: -1}, 'b': {DX: -1, DY: 1}, 'n': {DX: 1, DY: 1},
		'h': {DX: -1, DY: 0}, 'l': {DX: 1, DY: 0}}
	for r, want := range cases {
		if got := handleKey(tcell.NewEventKey(tcell.KeyRune, r, 0)); got != want {
			t.Errorf("key %q: want %+v, got %+v", r, want, got)
		}
	}
	if got := handleKey(tcell.NewEventKey(tcell.KeyUp, 0, 0)); got != (MovementAction{DX: 0, DY: -1}) {
		t.Errorf("arrow up broke: %+v", got)
	}
	if _, ok := handleKey(tcell.NewEventKey(tcell.KeyEscape, 0, 0)).(EscapeAction); !ok {
		t.Error("escape broke")
	}
	if got := handleKey(tcell.NewEventKey(tcell.KeyRune, 'x', 0)); got != nil {
		t.Errorf("'x' should return nil, got %+v", got)
	}
}
