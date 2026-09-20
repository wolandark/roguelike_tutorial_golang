package main

import (
	"strings"
	"testing"
)

func TestStairsHeal(t *testing.T) {
	e := NewGame()
	p := e.Player
	if err := (TakeStairsAction{}).Perform(e, p); err == nil {
		t.Fatal("stairs must still be Impossible away from them")
	}
	p.Fighter.HP = 10
	p.X, p.Y = e.GameMap.DownstairsX, e.GameMap.DownstairsY
	if err := (TakeStairsAction{}).Perform(e, p); err != nil {
		t.Fatal(err)
	}
	if e.GameWorld.CurrentFloor != 2 || p.Fighter.HP != 17 {
		t.Fatalf("want floor 2 and 17 HP, got floor %d and %d HP", e.GameWorld.CurrentFloor, p.Fighter.HP)
	}
	msgs := e.MessageLog.Messages
	if msgs[len(msgs)-2].Text != "You descend the staircase." {
		t.Fatalf("descend message missing or out of order: %q", msgs[len(msgs)-2].Text)
	}
	if last := msgs[len(msgs)-1]; !strings.Contains(last.Text, "recover 7 HP") || last.Color != colorHealthRecovered {
		t.Fatalf("bad heal message: %q", last.Text)
	}
	p.Fighter.HP = p.Fighter.MaxHP
	p.X, p.Y = e.GameMap.DownstairsX, e.GameMap.DownstairsY
	if err := (TakeStairsAction{}).Perform(e, p); err != nil {
		t.Fatal(err)
	}
	msgs = e.MessageLog.Messages
	if msgs[len(msgs)-1].Text != "You descend the staircase." {
		t.Fatalf("no heal message expected at full HP, got %q", msgs[len(msgs)-1].Text)
	}
}
