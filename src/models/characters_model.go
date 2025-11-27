package models

var Characterlist = []Stats{
	CharacterStats, ClassFighter, ClassBarbarian, ClassRogue, ClassWizard, ClassCleric, ClassWarlock, ClassBard, ClassDruid, ClassRanger, ClassSorcerer,
}

var CharacterStats = Stats{
	Strength:        1,
	Vigor:           1,
	Agility:         1,
	Dexterity:       1,
	Will:            1,
	Knowledge:       1,
	Resourcefulness: 1,
}

var ClassFighter = Stats{
	Strength:        15,
	Vigor:           15,
	Agility:         15,
	Dexterity:       15,
	Will:            15,
	Knowledge:       15,
	Resourcefulness: 15,
}

var ClassBarbarian = Stats{
	Strength:        20,
	Vigor:           20,
	Agility:         13,
	Dexterity:       12,
	Will:            18,
	Knowledge:       5,
	Resourcefulness: 12,
}

var ClassRogue = Stats{
	Strength:        9,
	Vigor:           10,
	Agility:         21,
	Dexterity:       25,
	Will:            10,
	Knowledge:       10,
	Resourcefulness: 20,
}

var ClassWizard = Stats{
	Strength:        6,
	Vigor:           7,
	Agility:         15,
	Dexterity:       17,
	Will:            20,
	Knowledge:       25,
	Resourcefulness: 15,
}

var ClassCleric = Stats{
	Strength:        11,
	Vigor:           13,
	Agility:         12,
	Dexterity:       14,
	Will:            23,
	Knowledge:       20,
	Resourcefulness: 12,
}

var ClassWarlock = Stats{
	Strength:        11,
	Vigor:           14,
	Agility:         14,
	Dexterity:       15,
	Will:            22,
	Knowledge:       15,
	Resourcefulness: 14,
}

var ClassBard = Stats{
	Strength:        13,
	Vigor:           13,
	Agility:         13,
	Dexterity:       20,
	Will:            11,
	Knowledge:       20,
	Resourcefulness: 15,
}

var ClassDruid = Stats{
	Strength:        12,
	Vigor:           13,
	Agility:         12,
	Dexterity:       12,
	Will:            18,
	Knowledge:       20,
	Resourcefulness: 18,
}

var ClassRanger = Stats{
	Strength:        12,
	Vigor:           10,
	Agility:         20,
	Dexterity:       18,
	Will:            10,
	Knowledge:       12,
	Resourcefulness: 23,
}

var ClassSorcerer = Stats{
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
	Fighter:   ClassFighter,
	Barbarian: ClassBarbarian,
	Rogue:     ClassRogue,
	Wizard:    ClassWizard,
	Cleric:    ClassCleric,
	Warlock:   ClassWarlock,
	Bard:      ClassBard,
	Ranger:    ClassRanger,
	Sorcerer:  ClassSorcerer,
}
