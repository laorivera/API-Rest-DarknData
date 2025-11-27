package models

var RaceStats = map[string]Stats{
	"Orc": {
		Strength: 1,
		Agility:  -1,
	},
	"Elf": {
		Agility:  1,
		Strength: -1,
	},
	"Dark Elf": {
		Will:      1,
		Dexterity: -1,
	},
	"Mummy": {
		Vigor:   1,
		Agility: -1,
	},
	"Panther": {
		Agility:  2,
		Strength: -1,
		Vigor:    -1,
	},
	"Felidian": {
		Dexterity: 2,
		Strength:  -1,
		Vigor:     -1,
	},
	"Lycan": {
		Vigor: 4,
	},
}

var RaceComputed = map[string]Computed_Stats{
	"Frost Walker": {
		Health:      3,
		ActionSpeed: -1.5,
	},
	"Zombie": {
		BuffDuration:   5,
		DebuffDuration: -5,
	},
	"Skeleton": {
		MagicRating:     10,
		FromArmorRating: -10,
	},
	"Elite Skeleton": {
		MagicRating:     10,
		FromArmorRating: -10,
	},
	"Nightmare Skeleton": {
		FromArmorRating: 10,
		MagicRating:     -10,
	},
	"Skeleton Champion": {
		FromArmorRating: 5,
		MagicRating:     -5,
	},
	"Lizardmen": {
		FromArmorRating:   20,
		HeadshotReduction: 10,
	},
	"Lycan": {
		HeadshotReduction: 10,
	},
}

var CharacterStat = map[string]Stats{
	"Fighter": {
		Strength:        15,
		Vigor:           15,
		Agility:         15,
		Dexterity:       15,
		Will:            15,
		Knowledge:       15,
		Resourcefulness: 15,
	},
	"Barbarian": {
		Strength:        20,
		Vigor:           20,
		Agility:         13,
		Dexterity:       12,
		Will:            18,
		Knowledge:       5,
		Resourcefulness: 12,
	},
	"Rogue": {
		Strength:        9,
		Vigor:           10,
		Agility:         21,
		Dexterity:       25,
		Will:            10,
		Knowledge:       10,
		Resourcefulness: 20,
	},
	"Wizard": {
		Strength:        6,
		Vigor:           7,
		Agility:         15,
		Dexterity:       17,
		Will:            20,
		Knowledge:       25,
		Resourcefulness: 15,
	},
	"Cleric": {
		Strength:        11,
		Vigor:           13,
		Agility:         12,
		Dexterity:       14,
		Will:            23,
		Knowledge:       20,
		Resourcefulness: 12,
	},
	"Warlock": {
		Strength:        11,
		Vigor:           14,
		Agility:         14,
		Dexterity:       15,
		Will:            22,
		Knowledge:       15,
		Resourcefulness: 14,
	},
	"Bard": {
		Strength:        13,
		Vigor:           13,
		Agility:         13,
		Dexterity:       20,
		Will:            11,
		Knowledge:       20,
		Resourcefulness: 15,
	},
	"Druid": {
		Strength:        12,
		Vigor:           13,
		Agility:         12,
		Dexterity:       12,
		Will:            18,
		Knowledge:       20,
		Resourcefulness: 18,
	},
	"Ranger": {
		Strength:        12,
		Vigor:           10,
		Agility:         20,
		Dexterity:       18,
		Will:            10,
		Knowledge:       12,
		Resourcefulness: 23,
	},
	"Sorcerer": {
		Strength:        10,
		Vigor:           10,
		Agility:         10,
		Dexterity:       18,
		Will:            25,
		Knowledge:       20,
		Resourcefulness: 12,
	},
}
