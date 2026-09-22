package main

import (
	"encoding/gob"
	"os"
)

const saveFile = "savegame.sav"

const saveVersion = 1

type saveData struct {
	Version     int // your turn! write it in SaveAs, check it in LoadGame
	GameMap     *GameMap
	MessageLog  *MessageLog
	PlayerIndex int
}

func init() {
	gob.Register(HostileEnemy{})
	gob.Register(&ConfusedEnemy{})
	gob.Register(HealingConsumable{})
	gob.Register(LightningDamageConsumable{})
	gob.Register(ConfusionConsumable{})
	gob.Register(FireballDamageConsumable{})
}

func (e *Engine) SaveAs(path string) error {
	data := saveData{GameMap: e.GameMap, MessageLog: e.MessageLog, PlayerIndex: -1}
	for i, ent := range e.GameMap.Entities {
		if ent == e.Player {
			data.PlayerIndex = i
		}
	}
	f, err := os.Create(path)
	if err != nil {
		return err
	}
	defer f.Close()
	return gob.NewEncoder(f).Encode(data)
}

func LoadGame(path string) (*Engine, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	var data saveData
	if err := gob.NewDecoder(f).Decode(&data); err != nil {
		return nil, err
	}
	engine := &Engine{GameMap: data.GameMap, MessageLog: data.MessageLog}
	engine.Player = data.GameMap.Entities[data.PlayerIndex]
	engine.UpdateFOV()
	return engine, nil
}
