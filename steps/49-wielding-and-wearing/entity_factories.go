package main

import "github.com/gdamore/tcell/v2"

var (
	playerTemplate = Entity{
		Char: '@', Color: tcell.ColorWhite, Name: "Player", BlocksMovement: true, Alive: true,
		RenderOrder: RenderActor,
		Fighter:     &Fighter{HP: 30, MaxHP: 30, BaseDefense: 1, BasePower: 2},
		Inventory:   &Inventory{Capacity: 26},
		Level:       &Level{CurrentLevel: 1, LevelUpBase: 200, LevelUpFactor: 150},
	}
	orc = Entity{
		Char: 'o', Color: tcell.NewRGBColor(63, 127, 63), Name: "Orc", BlocksMovement: true, Alive: true,
		RenderOrder: RenderActor,
		AI:          HostileEnemy{},
		Fighter:     &Fighter{HP: 10, MaxHP: 10, BaseDefense: 0, BasePower: 3},
		Level:       &Level{XPGiven: 35},
	}
	troll = Entity{
		Char: 'T', Color: tcell.NewRGBColor(0, 127, 0), Name: "Troll", BlocksMovement: true, Alive: true,
		RenderOrder: RenderActor,
		AI:          HostileEnemy{},
		Fighter:     &Fighter{HP: 16, MaxHP: 16, BaseDefense: 1, BasePower: 4},
		Level:       &Level{XPGiven: 100},
	}

	healthPotion = Entity{
		Char: '!', Color: tcell.NewRGBColor(127, 0, 255), Name: "Health Potion",
		RenderOrder: RenderItem,
		Consumable:  HealingConsumable{Amount: 4},
	}
	lightningScroll = Entity{
		Char: '~', Color: tcell.NewRGBColor(255, 255, 0), Name: "Lightning Scroll",
		RenderOrder: RenderItem,
		Consumable:  LightningDamageConsumable{Damage: 20, MaximumRange: 5},
	}
	confusionScroll = Entity{
		Char: '~', Color: tcell.NewRGBColor(207, 63, 255), Name: "Confusion Scroll",
		RenderOrder: RenderItem,
		Consumable:  ConfusionConsumable{NumberOfTurns: 10},
	}
	fireballScroll = Entity{
		Char: '~', Color: tcell.NewRGBColor(255, 0, 0), Name: "Fireball Scroll",
		RenderOrder: RenderItem,
		Consumable:  FireballDamageConsumable{Damage: 12, Radius: 3},
	}

	dagger = Entity{
		Char: '/', Color: tcell.NewRGBColor(0, 191, 255), Name: "Dagger",
		RenderOrder: RenderItem,
		Equippable:  &Equippable{Type: Weapon, PowerBonus: 2},
	}
	sword = Entity{
		Char: '/', Color: tcell.NewRGBColor(0, 191, 255), Name: "Sword",
		RenderOrder: RenderItem,
		Equippable:  &Equippable{Type: Weapon, PowerBonus: 4},
	}
	leatherArmor = Entity{
		Char: '[', Color: tcell.NewRGBColor(139, 69, 19), Name: "Leather Armor",
		RenderOrder: RenderItem,
		Equippable:  &Equippable{Type: Armor, DefenseBonus: 1},
	}
	chainMail = Entity{
		Char: '[', Color: tcell.NewRGBColor(139, 69, 19), Name: "Chain Mail",
		RenderOrder: RenderItem,
		Equippable:  &Equippable{Type: Armor, DefenseBonus: 3},
	}
)
