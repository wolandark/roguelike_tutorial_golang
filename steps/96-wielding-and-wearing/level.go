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

func (l *Level) increaseLevel() {
	l.CurrentXP -= l.ExperienceToNextLevel()
	l.CurrentLevel++
}

func (l *Level) IncreaseMaxHP(engine *Engine, entity *Entity, amount int) {
	entity.Fighter.MaxHP += amount
	entity.Fighter.HP += amount
	engine.Log("Your health improves!", colorWhite)
	l.increaseLevel()
}

func (l *Level) IncreasePower(engine *Engine, entity *Entity, amount int) {
	entity.Fighter.BasePower += amount
	engine.Log("You feel stronger!", colorWhite)
	l.increaseLevel()
}

func (l *Level) IncreaseDefense(engine *Engine, entity *Entity, amount int) {
	entity.Fighter.BaseDefense += amount
	engine.Log("Your movements are getting swifter!", colorWhite)
	l.increaseLevel()
}
