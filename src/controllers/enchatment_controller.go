package controllers

import (
	"builder/src/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Helmet_EnchantmentList_Handler(c *gin.Context) {

	var selection models.Selection
	c.BindJSON(&selection)

	//fmt.Println(selection)

	//listU := GetAvailableHelmetEnchantments(selection)
	helmet_EnchantmentNameU := GetEnchatmentLists_Armor_Base(selection)["helmet"]
	helmet_EnchantmentNameR := GetEnchatmentLists_Armor_TypeRare(selection)["helmet"]
	helmet_EnchantmentNameE := GetEnchatmentLists_Armor_TypeEpic(selection)["helmet"]
	helmet_EnchantmentNameL := GetEnchatmentLists_Armor_TypeLegend(selection)["helmet"]
	helmet_EnchantmentNameQ := GetEnchatmentLists_Armor_TypeUnique(selection)["helmet"]

	helmet_EnchantmentValuesU := GetEnchatmentLists_Armor_ValuesUncommon(selection)["helmet"]
	helmet_EnchantmentValuesR := GetEnchatmentLists_Armor_ValuesRare(selection)["helmet"]
	helmet_EnchantmentValuesE := GetEnchatmentLists_Armor_ValuesEpic(selection)["helmet"]
	helmet_EnchantmentValuesL := GetEnchatmentLists_Armor_ValuesLegend(selection)["helmet"]
	helmet_EnchantmentValuesQ := GetEnchatmentLists_Armor_ValuesUnique(selection)["helmet"]

	c.JSON(http.StatusOK, gin.H{
		"listname_uncommon":  helmet_EnchantmentNameU,
		"listvalue_uncommon": helmet_EnchantmentValuesU,
		"listname_rare":      helmet_EnchantmentNameR,
		"listvalue_rare":     helmet_EnchantmentValuesR,
		"listname_epic":      helmet_EnchantmentNameE,
		"listvalue_epic":     helmet_EnchantmentValuesE,
		"listname_legend":    helmet_EnchantmentNameL,
		"listvalue_legend":   helmet_EnchantmentValuesL,
		"listname_unique":    helmet_EnchantmentNameQ,
		"listvalue_unique":   helmet_EnchantmentValuesQ,
	})

}

func Chest_EnchantmentList_Handler(c *gin.Context) {
	var selection models.Selection
	c.BindJSON(&selection)
	//fmt.Println(selection)

	chest_EnchantmentNameU := GetEnchatmentLists_Armor_Base(selection)["chest"]
	chest_EnchantmentNameR := GetEnchatmentLists_Armor_TypeRare(selection)["chest"]
	chest_EnchantmentNameE := GetEnchatmentLists_Armor_TypeEpic(selection)["chest"]
	chest_EnchantmentNameL := GetEnchatmentLists_Armor_TypeLegend(selection)["chest"]
	chest_EnchantmentNameQ := GetEnchatmentLists_Armor_TypeUnique(selection)["chest"]

	chest_EnchantmentValuesU := GetEnchatmentLists_Armor_ValuesUncommon(selection)["chest"]
	chest_EnchantmentValuesR := GetEnchatmentLists_Armor_ValuesRare(selection)["chest"]
	chest_EnchantmentValuesE := GetEnchatmentLists_Armor_ValuesEpic(selection)["chest"]
	chest_EnchantmentValuesL := GetEnchatmentLists_Armor_ValuesLegend(selection)["chest"]
	chest_EnchantmentValuesQ := GetEnchatmentLists_Armor_ValuesUnique(selection)["chest"]

	c.JSON(http.StatusOK, gin.H{
		"listname_uncommon":  chest_EnchantmentNameU,
		"listvalue_uncommon": chest_EnchantmentValuesU,
		"listname_rare":      chest_EnchantmentNameR,
		"listvalue_rare":     chest_EnchantmentValuesR,
		"listname_epic":      chest_EnchantmentNameE,
		"listvalue_epic":     chest_EnchantmentValuesE,
		"listname_legend":    chest_EnchantmentNameL,
		"listvalue_legend":   chest_EnchantmentValuesL,
		"listname_unique":    chest_EnchantmentNameQ,
		"listvalue_unique":   chest_EnchantmentValuesQ,
	})

}

func Gloves_EnchantmentList_Handler(c *gin.Context) {
	var selection models.Selection
	c.BindJSON(&selection)

	gloves_EnchantmentNameU := GetEnchatmentLists_Armor_Base(selection)["gloves"]
	gloves_EnchantmentNameR := GetEnchatmentLists_Armor_TypeRare(selection)["gloves"]
	gloves_EnchantmentNameE := GetEnchatmentLists_Armor_TypeEpic(selection)["gloves"]
	gloves_EnchantmentNameL := GetEnchatmentLists_Armor_TypeLegend(selection)["gloves"]
	gloves_EnchantmentNameQ := GetEnchatmentLists_Armor_TypeUnique(selection)["gloves"]

	gloves_EnchantmentValuesU := GetEnchatmentLists_Armor_ValuesUncommon(selection)["gloves"]
	gloves_EnchantmentValuesR := GetEnchatmentLists_Armor_ValuesRare(selection)["gloves"]
	gloves_EnchantmentValuesE := GetEnchatmentLists_Armor_ValuesEpic(selection)["gloves"]
	gloves_EnchantmentValuesL := GetEnchatmentLists_Armor_ValuesLegend(selection)["gloves"]
	gloves_EnchantmentValuesQ := GetEnchatmentLists_Armor_ValuesUnique(selection)["gloves"]

	c.JSON(http.StatusOK, gin.H{
		"listname_uncommon":  gloves_EnchantmentNameU,
		"listvalue_uncommon": gloves_EnchantmentValuesU,
		"listname_rare":      gloves_EnchantmentNameR,
		"listvalue_rare":     gloves_EnchantmentValuesR,
		"listname_epic":      gloves_EnchantmentNameE,
		"listvalue_epic":     gloves_EnchantmentValuesE,
		"listname_legend":    gloves_EnchantmentNameL,
		"listvalue_legend":   gloves_EnchantmentValuesL,
		"listname_unique":    gloves_EnchantmentNameQ,
		"listvalue_unique":   gloves_EnchantmentValuesQ,
	})

}

func Pants_EnchantmentList_Handler(c *gin.Context) {
	var selection models.Selection
	c.BindJSON(&selection)

	pants_EnchantmentNameU := GetEnchatmentLists_Armor_Base(selection)["pants"]
	pants_EnchantmentNameR := GetEnchatmentLists_Armor_TypeRare(selection)["pants"]
	pants_EnchantmentNameE := GetEnchatmentLists_Armor_TypeEpic(selection)["pants"]
	pants_EnchantmentNameL := GetEnchatmentLists_Armor_TypeLegend(selection)["pants"]
	pants_EnchantmentNameQ := GetEnchatmentLists_Armor_TypeUnique(selection)["pants"]

	pants_EnchantmentValuesU := GetEnchatmentLists_Armor_ValuesUncommon(selection)["pants"]
	pants_EnchantmentValuesR := GetEnchatmentLists_Armor_ValuesRare(selection)["pants"]
	pants_EnchantmentValuesE := GetEnchatmentLists_Armor_ValuesEpic(selection)["pants"]
	pants_EnchantmentValuesL := GetEnchatmentLists_Armor_ValuesLegend(selection)["pants"]
	pants_EnchantmentValuesQ := GetEnchatmentLists_Armor_ValuesUnique(selection)["pants"]

	c.JSON(http.StatusOK, gin.H{
		"listname_uncommon":  pants_EnchantmentNameU,
		"listvalue_uncommon": pants_EnchantmentValuesU,
		"listname_rare":      pants_EnchantmentNameR,
		"listvalue_rare":     pants_EnchantmentValuesR,
		"listname_epic":      pants_EnchantmentNameE,
		"listvalue_epic":     pants_EnchantmentValuesE,
		"listname_legend":    pants_EnchantmentNameL,
		"listvalue_legend":   pants_EnchantmentValuesL,
		"listname_unique":    pants_EnchantmentNameQ,
		"listvalue_unique":   pants_EnchantmentValuesQ,
	})

}

func Boots_EnchantmentList_Handler(c *gin.Context) {
	var selection models.Selection
	c.BindJSON(&selection)

	boots_EnchantmentNameU := GetEnchatmentLists_Armor_Base(selection)["boots"]
	boots_EnchantmentNameR := GetEnchatmentLists_Armor_TypeRare(selection)["boots"]
	boots_EnchantmentNameE := GetEnchatmentLists_Armor_TypeEpic(selection)["boots"]
	boots_EnchantmentNameL := GetEnchatmentLists_Armor_TypeLegend(selection)["boots"]
	boots_EnchantmentNameQ := GetEnchatmentLists_Armor_TypeUnique(selection)["boots"]

	boots_EnchantmentValuesU := GetEnchatmentLists_Armor_ValuesUncommon(selection)["boots"]
	boots_EnchantmentValuesR := GetEnchatmentLists_Armor_ValuesRare(selection)["boots"]
	boots_EnchantmentValuesE := GetEnchatmentLists_Armor_ValuesEpic(selection)["boots"]
	boots_EnchantmentValuesL := GetEnchatmentLists_Armor_ValuesLegend(selection)["boots"]
	boots_EnchantmentValuesQ := GetEnchatmentLists_Armor_ValuesUnique(selection)["boots"]

	c.JSON(http.StatusOK, gin.H{
		"listname_uncommon":  boots_EnchantmentNameU,
		"listvalue_uncommon": boots_EnchantmentValuesU,
		"listname_rare":      boots_EnchantmentNameR,
		"listvalue_rare":     boots_EnchantmentValuesR,
		"listname_epic":      boots_EnchantmentNameE,
		"listvalue_epic":     boots_EnchantmentValuesE,
		"listname_legend":    boots_EnchantmentNameL,
		"listvalue_legend":   boots_EnchantmentValuesL,
		"listname_unique":    boots_EnchantmentNameQ,
		"listvalue_unique":   boots_EnchantmentValuesQ,
	})
}

func Cloak_EnchantmentList_Handler(c *gin.Context) {
	var selection models.Selection
	c.BindJSON(&selection)

	cloak_EnchantmentNameU := GetEnchatmentLists_Armor_Base(selection)["cloak"]
	cloak_EnchantmentNameR := GetEnchatmentLists_Armor_TypeRare(selection)["cloak"]
	cloak_EnchantmentNameE := GetEnchatmentLists_Armor_TypeEpic(selection)["cloak"]
	cloak_EnchantmentNameL := GetEnchatmentLists_Armor_TypeLegend(selection)["cloak"]
	cloak_EnchantmentNameQ := GetEnchatmentLists_Armor_TypeUnique(selection)["cloak"]

	cloak_EnchantmentValuesU := GetEnchatmentLists_Armor_ValuesUncommon(selection)["cloak"]
	cloak_EnchantmentValuesR := GetEnchatmentLists_Armor_ValuesRare(selection)["cloak"]
	cloak_EnchantmentValuesE := GetEnchatmentLists_Armor_ValuesEpic(selection)["cloak"]
	cloak_EnchantmentValuesL := GetEnchatmentLists_Armor_ValuesLegend(selection)["cloak"]
	cloak_EnchantmentValuesQ := GetEnchatmentLists_Armor_ValuesUnique(selection)["cloak"]

	c.JSON(http.StatusOK, gin.H{
		"listname_uncommon":  cloak_EnchantmentNameU,
		"listvalue_uncommon": cloak_EnchantmentValuesU,
		"listname_rare":      cloak_EnchantmentNameR,
		"listvalue_rare":     cloak_EnchantmentValuesR,
		"listname_epic":      cloak_EnchantmentNameE,
		"listvalue_epic":     cloak_EnchantmentValuesE,
		"listname_legend":    cloak_EnchantmentNameL,
		"listvalue_legend":   cloak_EnchantmentValuesL,
		"listname_unique":    cloak_EnchantmentNameQ,
		"listvalue_unique":   cloak_EnchantmentValuesQ,
	})

}

func Necklace_EnchantmentList_Handler(c *gin.Context) {
	var selection models.Selection
	c.BindJSON(&selection)

	necklace_EnchantmentNameU := GetEnchatmentLists_Accessory_Base(selection)["necklace"]
	necklace_EnchantmentNameR := GetEnchantmentLists_Accessory_TypeRare(selection)["necklace"]
	necklace_EnchantmentNameE := GetEnchantmentLists_Accessory_TypeEpic(selection)["necklace"]
	necklace_EnchantmentNameL := GetEnchantmentLists_Accessory_TypeLegend(selection)["necklace"]
	necklace_EnchantmentNameQ := GetEnchantmentLists_Accessory_TypeUnique(selection)["necklace"]

	necklace_EnchantmentValuesU := GetEnchantmentLists_Accessory_ValuesUncommon(selection)["necklace"]
	necklace_EnchantmentValuesR := GetEnchantmentLists_Accessory_ValuesRare(selection)["necklace"]
	necklace_EnchantmentValuesE := GetEnchantmentLists_Accessory_ValuesEpic(selection)["necklace"]
	necklace_EnchantmentValuesL := GetEnchantmentLists_Accessory_ValuesLegend(selection)["necklace"]
	necklace_EnchantmentValuesQ := GetEnchantmentLists_Accessory_ValuesUnique(selection)["necklace"]

	c.JSON(http.StatusOK, gin.H{
		"listname_uncommon":  necklace_EnchantmentNameU,
		"listvalue_uncommon": necklace_EnchantmentValuesU,
		"listname_rare":      necklace_EnchantmentNameR,
		"listvalue_rare":     necklace_EnchantmentValuesR,
		"listname_epic":      necklace_EnchantmentNameE,
		"listvalue_epic":     necklace_EnchantmentValuesE,
		"listname_legend":    necklace_EnchantmentNameL,
		"listvalue_legend":   necklace_EnchantmentValuesL,
		"listname_unique":    necklace_EnchantmentNameQ,
		"listvalue_unique":   necklace_EnchantmentValuesQ,
	})

}

func Ring_EnchantmentList_Handler(c *gin.Context) {
	var selection models.Selection
	c.BindJSON(&selection)

	ring_EnchantmentNameU := GetEnchatmentLists_Accessory_Base(selection)["ring"]
	ring_EnchantmentNameR := GetEnchantmentLists_Accessory_TypeRare(selection)["ring"]
	ring_EnchantmentNameE := GetEnchantmentLists_Accessory_TypeEpic(selection)["ring"]
	ring_EnchantmentNameL := GetEnchantmentLists_Accessory_TypeLegend(selection)["ring"]
	ring_EnchantmentNameQ := GetEnchantmentLists_Accessory_TypeUnique(selection)["ring"]

	ring_EnchantmentValuesU := GetEnchantmentLists_Accessory_ValuesUncommon(selection)["ring"]
	ring_EnchantmentValuesR := GetEnchantmentLists_Accessory_ValuesRare(selection)["ring"]
	ring_EnchantmentValuesE := GetEnchantmentLists_Accessory_ValuesEpic(selection)["ring"]
	ring_EnchantmentValuesL := GetEnchantmentLists_Accessory_ValuesLegend(selection)["ring"]
	ring_EnchantmentValuesQ := GetEnchantmentLists_Accessory_ValuesUnique(selection)["ring"]

	c.JSON(http.StatusOK, gin.H{
		"listname_uncommon":  ring_EnchantmentNameU,
		"listvalue_uncommon": ring_EnchantmentValuesU,
		"listname_rare":      ring_EnchantmentNameR,
		"listvalue_rare":     ring_EnchantmentValuesR,
		"listname_epic":      ring_EnchantmentNameE,
		"listvalue_epic":     ring_EnchantmentValuesE,
		"listname_legend":    ring_EnchantmentNameL,
		"listvalue_legend":   ring_EnchantmentValuesL,
		"listname_unique":    ring_EnchantmentNameQ,
		"listvalue_unique":   ring_EnchantmentValuesQ,
	})

}

func RingTwo_EnchantmentList_Handler(c *gin.Context) {
	var selection models.Selection
	c.BindJSON(&selection)

	ringTwo_EnchantmentNameU := GetEnchatmentLists_Accessory_Base(selection)["ringtwo"]
	ringTwo_EnchantmentNameR := GetEnchantmentLists_Accessory_TypeRare(selection)["ringtwo"]
	ringTwo_EnchantmentNameE := GetEnchantmentLists_Accessory_TypeEpic(selection)["ringtwo"]
	ringTwo_EnchantmentNameL := GetEnchantmentLists_Accessory_TypeLegend(selection)["ringtwo"]
	ringTwo_EnchantmentNameQ := GetEnchantmentLists_Accessory_TypeUnique(selection)["ringtwo"]

	ringTwo_EnchantmentValuesU := GetEnchantmentLists_Accessory_ValuesUncommon(selection)["ringtwo"]
	ringTwo_EnchantmentValuesR := GetEnchantmentLists_Accessory_ValuesRare(selection)["ringtwo"]
	ringTwo_EnchantmentValuesE := GetEnchantmentLists_Accessory_ValuesEpic(selection)["ringtwo"]
	ringTwo_EnchantmentValuesL := GetEnchantmentLists_Accessory_ValuesLegend(selection)["ringtwo"]
	ringTwo_EnchantmentValuesQ := GetEnchantmentLists_Accessory_ValuesUnique(selection)["ringtwo"]

	c.JSON(http.StatusOK, gin.H{
		"listname_uncommon":  ringTwo_EnchantmentNameU,
		"listvalue_uncommon": ringTwo_EnchantmentValuesU,
		"listname_rare":      ringTwo_EnchantmentNameR,
		"listvalue_rare":     ringTwo_EnchantmentValuesR,
		"listname_epic":      ringTwo_EnchantmentNameE,
		"listvalue_epic":     ringTwo_EnchantmentValuesE,
		"listname_legend":    ringTwo_EnchantmentNameL,
		"listvalue_legend":   ringTwo_EnchantmentValuesL,
		"listname_unique":    ringTwo_EnchantmentNameQ,
		"listvalue_unique":   ringTwo_EnchantmentValuesQ,
	})
}

func Pwo_EnchantmentList_Handler(c *gin.Context) {
	var selection models.Selection
	c.BindJSON(&selection)

	//primaryWeapon_EnchantmentName := map[string][]string{}
	//primaryWeapon_EnchantmentValues := map[string][]float32{}

	primaryWeapon_EnchantmentNameU := GetEnchantmentLists_Weapon_Base(selection)["pwo"]
	primaryWeapon_EnchantmentNameR := GetEnchantmentLists_Weapon_TypeRare(selection)["pwo"]
	primaryWeapon_EnchantmentNameE := GetEnchantmentLists_Weapon_TypeEpic(selection)["pwo"]
	primaryWeapon_EnchantmentNameL := GetEnchantmentLists_Weapon_TypeLegend(selection)["pwo"]
	primaryWeapon_EnchantmentNameQ := GetEnchantmentLists_Weapon_TypeUnique(selection)["pwo"]

	primaryWeapon_EnchantmentValuesU := GetEnchantmentLists_Weapon_ValuesUncommon(selection)["pwo"]
	primaryWeapon_EnchantmentValuesR := GetEnchantmentLists_Weapon_ValuesRare(selection)["pwo"]
	primaryWeapon_EnchantmentValuesE := GetEnchantmentLists_Weapon_ValuesEpic(selection)["pwo"]
	primaryWeapon_EnchantmentValuesL := GetEnchantmentLists_Weapon_ValuesLegend(selection)["pwo"]
	primaryWeapon_EnchantmentValuesQ := GetEnchantmentLists_Weapon_ValuesUnique(selection)["pwo"]

	c.JSON(http.StatusOK, gin.H{
		"listname_uncommon":  primaryWeapon_EnchantmentNameU,
		"listvalue_uncommon": primaryWeapon_EnchantmentValuesU,
		"listname_rare":      primaryWeapon_EnchantmentNameR,
		"listvalue_rare":     primaryWeapon_EnchantmentValuesR,
		"listname_epic":      primaryWeapon_EnchantmentNameE,
		"listvalue_epic":     primaryWeapon_EnchantmentValuesE,
		"listname_legend":    primaryWeapon_EnchantmentNameL,
		"listvalue_legend":   primaryWeapon_EnchantmentValuesL,
		"listname_unique":    primaryWeapon_EnchantmentNameQ,
		"listvalue_unique":   primaryWeapon_EnchantmentValuesQ,
	})

}

func Pwt_EnchantmentList_Handler(c *gin.Context) {
	var selection models.Selection
	c.BindJSON(&selection)

	//primaryOffhandWeapon_EnchantmentName := map[string][]string{}
	//primaryOffhandWeapon_EnchantmentValues := map[string][]float32{}

	primaryOffhandWeapon_EnchantmentNameU := GetEnchantmentLists_Weapon_Base(selection)["pwt"]
	primaryOffhandWeapon_EnchantmentNameR := GetEnchantmentLists_Weapon_TypeRare(selection)["pwt"]
	primaryOffhandWeapon_EnchantmentNameE := GetEnchantmentLists_Weapon_TypeEpic(selection)["pwt"]
	primaryOffhandWeapon_EnchantmentNameL := GetEnchantmentLists_Weapon_TypeLegend(selection)["pwt"]
	primaryOffhandWeapon_EnchantmentNameQ := GetEnchantmentLists_Weapon_TypeUnique(selection)["pwt"]

	primaryOffhandWeapon_EnchantmentValuesU := GetEnchantmentLists_Weapon_ValuesUncommon(selection)["pwt"]
	primaryOffhandWeapon_EnchantmentValuesR := GetEnchantmentLists_Weapon_ValuesRare(selection)["pwt"]
	primaryOffhandWeapon_EnchantmentValuesE := GetEnchantmentLists_Weapon_ValuesEpic(selection)["pwt"]
	primaryOffhandWeapon_EnchantmentValuesL := GetEnchantmentLists_Weapon_ValuesLegend(selection)["pwt"]
	primaryOffhandWeapon_EnchantmentValuesQ := GetEnchantmentLists_Weapon_ValuesUnique(selection)["pwt"]

	c.JSON(http.StatusOK, gin.H{
		"listname_uncommon":  primaryOffhandWeapon_EnchantmentNameU,
		"listvalue_uncommon": primaryOffhandWeapon_EnchantmentValuesU,
		"listname_rare":      primaryOffhandWeapon_EnchantmentNameR,
		"listvalue_rare":     primaryOffhandWeapon_EnchantmentValuesR,
		"listname_epic":      primaryOffhandWeapon_EnchantmentNameE,
		"listvalue_epic":     primaryOffhandWeapon_EnchantmentValuesE,
		"listname_legend":    primaryOffhandWeapon_EnchantmentNameL,
		"listvalue_legend":   primaryOffhandWeapon_EnchantmentValuesL,
		"listname_unique":    primaryOffhandWeapon_EnchantmentNameQ,
		"listvalue_unique":   primaryOffhandWeapon_EnchantmentValuesQ,
	})
}
