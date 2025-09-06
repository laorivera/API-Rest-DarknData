package main

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

func (im *ItemManager) ArmorByName(name string) Item_Armor {
	var result Item_Armor
	for i := 0; i < len(im.armorItems); i++ {
		if im.armorItems[i].Name == name {
			result = im.armorItems[i]
			return result
		}
	}
	return result
}

func (im *ItemManager) WeaponByName(name string) Item_Weapon {
	var result Item_Weapon
	for i := 0; i < len(im.weaponItems); i++ {
		if im.weaponItems[i].Name == name {
			result = im.weaponItems[i]

		}
	}
	return result
}

func (im *ItemManager) AccesoryByName(name string) Item_Accessory {
	var result Item_Accessory
	for i := 0; i < len(im.accessoryItems); i++ {
		if im.accessoryItems[i].Name == name {
			result = im.accessoryItems[i]

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
		base:     characterHolder,
		variable: Computed_Stats{},
	}
}

/*
func (im *ItemManager) ArmorsByName(slot string, name string) []string{
	class := InttoClass(name)
	var result []string
	for i := 0; i < len(Items.ItemsArmor); i++ {
		for _, c := range Items.ItemsArmor[i].Classes {
			if c == class && slot == Items.ItemsArmor[i].SlotType {
				result = append(result, Items.ItemsArmor[i].Name)
				break
			}
		}

	}
	return result

}
func (im *ItemManager) AccesoryByName(slot string, name string) []string{
	for i := 0; i < len(Items.ItemsAccessory); i++ {
		class := InttoClass(classID)
		var result []string
		for _, c := range Items.ItemsAccessory[i].Classes {
			if c == class && slot == Items.ItemsAccessory[i].SlotType {
				result = append(result, Items.ItemsAccessory[i].Name)
				break
			}
		}
		return result
	}
}
*/
