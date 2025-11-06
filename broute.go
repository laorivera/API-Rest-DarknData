package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

/*
func updateStatsHandler(c *gin.Context) {

	//class param
	class := c.Param("classSelection")
	classStatSelect := SelectClass(class)

	//queries
	raceSelectedStats := raceStats[GetSelectedRace(c)]
	raceSelectedComputed := raceComputed[GetSelectedRace(c)]

	itemsSelected_Armor := GetSelectedItems_Armor(c) //Items query selection Armor
	raritySelected_Armor := GetSelectedRarities_Armor(c)
	ratingSelected_Armor := GetSelectedRatings_Armor(c)

	itemsSelected_Accessory := GetSelectedItems_Accessory(c) //Items query selection Accessory
	raritySelected_Accessory := GetSelectedRarities_Accessory(c)

	itemSelected_Weapon := GetSelected_Weapons(c) //Items query selection Weapon
	raritySelected_Weapon := GetSelectedRarities_Weapons(c)
	ratingSelected_Weapon := GetSelectedRatings_Weapons(c)

	//Enchatment query selected Armor
	enchantmentSelected_ArmorUncommon := GetSelectedEnchantmentsBase_ArmorUncommon(c)
	enchantmentSelected_ArmorRare := GetSelectedEnchantmentsBase_ArmorRare(c)
	enchantmentSelected_ArmorEpic := GetSelectedEnchantmentsBase_ArmorEpic(c)
	enchantmentSelected_ArmorLegend := GetSelectedEnchantmentsBase_ArmorLegend(c)
	enchantmentSelected_ArmorUnique := GetSelectedEnchantmentsBase_ArmorUnique(c)

	enchantmentSelectedOther_ArmorUncommon := GetSelectedEnchantmentsOther_ArmorUncommon(c)
	enchantmentSelectedOther_ArmorRare := GetSelectedEnchantmentsOther_ArmorRare(c)
	enchantmentSelectedOther_ArmorEpic := GetSelectedEnchantmentsOther_ArmorEpic(c)
	enchantmentSelectedOther_ArmorLegend := GetSelectedEnchantmentsOther_ArmorLegend(c)
	enchantmentSelectedOther_ArmorUnique := GetSelectedEnchantmentsOther_ArmorUnique(c)

	// Enchatment query selected Accesory
	enchantmentSelected_AccessoryUncommon := GetSelectedEnchantmentsBase_AccessoryUncommon(c)
	enchantmentSelected_AccessoryRare := GetSelectedEnchantmentsBase_AccessoryRare(c)
	enchantmentSelected_AccessoryEpic := GetSelectedEnchantmentsBase_AccessoryEpic(c)
	enchantmentSelected_AccessoryLegend := GetSelectedEnchantmentsBase_AccessoryLegend(c)
	enchantmentSelected_AccessoryUnique := GetSelectedEnchantmentsBase_AccessoryUnique(c)

	enchantmentSelectedOther_AccessoryUncommon := GetSelectedEnchantmentsOther_AccessoryUncommon(c)
	enchantmentSelectedOther_AccessoryRare := GetSelectedEnchantmentsOther_AccessoryRare(c)
	enchantmentSelectedOther_AccessoryEpic := GetSelectedEnchantmentsOther_AccessoryEpic(c)
	enchantmentSelectedOther_AccessoryLegend := GetSelectedEnchantmentsOther_AccessoryLegend(c)
	enchantmentSelectedOther_AccessoryUnique := GetSelectedEnchantmentsOther_AccessoryUnique(c)

	computedStatsEnchant_Other_Armor := Computed_Stats{}
	computedStatsEnchant_Other_Armor = computedStatsEnchant_Other_Armor.AddEnchant(EnchantComputedOthers(enchantmentSelectedOther_ArmorUncommon),
		EnchantComputedOthers(enchantmentSelectedOther_ArmorRare), EnchantComputedOthers(enchantmentSelectedOther_ArmorEpic),
		EnchantComputedOthers(enchantmentSelectedOther_ArmorLegend), EnchantComputedOthers(enchantmentSelectedOther_ArmorUnique))

	computedStatsEnchant_Other_Accesory := Computed_Stats{}
	computedStatsEnchant_Other_Accesory = computedStatsEnchant_Other_Accesory.AddEnchant(EnchantComputedOthers(enchantmentSelectedOther_AccessoryUncommon),
		EnchantComputedOthers(enchantmentSelectedOther_AccessoryRare), EnchantComputedOthers(enchantmentSelectedOther_AccessoryEpic),
		EnchantComputedOthers(enchantmentSelectedOther_AccessoryLegend), EnchantComputedOthers(enchantmentSelectedOther_AccessoryUnique))

	updatedBaseEchant_StatsArmor := SetItemStats(classStatSelect, itemsSelected_Armor, raritySelected_Armor)

	updatedBaseEchant_StatsArmor = updatedBaseEchant_StatsArmor.AddStats(setEnchantStats(enchantmentSelected_ArmorUncommon), setEnchantStats(enchantmentSelected_ArmorRare),
		setEnchantStats(enchantmentSelected_ArmorEpic), setEnchantStats(enchantmentSelected_ArmorLegend), setEnchantStats(enchantmentSelected_ArmorUnique))

	updatedBaseEnchant_StatsAccessory := SetItemStatsAccessory(characterHolder, itemsSelected_Accessory, raritySelected_Accessory)
	updatedBaseEnchant_StatsAccessory = updatedBaseEnchant_StatsAccessory.AddStats(setEnchantStats(enchantmentSelected_AccessoryUncommon), setEnchantStats(enchantmentSelected_AccessoryRare),
		setEnchantStats(enchantmentSelected_AccessoryEpic), setEnchantStats(enchantmentSelected_AccessoryLegend), setEnchantStats(enchantmentSelected_AccessoryUnique))

	//updatedStatsEnchant_Weapon := SetItemStatsWeapon(itemSelected_Weapon, raritySelected_Weapon)

	updatedTotalStats := Stats{}
	updatedTotalStats = updatedTotalStats.AddStats(updatedBaseEchant_StatsArmor, updatedBaseEnchant_StatsAccessory, raceSelectedStats)

	totalRating := RatingCalc(ratingSelected_Armor)

	totalBaseItem := BaseItemCalc(itemsSelected_Armor, raritySelected_Armor)

	totalSpeed := SpeedCalc(itemsSelected_Armor, raritySelected_Armor)

	//totalDamagerating := DamageRatingCalc(itemSelected_Weapon, raritySelected_Weapon, ratingSelected_Armor, enchantment

	computedStatsCurve := CalculateComputedValues(updatedTotalStats, totalRating, totalSpeed, computedStatsEnchant_Other_Armor, computedStatsEnchant_Other_Accesory, totalBaseItem, raceSelectedComputed)

	computedStatsTotal := ComputedTotal(computedStatsCurve, computedStatsEnchant_Other_Armor, computedStatsEnchant_Other_Accesory, totalBaseItem, raceSelectedComputed)

	totalWeaponDamage := WeaponDamageCalc(itemSelected_Weapon, raritySelected_Weapon, computedStatsTotal.PhysicalPowerBonus, ratingSelected_Weapon)

	c.JSON(http.StatusOK, gin.H{
		"stats":               updatedTotalStats,
		"computedstats":       computedStatsTotal,
		"computedstatsweapon": totalWeaponDamage,
	})

}
*/
func postHandler(c *gin.Context) {
	var selection Selection

	if err := c.BindJSON(&selection); err != nil {
		fmt.Println("JSON Bind Error:", err)
		return
	}
	fmt.Println(selection)

	//var Im = NewItemManager()
	var Sm = NewStatsManager()

	//sele := Im.ArmorsByName(selection)

	Sm.ItemsBaseStats(selection)
	Sm.BaseItemCalc(selection)
	Sm.TotalSpeedRating(selection)
	Sm.TotalArmorRating(selection)

	Sm.base = ProcessBaseEnchantments(selection).AddStats(Sm.base)

	resultother := ProcessOtherEnchantments(selection).AddEnchant(Sm.variable)

	fmt.Println(Sm.base, "1")
	//fmt.Println(Sm.totalrating, "2")
	//fmt.Println(Sm.totalspeed, "3")
	//fmt.Println(resultother, "4")
	//fmt.Println(Sm.variable, "5")
	totally := Calculate(Sm.base, Sm.totalrating, Sm.totalspeed, resultother, Sm.variable)
	total := totally.AddEnchant(resultother)

	//fmt.Println(total.Luck)
	//fmt.Println(total)

	c.JSON(http.StatusOK, gin.H{
		//"test":   Sm.base,
		//"compu":  Sm.variable,
		//"test2":  test,
		//"test3":  Sm.totalrating,
		//"test4":  Sm.totalspeed,
		"stats":               Sm.base,
		"computedstats":       total,
		"computedstatsweapon": Sm.weapon,
	})

}
