package controllers

import (
	"builder/src/models"
)

// /////\\\\\\ --------------> ENCHCANTMENT LISTS ARMOR <------------------///////////\\\\\\\\\

func GetEnchatmentLists_Armor_Base(selection models.Selection) map[string][]string {

	lists := map[string][]string{}
	for i := 0; i < len(models.Slots); i++ {
		if models.Slots[i] == "helmet" {
			item := ItemsByNameArmor(selection.ItemSlot.Head.Name)
			echantments := EnchamentbySlot(models.Enchantments.Helmet)
			lists[models.Slots[i]] = EnchantBaseAttribExeption(echantments, *item)
		}
		if models.Slots[i] == "chest" {
			item := ItemsByNameArmor(selection.ItemSlot.Chest.Name)
			echantments := EnchamentbySlot(models.Enchantments.Chest)
			lists[models.Slots[i]] = EnchantBaseAttribExeption(echantments, *item)
		}
		if models.Slots[i] == "gloves" {
			item := ItemsByNameArmor(selection.ItemSlot.Hands.Name)
			echantments := EnchamentbySlot(models.Enchantments.Gloves)
			lists[models.Slots[i]] = EnchantBaseAttribExeption(echantments, *item)
		}
		if models.Slots[i] == "pants" {
			item := ItemsByNameArmor(selection.ItemSlot.Pants.Name)
			echantments := EnchamentbySlot(models.Enchantments.Pants)
			lists[models.Slots[i]] = EnchantBaseAttribExeption(echantments, *item)
		}
		if models.Slots[i] == "boots" {
			item := ItemsByNameArmor(selection.ItemSlot.Foot.Name)
			echantments := EnchamentbySlot(models.Enchantments.Boots)
			lists[models.Slots[i]] = EnchantBaseAttribExeption(echantments, *item)
		}
		if models.Slots[i] == "cloak" {
			item := ItemsByNameArmor(selection.ItemSlot.Back.Name)
			echantments := EnchamentbySlot(models.Enchantments.Cloak)
			lists[models.Slots[i]] = EnchantBaseAttribExeption(echantments, *item)
		}

	}
	return lists

}
func GetEnchatmentLists_Armor_ValuesUncommon(selection models.Selection) map[string][]float32 {

	lists := map[string][]float32{}
	for i := 0; i < len(models.Slots); i++ {

		if models.Slots[i] == "helmet" {
			lists[models.Slots[i]] = EnchantValuesCalc(selection.ItemSlot.Head.Enchant.TypeU, models.Enchantments.Helmet)
		}
		if models.Slots[i] == "chest" {
			lists[models.Slots[i]] = EnchantValuesCalc(selection.ItemSlot.Chest.Enchant.TypeU, models.Enchantments.Chest)
		}
		if models.Slots[i] == "gloves" {
			lists[models.Slots[i]] = EnchantValuesCalc(selection.ItemSlot.Hands.Enchant.TypeU, models.Enchantments.Gloves)
		}
		if models.Slots[i] == "pants" {
			lists[models.Slots[i]] = EnchantValuesCalc(selection.ItemSlot.Pants.Enchant.TypeU, models.Enchantments.Pants)
		}
		if models.Slots[i] == "boots" {
			lists[models.Slots[i]] = EnchantValuesCalc(selection.ItemSlot.Foot.Enchant.TypeU, models.Enchantments.Boots)
		}
		if models.Slots[i] == "cloak" {
			lists[models.Slots[i]] = EnchantValuesCalc(selection.ItemSlot.Back.Enchant.TypeU, models.Enchantments.Cloak)
		}
	}
	return lists
}
func GetEnchatmentLists_Armor_TypeRare(selection models.Selection) map[string][]string {

	lists := make(map[string][]string)

	baseLists := GetEnchatmentLists_Armor_Base(selection)

	for _, slot := range models.Slots {
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
	for i := 0; i < len(models.Slots); i++ {

		if models.Slots[i] == "helmet" {
			lists[models.Slots[i]] = EnchantValuesCalc(selection.ItemSlot.Head.Enchant.TypeR, models.Enchantments.Helmet)
		}
		if models.Slots[i] == "chest" {
			lists[models.Slots[i]] = EnchantValuesCalc(selection.ItemSlot.Chest.Enchant.TypeR, models.Enchantments.Chest)
		}
		if models.Slots[i] == "gloves" {
			lists[models.Slots[i]] = EnchantValuesCalc(selection.ItemSlot.Hands.Enchant.TypeR, models.Enchantments.Gloves)
		}
		if models.Slots[i] == "pants" {
			lists[models.Slots[i]] = EnchantValuesCalc(selection.ItemSlot.Pants.Enchant.TypeR, models.Enchantments.Pants)
		}
		if models.Slots[i] == "boots" {
			lists[models.Slots[i]] = EnchantValuesCalc(selection.ItemSlot.Foot.Enchant.TypeR, models.Enchantments.Boots)
		}
		if models.Slots[i] == "cloak" {
			lists[models.Slots[i]] = EnchantValuesCalc(selection.ItemSlot.Back.Enchant.TypeR, models.Enchantments.Cloak)
		}
	}
	return lists
}
func GetEnchatmentLists_Armor_TypeEpic(selection models.Selection) map[string][]string {

	lists := make(map[string][]string)

	baseLists := GetEnchatmentLists_Armor_Base(selection)

	for _, slot := range models.Slots {
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
	for i := 0; i < len(models.Slots); i++ {

		if models.Slots[i] == "helmet" {
			lists[models.Slots[i]] = EnchantValuesCalc(selection.ItemSlot.Head.Enchant.TypeE, models.Enchantments.Helmet)
		}
		if models.Slots[i] == "chest" {
			lists[models.Slots[i]] = EnchantValuesCalc(selection.ItemSlot.Chest.Enchant.TypeE, models.Enchantments.Chest)
		}
		if models.Slots[i] == "gloves" {
			lists[models.Slots[i]] = EnchantValuesCalc(selection.ItemSlot.Hands.Enchant.TypeE, models.Enchantments.Gloves)
		}
		if models.Slots[i] == "pants" {
			lists[models.Slots[i]] = EnchantValuesCalc(selection.ItemSlot.Pants.Enchant.TypeE, models.Enchantments.Pants)
		}
		if models.Slots[i] == "boots" {
			lists[models.Slots[i]] = EnchantValuesCalc(selection.ItemSlot.Foot.Enchant.TypeE, models.Enchantments.Boots)
		}
		if models.Slots[i] == "cloak" {
			lists[models.Slots[i]] = EnchantValuesCalc(selection.ItemSlot.Back.Enchant.TypeE, models.Enchantments.Cloak)
		}
	}
	return lists
}
func GetEnchatmentLists_Armor_TypeLegend(selection models.Selection) map[string][]string {

	lists := make(map[string][]string)

	baseLists := GetEnchatmentLists_Armor_Base(selection)

	for _, slot := range models.Slots {
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
	for i := 0; i < len(models.Slots); i++ {

		if models.Slots[i] == "helmet" {
			lists[models.Slots[i]] = EnchantValuesCalc(selection.ItemSlot.Head.Enchant.TypeL, models.Enchantments.Helmet)
		}
		if models.Slots[i] == "chest" {
			lists[models.Slots[i]] = EnchantValuesCalc(selection.ItemSlot.Chest.Enchant.TypeL, models.Enchantments.Chest)
		}
		if models.Slots[i] == "gloves" {
			lists[models.Slots[i]] = EnchantValuesCalc(selection.ItemSlot.Hands.Enchant.TypeL, models.Enchantments.Gloves)
		}
		if models.Slots[i] == "pants" {
			lists[models.Slots[i]] = EnchantValuesCalc(selection.ItemSlot.Pants.Enchant.TypeL, models.Enchantments.Pants)
		}
		if models.Slots[i] == "boots" {
			lists[models.Slots[i]] = EnchantValuesCalc(selection.ItemSlot.Foot.Enchant.TypeL, models.Enchantments.Boots)
		}
		if models.Slots[i] == "cloak" {
			lists[models.Slots[i]] = EnchantValuesCalc(selection.ItemSlot.Back.Enchant.TypeL, models.Enchantments.Cloak)
		}
	}
	return lists
}
func GetEnchatmentLists_Armor_TypeUnique(selection models.Selection) map[string][]string {

	lists := make(map[string][]string)

	baseLists := GetEnchatmentLists_Armor_Base(selection)

	for _, slot := range models.Slots {
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
	for i := 0; i < len(models.Slots); i++ {

		if models.Slots[i] == "helmet" {
			lists[models.Slots[i]] = EnchantValuesCalc(selection.ItemSlot.Head.Enchant.TypeQ, models.Enchantments.Helmet)
		}
		if models.Slots[i] == "chest" {
			lists[models.Slots[i]] = EnchantValuesCalc(selection.ItemSlot.Chest.Enchant.TypeQ, models.Enchantments.Chest)
		}
		if models.Slots[i] == "gloves" {
			lists[models.Slots[i]] = EnchantValuesCalc(selection.ItemSlot.Hands.Enchant.TypeQ, models.Enchantments.Gloves)
		}
		if models.Slots[i] == "pants" {
			lists[models.Slots[i]] = EnchantValuesCalc(selection.ItemSlot.Pants.Enchant.TypeQ, models.Enchantments.Pants)
		}
		if models.Slots[i] == "boots" {
			lists[models.Slots[i]] = EnchantValuesCalc(selection.ItemSlot.Foot.Enchant.TypeQ, models.Enchantments.Boots)
		}
		if models.Slots[i] == "cloak" {
			lists[models.Slots[i]] = EnchantValuesCalc(selection.ItemSlot.Back.Enchant.TypeQ, models.Enchantments.Cloak)
		}
	}
	return lists
}

/////////\\\\\\\\\ --------------> ENCHCANTMENT LISTS ACCESSORY <------------------ //////////\\\\\\\\\

func GetEnchatmentLists_Accessory_Base(selection models.Selection) map[string][]string {

	lists := map[string][]string{}
	for i := 0; i < len(models.Slotsacc); i++ {
		if models.Slotsacc[i] == "necklace" {
			lists[models.Slotsacc[i]] = EnchantBaseAttribExeptionAcc(EnchamentbySlot(models.Enchantments.Necklace), *ItemsByNameAccesory(selection.ItemSlot.Necklace.Name))
		}
		if models.Slotsacc[i] == "ring" {
			lists[models.Slotsacc[i]] = EnchantBaseAttribExeptionAcc(EnchamentbySlot(models.Enchantments.Ring), *ItemsByNameAccesory(selection.ItemSlot.RingOne.Name))
		}
		if models.Slotsacc[i] == "ringtwo" {
			lists[models.Slotsacc[i]] = EnchantBaseAttribExeptionAcc(EnchamentbySlot(models.Enchantments.Ring), *ItemsByNameAccesory(selection.ItemSlot.RingTwo.Name))
		}
	}
	return lists

}
func GetEnchantmentLists_Accessory_ValuesUncommon(selection models.Selection) map[string][]float32 {

	lists := map[string][]float32{}
	for i := 0; i < len(models.Slotsacc); i++ {
		//
		if models.Slotsacc[i] == "necklace" {
			lists[models.Slotsacc[i]] = EnchantValuesCalc(selection.ItemSlot.Necklace.Enchant.TypeU, models.Enchantments.Necklace)
		}
		if models.Slotsacc[i] == "ring" {
			lists[models.Slotsacc[i]] = EnchantValuesCalc(selection.ItemSlot.RingOne.Enchant.TypeU, models.Enchantments.Ring)
		}
		if models.Slotsacc[i] == "ringtwo" {
			lists[models.Slotsacc[i]] = EnchantValuesCalc(selection.ItemSlot.RingTwo.Enchant.TypeU, models.Enchantments.Ring)
		}
	}
	return lists
}
func GetEnchantmentLists_Accessory_TypeRare(selection models.Selection) map[string][]string {

	lists := make(map[string][]string)

	baseLists := GetEnchatmentLists_Accessory_Base(selection)

	for _, slot := range models.Slotsacc {
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
	for i := 0; i < len(models.Slotsacc); i++ {

		if models.Slotsacc[i] == "necklace" {
			lists[models.Slotsacc[i]] = EnchantValuesCalc(selection.ItemSlot.Necklace.Enchant.TypeR, models.Enchantments.Necklace)
		}
		if models.Slotsacc[i] == "ring" {
			lists[models.Slotsacc[i]] = EnchantValuesCalc(selection.ItemSlot.RingOne.Enchant.TypeR, models.Enchantments.Ring)
		}
		if models.Slotsacc[i] == "ringtwo" {
			lists[models.Slotsacc[i]] = EnchantValuesCalc(selection.ItemSlot.RingTwo.Enchant.TypeR, models.Enchantments.Ring)
		}
	}
	return lists
}
func GetEnchantmentLists_Accessory_TypeEpic(selection models.Selection) map[string][]string {

	lists := make(map[string][]string)

	baseLists := GetEnchatmentLists_Accessory_Base(selection)

	for _, slot := range models.Slotsacc {
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
	for i := 0; i < len(models.Slotsacc); i++ {

		if models.Slotsacc[i] == "necklace" {
			lists[models.Slotsacc[i]] = EnchantValuesCalc(selection.ItemSlot.Necklace.Enchant.TypeE, models.Enchantments.Necklace)
		}
		if models.Slotsacc[i] == "ring" {
			lists[models.Slotsacc[i]] = EnchantValuesCalc(selection.ItemSlot.RingOne.Enchant.TypeE, models.Enchantments.Ring)
		}
		if models.Slotsacc[i] == "ringtwo" {
			lists[models.Slotsacc[i]] = EnchantValuesCalc(selection.ItemSlot.RingTwo.Enchant.TypeE, models.Enchantments.Ring)
		}
	}
	return lists
}
func GetEnchantmentLists_Accessory_TypeLegend(selection models.Selection) map[string][]string {

	lists := make(map[string][]string)

	baseLists := GetEnchatmentLists_Accessory_Base(selection)

	for _, slot := range models.Slotsacc {
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
	for i := 0; i < len(models.Slotsacc); i++ {

		if models.Slotsacc[i] == "necklace" {
			lists[models.Slotsacc[i]] = EnchantValuesCalc(selection.ItemSlot.Necklace.Enchant.TypeL, models.Enchantments.Necklace)
		}
		if models.Slotsacc[i] == "ring" {
			lists[models.Slotsacc[i]] = EnchantValuesCalc(selection.ItemSlot.RingOne.Enchant.TypeL, models.Enchantments.Ring)
		}
		if models.Slotsacc[i] == "ringtwo" {
			lists[models.Slotsacc[i]] = EnchantValuesCalc(selection.ItemSlot.RingTwo.Enchant.TypeL, models.Enchantments.Ring)
		}
	}
	return lists
}
func GetEnchantmentLists_Accessory_TypeUnique(selection models.Selection) map[string][]string {

	lists := make(map[string][]string)

	baseLists := GetEnchatmentLists_Accessory_Base(selection)

	for _, slot := range models.Slotsacc {
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
	for i := 0; i < len(models.Slotsacc); i++ {

		if models.Slotsacc[i] == "necklace" {
			lists[models.Slotsacc[i]] = EnchantValuesCalc(selection.ItemSlot.Necklace.Enchant.TypeQ, models.Enchantments.Necklace)
		}
		if models.Slotsacc[i] == "ring" {
			lists[models.Slotsacc[i]] = EnchantValuesCalc(selection.ItemSlot.RingOne.Enchant.TypeQ, models.Enchantments.Ring)
		}
		if models.Slotsacc[i] == "ringtwo" {
			lists[models.Slotsacc[i]] = EnchantValuesCalc(selection.ItemSlot.RingTwo.Enchant.TypeQ, models.Enchantments.Ring)
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

	for i := 0; i < len(models.Slotsweapon); i++ {
		itemName := selection.ItemSlot.WeaponOne.Name
		itemNameTwo := selection.ItemSlot.WeaponTwo.Name
		if models.Slotsweapon[i] == "pwo" {
			if ItemsByNameWeapon(itemName).HandType == "One Handed" && ItemsByNameWeapon(itemName).SlotType == "Main Hand" {
				lists[models.Slotsweapon[i]] = EnchamentbySlot(models.Enchantments.WeaponOneHand)
			} else if ItemsByNameWeapon(itemName).HandType == "Two Handed" && ItemsByNameWeapon(itemName).SlotType == "Main Hand" {
				lists[models.Slotsweapon[i]] = EnchamentbySlot(models.Enchantments.PhysicalTwoHand)
			}
		}
		if models.Slotsweapon[i] == "pwt" {
			itemName := selection.ItemSlot.WeaponTwo.Name
			if ItemsByNameWeapon(itemNameTwo).SlotType == "Off Hand" && ItemsByNameWeapon(itemName).HandType == "One Handed" {
				lists[models.Slotsweapon[i]] = EnchamentbySlot(models.Enchantments.WeaponOneHand)
			}
			//println(ItemsByNameWeapon(c.Query("item" + models.Slotsweapon[i])).SlotType)
		}
		/*
			if models.Slotsweapon[i] == "swo" {
				if ItemsByNameWeapon(c.Query("item"+models.Slotsweapon[i])).SlotType == "Main Hand" {
					lists[models.Slotsweapon[i]] = EnchamentbySlot(models.Enchantments.WeaponOneHand)
				}
			}
			if models.Slotsweapon[i] == "swt" {
				if ItemsByNameWeapon(c.Query("item"+models.Slotsweapon[i])).SlotType == "Off Hand" {
					lists[models.Slotsweapon[i]] = EnchamentbySlot(models.Enchantments.WeaponOneHand)
				}
			}
		*/
	}
	return lists

}
func GetEnchantmentLists_Weapon_ValuesUncommon(selection models.Selection) map[string][]float32 {

	lists := map[string][]float32{}
	for i := 0; i < len(models.Slotsweapon); i++ {

		if models.Slotsweapon[i] == "pwo" {
			lists[models.Slotsweapon[i]] = EnchantValuesCalc(selection.ItemSlot.WeaponOne.Enchant.TypeU, models.Enchantments.WeaponOneHand)
		}

		if models.Slotsweapon[i] == "pwt" {
			lists[models.Slotsweapon[i]] = EnchantValuesCalc(selection.ItemSlot.WeaponTwo.Enchant.TypeU, models.Enchantments.WeaponOneHand)
		}
		/*
			if models.Slotsweapon[i] == "swo" {
				lists[models.Slotsweapon[i]] = EnchantValuesCalc(Query, models.Enchantments.WeaponOneHand)
			}
			if models.Slotsweapon[i] == "swt" {
				lists[models.Slotsweapon[i]] = EnchantValuesCalc(Query, models.Enchantments.WeaponOneHand)
			}
		*/
	}
	return lists
}

func GetEnchantmentLists_Weapon_TypeRare(selection models.Selection) map[string][]string {

	lists := make(map[string][]string)
	baseLists := GetEnchantmentLists_Weapon_Base(selection)

	for _, slot := range models.Slotsweapon {
		previousSelections := []string{
			selection.ItemSlot.WeaponOne.Enchant.TypeU, selection.ItemSlot.WeaponTwo.Enchant.TypeU,
		}
		lists[slot] = EnchantTypeExeption(baseLists[slot], previousSelections)
	}
	return lists

}

func GetEnchantmentLists_Weapon_ValuesRare(selection models.Selection) map[string][]float32 {
	//models.Slotsweapon := []string{"pwo", "pwt", "swo", "swt"}

	lists := make(map[string][]float32)
	for i := 0; i < len(models.Slotsweapon); i++ {
		//Query := c.Query("enchantment_" + models.Slotsweapon[i] + "type2")
		if models.Slotsweapon[i] == "pwo" {
			lists[models.Slotsweapon[i]] = EnchantValuesCalc(selection.ItemSlot.WeaponOne.Enchant.TypeR, models.Enchantments.WeaponOneHand)
		}
		if models.Slotsweapon[i] == "pwt" {
			lists[models.Slotsweapon[i]] = EnchantValuesCalc(selection.ItemSlot.WeaponTwo.Enchant.TypeR, models.Enchantments.WeaponOneHand)
		}
		/*
			if models.Slotsweapon[i] == "swo" {
				lists[models.Slotsweapon[i]] = EnchantValuesCalc(Query, models.Enchantments.WeaponOneHand)
			}
			if models.Slotsweapon[i] == "swt" {
				lists[models.Slotsweapon[i]] = EnchantValuesCalc(Query, models.Enchantments.WeaponOneHand)
			}
		*/

	}
	return lists
}

func GetEnchantmentLists_Weapon_TypeEpic(selection models.Selection) map[string][]string {

	lists := make(map[string][]string)

	baseLists := GetEnchantmentLists_Weapon_Base(selection)

	for _, slot := range models.Slotsweapon {
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
	for i := 0; i < len(models.Slotsweapon); i++ {
		//Query := c.Query("enchantment_" + models.Slotsweapon[i] + "type3")
		if models.Slotsweapon[i] == "pwo" {
			lists[models.Slotsweapon[i]] = EnchantValuesCalc(selection.ItemSlot.WeaponOne.Enchant.TypeE, models.Enchantments.WeaponOneHand)
		}

		if models.Slots[i] == "pwt" {
			lists[models.Slots[i]] = EnchantValuesCalc(selection.ItemSlot.WeaponTwo.Enchant.TypeE, models.Enchantments.WeaponOneHand)
		}
		/*
			if models.Slots[i] == "swo" {
				lists[models.Slots[i]] = EnchantValuesCalc(Query, models.Enchantments.SecondaryWeaponOneHand)
			}
			if models.Slots[i] == "swt" {
				lists[models.Slots[i]] = EnchantValuesCalc(Query, models.Enchantments.SecondaryWeaponTwoHand)
			}
		*/
	}
	return lists
}
func GetEnchantmentLists_Weapon_TypeLegend(selection models.Selection) map[string][]string {

	lists := make(map[string][]string)

	baseLists := GetEnchantmentLists_Weapon_Base(selection)

	for _, slot := range models.Slotsweapon {
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
	for i := 0; i < len(models.Slotsweapon); i++ {

		if models.Slotsweapon[i] == "pwo" {
			lists[models.Slotsweapon[i]] = EnchantValuesCalc(selection.ItemSlot.WeaponOne.Enchant.TypeL, models.Enchantments.WeaponOneHand)
		}

		if models.Slots[i] == "pwt" {
			lists[models.Slots[i]] = EnchantValuesCalc(selection.ItemSlot.WeaponTwo.Enchant.TypeL, models.Enchantments.WeaponOneHand)
		}
		/*
			if models.Slots[i] == "swo" {
				lists[models.Slots[i]] = EnchantValuesCalc(Query, models.Enchantments.SecondaryWeaponOneHand)
			}
			if models.Slots[i] == "swt" {
				lists[models.Slots[i]] = EnchantValuesCalc(Query, models.Enchantments.SecondaryWeaponTwoHand)
			}
		*/

	}
	return lists
}
func GetEnchantmentLists_Weapon_TypeUnique(selection models.Selection) map[string][]string {

	lists := make(map[string][]string)

	baseLists := GetEnchantmentLists_Weapon_Base(selection)

	for _, slot := range models.Slotsweapon {
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
	for i := 0; i < len(models.Slotsweapon); i++ {

		if models.Slotsweapon[i] == "pwo" {
			lists[models.Slotsweapon[i]] = EnchantValuesCalc(selection.ItemSlot.WeaponOne.Enchant.TypeQ, models.Enchantments.WeaponOneHand)
		}

		if models.Slots[i] == "pwt" {
			lists[models.Slots[i]] = EnchantValuesCalc(selection.ItemSlot.WeaponTwo.Enchant.TypeQ, models.Enchantments.WeaponOneHand)
		}
		/*
			if models.Slots[i] == "swo" {
				lists[models.Slots[i]] = EnchantValuesCalc(Query, models.Enchantments.SecondaryWeaponOneHand)
			}
			if models.Slots[i] == "swt" {
				lists[models.Slots[i]] = EnchantValuesCalc(Query, models.Enchantments.SecondaryWeaponTwoHand)
			}
		*/

	}
	return lists
}
