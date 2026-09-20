package main

import (
	"encoding/gob"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestSaveVersion(t *testing.T) {
	if saveVersion != 1 {
		t.Fatalf("saveVersion should be 1, got %d", saveVersion)
	}
	e := NewGame()
	dir := t.TempDir()
	good := filepath.Join(dir, "good.sav")
	if err := e.SaveAs(good); err != nil {
		t.Fatal(err)
	}
	f, _ := os.Open(good)
	var raw saveData
	if err := gob.NewDecoder(f).Decode(&raw); err != nil {
		t.Fatal(err)
	}
	f.Close()
	if raw.Version != saveVersion {
		t.Fatalf("SaveAs wrote version %d, want %d", raw.Version, saveVersion)
	}
	if loaded, err := LoadGame(good); err != nil || loaded.Player.X != e.Player.X {
		t.Fatalf("a current save must load: %v", err)
	}
	bad := filepath.Join(dir, "bad.sav")
	raw.Version = 99
	bf, _ := os.Create(bad)
	if err := gob.NewEncoder(bf).Encode(raw); err != nil {
		t.Fatal(err)
	}
	bf.Close()
	_, err := LoadGame(bad)
	if err == nil || !strings.Contains(strings.ToLower(err.Error()), "version") {
		t.Fatalf("loading a version 99 save should fail mentioning the version, got: %v", err)
	}
}
