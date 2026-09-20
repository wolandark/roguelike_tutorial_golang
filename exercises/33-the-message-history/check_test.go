package main

import (
	"reflect"
	"strings"
	"testing"
)

func TestWrapLongWords(t *testing.T) {
	got := wrap("the orc attacks you for 3 hit points", 12)
	want := []string{"the orc", "attacks you", "for 3 hit", "points"}
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("normal wrapping changed:\n got %q\nwant %q", got, want)
	}
	long := "Supercalifragilisticexpialidocious"
	for _, line := range wrap("a "+long+" scroll", 10) {
		if len([]rune(line)) > 10 {
			t.Fatalf("line %q is longer than 10", line)
		}
	}
	if joined := strings.Join(wrap(long, 10), ""); joined != long {
		t.Fatalf("characters lost or reordered: %q", joined)
	}
	if n := len(wrap(long, 10)); n != 4 {
		t.Fatalf("%d letters at width 10 should give 4 lines, got %d", len(long), n)
	}
	if lines := wrap("ab ééééééé cd", 5); len(lines) != 3 || lines[1] != "ééééé" {
		t.Fatalf("count runes, not bytes: %q", lines)
	}
}
