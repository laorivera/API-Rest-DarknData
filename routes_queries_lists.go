package main

import "builder/src/models"

// /////\\\\\\ --------------> ENCHCANTMENT LISTS ARMOR <------------------///////////\\\\\\\\\

func GetEnchatmentLists_Armor_Base(selection models.Selection) map[string][]string {

	lists := map[string][]string{}
	for i := 0; i < len(slots); i++ {
		if slots[i] == "helmet" {
			item := ItemsByNameArmor(selection.ItemSlot.Head.Name)
			echantments := EnchamentbySlot(Enchantments.Helmet)
			lists[slots[i]] = EnchantBaseAttribExeption(echantments, item)
		}
		if slots[i] == "chest" {
			item := ItemsByNameArmor(selection.ItemSlot.Chest.Name)
			echantments := EnchamentbySlot(Enchantments.Chest)
			lists[slots[i]] = EnchantBaseAttribExeption(echantments, item)
		}
		if slots[i] == "gloves" {
			item := ItemsByNameArmor(selection.ItemSlot.Hands.Name)
			echantments := EnchamentbySlot(Enchantments.Gloves)
			lists[slots[i]] = EnchantBaseAttribExeption(echantments, item)
		}
		if slots[i] == "pants" {
			item := ItemsByNameArmor(selection.ItemSlot.Pants.Name)
			echantments := EnchamentbySlot(Enchantments.Pants)
			lists[slots[i]] = EnchantBaseAttribExeption(echantments, item)
		}
		if slots[i] == "boots" {
			item := ItemsByNameArmor(selection.ItemSlot.Foot.Name)
			echantments := EnchamentbySlot(Enchantments.Boots)
			lists[slots[i]] = EnchantBaseAttribExeption(echantments, item)
		}
		if slots[i] == "cloak" {
			item := ItemsByNameArmor(selection.ItemSlot.Back.Name)
			echantments := EnchamentbySlot(Enchantments.Cloak)
			lists[slots[i]] = EnchantBaseAttribExeption(echantments, item)
		}

	}
	return lists

}
func GetEnchatmentLists_Armor_ValuesUncommon(selection models.Selection) map[string][]float32 {

	lists := map[string][]float32{}
	for i := 0; i < len(slots); i++ {

		if slots[i] == "helmet" {
			lists[slots[i]] = EnchantValuesCalc(selection.ItemSlot.Head.Enchant.TypeU, Enchantments.Helmet)
		}
		if slots[i] == "chest" {
			lists[slots[i]] = EnchantValuesCalc(selection.ItemSlot.Chest.Enchant.TypeU, Enchantments.Chest)
		}
		if slots[i] == "gloves" {
			lists[slots[i]] = EnchantValuesCalc(selection.ItemSlot.Hands.Enchant.TypeU, Enchantments.Gloves)
		}
		if slots[i] == "pants" {
			lists[slots[i]] = EnchantValuesCalc(selection.ItemSlot.Pants.Enchant.TypeU, Enchantments.Pants)
		}
		if slots[i] == "boots" {
			lists[slots[i]] = EnchantValuesCalc(selection.ItemSlot.Foot.Enchant.TypeU, Enchantments.Boots)
		}
		if slots[i] == "cloak" {
			lists[slots[i]] = EnchantValuesCalc(selection.ItemSlot.Back.Enchant.TypeU, Enchantments.Cloak)
		}
	}
	return lists
}
func GetEnchatmentLists_Armor_TypeRare(selection models.Selection) map[string][]string {

	lists := make(map[string][]string)

	baseLists := GetEnchatmentLists_Armor_Base(selection)

	for _, slot := range slots {
		// Pass all previous selections up to uncommon
		previousSelections := []string{
			selection.ItemSlot.Head.Enchant.TypeU, selection.ItemSlot.Chest.Enchant.TypeU, selection.ItemSlot.Hands.Enchant.TypeU,
			selection.ItemSlot.Pants.Enchant.TypeU, selection.ItemSlot.Foot.Enchant.TypeU, selection.ItemSlot.Back.Enchant.TypeU,
		}
		lists[slot] = EnchantTypeExeption(baseLists[slot], previousSelections)
	}

	return lists

}
func GetEnchatmentLists_Armor_ValuesRare(selection models.Selection) map[string][]float32 {

	lists := make(map[string][]float32)
	for i := 0; i < len(slots); i++ {

		if slots[i] == "helmet" {
			lists[slots[i]] = EnchantValuesCalc(selection.ItemSlot.Head.Enchant.TypeR, Enchantments.Helmet)
		}
		if slots[i] == "chest" {
			lists[slots[i]] = EnchantValuesCalc(selection.ItemSlot.Chest.Enchant.TypeR, Enchantments.Chest)
		}
		if slots[i] == "gloves" {
			lists[slots[i]] = EnchantValuesCalc(selection.ItemSlot.Hands.Enchant.TypeR, Enchantments.Gloves)
		}
		if slots[i] == "pants" {
			lists[slots[i]] = EnchantValuesCalc(selection.ItemSlot.Pants.Enchant.TypeR, Enchantments.Pants)
		}
		if slots[i] == "boots" {
			lists[slots[i]] = EnchantValuesCalc(selection.ItemSlot.Foot.Enchant.TypeR, Enchantments.Boots)
		}
		if slots[i] == "cloak" {
			lists[slots[i]] = EnchantValuesCalc(selection.ItemSlot.Back.Enchant.TypeR, Enchantments.Cloak)
		}
	}
	return lists
}
func GetEnchatmentLists_Armor_TypeEpic(selection models.Selection) map[string][]string {

	lists := make(map[string][]string)

	baseLists := GetEnchatmentLists_Armor_Base(selection)

	for _, slot := range slots {
		// Pass all previous selections up to rare
		previousSelections := []string{
			selection.ItemSlot.Head.Enchant.TypeU, selection.ItemSlot.Chest.Enchant.TypeU, selection.ItemSlot.Hands.Enchant.TypeU,
			selection.ItemSlot.Pants.Enchant.TypeU, selection.ItemSlot.Foot.Enchant.TypeU, selection.ItemSlot.Back.Enchant.TypeU,
			selection.ItemSlot.Head.Enchant.TypeR, selection.ItemSlot.Chest.Enchant.TypeR, selection.ItemSlot.Hands.Enchant.TypeR,
			selection.ItemSlot.Pants.Enchant.TypeR, selection.ItemSlot.Foot.Enchant.TypeR, selection.ItemSlot.Back.Enchant.TypeR,
		}
		lists[slot] = EnchantTypeExeption(baseLists[slot], previousSelections)
	}

	return lists
}
func GetEnchatmentLists_Armor_ValuesEpic(selection models.Selection) map[string][]float32 {

	lists := make(map[string][]float32)
	for i := 0; i < len(slots); i++ {

		if slots[i] == "helmet" {
			lists[slots[i]] = EnchantValuesCalc(selection.ItemSlot.Head.Enchant.TypeE, Enchantments.Helmet)
		}
		if slots[i] == "chest" {
			lists[slots[i]] = EnchantValuesCalc(selection.ItemSlot.Chest.Enchant.TypeE, Enchantments.Chest)
		}
		if slots[i] == "gloves" {
			lists[slots[i]] = EnchantValuesCalc(selection.ItemSlot.Hands.Enchant.TypeE, Enchantments.Gloves)
		}
		if slots[i] == "pants" {
			lists[slots[i]] = EnchantValuesCalc(selection.ItemSlot.Pants.Enchant.TypeE, Enchantments.Pants)
		}
		if slots[i] == "boots" {
			lists[slots[i]] = EnchantValuesCalc(selection.ItemSlot.Foot.Enchant.TypeE, Enchantments.Boots)
		}
		if slots[i] == "cloak" {
			lists[slots[i]] = EnchantValuesCalc(selection.ItemSlot.Back.Enchant.TypeE, Enchantments.Cloak)
		}
	}
	return lists
}
func GetEnchatmentLists_Armor_TypeLegend(selection models.Selection) map[string][]string {

	lists := make(map[string][]string)

	baseLists := GetEnchatmentLists_Armor_Base(selection)

	for _, slot := range slots {
		// Pass all previous selections up to Epic
		previousSelections := []string{
			selection.ItemSlot.Head.Enchant.TypeU, selection.ItemSlot.Chest.Enchant.TypeU, selection.ItemSlot.Hands.Enchant.TypeU,
			selection.ItemSlot.Pants.Enchant.TypeU, selection.ItemSlot.Foot.Enchant.TypeU, selection.ItemSlot.Back.Enchant.TypeU,
			selection.ItemSlot.Head.Enchant.TypeR, selection.ItemSlot.Chest.Enchant.TypeR, selection.ItemSlot.Hands.Enchant.TypeR,
			selection.ItemSlot.Pants.Enchant.TypeR, selection.ItemSlot.Foot.Enchant.TypeR, selection.ItemSlot.Back.Enchant.TypeR,
			selection.ItemSlot.Head.Enchant.TypeE, selection.ItemSlot.Chest.Enchant.TypeE, selection.ItemSlot.Hands.Enchant.TypeE,
			selection.ItemSlot.Pants.Enchant.TypeE, selection.ItemSlot.Foot.Enchant.TypeE, selection.ItemSlot.Back.Enchant.TypeE,
		}
		lists[slot] = EnchantTypeExeption(baseLists[slot], previousSelections)
	}

	return lists
}
func GetEnchatmentLists_Armor_ValuesLegend(selection models.Selection) map[string][]float32 {

	lists := make(map[string][]float32)
	for i := 0; i < len(slots); i++ {

		if slots[i] == "helmet" {
			lists[slots[i]] = EnchantValuesCalc(selection.ItemSlot.Head.Enchant.TypeL, Enchantments.Helmet)
		}
		if slots[i] == "chest" {
			lists[slots[i]] = EnchantValuesCalc(selection.ItemSlot.Chest.Enchant.TypeL, Enchantments.Chest)
		}
		if slots[i] == "gloves" {
			lists[slots[i]] = EnchantValuesCalc(selection.ItemSlot.Hands.Enchant.TypeL, Enchantments.Gloves)
		}
		if slots[i] == "pants" {
			lists[slots[i]] = EnchantValuesCalc(selection.ItemSlot.Pants.Enchant.TypeL, Enchantments.Pants)
		}
		if slots[i] == "boots" {
			lists[slots[i]] = EnchantValuesCalc(selection.ItemSlot.Foot.Enchant.TypeL, Enchantments.Boots)
		}
		if slots[i] == "cloak" {
			lists[slots[i]] = EnchantValuesCalc(selection.ItemSlot.Back.Enchant.TypeL, Enchantments.Cloak)
		}
	}
	return lists
}
func GetEnchatmentLists_Armor_TypeUnique(selection models.Selection) map[string][]string {

	lists := make(map[string][]string)

	baseLists := GetEnchatmentLists_Armor_Base(selection)

	for _, slot := range slots {
		// Pass all previous selections up to Epic
		previousSelections := []string{
			selection.ItemSlot.Head.Enchant.TypeU, selection.ItemSlot.Chest.Enchant.TypeU, selection.ItemSlot.Hands.Enchant.TypeU,
			selection.ItemSlot.Pants.Enchant.TypeU, selection.ItemSlot.Foot.Enchant.TypeU, selection.ItemSlot.Back.Enchant.TypeU,
			selection.ItemSlot.Head.Enchant.TypeR, selection.ItemSlot.Chest.Enchant.TypeR, selection.ItemSlot.Hands.Enchant.TypeR,
			selection.ItemSlot.Pants.Enchant.TypeR, selection.ItemSlot.Foot.Enchant.TypeR, selection.ItemSlot.Back.Enchant.TypeR,
			selection.ItemSlot.Head.Enchant.TypeE, selection.ItemSlot.Chest.Enchant.TypeE, selection.ItemSlot.Hands.Enchant.TypeE,
			selection.ItemSlot.Pants.Enchant.TypeE, selection.ItemSlot.Foot.Enchant.TypeE, selection.ItemSlot.Back.Enchant.TypeE,
			selection.ItemSlot.Head.Enchant.TypeL, selection.ItemSlot.Chest.Enchant.TypeL, selection.ItemSlot.Hands.Enchant.TypeL,
			selection.ItemSlot.Pants.Enchant.TypeL, selection.ItemSlot.Foot.Enchant.TypeL, selection.ItemSlot.Back.Enchant.TypeL,
		}
		lists[slot] = EnchantTypeExeption(baseLists[slot], previousSelections)
	}

	return lists
}
func GetEnchatmentLists_Armor_ValuesUnique(selection models.Selection) map[string][]float32 {

	lists := make(map[string][]float32)
	for i := 0; i < len(slots); i++ {

		if slots[i] == "helmet" {
			lists[slots[i]] = EnchantValuesCalc(selection.ItemSlot.Head.Enchant.TypeQ, Enchantments.Helmet)
		}
		if slots[i] == "chest" {
			lists[slots[i]] = EnchantValuesCalc(selection.ItemSlot.Chest.Enchant.TypeQ, Enchantments.Chest)
		}
		if slots[i] == "gloves" {
			lists[slots[i]] = EnchantValuesCalc(selection.ItemSlot.Hands.Enchant.TypeQ, Enchantments.Gloves)
		}
		if slots[i] == "pants" {
			lists[slots[i]] = EnchantValuesCalc(selection.ItemSlot.Pants.Enchant.TypeQ, Enchantments.Pants)
		}
		if slots[i] == "boots" {
			lists[slots[i]] = EnchantValuesCalc(selection.ItemSlot.Foot.Enchant.TypeQ, Enchantments.Boots)
		}
		if slots[i] == "cloak" {
			lists[slots[i]] = EnchantValuesCalc(selection.ItemSlot.Back.Enchant.TypeQ, Enchantments.Cloak)
		}
	}
	return lists
}

/////////\\\\\\\\\ --------------> ENCHCANTMENT LISTS ACCESSORY <------------------ //////////\\\\\\\\\

func GetEnchatmentLists_Accessory_Base(selection models.Selection) map[string][]string {

	lists := map[string][]string{}
	for i := 0; i < len(slotsacc); i++ {
		if slotsacc[i] == "necklace" {
			lists[slotsacc[i]] = EnchantBaseAttribExeptionAcc(EnchamentbySlot(Enchantments.Necklace), ItemsByNameAccesory(selection.ItemSlot.Necklace.Name))
		}
		if slotsacc[i] == "ring" {
			lists[slotsacc[i]] = EnchantBaseAttribExeptionAcc(EnchamentbySlot(Enchantments.Ring), ItemsByNameAccesory(selection.ItemSlot.RingOne.Name))
		}
		if slotsacc[i] == "ringtwo" {
			lists[slotsacc[i]] = EnchantBaseAttribExeptionAcc(EnchamentbySlot(Enchantments.Ring), ItemsByNameAccesory(selection.ItemSlot.RingTwo.Name))
		}
	}
	return lists

}
func GetEnchantmentLists_Accessory_ValuesUncommon(selection models.Selection) map[string][]float32 {

	lists := map[string][]float32{}
	for i := 0; i < len(slotsacc); i++ {
		//
		if slotsacc[i] == "necklace" {
			lists[slotsacc[i]] = EnchantValuesCalc(selection.ItemSlot.Necklace.Enchant.TypeU, Enchantments.Necklace)
		}
		if slotsacc[i] == "ring" {
			lists[slotsacc[i]] = EnchantValuesCalc(selection.ItemSlot.RingOne.Enchant.TypeU, Enchantments.Ring)
		}
		if slotsacc[i] == "ringtwo" {
			lists[slotsacc[i]] = EnchantValuesCalc(selection.ItemSlot.RingTwo.Enchant.TypeU, Enchantments.Ring)
		}
	}
	return lists
}
func GetEnchantmentLists_Accessory_TypeRare(selection models.Selection) map[string][]string {

	lists := make(map[string][]string)

	baseLists := GetEnchatmentLists_Accessory_Base(selection)

	for _, slot := range slotsacc {
		// Pass all previous selections up to Epic
		previousSelections := []string{
			selection.ItemSlot.Necklace.Enchant.TypeU, selection.ItemSlot.RingOne.Enchant.TypeU,
			selection.ItemSlot.RingTwo.Enchant.TypeU, // Uncommon
		}
		lists[slot] = EnchantTypeExeption(baseLists[slot], previousSelections)
	}
	return lists
}
func GetEnchantmentLists_Accessory_ValuesRare(selection models.Selection) map[string][]float32 {

	lists := make(map[string][]float32)
	for i := 0; i < len(slotsacc); i++ {

		if slotsacc[i] == "necklace" {
			lists[slotsacc[i]] = EnchantValuesCalc(selection.ItemSlot.Necklace.Enchant.TypeR, Enchantments.Necklace)
		}
		if slotsacc[i] == "ring" {
			lists[slotsacc[i]] = EnchantValuesCalc(selection.ItemSlot.RingOne.Enchant.TypeR, Enchantments.Ring)
		}
		if slotsacc[i] == "ringtwo" {
			lists[slotsacc[i]] = EnchantValuesCalc(selection.ItemSlot.RingTwo.Enchant.TypeR, Enchantments.Ring)
		}
	}
	return lists
}
func GetEnchantmentLists_Accessory_TypeEpic(selection models.Selection) map[string][]string {

	lists := make(map[string][]string)

	baseLists := GetEnchatmentLists_Accessory_Base(selection)

	for _, slot := range slotsacc {
		// Pass all previous selections up to Epic
		previousSelections := []string{
			selection.ItemSlot.Necklace.Enchant.TypeU, selection.ItemSlot.RingOne.Enchant.TypeU,
			selection.ItemSlot.RingTwo.Enchant.TypeU, selection.ItemSlot.Necklace.Enchant.TypeR, selection.ItemSlot.RingOne.Enchant.TypeR,
			selection.ItemSlot.RingTwo.Enchant.TypeR, // Rare
		}
		lists[slot] = EnchantTypeExeption(baseLists[slot], previousSelections)
	}

	return lists
}
func GetEnchantmentLists_Accessory_ValuesEpic(selection models.Selection) map[string][]float32 {

	lists := make(map[string][]float32)
	for i := 0; i < len(slotsacc); i++ {

		if slotsacc[i] == "necklace" {
			lists[slotsacc[i]] = EnchantValuesCalc(selection.ItemSlot.Necklace.Enchant.TypeE, Enchantments.Necklace)
		}
		if slotsacc[i] == "ring" {
			lists[slotsacc[i]] = EnchantValuesCalc(selection.ItemSlot.RingOne.Enchant.TypeE, Enchantments.Ring)
		}
		if slotsacc[i] == "ringtwo" {
			lists[slotsacc[i]] = EnchantValuesCalc(selection.ItemSlot.RingTwo.Enchant.TypeE, Enchantments.Ring)
		}
	}
	return lists
}
func GetEnchantmentLists_Accessory_TypeLegend(selection models.Selection) map[string][]string {

	lists := make(map[string][]string)

	baseLists := GetEnchatmentLists_Accessory_Base(selection)

	for _, slot := range slotsacc {
		// Pass all previous selections up to Epic
		previousSelections := []string{
			selection.ItemSlot.Necklace.Enchant.TypeU, selection.ItemSlot.RingOne.Enchant.TypeU,
			selection.ItemSlot.RingTwo.Enchant.TypeU, selection.ItemSlot.Necklace.Enchant.TypeR, selection.ItemSlot.RingOne.Enchant.TypeR,
			selection.ItemSlot.RingTwo.Enchant.TypeR, selection.ItemSlot.Necklace.Enchant.TypeE, selection.ItemSlot.RingOne.Enchant.TypeE,
			selection.ItemSlot.RingTwo.Enchant.TypeE, // Epic
		}
		lists[slot] = EnchantTypeExeption(baseLists[slot], previousSelections)
	}

	return lists
}
func GetEnchantmentLists_Accessory_ValuesLegend(selection models.Selection) map[string][]float32 {

	lists := make(map[string][]float32)
	for i := 0; i < len(slotsacc); i++ {

		if slotsacc[i] == "necklace" {
			lists[slotsacc[i]] = EnchantValuesCalc(selection.ItemSlot.Necklace.Enchant.TypeL, Enchantments.Necklace)
		}
		if slotsacc[i] == "ring" {
			lists[slotsacc[i]] = EnchantValuesCalc(selection.ItemSlot.RingOne.Enchant.TypeL, Enchantments.Ring)
		}
		if slotsacc[i] == "ringtwo" {
			lists[slotsacc[i]] = EnchantValuesCalc(selection.ItemSlot.RingTwo.Enchant.TypeL, Enchantments.Ring)
		}
	}
	return lists
}
func GetEnchantmentLists_Accessory_TypeUnique(selection models.Selection) map[string][]string {

	lists := make(map[string][]string)

	baseLists := GetEnchatmentLists_Accessory_Base(selection)

	for _, slot := range slotsacc {
		// Pass all previous selections up to Epic
		previousSelections := []string{
			selection.ItemSlot.Necklace.Enchant.TypeU, selection.ItemSlot.RingOne.Enchant.TypeU,
			selection.ItemSlot.RingTwo.Enchant.TypeU, selection.ItemSlot.Necklace.Enchant.TypeR, selection.ItemSlot.RingOne.Enchant.TypeR,
			selection.ItemSlot.RingTwo.Enchant.TypeR, selection.ItemSlot.Necklace.Enchant.TypeE, selection.ItemSlot.RingOne.Enchant.TypeE,
			selection.ItemSlot.RingTwo.Enchant.TypeE, selection.ItemSlot.Necklace.Enchant.TypeL, selection.ItemSlot.RingOne.Enchant.TypeL,
			selection.ItemSlot.RingTwo.Enchant.TypeL,
		}
		lists[slot] = EnchantTypeExeption(baseLists[slot], previousSelections)
	}

	return lists
}
func GetEnchantmentLists_Accessory_ValuesUnique(selection models.Selection) map[string][]float32 {

	lists := make(map[string][]float32)
	for i := 0; i < len(slotsacc); i++ {

		if slotsacc[i] == "necklace" {
			lists[slotsacc[i]] = EnchantValuesCalc(selection.ItemSlot.Necklace.Enchant.TypeQ, Enchantments.Necklace)
		}
		if slotsacc[i] == "ring" {
			lists[slotsacc[i]] = EnchantValuesCalc(selection.ItemSlot.RingOne.Enchant.TypeQ, Enchantments.Ring)
		}
		if slotsacc[i] == "ringtwo" {
			lists[slotsacc[i]] = EnchantValuesCalc(selection.ItemSlot.RingTwo.Enchant.TypeQ, Enchantments.Ring)
		}
	}
	return lists
}

/////////\\\\\\\\\ --------------> ENCHCANTMENT LISTS WEAPON <------------------ //////////\\\\\\\\\

func GetEnchantmentLists_Weapon_Base(selection models.Selection) map[string][]string {

	lists := map[string][]string{}
	//physicaltypes := []string{"axe", "sword", "mace"}
	//shieldtypes := []string{"shield"}
	//hybridtypes := []string{"polearm", "crossbow"}
	//magictypes := []string{"magicstuff", "staff"}

	for i := 0; i < len(slotsweapon); i++ {
		itemName := selection.ItemSlot.WeaponOne.Name
		itemNameTwo := selection.ItemSlot.WeaponTwo.Name
		if slotsweapon[i] == "pwo" {
			if ItemsByNameWeapon(itemName).HandType == "One Handed" && ItemsByNameWeapon(itemName).SlotType == "Main Hand" {
				lists[slotsweapon[i]] = EnchamentbySlot(Enchantments.WeaponOneHand)
			} else if ItemsByNameWeapon(itemName).HandType == "Two Handed" && ItemsByNameWeapon(itemName).SlotType == "Main Hand" {
				lists[slotsweapon[i]] = EnchamentbySlot(Enchantments.PhysicalTwoHand)
			}
		}
		if slotsweapon[i] == "pwt" {
			itemName := selection.ItemSlot.WeaponTwo.Name
			if ItemsByNameWeapon(itemNameTwo).SlotType == "Off Hand" && ItemsByNameWeapon(itemName).HandType == "One Handed" {
				lists[slotsweapon[i]] = EnchamentbySlot(Enchantments.WeaponOneHand)
			}
			//println(ItemsByNameWeapon(c.Query("item" + slotsweapon[i])).SlotType)
		}
		/*
			if slotsweapon[i] == "swo" {
				if ItemsByNameWeapon(c.Query("item"+slotsweapon[i])).SlotType == "Main Hand" {
					lists[slotsweapon[i]] = EnchamentbySlot(Enchantments.WeaponOneHand)
				}
			}
			if slotsweapon[i] == "swt" {
				if ItemsByNameWeapon(c.Query("item"+slotsweapon[i])).SlotType == "Off Hand" {
					lists[slotsweapon[i]] = EnchamentbySlot(Enchantments.WeaponOneHand)
				}
			}
		*/
	}
	return lists

}
func GetEnchantmentLists_Weapon_ValuesUncommon(selection models.Selection) map[string][]float32 {

	lists := map[string][]float32{}
	for i := 0; i < len(slotsweapon); i++ {

		if slotsweapon[i] == "pwo" {
			lists[slotsweapon[i]] = EnchantValuesCalc(selection.ItemSlot.WeaponOne.Enchant.TypeU, Enchantments.WeaponOneHand)
		}

		if slotsweapon[i] == "pwt" {
			lists[slotsweapon[i]] = EnchantValuesCalc(selection.ItemSlot.WeaponTwo.Enchant.TypeU, Enchantments.WeaponOneHand)
		}
		/*
			if slotsweapon[i] == "swo" {
				lists[slotsweapon[i]] = EnchantValuesCalc(Query, Enchantments.WeaponOneHand)
			}
			if slotsweapon[i] == "swt" {
				lists[slotsweapon[i]] = EnchantValuesCalc(Query, Enchantments.WeaponOneHand)
			}
		*/
	}
	return lists
}

func GetEnchantmentLists_Weapon_TypeRare(selection models.Selection) map[string][]string {

	lists := make(map[string][]string)

	baseLists := GetEnchantmentLists_Weapon_Base(selection)

	for _, slot := range slotsweapon {
		// Pass all previous selections up to Epic
		previousSelections := []string{
			selection.ItemSlot.WeaponOne.Enchant.TypeU, selection.ItemSlot.WeaponTwo.Enchant.TypeU, // Uncommon
		}
		lists[slot] = EnchantTypeExeption(baseLists[slot], previousSelections)
	}

	return lists

}

func GetEnchantmentLists_Weapon_ValuesRare(selection models.Selection) map[string][]float32 {
	//slotsweapon := []string{"pwo", "pwt", "swo", "swt"}

	lists := make(map[string][]float32)
	for i := 0; i < len(slotsweapon); i++ {
		//Query := c.Query("enchantment_" + slotsweapon[i] + "type2")
		if slotsweapon[i] == "pwo" {
			lists[slotsweapon[i]] = EnchantValuesCalc(selection.ItemSlot.WeaponOne.Enchant.TypeR, Enchantments.WeaponOneHand)
		}
		if slotsweapon[i] == "pwt" {
			lists[slotsweapon[i]] = EnchantValuesCalc(selection.ItemSlot.WeaponTwo.Enchant.TypeR, Enchantments.WeaponOneHand)
		}
		/*
			if slotsweapon[i] == "swo" {
				lists[slotsweapon[i]] = EnchantValuesCalc(Query, Enchantments.WeaponOneHand)
			}
			if slotsweapon[i] == "swt" {
				lists[slotsweapon[i]] = EnchantValuesCalc(Query, Enchantments.WeaponOneHand)
			}
		*/

	}
	return lists
}

func GetEnchantmentLists_Weapon_TypeEpic(selection models.Selection) map[string][]string {

	lists := make(map[string][]string)

	baseLists := GetEnchantmentLists_Weapon_Base(selection)

	for _, slot := range slotsweapon {
		// Pass all previous selections up to Epic
		previousSelections := []string{
			selection.ItemSlot.WeaponOne.Enchant.TypeU, selection.ItemSlot.WeaponTwo.Enchant.TypeU,
			selection.ItemSlot.WeaponOne.Enchant.TypeR, selection.ItemSlot.WeaponTwo.Enchant.TypeR,
		}
		lists[slot] = EnchantTypeExeption(baseLists[slot], previousSelections)
	}

	return lists

}
func GetEnchantmentLists_Weapon_ValuesEpic(selection models.Selection) map[string][]float32 {

	lists := make(map[string][]float32)
	for i := 0; i < len(slotsweapon); i++ {
		//Query := c.Query("enchantment_" + slotsweapon[i] + "type3")
		if slotsweapon[i] == "pwo" {
			lists[slotsweapon[i]] = EnchantValuesCalc(selection.ItemSlot.WeaponOne.Enchant.TypeE, Enchantments.WeaponOneHand)
		}

		if slots[i] == "pwt" {
			lists[slots[i]] = EnchantValuesCalc(selection.ItemSlot.WeaponTwo.Enchant.TypeE, Enchantments.WeaponOneHand)
		}
		/*
			if slots[i] == "swo" {
				lists[slots[i]] = EnchantValuesCalc(Query, Enchantments.SecondaryWeaponOneHand)
			}
			if slots[i] == "swt" {
				lists[slots[i]] = EnchantValuesCalc(Query, Enchantments.SecondaryWeaponTwoHand)
			}
		*/
	}
	return lists
}
func GetEnchantmentLists_Weapon_TypeLegend(selection models.Selection) map[string][]string {

	lists := make(map[string][]string)

	baseLists := GetEnchantmentLists_Weapon_Base(selection)

	for _, slot := range slotsweapon {
		// Pass all previous selections up to Epic
		previousSelections := []string{
			selection.ItemSlot.WeaponOne.Enchant.TypeU, selection.ItemSlot.WeaponTwo.Enchant.TypeU,
			selection.ItemSlot.WeaponOne.Enchant.TypeR, selection.ItemSlot.WeaponTwo.Enchant.TypeR,
			selection.ItemSlot.WeaponOne.Enchant.TypeE, selection.ItemSlot.WeaponTwo.Enchant.TypeE,
		}
		lists[slot] = EnchantTypeExeption(baseLists[slot], previousSelections)
	}

	return lists

}
func GetEnchantmentLists_Weapon_ValuesLegend(selection models.Selection) map[string][]float32 {

	lists := make(map[string][]float32)
	for i := 0; i < len(slotsweapon); i++ {

		if slotsweapon[i] == "pwo" {
			lists[slotsweapon[i]] = EnchantValuesCalc(selection.ItemSlot.WeaponOne.Enchant.TypeL, Enchantments.WeaponOneHand)
		}

		if slots[i] == "pwt" {
			lists[slots[i]] = EnchantValuesCalc(selection.ItemSlot.WeaponTwo.Enchant.TypeL, Enchantments.WeaponOneHand)
		}
		/*
			if slots[i] == "swo" {
				lists[slots[i]] = EnchantValuesCalc(Query, Enchantments.SecondaryWeaponOneHand)
			}
			if slots[i] == "swt" {
				lists[slots[i]] = EnchantValuesCalc(Query, Enchantments.SecondaryWeaponTwoHand)
			}
		*/

	}
	return lists
}
func GetEnchantmentLists_Weapon_TypeUnique(selection models.Selection) map[string][]string {

	lists := make(map[string][]string)

	baseLists := GetEnchantmentLists_Weapon_Base(selection)

	for _, slot := range slotsweapon {
		// Pass all previous selections up to Epic
		previousSelections := []string{
			selection.ItemSlot.WeaponOne.Enchant.TypeU, selection.ItemSlot.WeaponTwo.Enchant.TypeU,
			selection.ItemSlot.WeaponOne.Enchant.TypeR, selection.ItemSlot.WeaponTwo.Enchant.TypeR,
			selection.ItemSlot.WeaponOne.Enchant.TypeE, selection.ItemSlot.WeaponTwo.Enchant.TypeE,
			selection.ItemSlot.WeaponOne.Enchant.TypeL, selection.ItemSlot.WeaponTwo.Enchant.TypeL,
		}
		lists[slot] = EnchantTypeExeption(baseLists[slot], previousSelections)
	}

	return lists

}
func GetEnchantmentLists_Weapon_ValuesUnique(selection models.Selection) map[string][]float32 {

	lists := make(map[string][]float32)
	for i := 0; i < len(slotsweapon); i++ {

		if slotsweapon[i] == "pwo" {
			lists[slotsweapon[i]] = EnchantValuesCalc(selection.ItemSlot.WeaponOne.Enchant.TypeQ, Enchantments.WeaponOneHand)
		}

		if slots[i] == "pwt" {
			lists[slots[i]] = EnchantValuesCalc(selection.ItemSlot.WeaponTwo.Enchant.TypeQ, Enchantments.WeaponOneHand)
		}
		/*
			if slots[i] == "swo" {
				lists[slots[i]] = EnchantValuesCalc(Query, Enchantments.SecondaryWeaponOneHand)
			}
			if slots[i] == "swt" {
				lists[slots[i]] = EnchantValuesCalc(Query, Enchantments.SecondaryWeaponTwoHand)
			}
		*/

	}
	return lists
}
