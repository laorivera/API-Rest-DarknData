package controllers

import (
	//"fmt"

	//"builder/src/controllers"
	"builder/src/models"
	"fmt"
	"strconv"
)

func Selection_Class(item models.Selection) models.Stats {
	class := item.Class
	var result models.Stats
	if x, exists := models.CharacterStat[class]; exists {
		result = x
	}
	return result
}

func Selection_Race(item models.Selection) models.Stats {
	race := item.Race
	var result models.Stats
	if x, exists := models.RaceStats[race]; exists {
		result = x
	}
	return result
}

func Selection_Brace(item models.Selection) models.Computed_Stats {
	race := item.Race
	var result models.Computed_Stats
	if x, exists := models.RaceComputed[race]; exists {
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

func Speed_Calc(items []models.Item_Armor, rarity []int) int {
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
	armorItems     []*models.Item_Armor
	weaponItems    []*models.Item_Weapon
	accessoryItems []*models.Item_Accessory
}

func NewItemManager() *ItemManager {
	return &ItemManager{
		armorItems:     Items.ItemsArmor,
		weaponItems:    Items.ItemsWeapon,
		accessoryItems: Items.ItemsAccessory,
	}
}

func (im *ItemManager) ArmorsBySlot(slot string, class string) []string {
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

func (im *ItemManager) WeaponsBySlot(slot string, class string) []string {
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
		}
	}

	return result
}

func (im *ItemManager) ArmorsByName(item models.Selection) []*models.Item_Armor {
	slots := []struct {
		name   string
		rarity string
	}{
		{item.ItemSlot.Head.Name, item.ItemSlot.Head.Rarity},
		{item.ItemSlot.Chest.Name, item.ItemSlot.Chest.Rarity},
		{item.ItemSlot.Foot.Name, item.ItemSlot.Foot.Rarity},
		{item.ItemSlot.Hands.Name, item.ItemSlot.Hands.Rarity},
		{item.ItemSlot.Pants.Name, item.ItemSlot.Pants.Rarity},
		{item.ItemSlot.Back.Name, item.ItemSlot.Back.Rarity},
	}

	var result []*models.Item_Armor
	for _, slot := range slots {
		if slot.name != "" && slot.rarity != "" { // ADD THIS CHECK!
			for i := 0; i < len(im.armorItems); i++ {
				if im.armorItems[i].Name == slot.name {
					result = append(result, im.armorItems[i])
					break
				}
			}
		}
	}
	return result
}

func (im *ItemManager) ArmorsByNameD(item models.Selection) []*models.Item_Armor {
	var selection []string
	selection = append(selection, item.ItemSlot.Head.Name, item.ItemSlot.Chest.Name, item.ItemSlot.Foot.Name, item.ItemSlot.Hands.Name, item.ItemSlot.Pants.Name, item.ItemSlot.Back.Name)
	var result []*models.Item_Armor
	for i := 0; i < len(im.armorItems); i++ {
		for j := 0; j < len(selection); j++ {
			if im.armorItems[i].Name == selection[j] && selection[j] != "" {
				result = append(result, im.armorItems[i])
			}
		}
	}

	return result
}

func (im *ItemManager) WeaponsByNameD(item models.Selection) []*models.Item_Weapon {
	var selection []string
	selection = append(selection,
		item.ItemSlot.WeaponOne.Name, item.ItemSlot.WeaponTwo.Name)

	var result []*models.Item_Weapon
	for i := 0; i < len(im.weaponItems); i++ {
		for j := 0; j < len(selection); j++ {
			if im.weaponItems[i].Name == selection[j] && selection[j] != "" {
				result = append(result, im.weaponItems[i])
			}
		}

	}
	return result
}

func (im *ItemManager) AccesoriesByNameD(item models.Selection) []*models.Item_Accessory {
	var selection []string
	selection = append(selection,
		item.ItemSlot.Necklace.Name, item.ItemSlot.RingOne.Name, item.ItemSlot.RingTwo.Name)

	var result []*models.Item_Accessory
	for i := 0; i < len(im.accessoryItems); i++ {
		for j := 0; j < len(selection); j++ {
			if im.accessoryItems[i].Name == selection[j] && selection[j] != "" {
				result = append(result, im.accessoryItems[i])
			}
		}
	}
	return result
}

func (im *ItemManager) WeaponsByName(item models.Selection) []*models.Item_Weapon {
	var selection []string
	selection = append(selection,
		item.ItemSlot.WeaponOne.Name, item.ItemSlot.WeaponTwo.Name)

	var result []*models.Item_Weapon
	for i := 0; i < len(im.weaponItems); i++ {
		for j := 0; j < len(selection); j++ {
			if im.weaponItems[i].Name == selection[j] && selection[j] != "" {
				result = append(result, im.weaponItems[i])
			}
		}

	}
	return result
}

func (im *ItemManager) AccesoriesByName(item models.Selection) []*models.Item_Accessory {
	var selection []string
	selection = append(selection,
		item.ItemSlot.Necklace.Name, item.ItemSlot.RingOne.Name, item.ItemSlot.RingTwo.Name)

	var result []*models.Item_Accessory
	for i := 0; i < len(im.accessoryItems); i++ {
		for j := 0; j < len(selection); j++ {
			if im.accessoryItems[i].Name == selection[j] && selection[j] != "" {
				result = append(result, im.accessoryItems[i])
			}
		}
	}
	return result
}

func (im *ItemManager) SelByRarity(item models.Selection) []int {
	slots := []struct {
		name   string
		rarity string
	}{
		{item.ItemSlot.Head.Name, item.ItemSlot.Head.Rarity},
		{item.ItemSlot.Chest.Name, item.ItemSlot.Chest.Rarity},
		{item.ItemSlot.Foot.Name, item.ItemSlot.Foot.Rarity},
		{item.ItemSlot.Hands.Name, item.ItemSlot.Hands.Rarity},
		{item.ItemSlot.Pants.Name, item.ItemSlot.Pants.Rarity},
		{item.ItemSlot.Back.Name, item.ItemSlot.Back.Rarity},
	}

	var result []int
	for _, slot := range slots {
		if slot.name != "" && slot.rarity != "" {
			if num, err := strconv.Atoi(slot.rarity); err == nil {
				result = append(result, num)
			}
		}
	}
	//fmt.Println(result)
	return result
}

func (im *ItemManager) SelByRarityAcc(item models.Selection) []int {
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

func (im *ItemManager) SelByRating(item models.Selection) []int {
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

func (im *ItemManager) SelByRarityWeapon(item models.Selection) []int {
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

func (im *ItemManager) SelByRatingWeapon(item models.Selection) []int {
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
	Base        models.Stats
	Variable    models.Computed_Stats
	Weapon      models.Computed_Stats_Weapon
	Totalrating int
	Totalspeed  int
}

func NewStatsManager() *StatsManager {
	return &StatsManager{
		Base:     models.CharacterStats,
		Variable: models.Computed_Stats{},
	}
}

func (sm *StatsManager) ItemsBaseStats(selection models.Selection) {
	// CALCULATE TOTAL BASE ATTRIBUTE STATS OF ITEMS=
	Im := NewItemManager()

	var selrace models.Stats = Selection_Race(selection)
	var selclass models.Stats = Selection_Class(selection)
	selclass = selclass.AddStats(selrace)

	//fmt.Printf("=== DEBUG: After race/class selection ===\n")
	//fmt.Printf("selclass: %+v\n", selclass)
	//fmt.Printf("selrace: %+v\n", selrace)

	var selarmor []*models.Item_Armor = Im.ArmorsByName(selection)
	var selrarity []int = Im.SelByRarity(selection)
	var totalacc models.Stats
	var selacc []*models.Item_Accessory = Im.AccesoriesByName(selection)
	var selrarityacc []int = Im.SelByRarityAcc(selection)

	//fmt.Printf("=== DEBUG: Items and rarities ===\n")
	//fmt.Printf("selarmor length: %d\n", len(selarmor))
	//fmt.Printf("selrarity: %v\n", selrarity)
	/*
		for i, item := range selarmor {
			fmt.Printf("Item %d: %s, Rarity: %d\n", i, item.Name, selrarity[i])
		}*/

	for i := 0; i < len(selarmor) && i < len(selrarity); i++ {
		rarity := selrarity[i]
		BaseAttr := selarmor[i].BaseAttribute

		//fmt.Printf("\n=== DEBUG: Processing %s (rarity %d) ===\n", selarmor[i].Name, rarity)
		//fmt.Printf("BaseAttribute: %+v\n", BaseAttr)

		// Check and add each attribute only if it exists for this rarity
		if val, exists := BaseAttr.Strength[rarity]; exists {
			//fmt.Printf("Adding Strength: %d\n", val)
			selclass.Strength += val
		} else {
			//fmt.Printf("No Strength at rarity %d\n", rarity)
		}
		if val, exists := BaseAttr.Vigor[rarity]; exists {
			//fmt.Printf("Adding Vigor: %d\n", val)
			selclass.Vigor += val
		} else {
			//fmt.Printf("No Vigor at rarity %d\n", rarity)
		}
		if val, exists := BaseAttr.Agility[rarity]; exists {
			//fmt.Printf("Adding Agility: %d\n", val)
			selclass.Agility += val
		} else {
			//fmt.Printf("No Agility at rarity %d\n", rarity)
		}
		if val, exists := BaseAttr.Dexterity[rarity]; exists {
			//fmt.Printf("Adding Dexterity: %d\n", val)
			selclass.Dexterity += val
		} else {
			//fmt.Printf("No Dexterity at rarity %d\n", rarity)
		}
		if val, exists := BaseAttr.Will[rarity]; exists {
			//fmt.Printf("Adding Will: %d\n", val)
			selclass.Will += val
		} else {
			//fmt.Printf("No Will at rarity %d\n", rarity)
		}
		if val, exists := BaseAttr.Knowledge[rarity]; exists {
			//fmt.Printf("Adding Knowledge: %d\n", val)
			selclass.Knowledge += val
		} else {
			//fmt.Printf("No Knowledge at rarity %d\n", rarity)
		}
		if val, exists := BaseAttr.Resourcefulness[rarity]; exists {
			//fmt.Printf("Adding Resourcefulness: %d\n", val)
			selclass.Resourcefulness += val
		} else {
			//fmt.Printf("No Resourcefulness at rarity %d\n", rarity)
		}

		//fmt.Printf("Current selclass after %s: %+v\n", selarmor[i].Name, selclass)
	}

	//fmt.Printf("\n=== DEBUG: After armor processing ===\n")
	//fmt.Printf("selclass: %+v\n", selclass)

	for i := 0; i < len(selacc) && i < len(selrarityacc); i++ {
		rarity := selrarityacc[i]
		BaseAttr := selacc[i].BaseAttribute

		//fmt.Printf("\n=== DEBUG: Processing accessory %s (rarity %d) ===\n", selacc[i].Name, rarity)

		// Check and add each attribute only if it exists for this rarity
		if val, exists := BaseAttr.Strength[rarity]; exists {
			//fmt.Printf("Adding Strength: %d\n", val)
			totalacc.Strength += val
		}
		if val, exists := BaseAttr.Vigor[rarity]; exists {
			//fmt.Printf("Adding Vigor: %d\n", val)
			totalacc.Vigor += val
		}
		if val, exists := BaseAttr.Agility[rarity]; exists {
			//fmt.Printf("Adding Agility: %d\n", val)
			totalacc.Agility += val
		}
		if val, exists := BaseAttr.Dexterity[rarity]; exists {
			//fmt.Printf("Adding Dexterity: %d\n", val)
			totalacc.Dexterity += val
		}
		if val, exists := BaseAttr.Will[rarity]; exists {
			//fmt.Printf("Adding Will: %d\n", val)
			totalacc.Will += val
		}
		if val, exists := BaseAttr.Knowledge[rarity]; exists {
			//fmt.Printf("Adding Knowledge: %d\n", val)
			totalacc.Knowledge += val
		}
		if val, exists := BaseAttr.Resourcefulness[rarity]; exists {
			//fmt.Printf("Adding Resourcefulness: %d\n", val)
			totalacc.Resourcefulness += val
		}
	}

	//fmt.Printf("\n=== DEBUG: Final totals before adding ===\n")
	//fmt.Printf("selclass: %+v\n", selclass)
	//fmt.Printf("totalacc: %+v\n", totalacc)

	sm.Base = selclass.AddStats(totalacc)

	//fmt.Printf("\n=== DEBUG: Final result ===\n")
	//fmt.Printf("sm.Base: %+v\n", sm.Base)
}

func (sm *StatsManager) BaseItemCalc(selection models.Selection) { //PENDING!

	Im := NewItemManager()

	var computed models.Computed_Stats = Selection_Brace(selection) // race selection passed into computed
	var selarmor []*models.Item_Armor = Im.ArmorsByName(selection)
	var selrarity []int = Im.SelByRarity(selection)
	var totalacc models.Computed_Stats
	var selacc []*models.Item_Accessory = Im.AccesoriesByName(selection)
	var selrarityacc []int = Im.SelByRarityAcc(selection)

	///////////////////////////PENDING TO REFACTOR /.////////////////////////////////
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
		totalacc.Luck += selacc[i].Luck[selrarityacc[i]]
	}

	sm.Variable = computed.AddEnchant()
	fmt.Println(sm.Variable)

	/////////////////////////////////////////////////////////////////////////////////////
}

func (sm *StatsManager) TotalSpeedRating(selection models.Selection) {
	Im := NewItemManager()
	var selarmor []*models.Item_Armor = Im.ArmorsByName(selection)
	var selrarity []int = Im.SelByRarity(selection)
	sm.Totalspeed = SpeedCalc(selarmor, selrarity)
}

func (sm *StatsManager) TotalArmorRating(selection models.Selection) {
	Im := NewItemManager()
	var selrating = Im.SelByRating(selection)
	sm.Totalrating = RatingCalc(selrating)
}

func (sm *StatsManager) WeaponDamageCalcx(selection models.Selection, statsother models.Computed_Stats) {

	Im := NewItemManager()
	selweapons := Im.WeaponsByName(selection)
	selrarity := Im.SelByRarityWeapon(selection)
	selrating := Im.SelByRatingWeapon(selection)

	var result models.Computed_Stats_Weapon
	for i := 0; i < len(selweapons) && i < len(selrarity); i++ {
		if selweapons[i].SlotType == "Main Hand" {
			result.PrimaryWeapon.Attackone += ((float64(selrating[i]) * (statsother.PhysicalPowerBonus / 100)) + float64(selrating[i])) * (float64((selweapons[i].ComboDamage[0])) / 100) // adjust % of power bonus
			result.PrimaryWeapon.Attacktwo += ((float64(selrating[i]) * (statsother.PhysicalPowerBonus / 100)) + float64(selrating[i])) * (float64((selweapons[i].ComboDamage[1])) / 100)
			result.PrimaryWeapon.Attackthree += ((float64(selrating[i]) * (statsother.PhysicalPowerBonus / 100)) + float64(selrating[i])) * (float64((selweapons[i].ComboDamage[2])) / 100)
		}
		if selweapons[i].SlotType == "Main Hand" && len(selweapons[i].ComboDamage) >= 4 {
			result.PrimaryWeapon.Attackfour += ((float64(selrating[i]) * (statsother.PhysicalPowerBonus / 100)) + float64(selrating[i])) * (float64((selweapons[i].ComboDamage[3])) / 100)
		}

		if selweapons[i].SlotType == "Off Hand" && len(selweapons[i].ComboDamage) <= 3 {
			result.SecondaryWeapon.Attackone += ((float64(selrating[i]) * (statsother.PhysicalPowerBonus / 100)) + float64(selrating[i])) * (float64((selweapons[i].ComboDamage[0])) / 100)
			result.SecondaryWeapon.Attacktwo += ((float64(selrating[i]) * (statsother.PhysicalPowerBonus / 100)) + float64(selrating[i])) * (float64((selweapons[i].ComboDamage[1])) / 100)
			result.SecondaryWeapon.Attackthree += ((float64(selrating[i]) * (statsother.PhysicalPowerBonus / 100)) + float64(selrating[i])) * (float64((selweapons[i].ComboDamage[2])) / 100)
		}
		if selweapons[i].SlotType == "Off Hand" && len(selweapons[i].ComboDamage) >= 4 {
			result.SecondaryWeapon.Attackfour += ((float64(selrating[i]) * (statsother.PhysicalPowerBonus / 100)) + float64(selrating[i])) * (float64((selweapons[i].ComboDamage[3])) / 100)
		}
	}

	sm.Weapon = result
}
