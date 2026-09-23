package main

import "fmt"

type Level struct {
	CurrentLevel  int
	CurrentXP     int
	LevelUpBase   int // XP needed for every level
	LevelUpFactor int // extra XP for each level already reached
	XPGiven       int // XP awarded to whoever kills this entity
}

func (l *Level) ExperienceToNextLevel() int {
	return l.LevelUpBase + l.CurrentLevel*l.LevelUpFactor
}

func (l *Level) RequiresLevelUp() bool {
	return l.CurrentXP >= l.ExperienceToNextLevel()
}

func (l *Level) AddXP(engine *Engine, xp int) {
	if xp == 0 || l.LevelUpBase == 0 {
		return
	}
	l.CurrentXP += xp
	engine.Log(fmt.Sprintf("You gain %d experience points.", xp), colorWhite)
	if l.RequiresLevelUp() {
		engine.Log(fmt.Sprintf("You advance to level %d!", l.CurrentLevel+1), colorWhite)
	}
}
