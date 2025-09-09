package main

type Attributes struct {
	Strenght        []float64
	Vigor           []float32
	Agility         []float32
	Dexterity       []float32
	Will            []float32
	Knowledge       []float32
	Resourcefulness []float32
}

type PhysicalDamage struct {
	TruePhyisicalDamage    []float32
	PhysicalPowerBonus     []float32
	PhysicalPower          []float32
	AdditionalWeaponDamage []float32
	PhysicalDamage         []float32
	ArmorPenetration       []float32
}

type MagicalDamage struct {
	MagicalPowerBonus []float32
	MagicalPower      []float32
	MagicalDamage     []float32
	TrueMagicalDamage []float32
	MagicPenetration  []float32
}

type Reductions struct {
	ArmorRating             []float32
	PhysicalDamageReduction []float32
	ProjectileReduction     []float32
	MagicalDamageReduction  []float32
	MagicResistance         []float32
	CooldownReduction       []float32
}

type Actions struct {
	ActionSpeed             []float32
	RegularInteractionSpeed []float32
	MagicalInteractionSpeed []float32
	SpellCastingSpeed       []float32
	MoveSpeed               []float32
	MoveSpeedBonus          []float32
}

type Health struct {
	MaxHealthAdd    []float32
	MaxHealthBonus  []float32
	PhysicalHealing []float32
	MagicalHealing  []float32
}
type Statuses struct {
	BuffDurationBonus []float32
	DebuffDuration    []float32
}

type Memory struct {
	MemoryCapacityAdd   []float32
	MemoryCapacityBonus []float32
}

type Misc struct {
	Luck []float32
}

type Enchantmentx struct {
	Base      Attributes
	Physical  PhysicalDamage
	Magical   MagicalDamage
	Reduction Reductions
	Action    Actions
	Health    Health
	Status    Statuses
	Memory    Memory
	Misc      Misc
}

type Enchanmentcat struct {
	Attributes map[string][]float32
	Others     map[string][]float32
}

type Enchantment struct {
	Helmet          map[string][]float32
	Chest           map[string][]float32
	Gloves          map[string][]float32
	Pants           map[string][]float32
	Boots           map[string][]float32
	Cloak           map[string][]float32
	Necklace        map[string][]float32
	Ring            map[string][]float32
	WeaponOneHand   map[string][]float32
	ShieldOneHand   map[string][]float32
	ShieldTwoHand   map[string][]float32
	PhysicalTwoHand map[string][]float32
	HybridTwoHand   map[string][]float32
	MagicalTwoHand  map[string][]float32
}

var Enchantments = Enchantment{
	// Armor
	Helmet: map[string][]float32{
		"Strength": {1, 2}, "Vigor": {1, 2}, "Agility": {1, 2}, "Dexterity": {1, 2}, "Will": {1, 2}, "Knowledge": {1, 2}, "Resourcefulness": {1, 2},
		"ArmorPenetration": {0.6, 1.2}, "PhysicalPowerBonus": {0.6, 1.2}, "MagicalPowerBonus": {0.7, 1.2}, "MagicPenetration": {0.7, 1.2},
		"ArmorRating": {3, 6}, "MagicResistance": {3, 6}, "PhysicalDamageReduction": {0.3, 0.6}, "MagicalDamageReduction": {0.2, 0.5}, "ProjectileReduction": {1, 1.5},
		"ActionSpeed": {1, 1.2}, "SpellCastingSpeed": {0.7, 1.5}, "RegularInteractionSpeed": {1, 2}, "MagicalInteractionSpeed": {1, 2},
		"BuffDurationBonus": {1, 1.5}, "DebuffDurationBonus": {0.7, 1.5}, "CooldownReduction": {0.7, 1.4},
		"MemoryCapacityBonus": {1.5, 3},
		"Luck":                {15, 30},
	},
	Chest: map[string][]float32{
		"Strength": {1, 3}, "Vigor": {1, 3}, "Agility": {1, 3}, "Dexterity": {1, 3}, "Will": {1, 3}, "Knowledge": {1, 3}, "Resourcefulness": {1, 3},
		"ArmorPenetration": {0.9, 1.8}, "PhysicalPowerBonus": {0.9, 1.8}, "MagicalPowerBonus": {0.7, 1.9}, "MagicPenetration": {1, 1.9},
		"ArmorRating": {5, 10}, "MagicResistance": {5, 10}, "PhysicalDamageReduction": {0.5, 1}, "MagicalDamageReduction": {0.5, 1}, "ProjectileReduction": {1, 2},
		"ActionSpeed": {0.9, 1.8}, "SpellCastingSpeed": {1, 2}, "RegularInteractionSpeed": {1.5, 3}, "MagicalInteractionSpeed": {1.5, 3},
		"BuffDurationBonus": {1, 2}, "DebuffDurationBonus": {1, 2}, "CooldownReduction": {1, 2},
		"MemoryCapacityAdd": {1, 4}, "MemoryCapacityBonus": {2, 4.5},
		"Luck": {20, 45},
	},
	Gloves: map[string][]float32{
		"Strength": {1, 2}, "Vigor": {1, 2}, "Agility": {1, 2}, "Dexterity": {1, 2}, "Will": {1, 2}, "Knowledge": {1, 2}, "Resourcefulness": {1, 2},
		"ArmorPenetration": {0.6, 1.2}, "PhysicalPowerBonus": {0.6, 1.2}, "MagicalPowerBonus": {0.6, 1.2}, "MagicPenetration": {0.7, 1.2},
		"ArmorRating": {3, 6}, "MagicResistance": {3, 6}, "PhysicalDamageReduction": {0.3, 0.6}, "MagicalDamageReduction": {0.2, 0.5}, "ProjectileReduction": {1, 1.5},
		"ActionSpeed": {1, 1.2}, "SpellCastingSpeed": {0.7, 1.5}, "RegularInteractionSpeed": {1, 2}, "MagicalInteractionSpeed": {1, 2},
		"BuffDurationBonus": {1, 1.5}, "DebuffDurationBonus": {0.7, 1.5}, "CooldownReduction": {0.7, 1.4},
		"MemoryCapacityBonus": {1.5, 3}, "MemoryCapacityAdd": {1, 2},
		"Luck": {15, 30},
	},
	Pants: map[string][]float32{
		"Strength": {1, 3}, "Vigor": {1, 3}, "Agility": {1, 3}, "Dexterity": {1, 3}, "Will": {1, 3}, "Knowledge": {1, 3}, "Resourcefulness": {1, 3},
		"ArmorPenetration": {0.9, 1.8}, "PhysicalPowerBonus": {0.9, 1.8}, "MagicalPowerBonus": {0.7, 1.8}, "MagicPenetration": {1, 1.8},
		"ArmorRating": {5, 10}, "MagicResistance": {5, 10}, "PhysicalDamageReduction": {0.5, 1}, "MagicalDamageReduction": {0.5, 1}, "ProjectileReduction": {1, 2},
		"ActionSpeed": {0.9, 1.8}, "SpellCastingSpeed": {1, 2}, "RegularInteractionSpeed": {1.5, 3}, "MagicalInteractionSpeed": {1.5, 3},
		"BuffDurationBonus": {1, 2}, "DebuffDurationBonus": {1, 2}, "CooldownReduction": {1, 2},
		"MemoryCapacityAdd": {1, 4}, "MemoryCapacityBonus": {2, 4.5},
		"Luck": {20, 45},
	},
	Boots: map[string][]float32{
		"Strength": {1, 2}, "Vigor": {1, 2}, "Agility": {1, 2}, "Dexterity": {1, 2}, "Will": {1, 2}, "Knowledge": {1, 2}, "Resourcefulness": {1, 2},
		"ArmorPenetration": {0.6, 1.2}, "PhysicalPowerBonus": {0.6, 1.2}, "MagicalPowerBonus": {0.6, 1.2}, "MagicPenetration": {0.7, 1.2},
		"ArmorRating": {3, 6}, "MagicResistance": {3, 6}, "PhysicalDamageReduction": {0.3, 0.6}, "MagicalDamageReduction": {0.2, 0.5}, "ProjectileReduction": {1, 1.5},
		"ActionSpeed": {1, 1.2}, "SpellCastingSpeed": {0.7, 1.5}, "RegularInteractionSpeed": {1, 2}, "MagicalInteractionSpeed": {1, 2},
		"BuffDurationBonus": {1, 1.5}, "DebuffDurationBonus": {0.7, 1.5}, "CooldownReduction": {0.7, 1.4},
		"MemoryCapacityBonus": {1.5, 3}, "MemoryCapacityAdd": {1, 2},
		"Luck": {15, 30},
	},
	Cloak: map[string][]float32{
		"ArmorPenetration": {0.6, 1.2}, "PhysicalPowerBonus": {0.6, 1.2}, "MagicalPowerBonus": {0.6, 1.2}, "MagicPenetration": {0.7, 1.2},
		"ArmorRating": {3, 6}, "MagicResistance": {3, 6}, "PhysicalDamageReduction": {0.3, 0.6}, "MagicalDamageReduction": {0.2, 0.5}, "ProjectileReduction": {1, 1.5},
		"ActionSpeed": {1, 1.2}, "SpellCastingSpeed": {0.7, 1.5}, "RegularInteractionSpeed": {1, 2}, "MagicalInteractionSpeed": {1, 2},
		"BuffDurationBonus": {1, 1.5}, "DebuffDurationBonus": {0.7, 1.5}, "CooldownReduction": {0.7, 1.4},
		"MemoryCapacityBonus": {1.5, 3}, "MemoryCapacityAdd": {1, 2},
		"MagicalHealing": {1, 2}, "MaxHealthBonus": {1, 2}, "MaxHealthAdd": {1, 3}, "PhysicalHealing": {1, 2},
		"Luck": {15, 30},
	},
	// Accessories
	Necklace: map[string][]float32{
		"AllAttributes":    {1},
		"ArmorPenetration": {0.9, 1.8}, "PhysicalPowerBonus": {0.9, 1.8}, "MagicalPowerBonus": {0.7, 1.8}, "MagicPenetration": {1, 1.8},
		"ArmorRating": {5, 10}, "MagicResistance": {5, 10}, "PhysicalDamageReduction": {0.5, 1}, "MagicalDamageReduction": {0.5, 1}, "ProjectileReduction": {1, 2},
		"ActionSpeed": {0.9, 1.8}, "SpellCastingSpeed": {1, 2}, "RegularInteractionSpeed": {1.5, 3}, "MagicalInteractionSpeed": {1.5, 3},
		"BuffDurationBonus": {1, 2}, "DebuffDurationBonus": {1, 2}, "CooldownReduction": {1, 2},
		"MemoryCapacityAdd": {1, 4}, "MemoryCapacityBonus": {2, 4.5},
		"MagicalHealing": {1, 2}, "MaxHealthBonus": {1, 2}, "MaxHealthAdd": {1, 3}, "PhysicalHealing": {1, 2},
		"Luck": {15, 30},
	},
	Ring: map[string][]float32{
		"AllAttributes":    {1},
		"ArmorPenetration": {0.6, 1.2}, "PhysicalPowerBonus": {0.6, 1.2}, "MagicalPowerBonus": {0.6, 1.2}, "MagicPenetration": {0.7, 1.2},
		"ArmorRating": {3, 6}, "MagicResistance": {3, 6}, "PhysicalDamageReduction": {0.3, 0.6}, "MagicalDamageReduction": {0.2, 0.5}, "ProjectileReduction": {1, 1.5},
		"ActionSpeed": {1, 1.2}, "SpellCastingSpeed": {0.7, 1.5}, "RegularInteractionSpeed": {1, 2}, "MagicalInteractionSpeed": {1, 2},
		"BuffDurationBonus": {1, 1.5}, "DebuffDurationBonus": {0.7, 1.5}, "CooldownReduction": {0.7, 1.4},
		"MemoryCapacityBonus": {1.5, 3}, "MemoryCapacityAdd": {1, 2},
		"MagicalHealing": {1, 2}, "MaxHealthBonus": {1, 2}, "MaxHealthAdd": {1, 3}, "PhysicalHealing": {1, 2},
		"Luck": {15, 30},
	},
	// Weapons
	WeaponOneHand: map[string][]float32{
		"PhysicalPower": {1, 2}, "ArmorPenetration": {2, 3}, "PhysicalPowerBonus": {1, 2}, "AdditionalWeaponDamage": {1},
		"MagicalPower": {1, 2}, "MagicPenetration": {2, 3}, "MagicalPowerBonus": {1, 2},
		"CooldownReduction": {1, 2},
		"ActionSpeed":       {1, 2}, "RegularIntractionSpeed": {2, 3}, "MagicalInteractionSpeed": {2, 3}, "SpellCastingSpeed": {1, 2},
		"MaxHealthBonus": {1, 2}, "MaxHealthAdd": {1, 3},
		"MagicalHealing":    {1, 2},
		"BuffDurationBonus": {2, 3}, "DebuffDurationBonus": {2, 3},
		"MemoryCapacityAdd": {1, 3}, "MemoryCapacityBonus": {3, 6},
		"MoveSpeed": {1, 3}, "MoveSpeedBonus": {0.5, 1},
	},
	ShieldOneHand: map[string][]float32{
		"PhysicalPowerBonus": {1, 2}, "AdditionalWeaponDamage": {1},
		"ArmorRating": {5, 10}, "MagicResistance": {5, 10}, "PhysicalDamageReduction": {1, 2}, "MagicalDamageReduction": {1, 2}, "ProjectileReduction": {1, 2}, "CooldownReduction": {1, 2},
		"ActionSpeed": {1, 2}, "RegularInteractionSpeed": {2, 3}, "MagicalInteractionSpeed": {2, 3}, "SpellCastingSpeed": {1, 2},
		"MaxHealthBonus": {1, 2}, "MaxHealthAdd": {1, 3},
		"MagicalHealing":    {1, 2},
		"BuffDurationBonus": {2, 3}, "DebuffDurationBonus": {2, 3},
		"MoveSpeed": {1, 3}, "MoveSpeedBonus": {0.5, 1},
	},
	ShieldTwoHand: map[string][]float32{
		"ArmorRating": {5, 15}, "MagicResistance": {10, 20}, "PhysicalDamageReduction": {2, 3}, "MagicalDamageReduction": {2, 3}, "ProjectileReduction": {2, 3}, "CooldownReduction": {1, 2},
		"ActionSpeed": {2, 3}, "RegularInteractionSpeed": {3, 4}, "MagicalInteractionSpeed": {3, 4},
		"MaxHealthBonus": {2, 3}, "MaxHealthAdd": {1, 4},
		"BuffDurationBonus": {3, 4}, "DebuffDurationBonus": {3, 4},
		"MoveSpeed": {1, 4}, "MoveSpeedBonus": {1, 2},
	},
	PhysicalTwoHand: map[string][]float32{
		"PhysicalPower": {1, 3}, "ArmorPenetration": {2, 3}, "PhysicalPowerBonus": {2, 3}, "AdditionalWeaponDamage": {1, 2},
		"CooldownReduction": {2, 3},
		"ActionSpeed":       {2, 3}, "RegularInteractionSpeed": {3, 4}, "MagicalInteractionSpeed": {3, 4},
		"MaxHealthBonus": {2, 3}, "MaxHealthAdd": {1, 4},
		"BuffDurationBonus": {3, 4}, "DebuffDurationBonus": {3, 4},
		"MoveSpeed": {1, 4}, "MoveSpeedBonus": {1, 2},
	},
	HybridTwoHand: map[string][]float32{
		"PhysicalPower": {1, 3}, "ArmorPenetration": {2, 3}, "PhysicalPowerBonus": {2, 3}, "AdditionalWeaponDamage": {1, 2},
		"MagicalPower": {1, 3}, "MagicPenetration": {3, 4}, "MagicalPowerBonus": {2, 3}, "CooldownReduction": {2, 3},
		"ActionSpeed": {2, 3}, "RegularInteractionSpeed": {3, 4}, "MagicalInteractionSpeed": {3, 4}, "SpellCastingSpeed": {2, 3},
		"MaxHealthBonus": {2, 3}, "MaxHealthAdd": {1, 4},
		"BuffDurationBonus": {3, 4}, "DebuffDurationBonus": {3, 4},
		"MemoryCapacityAdd": {1, 4}, "MemoryCapacityBonus": {4, 8},
		"MoveSpeed": {1, 4}, "MoveSpeedBonus": {1, 2},
	},
	MagicalTwoHand: map[string][]float32{
		"MaigcalPower": {1, 3}, "MagicPenetration": {3, 4}, "MagicalPowerBonus": {2, 3},
		"CooldownReduction": {2, 3},
		"ActionSpeed":       {2, 3}, "RegularInteractionSpeed": {3, 4}, "MagicalInteractionSpeed": {3, 4}, "SpellCastingSpeed": {2, 3},
		"MaxHealthBonus": {2, 3}, "MaxHealthAdd": {1, 4},
		"MagicalHealing":    {1, 3},
		"BuffDurationBonus": {3, 4}, "DebuffDurationBonus": {3, 4},
		"MemoryCapacityAdd": {1, 4}, "MemoryCapacityBonus": {4, 8},
		"MoveSpeed": {1, 4}, "MoveSpeedBonus": {1, 2},
	},
}

var EnchantmentMap = map[string]Enchantmentx{
	"chest":    chestEnchantment,
	"cloak":    cloakEnchantment,
	"foot":     bootsEnchantment,
	"hands":    glovesEnchantment,
	"head":     helmetEnchantment,
	"pants":    pantsEnchantment,
	"necklace": necklaceEnchantment,
	"ring":     ringEnchantment,
	"onehand":  onehandEnchantment,
	"twohand":  twohandEnchantment,
}

var chestEnchantment = Enchantmentx{
	Base: Attributes{
		Strenght:        []float64{1, 3},
		Vigor:           []float32{1, 3},
		Agility:         []float32{1, 3},
		Dexterity:       []float32{1, 3},
		Will:            []float32{1, 3},
		Knowledge:       []float32{1, 3},
		Resourcefulness: []float32{1, 3},
	},
	Physical: PhysicalDamage{
		TruePhyisicalDamage:    []float32{0}, // Not specified
		PhysicalPowerBonus:     []float32{0.9, 1.8},
		PhysicalPower:          []float32{1, 2},
		AdditionalWeaponDamage: []float32{0}, // Not specified
		PhysicalDamage:         []float32{0.9, 1.8},
		ArmorPenetration:       []float32{1.5, 3.0},
	},
	Magical: MagicalDamage{
		MagicalPowerBonus: []float32{0.7, 1.8},
		MagicalPower:      []float32{1, 2},
		MagicalDamage:     []float32{0.7, 1.8},
		TrueMagicalDamage: []float32{0}, // Not specified
		MagicPenetration:  []float32{1.5, 3.0},
	},
	Reduction: Reductions{
		ArmorRating:             []float32{5, 10},
		PhysicalDamageReduction: []float32{0.7, 1.5},
		ProjectileReduction:     []float32{1.0, 2.0},
		MagicalDamageReduction:  []float32{0.7, 1.5},
		MagicResistance:         []float32{5, 10},
		CooldownReduction:       []float32{1.0, 2.0},
	},
	Action: Actions{
		ActionSpeed:             []float32{0.9, 1.8},
		RegularInteractionSpeed: []float32{1.5, 3.0},
		MagicalInteractionSpeed: []float32{1.5, 3.0},
		SpellCastingSpeed:       []float32{1.2, 2.5},
		MoveSpeed:               []float32{0}, // Not specified
		MoveSpeedBonus:          []float32{0}, // Not specified
	},
	Health: Health{
		MaxHealthAdd:    []float32{2, 4},
		MaxHealthBonus:  []float32{1.5, 3.0},
		PhysicalHealing: []float32{1, 1},
		MagicalHealing:  []float32{1, 1},
	},
	Status: Statuses{
		BuffDurationBonus: []float32{1.0, 2.0},
		DebuffDuration:    []float32{1.0, 2.0},
	},
	Memory: Memory{
		MemoryCapacityAdd:   []float32{2, 4},
		MemoryCapacityBonus: []float32{2.2, 4.5},
	},
	Misc: Misc{
		Luck: []float32{22, 45},
	},
}

var pantsEnchantment = Enchantmentx{
	Base: Attributes{
		Strenght:        []float64{1, 3},
		Vigor:           []float32{1, 3},
		Agility:         []float32{1, 3},
		Dexterity:       []float32{1, 3},
		Will:            []float32{1, 3},
		Knowledge:       []float32{1, 3},
		Resourcefulness: []float32{1, 3},
	},
	Physical: PhysicalDamage{
		TruePhyisicalDamage:    []float32{0}, // Not specified
		PhysicalPowerBonus:     []float32{0.9, 1.8},
		PhysicalPower:          []float32{1, 2},
		AdditionalWeaponDamage: []float32{0}, // Not specified
		PhysicalDamage:         []float32{0.9, 1.8},
		ArmorPenetration:       []float32{1.5, 3.0},
	},
	Magical: MagicalDamage{
		MagicalPowerBonus: []float32{0.7, 1.8},
		MagicalPower:      []float32{1, 2},
		MagicalDamage:     []float32{0.7, 1.8},
		TrueMagicalDamage: []float32{0}, // Not specified
		MagicPenetration:  []float32{1.5, 3.0},
	},
	Reduction: Reductions{
		ArmorRating:             []float32{5, 10},
		PhysicalDamageReduction: []float32{0.7, 1.5},
		ProjectileReduction:     []float32{1.0, 2.0},
		MagicalDamageReduction:  []float32{0.7, 1.5},
		MagicResistance:         []float32{5, 10},
		CooldownReduction:       []float32{1.0, 2.0},
	},
	Action: Actions{
		ActionSpeed:             []float32{0.9, 1.8},
		RegularInteractionSpeed: []float32{1.5, 3.0},
		MagicalInteractionSpeed: []float32{1.5, 3.0},
		SpellCastingSpeed:       []float32{1.2, 2.5},
		MoveSpeed:               []float32{0}, // Not specified
		MoveSpeedBonus:          []float32{0}, // Not specified
	},
	Health: Health{
		MaxHealthAdd:    []float32{2, 4},
		MaxHealthBonus:  []float32{1.5, 3.0},
		PhysicalHealing: []float32{1, 1},
		MagicalHealing:  []float32{1, 1},
	},
	Status: Statuses{
		BuffDurationBonus: []float32{1.0, 2.0},
		DebuffDuration:    []float32{1.0, 2.0},
	},
	Memory: Memory{
		MemoryCapacityAdd:   []float32{2, 4},
		MemoryCapacityBonus: []float32{2.2, 4.5},
	},
	Misc: Misc{
		Luck: []float32{22, 45},
	},
}

var cloakEnchantment = Enchantmentx{
	Base: Attributes{
		Strenght:        []float64{1, 2},
		Vigor:           []float32{1, 2},
		Agility:         []float32{1, 2},
		Dexterity:       []float32{1, 2},
		Will:            []float32{1, 2},
		Knowledge:       []float32{1, 2},
		Resourcefulness: []float32{1, 2},
	},
	Physical: PhysicalDamage{
		TruePhyisicalDamage:    []float32{1, 1},
		PhysicalPowerBonus:     []float32{0.6, 1.2},
		PhysicalPower:          []float32{1, 1},
		AdditionalWeaponDamage: []float32{1, 1},
		PhysicalDamage:         []float32{0.6, 1.2},
		ArmorPenetration:       []float32{1.0, 2.0},
	},
	Magical: MagicalDamage{
		MagicalPowerBonus: []float32{0.5, 1.2},
		MagicalPower:      []float32{1, 1},
		MagicalDamage:     []float32{0.5, 1.2},
		TrueMagicalDamage: []float32{1, 1},
		MagicPenetration:  []float32{1.0, 2.0},
	},
	Reduction: Reductions{
		ArmorRating:             []float32{3, 6},
		PhysicalDamageReduction: []float32{0.5, 1.0},
		ProjectileReduction:     []float32{0.7, 1.5},
		MagicalDamageReduction:  []float32{0.5, 1.0},
		MagicResistance:         []float32{3, 6},
		CooldownReduction:       []float32{0.7, 1.5},
	},
	Action: Actions{
		ActionSpeed:             []float32{0.6, 1.2},
		RegularInteractionSpeed: []float32{1.0, 2.0},
		MagicalInteractionSpeed: []float32{1.0, 2.0},
		SpellCastingSpeed:       []float32{1.0, 2.0},
		MoveSpeed:               []float32{0}, // Not specified
		MoveSpeedBonus:          []float32{0}, // Not specified
	},
	Health: Health{
		MaxHealthAdd:    []float32{1, 3},
		MaxHealthBonus:  []float32{1.0, 2.0},
		PhysicalHealing: []float32{1, 1},
		MagicalHealing:  []float32{1, 1},
	},
	Status: Statuses{
		BuffDurationBonus: []float32{0.7, 1.5},
		DebuffDuration:    []float32{0.7, 1.5},
	},
	Memory: Memory{
		MemoryCapacityAdd:   []float32{1, 2},
		MemoryCapacityBonus: []float32{1.5, 3.0},
	},
	Misc: Misc{
		Luck: []float32{15, 30},
	},
}

var bootsEnchantment = Enchantmentx{
	Base: Attributes{
		Strenght:        []float64{1, 2},
		Vigor:           []float32{1, 2},
		Agility:         []float32{1, 2},
		Dexterity:       []float32{1, 2},
		Will:            []float32{1, 2},
		Knowledge:       []float32{1, 2},
		Resourcefulness: []float32{1, 2},
	},
	Physical: PhysicalDamage{
		TruePhyisicalDamage:    []float32{0}, // Not specified
		PhysicalPowerBonus:     []float32{0.6, 1.2},
		PhysicalPower:          []float32{1, 1},
		AdditionalWeaponDamage: []float32{0}, // Not specified
		PhysicalDamage:         []float32{0.6, 1.2},
		ArmorPenetration:       []float32{1.0, 2.0},
	},
	Magical: MagicalDamage{
		MagicalPowerBonus: []float32{0.5, 1.2},
		MagicalPower:      []float32{1, 1},
		MagicalDamage:     []float32{0.5, 1.2},
		TrueMagicalDamage: []float32{0}, // Not specified
		MagicPenetration:  []float32{1.0, 2.0},
	},
	Reduction: Reductions{
		ArmorRating:             []float32{3, 6},
		PhysicalDamageReduction: []float32{0.5, 1.0},
		ProjectileReduction:     []float32{0.7, 1.5},
		MagicalDamageReduction:  []float32{0.5, 1.0},
		MagicResistance:         []float32{3, 6},
		CooldownReduction:       []float32{0.7, 1.5},
	},
	Action: Actions{
		ActionSpeed:             []float32{0.6, 1.2},
		RegularInteractionSpeed: []float32{1.0, 2.0},
		MagicalInteractionSpeed: []float32{1.0, 2.0},
		SpellCastingSpeed:       []float32{1.0, 2.0},
		MoveSpeed:               []float32{1, 5},
		MoveSpeedBonus:          []float32{3, 15},
	},
	Health: Health{
		MaxHealthAdd:    []float32{1, 3},
		MaxHealthBonus:  []float32{1.0, 2.0},
		PhysicalHealing: []float32{1, 1},
		MagicalHealing:  []float32{1, 1},
	},
	Status: Statuses{
		BuffDurationBonus: []float32{0.7, 1.5},
		DebuffDuration:    []float32{0.7, 1.5},
	},
	Memory: Memory{
		MemoryCapacityAdd:   []float32{1, 2},
		MemoryCapacityBonus: []float32{1.5, 3.0},
	},
	Misc: Misc{
		Luck: []float32{15, 30},
	},
}

var glovesEnchantment = Enchantmentx{
	Base: Attributes{
		Strenght:        []float64{1, 2},
		Vigor:           []float32{1, 2},
		Agility:         []float32{1, 2},
		Dexterity:       []float32{1, 2},
		Will:            []float32{1, 2},
		Knowledge:       []float32{1, 2},
		Resourcefulness: []float32{1, 2},
	},
	Physical: PhysicalDamage{
		TruePhyisicalDamage:    []float32{1, 1},
		PhysicalPowerBonus:     []float32{0.6, 1.2},
		PhysicalPower:          []float32{1, 1},
		AdditionalWeaponDamage: []float32{1, 1},
		PhysicalDamage:         []float32{0.6, 1.2},
		ArmorPenetration:       []float32{1.0, 2.0},
	},
	Magical: MagicalDamage{
		MagicalPowerBonus: []float32{0.5, 1.2},
		MagicalPower:      []float32{1, 1},
		MagicalDamage:     []float32{0.5, 1.2},
		TrueMagicalDamage: []float32{0}, // Not specified
		MagicPenetration:  []float32{1.0, 2.0},
	},
	Reduction: Reductions{
		ArmorRating:             []float32{3, 6},
		PhysicalDamageReduction: []float32{0.5, 1.0},
		ProjectileReduction:     []float32{0.7, 1.5},
		MagicalDamageReduction:  []float32{0.5, 1.0},
		MagicResistance:         []float32{3, 6},
		CooldownReduction:       []float32{0.7, 1.5},
	},
	Action: Actions{
		ActionSpeed:             []float32{0.6, 1.2},
		RegularInteractionSpeed: []float32{1.0, 2.0},
		MagicalInteractionSpeed: []float32{1.0, 2.0},
		SpellCastingSpeed:       []float32{1.0, 2.0},
		MoveSpeed:               []float32{0}, // Not specified
		MoveSpeedBonus:          []float32{0}, // Not specified
	},
	Health: Health{
		MaxHealthAdd:    []float32{1, 3},
		MaxHealthBonus:  []float32{1.0, 2.0},
		PhysicalHealing: []float32{1, 1},
		MagicalHealing:  []float32{1, 1},
	},
	Status: Statuses{
		BuffDurationBonus: []float32{0.7, 1.5},
		DebuffDuration:    []float32{0.7, 1.5},
	},
	Memory: Memory{
		MemoryCapacityAdd:   []float32{1, 2},
		MemoryCapacityBonus: []float32{1.5, 3.0},
	},
	Misc: Misc{
		Luck: []float32{15, 30},
	},
}

var helmetEnchantment = Enchantmentx{
	Base: Attributes{
		Strenght:        []float64{1, 2},
		Vigor:           []float32{1, 2},
		Agility:         []float32{1, 2},
		Dexterity:       []float32{1, 2},
		Will:            []float32{1, 2},
		Knowledge:       []float32{1, 2},
		Resourcefulness: []float32{1, 2},
	},
	Physical: PhysicalDamage{
		TruePhyisicalDamage:    []float32{0}, // Not specified
		PhysicalPowerBonus:     []float32{0.6, 1.2},
		PhysicalPower:          []float32{1, 1},
		AdditionalWeaponDamage: []float32{1, 1},
		PhysicalDamage:         []float32{0.6, 1.2},
		ArmorPenetration:       []float32{1.0, 2.0},
	},
	Magical: MagicalDamage{
		MagicalPowerBonus: []float32{0.5, 1.2},
		MagicalPower:      []float32{1, 1},
		MagicalDamage:     []float32{0.5, 1.2},
		TrueMagicalDamage: []float32{1, 1},
		MagicPenetration:  []float32{1.0, 2.0},
	},
	Reduction: Reductions{
		ArmorRating:             []float32{3, 6},
		PhysicalDamageReduction: []float32{0.5, 1.0},
		ProjectileReduction:     []float32{0.7, 1.5},
		MagicalDamageReduction:  []float32{0.5, 1.0},
		MagicResistance:         []float32{3, 6},
		CooldownReduction:       []float32{0.7, 1.5},
	},
	Action: Actions{
		ActionSpeed:             []float32{0.6, 1.2},
		RegularInteractionSpeed: []float32{1.0, 2.0},
		MagicalInteractionSpeed: []float32{1.0, 2.0},
		SpellCastingSpeed:       []float32{1.0, 2.0},
		MoveSpeed:               []float32{0}, // Not specified
		MoveSpeedBonus:          []float32{0}, // Not specified
	},
	Health: Health{
		MaxHealthAdd:    []float32{1, 3},
		MaxHealthBonus:  []float32{1.0, 2.0},
		PhysicalHealing: []float32{1, 1},
		MagicalHealing:  []float32{1, 1},
	},
	Status: Statuses{
		BuffDurationBonus: []float32{0.7, 1.5},
		DebuffDuration:    []float32{0.7, 1.5},
	},
	Memory: Memory{
		MemoryCapacityAdd:   []float32{1, 2},
		MemoryCapacityBonus: []float32{1.5, 3.0},
	},
	Misc: Misc{
		Luck: []float32{15, 30},
	},
}

var necklaceEnchantment = Enchantmentx{
	Base: Attributes{
		Strenght:        []float64{1, 3},
		Vigor:           []float32{1, 3},
		Agility:         []float32{1, 3},
		Dexterity:       []float32{1, 3},
		Will:            []float32{1, 3},
		Knowledge:       []float32{1, 3},
		Resourcefulness: []float32{1, 3},
	},
	Physical: PhysicalDamage{
		TruePhyisicalDamage:    []float32{1, 1},
		PhysicalPowerBonus:     []float32{0.9, 1.8},
		PhysicalPower:          []float32{1, 2},
		AdditionalWeaponDamage: []float32{1, 1},
		PhysicalDamage:         []float32{0.9, 1.8},
		ArmorPenetration:       []float32{1.5, 3.0},
	},
	Magical: MagicalDamage{
		MagicalPowerBonus: []float32{0.7, 1.8},
		MagicalPower:      []float32{1, 2},
		MagicalDamage:     []float32{0.7, 1.8},
		TrueMagicalDamage: []float32{1, 1},
		MagicPenetration:  []float32{1.5, 3.0},
	},
	Reduction: Reductions{
		ArmorRating:             []float32{5, 10},
		PhysicalDamageReduction: []float32{0.7, 1.5},
		ProjectileReduction:     []float32{1.0, 2.0},
		MagicalDamageReduction:  []float32{0.7, 1.5},
		MagicResistance:         []float32{5, 10},
		CooldownReduction:       []float32{1.0, 2.0},
	},
	Action: Actions{
		ActionSpeed:             []float32{0.9, 1.8},
		RegularInteractionSpeed: []float32{1.5, 3.0},
		MagicalInteractionSpeed: []float32{1.5, 3.0},
		SpellCastingSpeed:       []float32{1.2, 2.5},
		MoveSpeed:               []float32{0}, // Not specified
		MoveSpeedBonus:          []float32{0}, // Not specified
	},
	Health: Health{
		MaxHealthAdd:    []float32{2, 4},
		MaxHealthBonus:  []float32{1.5, 3.0},
		PhysicalHealing: []float32{1, 1},
		MagicalHealing:  []float32{1, 1},
	},
	Status: Statuses{
		BuffDurationBonus: []float32{1.0, 2.0},
		DebuffDuration:    []float32{1.0, 2.0},
	},
	Memory: Memory{
		MemoryCapacityAdd:   []float32{2, 4},
		MemoryCapacityBonus: []float32{2.2, 4.5},
	},
	Misc: Misc{
		Luck: []float32{22, 45},
	},
}

var ringEnchantment = Enchantmentx{
	Base: Attributes{
		Strenght:        []float64{1, 2},
		Vigor:           []float32{1, 2},
		Agility:         []float32{1, 2},
		Dexterity:       []float32{1, 2},
		Will:            []float32{1, 2},
		Knowledge:       []float32{1, 2},
		Resourcefulness: []float32{1, 2},
	},
	Physical: PhysicalDamage{
		TruePhyisicalDamage:    []float32{1, 1},
		PhysicalPowerBonus:     []float32{0.6, 1.2},
		PhysicalPower:          []float32{1, 1},
		AdditionalWeaponDamage: []float32{1, 1},
		PhysicalDamage:         []float32{0.6, 1.2},
		ArmorPenetration:       []float32{1.0, 2.0},
	},
	Magical: MagicalDamage{
		MagicalPowerBonus: []float32{0.5, 1.2},
		MagicalPower:      []float32{1, 1},
		MagicalDamage:     []float32{0.5, 1.2},
		TrueMagicalDamage: []float32{1, 1},
		MagicPenetration:  []float32{1.0, 2.0},
	},
	Reduction: Reductions{
		ArmorRating:             []float32{3, 6},
		PhysicalDamageReduction: []float32{0.5, 1.0},
		ProjectileReduction:     []float32{0.7, 1.5},
		MagicalDamageReduction:  []float32{0.5, 1.0},
		MagicResistance:         []float32{3, 6},
		CooldownReduction:       []float32{0.7, 1.5},
	},
	Action: Actions{
		ActionSpeed:             []float32{0.6, 1.2},
		RegularInteractionSpeed: []float32{1.0, 2.0},
		MagicalInteractionSpeed: []float32{1.0, 2.0},
		SpellCastingSpeed:       []float32{1.0, 2.0},
		MoveSpeed:               []float32{0}, // Not specified
		MoveSpeedBonus:          []float32{0}, // Not specified
	},
	Health: Health{
		MaxHealthAdd:    []float32{1, 3},
		MaxHealthBonus:  []float32{1.0, 2.0},
		PhysicalHealing: []float32{1, 1},
		MagicalHealing:  []float32{1, 1},
	},
	Status: Statuses{
		BuffDurationBonus: []float32{0.7, 1.5},
		DebuffDuration:    []float32{0.7, 1.5},
	},
	Memory: Memory{
		MemoryCapacityAdd:   []float32{1, 2},
		MemoryCapacityBonus: []float32{1.5, 3.0},
	},
	Misc: Misc{
		Luck: []float32{15, 30},
	},
}

var onehandEnchantment = Enchantmentx{
	Base: Attributes{
		Strenght:        []float64{1, 1},
		Vigor:           []float32{1, 1},
		Agility:         []float32{1, 1},
		Dexterity:       []float32{1, 1},
		Will:            []float32{1, 1},
		Knowledge:       []float32{1, 1},
		Resourcefulness: []float32{1, 1},
	},
	Physical: PhysicalDamage{
		TruePhyisicalDamage:    []float32{0, 0}, // Not specified
		PhysicalPowerBonus:     []float32{0.6, 1.2},
		PhysicalPower:          []float32{1, 1},
		AdditionalWeaponDamage: []float32{1, 1},
		PhysicalDamage:         []float32{0.6, 1.2},
		ArmorPenetration:       []float32{1.0, 2.0},
	},
	Magical: MagicalDamage{
		MagicalPowerBonus: []float32{0.5, 1.2},
		MagicalPower:      []float32{1, 1},
		MagicalDamage:     []float32{0.5, 1.2},
		TrueMagicalDamage: []float32{0}, // Not specified
		MagicPenetration:  []float32{1.0, 2.0},
	},
	Reduction: Reductions{
		ArmorRating:             []float32{3, 6},
		PhysicalDamageReduction: []float32{0.5, 1.0},
		ProjectileReduction:     []float32{0.7, 1.5},
		MagicalDamageReduction:  []float32{0.5, 1.0},
		MagicResistance:         []float32{3, 6},
		CooldownReduction:       []float32{0.7, 1.5},
	},
	Action: Actions{
		ActionSpeed:             []float32{0.6, 1.2},
		RegularInteractionSpeed: []float32{1.0, 2.0},
		MagicalInteractionSpeed: []float32{1.0, 2.0},
		SpellCastingSpeed:       []float32{1.0, 2.0},
		MoveSpeed:               []float32{0}, // Not specified
		MoveSpeedBonus:          []float32{0}, // Not specified
	},
	Health: Health{
		MaxHealthAdd:    []float32{1, 2},
		MaxHealthBonus:  []float32{0.7, 1.5},
		PhysicalHealing: []float32{1, 1},
		MagicalHealing:  []float32{1, 1},
	},
	Status: Statuses{
		BuffDurationBonus: []float32{0.7, 1.5},
		DebuffDuration:    []float32{0.7, 1.5},
	},
	Memory: Memory{
		MemoryCapacityAdd:   []float32{1, 2},
		MemoryCapacityBonus: []float32{1.5, 3.0},
	},
	Misc: Misc{
		Luck: []float32{15, 30},
	},
}

var twohandEnchantment = Enchantmentx{
	Base: Attributes{
		Strenght:        []float64{1, 2},
		Vigor:           []float32{1, 2},
		Agility:         []float32{1, 2},
		Dexterity:       []float32{1, 2},
		Will:            []float32{1, 2},
		Knowledge:       []float32{1, 2},
		Resourcefulness: []float32{1, 2},
	},
	Physical: PhysicalDamage{
		TruePhyisicalDamage:    []float32{0, 0}, // Not specified
		PhysicalPowerBonus:     []float32{1.2, 2.4},
		PhysicalPower:          []float32{1, 2},
		AdditionalWeaponDamage: []float32{1, 2},
		PhysicalDamage:         []float32{1.2, 2.4},
		ArmorPenetration:       []float32{2.0, 4.0},
	},
	Magical: MagicalDamage{
		MagicalPowerBonus: []float32{1.2, 2.4},
		MagicalPower:      []float32{1, 2},
		MagicalDamage:     []float32{1.2, 2.4},
		TrueMagicalDamage: []float32{0, 0}, // Not specified
		MagicPenetration:  []float32{2.0, 4.0},
	},
	Reduction: Reductions{
		ArmorRating:             []float32{6, 12},
		PhysicalDamageReduction: []float32{1.0, 2.0},
		ProjectileReduction:     []float32{2.0, 4.0},
		MagicalDamageReduction:  []float32{1.0, 2.0},
		MagicResistance:         []float32{6, 12},
		CooldownReduction:       []float32{1.5, 3.0},
	},
	Action: Actions{
		ActionSpeed:             []float32{1.2, 2.4},
		RegularInteractionSpeed: []float32{2.0, 4.0},
		MagicalInteractionSpeed: []float32{2.0, 4.0},
		SpellCastingSpeed:       []float32{2.0, 4.0},
		MoveSpeed:               []float32{0, 0}, // Not specified
		MoveSpeedBonus:          []float32{0, 0}, // Not specified
	},
	Health: Health{
		MaxHealthAdd:    []float32{2, 4},
		MaxHealthBonus:  []float32{1.5, 3.0},
		PhysicalHealing: []float32{1, 2},
		MagicalHealing:  []float32{1, 2},
	},
	Status: Statuses{
		BuffDurationBonus: []float32{1.5, 3.0},
		DebuffDuration:    []float32{1.5, 3.0},
	},
	Memory: Memory{
		MemoryCapacityAdd:   []float32{2, 4},
		MemoryCapacityBonus: []float32{3.0, 6.0},
	},
	Misc: Misc{
		Luck: []float32{30, 60},
	},
}
