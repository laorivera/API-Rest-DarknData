package main

var Characterlist = []Stats{
	characterStats, classFighter, classBarbarian, classRogue, classWizard, classCleric, classWarlock, classBard, classDruid, classRanger, classSorcerer,
}

var characterStats = Stats{
	Strength:        1,
	Vigor:           1,
	Agility:         1,
	Dexterity:       1,
	Will:            1,
	Knowledge:       1,
	Resourcefulness: 1,
}

var classFighter = Stats{
	Strength:        15,
	Vigor:           15,
	Agility:         15,
	Dexterity:       15,
	Will:            15,
	Knowledge:       15,
	Resourcefulness: 15,
}

var classBarbarian = Stats{
	Strength:        20,
	Vigor:           20,
	Agility:         13,
	Dexterity:       12,
	Will:            18,
	Knowledge:       5,
	Resourcefulness: 12,
}

var classRogue = Stats{
	Strength:        9,
	Vigor:           10,
	Agility:         21,
	Dexterity:       25,
	Will:            10,
	Knowledge:       10,
	Resourcefulness: 20,
}

var classWizard = Stats{
	Strength:        6,
	Vigor:           7,
	Agility:         15,
	Dexterity:       17,
	Will:            20,
	Knowledge:       25,
	Resourcefulness: 15,
}

var classCleric = Stats{
	Strength:        11,
	Vigor:           13,
	Agility:         12,
	Dexterity:       14,
	Will:            23,
	Knowledge:       20,
	Resourcefulness: 12,
}

var classWarlock = Stats{
	Strength:        11,
	Vigor:           14,
	Agility:         14,
	Dexterity:       15,
	Will:            22,
	Knowledge:       15,
	Resourcefulness: 14,
}

var classBard = Stats{
	Strength:        13,
	Vigor:           13,
	Agility:         13,
	Dexterity:       20,
	Will:            11,
	Knowledge:       20,
	Resourcefulness: 15,
}

var classDruid = Stats{
	Strength:        12,
	Vigor:           13,
	Agility:         12,
	Dexterity:       12,
	Will:            18,
	Knowledge:       20,
	Resourcefulness: 18,
}

var classRanger = Stats{
	Strength:        12,
	Vigor:           10,
	Agility:         20,
	Dexterity:       18,
	Will:            10,
	Knowledge:       12,
	Resourcefulness: 23,
}

var classSorcerer = Stats{
	Strength:        10,
	Vigor:           10,
	Agility:         10,
	Dexterity:       18,
	Will:            25,
	Knowledge:       20,
	Resourcefulness: 12,
}

type Character struct {
	Fighter   Stats
	Barbarian Stats
	Rogue     Stats
	Wizard    Stats
	Cleric    Stats
	Warlock   Stats
	Bard      Stats
	Druid     Stats
	Ranger    Stats
	Sorcerer  Stats
}

var SCharacter = Character{
	Fighter:   classFighter,
	Barbarian: classBarbarian,
	Rogue:     classRogue,
	Wizard:    classWizard,
	Cleric:    classCleric,
	Warlock:   classWarlock,
	Bard:      classBard,
	Ranger:    classRanger,
	Sorcerer:  classSorcerer,
}
