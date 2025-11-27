package models

// STRUCT STATS
type Stats struct {
	Strength        int
	Vigor           int
	Agility         int
	Dexterity       int
	Will            int
	Knowledge       int
	Resourcefulness int
}

// STRUCT COMPLETE STATS
type Computed_Stats struct {
	Health                       float64
	MaxHealthBonus               float64
	ActionSpeed                  float64
	RegularInteractionSpeed      float64
	MoveSpeed                    float64
	MoveSpeedBonus               float64
	PhysicalPower                float64
	PhysicalPowerBonus           float64
	HealthRecovery               float64
	ManualDexterity              float64
	EquipSpeed                   float64
	MagicalPower                 float64
	MagicalPowerBonus            float64
	BuffDuration                 float64
	MagicRating                  float64
	MagicalDamageReduction       float64
	DebuffDuration               float64
	MemoryCapacity               float64
	MemoryCapacityBonus          float64
	SpellRecovery                float64
	SpellCastingSpeed            float64
	MagicalInteractionSpeed      float64
	Persuasiveness               float64
	CooldownReduction            float64
	PhysicalDamageReduction      float64
	SpellRecoveryBonus           float64
	PhysicalHealing              int
	MagicalHealing               int
	MemorySpellPayload           int
	MemoryMusicPayload           int
	UtilityEffectiveness         int
	Luck                         int
	ArmorPenetration             float64
	MagicPenetration             float64
	HeadshotReduction            float64
	ProjectileReduction          float64
	FromArmorRating              int
	BonusPhysicalDamageReduction float64
	BonusMagicalDamageReduction  float64
	BonusPhysicalPower           float64
	BonusMagicalPower            float64
	TruePhysicalDamage           int
	TrueMagicalDamage            int
}

type Computed_Stats_Weapon struct {
	PrimaryWeapon        WeaponStats
	SecondaryWeapon      WeaponStats
	PrimaryImpactPower   int
	SecondaryImpactPower int
}

type WeaponStats struct {
	Attackone   float64
	Attacktwo   float64
	Attackthree float64
	Attackfour  float64
}

type BaseAttribute struct {
	Strength        map[int]int `json:"strength"`
	Vigor           map[int]int `json:"vigor"`
	Agility         map[int]int `json:"agility"`
	Dexterity       map[int]int `json:"dexterity"`
	Will            map[int]int `json:"will"`
	Knowledge       map[int]int `json:"knowledge"`
	Resourcefulness map[int]int `json:"resourcefulness"`
}

type Item_Armor struct {
	Name                    string            `json:"name"`
	ArmorRatings            map[int][]int     `json:"armorRatings"`
	MaxHealthAdd            map[int]float64   `json:"maxHealthAdd"`
	MaxHealthBonus          map[int][]float64 `json:"maxHealthBonus"` // exception
	MagicResistance         map[int][]float64 `json:"magicalResistance"`
	ProjectileReduction     float64           `json:"projectileReduction"` //
	ProjectileReductionRate map[int]float64   `json:"projectileReductionRate"`
	HeadshotReduction       float64           `json:"headshotReduction"` //
	BaseAttribute           BaseAttribute     `json:"BaseAttribute"`     //
	PhysicalPower           float64           `json:"physicalPower"`
	TruePhysicalDamage      map[int]int       `json:"truePhysicalDamage"`
	TrueMagicalDamage       map[int]int       `json:"trueMagicalDamage"`
	MoveSpeed               map[int]int       `json:"moveSpeed"` //
	MoveSpeedBonus          float64           `json:"moveSpeedbonus"`
	ArmorPenetration        float64           `json:"armorpPenetration"`
	MagicPenetration        float64           `json:"magicPenetration"`
	BuffDuration            float64           `json:"buffDuration"`
	MagicalPower            map[int]float64   `json:"magicalPower"`
	MagicalPowerBonus       map[int]float64   `json:"magicalPowerBonus"`
	RegularInteractionSpeed map[int]float64   `json:"regularInteraction"`
	MagicalHealing          map[int]int       `json:"magicalHealing"`
	ActionSpeed             map[int]float64   `json:"actionSpeed"`
	MagicalDamageReduction  float64           `json:"magicalDamageReduction"`
	PhysicalDamageReduction float64           `json:"physicalDamageReduction"`
	Luck                    int               `json:"luck"`
	ArmorType               string            `json:"armorType"`   //
	Rarities                []int             `json:"rarities"`    //
	SlotType                string            `json:"slotType"`    //
	InvWidth                int               `json:"invWidth"`    //
	InvHeight               int               `json:"invHeight"`   //
	FlavorText              string            `json:"flavorText"`  //
	AP                      map[int]int       `json:"ap"`          //
	XP                      map[int]int       `json:"xp"`          //
	GearScore               map[int]int       `json:"gearScore"`   //
	Classes                 []string          `json:"classes"`     //
	NumEnchants             map[int]int       `json:"numEnchants"` //
}

type Item_Weapon struct {
	Name                    string            `json:"name"`
	Classes                 []string          `json:"classes"` // List of classes (e.g., "Fighter", "Barbarian")
	DamageRatings           map[int][]int     `json:"damageRatings"`
	ArmorRatings            map[int][]int     `json:"armorRatings"`
	MoveSpeed               map[int]int       `json:"moveSpeed"`
	MoveSpeedBonus          float64         `json:"moveSpeedbonus"`
	TruePhysicalDamage      map[int]int       `json:"truePhysicalDamage"`
	TrueMagicalDamage       map[int]int       `json:"trueMagicalDamage"`
	DamageTypes             []string          `json:"damageTypes"`
	ImpactZones             []int             `json:"impactZones"`
	ImpactPower             int               `json:"impactPower"`
	ComboDamage             []int             `json:"comboDamage"`
	WeaponType              []string          `json:"type"`        // Weapon type (e.g., "Axe")
	Rarities                []int             `json:"rarities"`    // List of rarities
	HandType                string            `json:"handType"`    // Hand type (e.g., "Two Handed")
	SlotType                string            `json:"slotType"`    // Slot type (e.g., "Main Hand")
	GearScore               map[int]int       `json:"gearScore"`   // Gear score for levels 0 to 7
	InvWidth                int               `json:"invWidth"`    // Inventory width
	InvHeight               int               `json:"invHeight"`   // Inventory height
	FlavorText              string            `json:"flavorText"`  // Description text
	AP                      map[int]int       `json:"ap"`          // Action points for levels 3 to 7
	XP                      map[int]int       `json:"xp"`          // XP for levels 1 to 7
	NumEnchants             map[int]int       `json:"numEnchants"` // Number of enchants for levels 3 to 7
	MaxHealthAdd            map[int]float64   `json:"maxHealthAdd"`
	MaxHealthBonus          map[int][]float64 `json:"maxHealthBonus"`
	MagicResistance         map[int][]float64 `json:"magicalResistance"`
	ProjectileReduction     float64           `json:"projectileReduction"`
	ProjectileReductionRate map[int]float64   `json:"projectileReductionRate"`
	HeadshotReduction       float64           `json:"headshotReduction"`
	BaseAttribute           BaseAttribute     `json:"BaseAttribute"`
	PhysicalPower           float64           `json:"physicalPower"`
	ArmorPenetration        float64           `json:"armorpPenetration"`
	MagicPenetration       float64           `json:"magicPenetration"`
	BuffDuration            float64           `json:"buffDuration"`
	MagicalPower            map[int]float64   `json:"magicalPower"`
	MagicalPowerBonus       map[int]float64   `json:"magicalPowerBonus"`
	RegularInteractionSpeed map[int]float64   `json:"regularInteraction"`
	MagicalHealing          map[int]int       `json:"magicalHealing"`
	ActionSpeed             map[int]float64   `json:"actionSpeed"`
	MagicalDamageReduction  float64           `json:"magicalDamageReduction"`
	PhysicalDamageReduction float64           `json:"physicalDamageReduction"`
	Luck                    int               `json:"luck"`
}

type Item_Accessory struct {
	Name               string          `json:"name"`
	Classes            []string        `json:"classes"` // List of classes (e.g., "Fighter", "Barbarian")
	MaxHealthAdd       map[int]float64 `json:"maxHealthAdd"`
	MoveSpeed          map[int]int     `json:"moveSpeed"` //
	MoveSpeedBonus     float64         `json:"moveSpeedbonus"`
	BaseAttribute      BaseAttribute   `json:"BaseAttribute"` // BaseStats 1 to 7
	TruePhysicalDamage map[int]int     `json:"truePhysicalDamage"`
	TrueMagicalDamage  map[int]int     `json:"trueMagicalDamage"`
	Luck               map[int]int     `json:"luck"`
	Type               string          `json:"type"`         // Accessory type
	Rarities           []int           `json:"rarities"`     // List of rarities
	SlotType           string          `json:"slotType"`     // Slot type (e.g., "Necklace")
	InvWidth           int             `json:"invWidth"`     // Inventory width
	InvHeight          int             `json:"invHeight"`    // Inventory height
	FlavorText         string          `json:"flavorText"`   // Description text
	MaxAmmoCount       int             `json:"maxAmmoCount"` // Max ammo count
	MaxCount           int             `json:"maxCount"`     // Max count
	AP                 map[int]int     `json:"ap"`           // Action points for levels 3 to 7
	XP                 map[int]int     `json:"xp"`           // XP for levels 1 to 7
	GearScore          map[int]int     `json:"gearScore"`    // Gear score for levels 0 to 7
	NumEnchants        map[int]int     `json:"numEnchants"`  // Number of enchants for levels 3 to 7

}

// AddStats adds multiple Stats together
func (s Stats) AddStats(others ...Stats) Stats {
	result := s
	for _, other := range others {
		result.Strength += other.Strength
		result.Vigor += other.Vigor
		result.Agility += other.Agility
		result.Dexterity += other.Dexterity
		result.Will += other.Will
		result.Knowledge += other.Knowledge
		result.Resourcefulness += other.Resourcefulness
	}
	return result
}

// AddEnchant adds multiple Computed_Stats together
func (cs Computed_Stats) AddEnchant(others ...Computed_Stats) Computed_Stats {
	result := cs
	for _, other := range others {
		result.Health += other.Health
		result.MaxHealthBonus += other.MaxHealthBonus
		result.ActionSpeed += other.ActionSpeed
		result.RegularInteractionSpeed += other.RegularInteractionSpeed
		result.MoveSpeed += other.MoveSpeed
		result.MoveSpeedBonus += other.MoveSpeedBonus
		result.PhysicalPower += other.PhysicalPower
		result.PhysicalPowerBonus += other.PhysicalPowerBonus
		result.HealthRecovery += other.HealthRecovery
		result.ManualDexterity += other.ManualDexterity
		result.EquipSpeed += other.EquipSpeed
		result.MagicalPower += other.MagicalPower
		result.MagicalPowerBonus += other.MagicalPowerBonus
		result.BuffDuration += other.BuffDuration
		result.MagicRating += other.MagicRating
		result.MagicalDamageReduction += other.MagicalDamageReduction
		result.DebuffDuration += other.DebuffDuration
		result.MemoryCapacity += other.MemoryCapacity
		result.MemoryCapacityBonus += other.MemoryCapacityBonus
		result.SpellRecovery += other.SpellRecovery
		result.SpellCastingSpeed += other.SpellCastingSpeed
		result.MagicalInteractionSpeed += other.MagicalInteractionSpeed
		result.Persuasiveness += other.Persuasiveness
		result.CooldownReduction += other.CooldownReduction
		result.PhysicalDamageReduction += other.PhysicalDamageReduction
		result.SpellRecoveryBonus += other.SpellRecoveryBonus
		result.PhysicalHealing += other.PhysicalHealing
		result.MagicalHealing += other.MagicalHealing
		result.MemorySpellPayload += other.MemorySpellPayload
		result.MemoryMusicPayload += other.MemoryMusicPayload
		result.UtilityEffectiveness += other.UtilityEffectiveness
		result.Luck += other.Luck
		result.ArmorPenetration += other.ArmorPenetration
		result.MagicPenetration += other.MagicPenetration
		result.HeadshotReduction += other.HeadshotReduction
		result.ProjectileReduction += other.ProjectileReduction
		result.FromArmorRating += other.FromArmorRating
		result.BonusPhysicalDamageReduction += other.BonusPhysicalDamageReduction
		result.BonusMagicalDamageReduction += other.BonusMagicalDamageReduction
	}
	return result
}
