package main

import "testing"

func TestHeal(t *testing.T) {
	f := &Fighter{HP: 10, MaxHP: 30, Defense: 1, Power: 2}
	if got := f.Heal(5); got != 5 || f.HP != 15 {
		t.Fatalf("heal 5 from 10/30: got %d recovered, hp %d", got, f.HP)
	}
	if got := f.Heal(100); got != 15 || f.HP != 30 {
		t.Fatalf("overheal: got %d recovered, hp %d (want 15, 30)", got, f.HP)
	}
	if got := f.Heal(4); got != 0 || f.HP != 30 {
		t.Fatalf("heal at full: got %d recovered, hp %d", got, f.HP)
	}
	f.HP = 28
	if got := f.Heal(4); got != 2 || f.HP != 30 {
		t.Fatalf("heal 4 at 28/30: got %d, hp %d", got, f.HP)
	}
}
