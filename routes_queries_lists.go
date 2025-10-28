package main

//import "fmt"

//"net/http"
//"strconv"

//type amor_manager struct {
//	slotarmor   []string
//	slot
//}

/*
// get item lists

	func GetItemLists_Armor_Json(class string) map[string][]string {
		lists := map[string][]string{}
		for i := 0; i < len(slotsm); i++ {
			lists[slotsm[i]] = ItemsBySlotType_Json(class, slotsm[i])
		}
		return lists
	}

func GetItemLists_Accesory_Json() map[string][]string {

		lists := map[string][]string{}
		for i := 0; i < len(slotsmacc); i++ {
			lists[slotsmacc[i]] = AccessoryBySlotType_Json(slotsmacc[i])
		}
		return lists
	}

// get rating lists

	func GetRatingLists_Armor(c *gin.Context) map[string][]int {
		ratings := map[string][]int{}

		for i := 0; i < len(slots); i++ {
			itemName := c.Query("item" + slots[i])
			rarityIndex, _ := strconv.Atoi((c.Query("rarityselect_" + slots[i])))

			item := ItemsByNameArmor(itemName)
			ratings[slots[i]] = CompleteArrayInt(item.ArmorRatings[rarityIndex])
		}

		return ratings
	}

	func GetItemLists_Weapon_Json(class string) map[string][]string {
		lists := map[string][]string{}

		for i := 0; i < len(slotshand); i++ {
			lists[slotshand[i]] = WeaponsBySlotType_Json(class, slotshand[i])
		}

		return lists
	}

func GetRatingLists_Weapon(c *gin.Context) map[string][]int {

		ratings := map[string][]int{}
		for i := 0; i < len(slotsweapon); i++ {
			itemname := c.Query("item" + slotsweapon[i])
			rarityindex, _ := strconv.Atoi((c.Query("rarityselect_" + slotsweapon[i])))

			item := ItemsByNameWeapon(itemname)
			ratings[slotsweapon[i]] = CompleteArrayInt(item.DamageRatings[rarityindex])

		}

		return ratings
	}

// /////\\\\\\ --------------> ENCHCANTMENT LISTS ARMOR <------------------///////////\\\\\\\\\
*/
func GetEnchatmentLists_Armor_Base(selection Selection) map[string][]string {

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
func GetEnchatmentLists_Armor_ValuesUncommon(selection Selection) map[string][]float32 {

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
func GetEnchatmentLists_Armor_TypeRare(selection Selection) map[string][]string {

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
func GetEnchatmentLists_Armor_ValuesRare(selection Selection) map[string][]float32 {

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
func GetEnchatmentLists_Armor_TypeEpic(selection Selection) map[string][]string {

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
func GetEnchatmentLists_Armor_ValuesEpic(selection Selection) map[string][]float32 {

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
func GetEnchatmentLists_Armor_TypeLegend(selection Selection) map[string][]string {

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
func GetEnchatmentLists_Armor_ValuesLegend(selection Selection) map[string][]float32 {

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
func GetEnchatmentLists_Armor_TypeUnique(selection Selection) map[string][]string {

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
func GetEnchatmentLists_Armor_ValuesUnique(selection Selection) map[string][]float32 {

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

func GetEnchatmentLists_Accessory_Base(selection Selection) map[string][]string {

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
func GetEnchantmentLists_Accessory_ValuesUncommon(selection Selection) map[string][]float32 {

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
func GetEnchantmentLists_Accessory_TypeRare(selection Selection) map[string][]string {

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
func GetEnchantmentLists_Accessory_ValuesRare(selection Selection) map[string][]float32 {

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
func GetEnchantmentLists_Accessory_TypeEpic(selection Selection) map[string][]string {

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
func GetEnchantmentLists_Accessory_ValuesEpic(selection Selection) map[string][]float32 {

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
func GetEnchantmentLists_Accessory_TypeLegend(selection Selection) map[string][]string {

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
func GetEnchantmentLists_Accessory_ValuesLegend(selection Selection) map[string][]float32 {

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
func GetEnchantmentLists_Accessory_TypeUnique(selection Selection) map[string][]string {

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
func GetEnchantmentLists_Accessory_ValuesUnique(selection Selection) map[string][]float32 {

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

/*

/////////\\\\\\\\\ --------------> ENCHCANTMENT LISTS WEAPON <------------------ //////////\\\\\\\\\

func GetEnchantmentLists_Weapon_Base(c *gin.Context) map[string][]string {

	lists := map[string][]string{}
	for i := 0; i < len(slotsweapon); i++ {
		if slotsweapon[i] == "pwo" {
			if ItemsByNameWeapon(c.Query("item"+slotsweapon[i])).SlotType == "Main Hand" {
				lists[slotsweapon[i]] = EnchamentbySlot(Enchantments.WeaponOneHand)
			}
			if ItemsByNameWeapon(c.Query("item"+slotsweapon[i])).HandType == "Two Handed" {
				lists[slotsweapon[i]] = EnchamentbySlot(Enchantments.PhysicalTwoHand)
			}
		}
		if slotsweapon[i] == "pwt" {
			if ItemsByNameWeapon(c.Query("item"+slotsweapon[i])).SlotType == "Off Hand" {
				lists[slotsweapon[i]] = EnchamentbySlot(Enchantments.WeaponOneHand)
			}
			println(ItemsByNameWeapon(c.Query("item" + slotsweapon[i])).SlotType)
		}
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

	}
	return lists

}
func GetEnchantmentLists_Weapon_ValuesUncommon(c *gin.Context) map[string][]float32 {

	lists := map[string][]float32{}
	for i := 0; i < len(slotsweapon); i++ {
		Query := c.Query("enchantment_" + slotsweapon[i] + "type")
		if slotsweapon[i] == "pwo" {
			lists[slotsweapon[i]] = EnchantValuesCalc(Query, Enchantments.WeaponOneHand)
		}

		if slotsweapon[i] == "pwt" {
			lists[slotsweapon[i]] = EnchantValuesCalc(Query, Enchantments.WeaponOneHand)
		}
		if slotsweapon[i] == "swo" {
			lists[slotsweapon[i]] = EnchantValuesCalc(Query, Enchantments.WeaponOneHand)
		}
		if slotsweapon[i] == "swt" {
			lists[slotsweapon[i]] = EnchantValuesCalc(Query, Enchantments.WeaponOneHand)
		}

	}
	return lists
}
func GetEnchantmentLists_Weapon_TypeRare(c *gin.Context) map[string][]string {

	lists := make(map[string][]string)

	baseLists := GetEnchantmentLists_Weapon_Base(c)

	for _, slot := range slotsweapon {
		// Pass all previous selections up to Epic
		previousSelections := []string{
			c.Query("enchantment_" + slot + "type"), // Uncommon
		}
		lists[slot] = EnchantTypeExeption(baseLists[slot], previousSelections)
	}

	return lists

}
func GetEnchantmentLists_Weapon_ValuesRare(c *gin.Context) map[string][]float32 {
	slotsweapon := []string{"pwo", "pwt", "swo", "swt"}
	lists := make(map[string][]float32)
	for i := 0; i < len(slotsweapon); i++ {
		Query := c.Query("enchantment_" + slotsweapon[i] + "type2")
		if slotsweapon[i] == "pwo" {
			lists[slotsweapon[i]] = EnchantValuesCalc(Query, Enchantments.WeaponOneHand)
		}
		if slotsweapon[i] == "pwt" {
			lists[slotsweapon[i]] = EnchantValuesCalc(Query, Enchantments.WeaponOneHand)
		}
		if slotsweapon[i] == "swo" {
			lists[slotsweapon[i]] = EnchantValuesCalc(Query, Enchantments.WeaponOneHand)
		}
		if slotsweapon[i] == "swt" {
			lists[slotsweapon[i]] = EnchantValuesCalc(Query, Enchantments.WeaponOneHand)
		}

	}
	return lists
}
func GetEnchantmentLists_Weapon_TypeEpic(c *gin.Context) map[string][]string {

	lists := make(map[string][]string)

	baseLists := GetEnchantmentLists_Weapon_Base(c)

	for _, slot := range slotsweapon {
		// Pass all previous selections up to Epic
		previousSelections := []string{
			c.Query("enchantment_" + slot + "type"),  // Uncommon
			c.Query("enchantment_" + slot + "type2"), // Rare
		}
		lists[slot] = EnchantTypeExeption(baseLists[slot], previousSelections)
	}

	return lists

}
func GetEnchantmentLists_Weapon_ValuesEpic(c *gin.Context) map[string][]float32 {

	lists := make(map[string][]float32)
	for i := 0; i < len(slotsweapon); i++ {
		Query := c.Query("enchantment_" + slotsweapon[i] + "type3")
		if slotsweapon[i] == "pwo" {
			lists[slotsweapon[i]] = EnchantValuesCalc(Query, Enchantments.WeaponOneHand)
		}

			if slots[i] == "pwt" {
				lists[slots[i]] = EnchantValuesCalc(Query, Enchantments.PrimaryWeaponTwoHand)
			}
			if slots[i] == "swo" {
				lists[slots[i]] = EnchantValuesCalc(Query, Enchantments.SecondaryWeaponOneHand)
			}
			if slots[i] == "swt" {
				lists[slots[i]] = EnchantValuesCalc(Query, Enchantments.SecondaryWeaponTwoHand)
			}

	}
	return lists
}
func GetEnchantmentLists_Weapon_TypeLegend(c *gin.Context) map[string][]string {

	lists := make(map[string][]string)

	baseLists := GetEnchantmentLists_Weapon_Base(c)

	for _, slot := range slotsweapon {
		// Pass all previous selections up to Epic
		previousSelections := []string{
			c.Query("enchantment_" + slot + "type"),  // Uncommon
			c.Query("enchantment_" + slot + "type2"), // Rare
			c.Query("enchantment_" + slot + "type3"), // Epic
		}
		lists[slot] = EnchantTypeExeption(baseLists[slot], previousSelections)
	}

	return lists

}
func GetEnchantmentLists_Weapon_ValuesLegend(c *gin.Context) map[string][]float32 {

	lists := make(map[string][]float32)
	for i := 0; i < len(slotsweapon); i++ {
		Query := c.Query("enchantment_" + slotsweapon[i] + "type4")
		if slotsweapon[i] == "pwo" {
			lists[slotsweapon[i]] = EnchantValuesCalc(Query, Enchantments.WeaponOneHand)
		}

			if slots[i] == "pwt" {
				lists[slots[i]] = EnchantValuesCalc(Query, Enchantments.PrimaryWeaponTwoHand)
			}
			if slots[i] == "swo" {
				lists[slots[i]] = EnchantValuesCalc(Query, Enchantments.SecondaryWeaponOneHand)
			}
			if slots[i] == "swt" {
				lists[slots[i]] = EnchantValuesCalc(Query, Enchantments.SecondaryWeaponTwoHand)
			}

	}
	return lists
}
func GetEnchantmentLists_Weapon_TypeUnique(c *gin.Context) map[string][]string {

	lists := make(map[string][]string)

	baseLists := GetEnchantmentLists_Weapon_Base(c)

	for _, slot := range slotsweapon {
		// Pass all previous selections up to Epic
		previousSelections := []string{
			c.Query("enchantment_" + slot + "type"),  // Uncommon
			c.Query("enchantment_" + slot + "type2"), // Rare
			c.Query("enchantment_" + slot + "type3"), // Epic
			c.Query("enchantment_" + slot + "type4"), // Legend
		}
		lists[slot] = EnchantTypeExeption(baseLists[slot], previousSelections)
	}

	return lists

}
func GetEnchantmentLists_Weapon_ValuesUnique(c *gin.Context) map[string][]float32 {

	lists := make(map[string][]float32)
	for i := 0; i < len(slotsweapon); i++ {
		Query := c.Query("enchantment_" + slotsweapon[i] + "type5")
		if slotsweapon[i] == "pwo" {
			lists[slotsweapon[i]] = EnchantValuesCalc(Query, Enchantments.WeaponOneHand)
		}

			if slots[i] == "pwt" {
				lists[slots[i]] = EnchantValuesCalc(Query, Enchantments.PrimaryWeaponTwoHand)
			}
			if slots[i] == "swo" {
				lists[slots[i]] = EnchantValuesCalc(Query, Enchantments.SecondaryWeaponOneHand)
			}
			if slots[i] == "swt" {
				lists[slots[i]] = EnchantValuesCalc(Query, Enchantments.SecondaryWeaponTwoHand)
			}

	}
	return lists
}

*/
