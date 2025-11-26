package main

type CharSkill struct {
	Fighter PassiveSkill
}

type PassiveSkill struct {
	AdrenalineSpike     []Computed_Stats
	Barricade           []Computed_Stats
	ComboAttack         []Computed_Stats
	CounterAttack       []Computed_Stats
	DefenseMastery      []Computed_Stats
	DualWield           []Computed_Stats
	ProjectileReduction []Computed_Stats
	ShieldMastery       []Computed_Stats
	Slayer              []Computed_Stats
	Swift               []Computed_Stats
	SwordMastery        []Computed_Stats
	WeaponMastery       bool
}
