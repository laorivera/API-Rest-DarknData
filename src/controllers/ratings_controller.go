package controllers

import (
	"builder/src/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// ////////\\\\\\\\\ ------->   RATING LISTS HANDLER <------- //////////\\\\\\\\\

func Helmet_RatingList_Handler(c *gin.Context) {
	var selection models.Selection
	c.BindJSON(&selection)

	list := Im.ArmorsByName(selection)
	rarity, _ := strconv.Atoi(selection.ItemSlot.Head.Rarity)

	var result = []int{}

	for i := 0; i < len(list); i++ {
		if list[i].Name == selection.ItemSlot.Head.Name {
			result = list[i].ArmorRatings[rarity]
		}
	}
	helmetRatingList := CompleteArrayInt(result)
	c.JSON(http.StatusOK, gin.H{
		"list": helmetRatingList},
	)
}

func Chest_RatingList_Handler(c *gin.Context) {
	var selection models.Selection
	c.BindJSON(&selection)

	list := Im.ArmorsByName(selection)
	rarity, _ := strconv.Atoi(selection.ItemSlot.Chest.Rarity)
	var result = []int{}

	for i := 0; i < len(list); i++ {
		if list[i].Name == selection.ItemSlot.Chest.Name {
			result = list[i].ArmorRatings[rarity]
		}
	}
	chestRatingList := CompleteArrayInt(result)
	c.JSON(http.StatusOK, gin.H{
		"list": chestRatingList},
	)
}

func Gloves_RatingList_Handler(c *gin.Context) {
	var selection models.Selection
	c.BindJSON(&selection)

	list := Im.ArmorsByName(selection)
	rarity, _ := strconv.Atoi(selection.ItemSlot.Hands.Rarity)
	var result = []int{}

	for i := 0; i < len(list); i++ {
		if list[i].Name == selection.ItemSlot.Hands.Name {
			result = list[i].ArmorRatings[rarity]
		}
	}
	glovesRatingList := CompleteArrayInt(result)

	c.JSON(http.StatusOK, gin.H{
		"list": glovesRatingList},
	)
}

func Pants_RatingList_Handler(c *gin.Context) {
	var selection models.Selection
	c.BindJSON(&selection)

	list := Im.ArmorsByName(selection)
	rarity, _ := strconv.Atoi(selection.ItemSlot.Pants.Rarity)
	var result = []int{}

	for i := 0; i < len(list); i++ {
		if list[i].Name == selection.ItemSlot.Pants.Name {
			result = list[i].ArmorRatings[rarity]
		}
	}
	pantsRatingList := CompleteArrayInt(result)
	c.JSON(http.StatusOK, gin.H{
		"list": pantsRatingList},
	)
}

func Boots_RatingList_Handler(c *gin.Context) {
	var selection models.Selection
	c.BindJSON(&selection)

	list := Im.ArmorsByName(selection)
	rarity, _ := strconv.Atoi(selection.ItemSlot.Foot.Rarity)
	var result = []int{}

	for i := 0; i < len(list); i++ {
		if list[i].Name == selection.ItemSlot.Foot.Name {
			result = list[i].ArmorRatings[rarity]
		}
	}
	bootsRatingList := CompleteArrayInt(result)
	c.JSON(http.StatusOK, gin.H{
		"list": bootsRatingList},
	)
}

func Cloak_RatingList_Handler(c *gin.Context) {
	var selection models.Selection
	c.BindJSON(&selection)

	list := Im.ArmorsByName(selection)
	rarity, _ := strconv.Atoi(selection.ItemSlot.Back.Rarity)
	var result = []int{}

	for i := 0; i < len(list); i++ {
		if list[i].Name == selection.ItemSlot.Back.Name {
			result = list[i].ArmorRatings[rarity]
		}
	}
	cloakRatingList := CompleteArrayInt(result)
	c.JSON(http.StatusOK, gin.H{
		"list": cloakRatingList},
	)
}

func Pwo_RatingList_Handler(c *gin.Context) {

	var selection models.Selection
	c.BindJSON(&selection)

	list := Im.WeaponsByName(selection)
	rarity, _ := strconv.Atoi(selection.ItemSlot.WeaponOne.Rarity)
	var result = []int{}

	for i := 0; i < len(list); i++ {
		if list[i].Name == selection.ItemSlot.WeaponOne.Name {
			result = list[i].DamageRatings[rarity]
		}
	}

	primaryWeaponRatingList := CompleteArrayInt(result)

	c.JSON(http.StatusOK, gin.H{
		"list": primaryWeaponRatingList},
	)
}

func Pwt_RatingList_Handler(c *gin.Context) {

	var selection models.Selection
	c.BindJSON(&selection)

	list := Im.WeaponsByName(selection)
	rarity, _ := strconv.Atoi(selection.ItemSlot.WeaponTwo.Rarity)
	var result = []int{}

	for i := 0; i < len(list); i++ {
		if list[i].Name == selection.ItemSlot.WeaponTwo.Name {
			result = list[i].DamageRatings[rarity]
		}
	}

	primaryoffhandWeaponRatingList := CompleteArrayInt(result)

	c.JSON(http.StatusOK, gin.H{
		"list": primaryoffhandWeaponRatingList},
	)
}
