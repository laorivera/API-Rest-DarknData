package main

import (
	"strconv"
)

func Selection_Class(item Selection) Stats {
	var result Stats
	class := item.Class
	switch class {
	case "Fighter":
		result = SCharacter.Fighter
	case "Barbarian":
		result = SCharacter.Barbarian
	case "Rogue":
		result = SCharacter.Rogue
	case "Wizard":
		result = SCharacter.Wizard
	case "Cleric":
		result = SCharacter.Cleric
	case "Warlock":
		result = SCharacter.Warlock
	case "Bard":
		result = SCharacter.Bard
	case "Ranger":
		result = SCharacter.Ranger
	case "Sorcerer":
		result = SCharacter.Sorcerer
	default:
		result = Stats{}
	}
	return result
}

// ItemManager
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

func (im *ItemManager) ArmorsBySlot(slot string, classID string) []Item_Armor {
	class := InttoClass(classID)
	var result []Item_Armor
	for i := 0; i < len(im.armorItems); i++ {
		for _, c := range im.armorItems[i].Classes {
			if c == class && slot == im.armorItems[i].SlotType {
				result = append(result, im.armorItems[i])
				break
			}
		}

	}
	return result
}

func (im *ItemManager) WeaponsBySlot(slot string, classID string) []Item_Weapon {
	class := InttoClass(classID)
	var result []Item_Weapon
	for i := 0; i < len(im.weaponItems); i++ {
		for _, c := range im.weaponItems[i].Classes {
			if c == class && slot == im.weaponItems[i].SlotType {
				result = append(result, im.weaponItems[i])
				break
			}
		}

	}
	return result
}

func (im *ItemManager) AccesoryBySlot(slot string, classID string) []Item_Accessory {
	class := InttoClass(classID)
	var result []Item_Accessory
	for i := 0; i < len(im.accessoryItems); i++ {
		for _, c := range im.accessoryItems[i].Classes {
			if c == class && slot == im.accessoryItems[i].SlotType {
				result = append(result, im.accessoryItems[i])
				break
			}
		}

	}
	return result
}

func (im *ItemManager) ArmorsByName(item Selection) []Item_Armor { // selected weapons from json
	var selection []string
	selection = append(selection,
		item.ItemSlot.Head.Name, item.ItemSlot.Chest.Name, item.ItemSlot.Foot.Name,
		item.ItemSlot.Hands.Name, item.ItemSlot.Pants.Name, item.ItemSlot.Back.Name)

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
	for i := 0; i < len(im.armorItems); i++ {
		for j := 0; j < len(selection); j++ {
			if im.armorItems[i].Name == selection[j] {
				result = append(result, im.weaponItems[i])
			}
		}

	}
	return result
}

func (im *ItemManager) AccesorysByName() []Item_Accessory {

	var result []Item_Accessory
	for i := 0; i < len(im.accessoryItems); i++ {
		result = append(result, im.accessoryItems[i])
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

type StatsManager struct {
	base     Stats
	variable Computed_Stats
}

func NewStatsManager() *StatsManager {
	return &StatsManager{
		base:     characterStats,
		variable: Computed_Stats{},
	}
}

func (sm *StatsManager) ItemsBaseStats(selection Selection) Stats {
	// CALCULATE TOTAL BASE ATTRIBUTE STATS OF ITEMS=
	Im := NewItemManager()
	var selclass Stats = Selection_Class(selection)
	var selarmor []Item_Armor = Im.ArmorsByName(selection)
	var selrarity []int = Im.SelByRarity(selection)
	// Calculate total by sum base stats and item stats
	for i := 0; i < len(selarmor) && i < len(selrarity); i++ {
		selclass = Stats{
			Strength:        selclass.Strength + selarmor[i].BaseAttribute.Strength[selrarity[i]],
			Vigor:           selclass.Vigor + selarmor[i].BaseAttribute.Vigor[selrarity[i]],
			Agility:         selclass.Agility + selarmor[i].BaseAttribute.Agility[selrarity[i]],
			Dexterity:       selclass.Dexterity + selarmor[i].BaseAttribute.Dexterity[selrarity[i]],
			Will:            selclass.Will + selarmor[i].BaseAttribute.Will[selrarity[i]],
			Knowledge:       selclass.Knowledge + selarmor[i].BaseAttribute.Knowledge[selrarity[i]],
			Resourcefulness: selclass.Resourcefulness + selarmor[i].BaseAttribute.Resourcefulness[selrarity[i]],
		}
	}
	// Return the total stats struct
	return selclass
}

func (sm *StatsManager) BaseItemCalc(selection Selection) Computed_Stats {
	Im := NewItemManager()
	var selarmor []Item_Armor = Im.ArmorsByName(selection)
	var selrarity []int = Im.SelByRarity(selection)
	computed := Computed_Stats{}
	for i := 0; i < len(selarmor); i++ {
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
	return computed
}
