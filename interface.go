package main

import (
	"strconv"
)

func Selection_Class(item Selection) Stats {
	class := item.Class
	var result Stats
	if x, exists := characterStat[class]; exists {
		result = x
	}
	return result
}

func Selection_Race(item Selection) Stats {
	race := item.Race
	var result Stats
	if x, exists := raceStats[race]; exists {
		result = x
	}
	return result
}

func Selection_Brace(item Selection) Computed_Stats {
	race := item.Race
	var result Computed_Stats
	if x, exists := raceComputed[race]; exists {
		result = x
	}
	return result
}

func Rating_Calc(ratings []int) int {
	result := 0
	for i := 0; i < len(ratings); i++ {
		result = result + ratings[i]
	}
	return result
}

func Speed_Calc(items []Item_Armor, rarity []int) int {
	speedrating := 0
	for i := 0; i < len(items); i++ {

		if items[i].SlotType != "Foot" {
			speedrating += items[i].MoveSpeed[1]
		}
		if items[i].SlotType == "Foot" {
			speedrating += items[i].MoveSpeed[rarity[i]]
		}
	}
	return speedrating
}

// //////////////////////////////////////////////ItemManager////////////////////////////////////////////////////////////////
type ItemManager struct {
	armorItems     []Item_Armor
	weaponItems    []Item_Weapon
	accessoryItems []Item_Accessory
}

func NewItemManager() *ItemManager {
	return &ItemManager{
		armorItems:     Items.ItemsArmor,
		weaponItems:    Items.ItemsWeapon,
		accessoryItems: Items.ItemsAccessory,
	}
}

func (im *ItemManager) ArmorsBySlot(slot string, classID string) []string {
	class := classID
	var result []string
	for i := 0; i < len(im.armorItems); i++ {
		for _, c := range im.armorItems[i].Classes {
			if c == class && slot == im.armorItems[i].SlotType {
				result = append(result, im.armorItems[i].Name)
				break
			}
		}

	}
	return result
}

func (im *ItemManager) WeaponsBySlot(slot string, classID string) []string {
	class := classID
	var result []string
	for i := 0; i < len(im.weaponItems); i++ {
		for _, c := range im.weaponItems[i].Classes {
			if c == class && slot == im.weaponItems[i].SlotType {
				result = append(result, im.weaponItems[i].Name)
				break
			}
		}

	}
	return result
}

func (im *ItemManager) AccesoryBySlot(slot string) []string {
	var result []string
	for i := 0; i < len(im.accessoryItems); i++ {
		if slot == im.accessoryItems[i].SlotType {
			result = append(result, im.accessoryItems[i].Name)
			break
		}
	}

	return result
}

func (im *ItemManager) ArmorsByName(item Selection) []Item_Armor { //select items from json
	var selection []string
	selection = append(selection, item.ItemSlot.Head.Name, item.ItemSlot.Chest.Name, item.ItemSlot.Foot.Name, item.ItemSlot.Hands.Name, item.ItemSlot.Pants.Name, item.ItemSlot.Back.Name)
	var result []Item_Armor
	for i := 0; i < len(im.armorItems); i++ {
		for j := 0; j < len(selection); j++ {
			if im.armorItems[i].Name == selection[j] && selection[j] != "" {
				result = append(result, im.armorItems[i])
			}
		}
	}

	return result
}

func (im *ItemManager) WeaponsByName(item Selection) []Item_Weapon {
	var selection []string
	selection = append(selection,
		item.ItemSlot.WeaponOne.Name, item.ItemSlot.WeaponTwo.Name)

	var result []Item_Weapon
	for i := 0; i < len(im.weaponItems); i++ {
		for j := 0; j < len(selection); j++ {
			if im.weaponItems[i].Name == selection[j] && selection[j] != "" {
				result = append(result, im.weaponItems[i])
			}
		}

	}
	return result
}

func (im *ItemManager) AccesoriesByName(item Selection) []Item_Accessory {
	var selection []string
	selection = append(selection,
		item.ItemSlot.Necklace.Name, item.ItemSlot.RingOne.Name, item.ItemSlot.RingTwo.Name)

	var result []Item_Accessory
	for i := 0; i < len(im.accessoryItems); i++ {
		for j := 0; j < len(selection); j++ {
			if im.accessoryItems[i].Name == selection[j] && selection[j] != "" {
				result = append(result, im.accessoryItems[i])
			}
		}
	}
	return result
}

func (im *ItemManager) SelByRarity(item Selection) []int {
	var selection []string
	selection = append(selection,
		item.ItemSlot.Head.Rarity, item.ItemSlot.Chest.Rarity, item.ItemSlot.Foot.Rarity,
		item.ItemSlot.Hands.Rarity, item.ItemSlot.Pants.Rarity, item.ItemSlot.Back.Rarity)
	var result []int
	for i := 0; i < len(selection); i++ {
		if num, err := strconv.Atoi(selection[i]); err == nil {
			result = append(result, num)
		}
	}

	return result
}

func (im *ItemManager) SelByRarityAcc(item Selection) []int {
	var selection []string
	selection = append(selection,
		item.ItemSlot.Necklace.Rarity, item.ItemSlot.RingOne.Rarity, item.ItemSlot.RingTwo.Rarity)
	var result []int
	for i := 0; i < len(selection); i++ {
		if num, err := strconv.Atoi(selection[i]); err == nil {
			result = append(result, num)
		}
	}

	return result
}

func (im *ItemManager) SelByRating(item Selection) []int {
	var selection []string
	selection = append(selection,
		item.ItemSlot.Head.Rating, item.ItemSlot.Chest.Rating, item.ItemSlot.Foot.Rating,
		item.ItemSlot.Hands.Rating, item.ItemSlot.Pants.Rating, item.ItemSlot.Back.Rating)
	var result []int
	for i := 0; i < len(selection); i++ {
		if num, err := strconv.Atoi(selection[i]); err == nil {
			result = append(result, num)
		}
	}
	return result
}

func (im *ItemManager) SelByRarityWeapon(item Selection) []int {
	var selection []string
	selection = append(selection,
		item.ItemSlot.WeaponOne.Rating, item.ItemSlot.WeaponTwo.Rating)
	var result []int
	for i := 0; i < len(selection); i++ {
		if num, err := strconv.Atoi(selection[i]); err == nil {
			result = append(result, num)
		}
	}
	return result
}

func (im *ItemManager) SelByRatingWeapon(item Selection) []int {
	var selection []string
	selection = append(selection,
		item.ItemSlot.WeaponOne.Rating, item.ItemSlot.WeaponTwo.Rating)
	var result []int
	for i := 0; i < len(selection); i++ {
		if num, err := strconv.Atoi(selection[i]); err == nil {
			result = append(result, num)
		}
	}
	return result
}

// //////////////////////////////////////////////StatsManager////////////////////////////////////////////////////////////////
type StatsManager struct {
	base        Stats
	variable    Computed_Stats
	weapon      Computed_Stats_Weapon
	totalrating int
	totalspeed  int
}

func NewStatsManager() *StatsManager {
	return &StatsManager{
		base:     characterStats,
		variable: Computed_Stats{},
	}
}

func (sm *StatsManager) ItemsBaseStats(selection Selection) {
	// CALCULATE TOTAL BASE ATTRIBUTE STATS OF ITEMS=
	Im := NewItemManager()

	var selrace Stats = Selection_Race(selection)
	var selclass Stats = Selection_Class(selection)
	selclass = selclass.AddStats(selrace)

	var selarmor []Item_Armor = Im.ArmorsByName(selection)
	var selrarity []int = Im.SelByRarity(selection)
	var totalacc Stats
	var selacc []Item_Accessory = Im.AccesoriesByName(selection)
	var selrarityacc []int = Im.SelByRarityAcc(selection)

	for i := 0; i < len(selarmor) && i < len(selrarity); i++ {
		selclass.Strength += selarmor[i].BaseAttribute.Strength[selrarity[i]]
		selclass.Vigor += selarmor[i].BaseAttribute.Vigor[selrarity[i]]
		selclass.Agility += selarmor[i].BaseAttribute.Agility[selrarity[i]]
		selclass.Dexterity += selarmor[i].BaseAttribute.Dexterity[selrarity[i]]
		selclass.Will += selarmor[i].BaseAttribute.Will[selrarity[i]]
		selclass.Knowledge += selarmor[i].BaseAttribute.Knowledge[selrarity[i]]
		selclass.Resourcefulness += selarmor[i].BaseAttribute.Resourcefulness[selrarity[i]]
	}

	for i := 0; i < len(selacc) && i < len(selrarityacc); i++ {
		totalacc.Strength += selacc[i].BaseAttribute.Strength[selrarityacc[i]]
		totalacc.Vigor += selacc[i].BaseAttribute.Vigor[selrarityacc[i]]
		totalacc.Agility += selacc[i].BaseAttribute.Agility[selrarityacc[i]]
		totalacc.Dexterity += selacc[i].BaseAttribute.Dexterity[selrarityacc[i]]
		totalacc.Will += selacc[i].BaseAttribute.Will[selrarityacc[i]]
		totalacc.Knowledge += selacc[i].BaseAttribute.Knowledge[selrarityacc[i]]
		totalacc.Resourcefulness += selacc[i].BaseAttribute.Resourcefulness[selrarityacc[i]]
	}

	sm.base = selclass.AddStats(totalacc)

}

func (sm *StatsManager) BaseItemCalc(selection Selection) {

	Im := NewItemManager()

	var computed Computed_Stats = Selection_Brace(selection) // race selection passed into computed
	var selarmor []Item_Armor = Im.ArmorsByName(selection)
	var selrarity []int = Im.SelByRarity(selection)
	var totalacc Computed_Stats
	var selacc []Item_Accessory = Im.AccesoriesByName(selection)
	var selrarityacc []int = Im.SelByRarityAcc(selection)

	for i := 0; i < len(selarmor) && i < len(selrarity); i++ {
		computed.Health += selarmor[i].MaxHealthAdd[selrarity[i]]
		computed.ProjectileReduction += selarmor[i].ProjectileReduction
		computed.ProjectileReduction += selarmor[i].ProjectileReductionRate[selrarity[i]]
		computed.HeadshotReduction += selarmor[i].HeadshotReduction
		computed.MoveSpeedBonus += selarmor[i].MoveSpeedBonus
		computed.ArmorPenetration += selarmor[i].ArmorPenetration
		computed.MagicPenetration += selarmor[i].MagicPenetration
		computed.ActionSpeed += selarmor[i].ActionSpeed[0]
		computed.MagicalDamageReduction += selarmor[i].MagicalDamageReduction
		computed.PhysicalDamageReduction += selarmor[i].PhysicalDamageReduction
		computed.Luck += selarmor[i].Luck
		computed.MagicalHealing += selarmor[i].MagicalHealing[selrarity[i]]
		computed.MagicalPower += selarmor[i].MagicalPower[selrarity[i]]
		computed.PhysicalPower += selarmor[i].PhysicalPower
		if len(selarmor[i].MagicResistance[selrarity[i]]) > 0 {
			computed.MagicRating += selarmor[i].MagicResistance[selrarity[i]][0]
		}
		if len(selarmor[i].MaxHealthBonus[selrarity[i]]) > 0 {
			computed.MaxHealthBonus += selarmor[i].MaxHealthBonus[selrarity[i]][i]
		}

	}

	for i := 0; i < len(selacc) && i < len(selrarityacc); i++ {
		totalacc.Health += selacc[i].MaxHealthAdd[selrarityacc[i]]
		totalacc.MoveSpeedBonus += selacc[i].MoveSpeedBonus
		totalacc.Luck += selarmor[i].Luck
	}

	sm.variable = computed.AddEnchant()
}

func (sm *StatsManager) TotalSpeedRating(selection Selection) {
	Im := NewItemManager()
	var selarmor []Item_Armor = Im.ArmorsByName(selection)
	var selrarity []int = Im.SelByRarity(selection)
	sm.totalspeed = SpeedCalc(selarmor, selrarity)
}

func (sm *StatsManager) TotalArmorRating(selection Selection) {
	Im := NewItemManager()
	var selrating = Im.SelByRating(selection)
	sm.totalrating = RatingCalc(selrating)
}

func (sm *StatsManager) WeaponDamageCalcx(selection Selection) {

	Im := NewItemManager()
	selweapons := Im.WeaponsByName(selection)
	selrarity := Im.SelByRarityWeapon(selection)
	selrating := Im.SelByRatingWeapon(selection)

	var result Computed_Stats_Weapon
	for i := 0; i < len(selweapons) && i < len(selrarity); i++ {
		if selweapons[i].SlotType == "Main Hand" {
			result.PrimaryWeapon.Attackone += ((float64(selrating[i]) * (sm.variable.PhysicalPowerBonus / 100)) + float64(selrating[i])) * (float64((selweapons[i].ComboDamage[0])) / 100) // adjust % of power bonus
			result.PrimaryWeapon.Attacktwo += ((float64(selrating[i]) * (sm.variable.PhysicalPowerBonus / 100)) + float64(selrating[i])) * (float64((selweapons[i].ComboDamage[1])) / 100)
			result.PrimaryWeapon.Attackthree += ((float64(selrating[i]) * (sm.variable.PhysicalPowerBonus / 100)) + float64(selrating[i])) * (float64((selweapons[i].ComboDamage[2])) / 100)
		}
		if selweapons[i].SlotType == "Main Hand" && len(selweapons[i].ComboDamage) >= 4 {
			result.PrimaryWeapon.Attackfour += ((float64(selrating[i]) * (sm.variable.PhysicalPowerBonus / 100)) + float64(selrating[i])) * (float64((selweapons[i].ComboDamage[3])) / 100)
		}

		if selweapons[i].SlotType == "Off Hand" && len(selweapons[i].ComboDamage) <= 3 {
			result.SecondaryWeapon.Attackone += ((float64(selrating[i]) * (sm.variable.PhysicalPowerBonus / 100)) + float64(selrating[i])) * (float64((selweapons[i].ComboDamage[0])) / 100)
			result.SecondaryWeapon.Attacktwo += ((float64(selrating[i]) * (sm.variable.PhysicalPowerBonus / 100)) + float64(selrating[i])) * (float64((selweapons[i].ComboDamage[1])) / 100)
			result.SecondaryWeapon.Attackthree += ((float64(selrating[i]) * (sm.variable.PhysicalPowerBonus / 100)) + float64(selrating[i])) * (float64((selweapons[i].ComboDamage[2])) / 100)
		}
		if selweapons[i].SlotType == "Off Hand" && len(selweapons[i].ComboDamage) >= 4 {
			result.SecondaryWeapon.Attackfour += ((float64(selrating[i]) * (sm.variable.PhysicalPowerBonus / 100)) + float64(selrating[i])) * (float64((selweapons[i].ComboDamage[3])) / 100)
		}
	}

	sm.weapon = result
}
