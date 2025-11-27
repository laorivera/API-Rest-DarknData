package controllers

import (
	"builder/src/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// ////////\\\\\\\\\ ------->  LISTS ITEMS HANDLER <------- //////////\\\\\\\\\
var Im = NewItemManager()

func itemDisplayHandler(c *gin.Context) {
	var selection models.Selection
	c.BindJSON(&selection)

	itemArmor := Im.ArmorsByNameD(selection)
	itemAccesory := Im.AccesoriesByNameD(selection)
	itemWeapon := Im.WeaponsByNameD(selection)

	c.JSON(http.StatusOK, gin.H{

		"itemdata":       itemArmor,
		"itemdataacc":    itemAccesory,
		"itemdataweapon": itemWeapon,
	})
}

func Helmet_List_Handler(c *gin.Context) {
	var selection models.Selection
	c.BindJSON(&selection)

	helmetList := Im.ArmorsBySlot("Head", selection.Class)
	imageHelmet := ImageLocation("Head", helmetList)

	c.JSON(http.StatusOK, gin.H{
		"list": imageHelmet},
	)
}

func Chest_List_Handler(c *gin.Context) {
	var selection models.Selection
	c.BindJSON(&selection)

	chestList := Im.ArmorsBySlot("Chest", selection.Class)
	imageChest := ImageLocation("Chest", chestList)

	c.JSON(http.StatusOK, gin.H{
		"list": imageChest},
	)
}

func Gloves_List_Handler(c *gin.Context) {
	var selection models.Selection
	c.BindJSON(&selection)

	glovesList := Im.ArmorsBySlot("Hands", selection.Class)
	imageGloves := ImageLocation("Hands", glovesList)

	c.JSON(http.StatusOK, gin.H{
		"list": imageGloves},
	)
}

func Pants_List_Handler(c *gin.Context) {
	var selection models.Selection
	c.BindJSON(&selection)

	pantsList := Im.ArmorsBySlot("Legs", selection.Class)
	imagePants := ImageLocation("Legs", pantsList)

	c.JSON(http.StatusOK, gin.H{
		"list": imagePants,
	})

}

func Boots_List_Handler(c *gin.Context) {
	var selection models.Selection
	c.BindJSON(&selection)

	bootsList := Im.ArmorsBySlot("Foot", selection.Class)
	imageBoots := ImageLocation("Foot", bootsList)

	c.JSON(http.StatusOK, gin.H{
		"list": imageBoots},
	)
}

func Cloak_List_Handler(c *gin.Context) {
	var selection models.Selection
	c.BindJSON(&selection)

	cloakList := Im.ArmorsBySlot("Back", selection.Class)
	imageCloak := ImageLocation("Back", cloakList)

	c.JSON(http.StatusOK, gin.H{
		"list": imageCloak},
	)
}

func Necklace_List_Handler(c *gin.Context) {
	var selection models.Selection
	c.BindJSON(&selection)

	necklaceList := Im.AccesoryBySlot("Necklace")
	imageNecklace := ImageLocation("Necklace", necklaceList)

	c.JSON(http.StatusOK, gin.H{
		"list": imageNecklace},
	)
}

func Ring_List_Handler(c *gin.Context) {

	ringList := Im.AccesoryBySlot("Ring")
	imageRing := ImageLocation("Ring", ringList)

	c.JSON(http.StatusOK, gin.H{
		"list": imageRing},
	)
}

func Pwo_List_Handler(c *gin.Context) {
	var selection models.Selection
	c.BindJSON(&selection)

	primaryWeaponList := Im.WeaponsBySlot("Main Hand", selection.Class)
	imagePwo := ImageLocation("Main Hand", primaryWeaponList)

	c.JSON(http.StatusOK, gin.H{
		"list": imagePwo},
	)
}

func Pwt_List_Handler(c *gin.Context) {
	var selection models.Selection
	c.BindJSON(&selection)

	primaryOffhandWeaponList := Im.WeaponsBySlot("Off Hand", selection.Class)
	imagePwt := ImageLocation("Off Hand", primaryOffhandWeaponList)

	c.JSON(http.StatusOK, gin.H{
		"list": imagePwt},
	)
}
