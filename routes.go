package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

//////////\\\\\\\\\ ------->  LISTS ITEMS HANDLER <------- //////////\\\\\\\\\

var Im = NewItemManager()

func itemDisplayHandler(c *gin.Context) {
	var selection Selection
	c.BindJSON(&selection)

	itemArmor := Im.ArmorsByName(selection)
	itemAccesory := Im.AccesoriesByName(selection)
	itemWeapon := Im.WeaponsByName(selection)

	c.JSON(http.StatusOK, gin.H{

		"itemdata":       itemArmor,
		"itemdataacc":    itemAccesory,
		"itemdataweapon": itemWeapon,
	})
}

func Helmet_List_Handler(c *gin.Context) {
	var selection Selection
	c.BindJSON(&selection)

	helmetList := Im.ArmorsBySlot("Head", selection.Class)
	imageHelmet := ImageLocation("Head", helmetList)

	c.JSON(http.StatusOK, gin.H{
		"list": imageHelmet},
	)
}

func Chest_List_Handler(c *gin.Context) {
	var selection Selection
	c.BindJSON(&selection)

	chestList := Im.ArmorsBySlot("Chest", selection.Class)
	imageChest := ImageLocation("Chest", chestList)

	c.JSON(http.StatusOK, gin.H{
		"list": imageChest},
	)
}

func Gloves_List_Handler(c *gin.Context) {
	var selection Selection
	c.BindJSON(&selection)

	glovesList := Im.ArmorsBySlot("Hands", selection.Class)
	imageGloves := ImageLocation("Hands", glovesList)

	c.JSON(http.StatusOK, gin.H{
		"list": imageGloves},
	)
}

func Pants_List_Handler(c *gin.Context) {
	var selection Selection
	c.BindJSON(&selection)

	pantsList := Im.ArmorsBySlot("Legs", selection.Class)
	imagePants := ImageLocation("Legs", pantsList)

	c.JSON(http.StatusOK, gin.H{
		"list": imagePants,
	})

}

func Boots_List_Handler(c *gin.Context) {
	var selection Selection
	c.BindJSON(&selection)

	bootsList := Im.ArmorsBySlot("Foot", selection.Class)
	imageBoots := ImageLocation("Foot", bootsList)

	c.JSON(http.StatusOK, gin.H{
		"list": imageBoots},
	)
}

func Cloak_List_Handler(c *gin.Context) {
	var selection Selection
	c.BindJSON(&selection)

	cloakList := Im.ArmorsBySlot("Back", selection.Class)
	imageCloak := ImageLocation("Back", cloakList)

	c.JSON(http.StatusOK, gin.H{
		"list": imageCloak},
	)
}

func Necklace_List_Handler(c *gin.Context) {

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
	var selection Selection
	c.BindJSON(&selection)

	primaryWeaponList := Im.WeaponsBySlot("Main Hand", selection.Class)
	imagePwo := ImageLocation("Main Hand", primaryWeaponList)

	c.JSON(http.StatusOK, gin.H{
		"list": imagePwo},
	)
}

func Pwt_List_Handler(c *gin.Context) {
	var selection Selection
	c.BindJSON(&selection)

	primaryOffhandWeaponList := Im.WeaponsBySlot("Off Hand", selection.Class)
	imagePwt := ImageLocation("Off Hand", primaryOffhandWeaponList)

	c.JSON(http.StatusOK, gin.H{
		"list": imagePwt},
	)
}

//////////\\\\\\\\\ ------->   RATING LISTS HANDLER <------- //////////\\\\\\\\\

func Rating_List(c *gin.Context) {

}

func Helmet_RatingList_Handler(c *gin.Context) {
	var selection Selection
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
	var selection Selection
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
	var selection Selection
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
	var selection Selection
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
	var selection Selection
	c.BindJSON(&selection)

	list := Im.ArmorsByName(selection)
	rarity, _ := strconv.Atoi(selection.ItemSlot.Foot.Rarity)
	var result = []int{}

	for i := 0; i < len(list); i++ {
		if list[i].Name == selection.ItemSlot.Foot.Id {
			result = list[i].ArmorRatings[rarity]
		}
	}
	bootsRatingList := CompleteArrayInt(result)
	c.JSON(http.StatusOK, gin.H{
		"list": bootsRatingList},
	)
}

func Cloak_RatingList_Handler(c *gin.Context) {
	var selection Selection
	c.BindJSON(&selection)

	list := Im.ArmorsByName(selection)
	rarity, _ := strconv.Atoi(selection.ItemSlot.Foot.Rarity)
	var result = []int{}

	for i := 0; i < len(list); i++ {
		if list[i].Name == selection.ItemSlot.Foot.Id {
			result = list[i].ArmorRatings[rarity]
		}
	}
	cloakRatingList := CompleteArrayInt(result)
	c.JSON(http.StatusOK, gin.H{
		"list": cloakRatingList},
	)
}

func Pwo_RatingList_Handler(c *gin.Context) {

	var selection Selection
	c.BindJSON(&selection)

	list := Im.WeaponsByName(selection)
	rarity, _ := strconv.Atoi(selection.ItemSlot.WeaponOne.Rarity)
	var result = []int{}

	for i := 0; i < len(list); i++ {
		if list[i].Name == selection.ItemSlot.WeaponOne.Id {
			result = list[i].DamageRatings[rarity]
		}
	}

	primaryWeaponRatingList := CompleteArrayInt(result)

	c.JSON(http.StatusOK, gin.H{
		"list": primaryWeaponRatingList},
	)
}

func Pwt_RatingList_Handler(c *gin.Context) {

	var selection Selection
	c.BindJSON(&selection)

	list := Im.WeaponsByName(selection)
	rarity, _ := strconv.Atoi(selection.ItemSlot.WeaponTwo.Rarity)
	var result = []int{}

	for i := 0; i < len(list); i++ {
		if list[i].Name == selection.ItemSlot.WeaponTwo.Id {
			result = list[i].DamageRatings[rarity]
		}
	}

	primaryoffhandWeaponRatingList := CompleteArrayInt(result)

	c.JSON(http.StatusOK, gin.H{
		"list": primaryoffhandWeaponRatingList},
	)
}

//////////\\\\\\\\\ ------->   ENCHANTMENT LISTS HANDLER   <------- //////////\\\\\\\\\

func Helmet_EnchantmentList_Handler(c *gin.Context) {

	var selection Selection
	c.BindJSON(&selection)

	fmt.Println(selection)

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

/*

func Chest_EnchantmentList_Handler(c *gin.Context) {
	chest_EnchantmentName := map[string][]string{}
	chest_EnchantmentValues := map[string][]float32{}

	chest_EnchantmentName["Uncommon"] = GetEnchatmentLists_Armor_Base(c)["chest"]
	chest_EnchantmentName["Rare"] = GetEnchatmentLists_Armor_TypeRare(c)["chest"]
	chest_EnchantmentName["Epic"] = GetEnchatmentLists_Armor_TypeEpic(c)["chest"]
	chest_EnchantmentName["Legend"] = GetEnchatmentLists_Armor_TypeLegend(c)["chest"]
	chest_EnchantmentName["Unique"] = GetEnchatmentLists_Armor_TypeUnique(c)["chest"]

	chest_EnchantmentValues["Uncommon"] = GetEnchatmentLists_Armor_ValuesUncommon(c)["chest"]
	chest_EnchantmentValues["Rare"] = GetEnchatmentLists_Armor_ValuesRare(c)["chest"]
	chest_EnchantmentValues["Epic"] = GetEnchatmentLists_Armor_ValuesEpic(c)["chest"]
	chest_EnchantmentValues["Legend"] = GetEnchatmentLists_Armor_ValuesLegend(c)["chest"]
	chest_EnchantmentValues["Unique"] = GetEnchatmentLists_Armor_ValuesUnique(c)["chest"]

	c.JSON(http.StatusOK, gin.H{
		"listname_uncommon":  chest_EnchantmentName["Uncommon"],
		"listvalue_uncommon": chest_EnchantmentValues["Uncommon"],
		"listname_rare":      chest_EnchantmentName["Rare"],
		"listvalue_rare":     chest_EnchantmentValues["Rare"],
		"listname_epic":      chest_EnchantmentName["Epic"],
		"listvalue_epic":     chest_EnchantmentValues["Epic"],
		"listname_legend":    chest_EnchantmentName["Legend"],
		"listvalue_legend":   chest_EnchantmentValues["Legend"],
		"listname_unique":    chest_EnchantmentName["Unique"],
		"listvalue_unique":   chest_EnchantmentValues["Unique"],
	})

}

func Gloves_EnchantmentList_Handler(c *gin.Context) {
	gloves_EnchantmentName := map[string][]string{}
	gloves_EnchantmentValues := map[string][]float32{}

	gloves_EnchantmentName["Uncommon"] = GetEnchatmentLists_Armor_Base(c)["gloves"]
	gloves_EnchantmentName["Rare"] = GetEnchatmentLists_Armor_TypeRare(c)["gloves"]
	gloves_EnchantmentName["Epic"] = GetEnchatmentLists_Armor_TypeEpic(c)["gloves"]
	gloves_EnchantmentName["Legend"] = GetEnchatmentLists_Armor_TypeLegend(c)["gloves"]
	gloves_EnchantmentName["Unique"] = GetEnchatmentLists_Armor_TypeUnique(c)["gloves"]

	gloves_EnchantmentValues["Uncommon"] = GetEnchatmentLists_Armor_ValuesUncommon(c)["gloves"]
	gloves_EnchantmentValues["Rare"] = GetEnchatmentLists_Armor_ValuesRare(c)["gloves"]
	gloves_EnchantmentValues["Epic"] = GetEnchatmentLists_Armor_ValuesEpic(c)["gloves"]
	gloves_EnchantmentValues["Legend"] = GetEnchatmentLists_Armor_ValuesLegend(c)["gloves"]
	gloves_EnchantmentValues["Unique"] = GetEnchatmentLists_Armor_ValuesUnique(c)["gloves"]

	c.JSON(http.StatusOK, gin.H{
		"listname_uncommon":  gloves_EnchantmentName["Uncommon"],
		"listvalue_uncommon": gloves_EnchantmentValues["Uncommon"],
		"listname_rare":      gloves_EnchantmentName["Rare"],
		"listvalue_rare":     gloves_EnchantmentValues["Rare"],
		"listname_epic":      gloves_EnchantmentName["Epic"],
		"listvalue_epic":     gloves_EnchantmentValues["Epic"],
		"listname_legend":    gloves_EnchantmentName["Legend"],
		"listvalue_legend":   gloves_EnchantmentValues["Legend"],
		"listname_unique":    gloves_EnchantmentName["Unique"],
		"listvalue_unique":   gloves_EnchantmentValues["Unique"],
	})

}

func Pants_EnchantmentList_Handler(c *gin.Context) {
	pants_EnchantmentName := map[string][]string{}
	pants_EnchantmentValues := map[string][]float32{}

	pants_EnchantmentName["Uncommon"] = GetEnchatmentLists_Armor_Base(c)["pants"]
	pants_EnchantmentName["Rare"] = GetEnchatmentLists_Armor_TypeRare(c)["pants"]
	pants_EnchantmentName["Epic"] = GetEnchatmentLists_Armor_TypeEpic(c)["pants"]
	pants_EnchantmentName["Legend"] = GetEnchatmentLists_Armor_TypeLegend(c)["pants"]
	pants_EnchantmentName["Unique"] = GetEnchatmentLists_Armor_TypeUnique(c)["pants"]

	pants_EnchantmentValues["Uncommon"] = GetEnchatmentLists_Armor_ValuesUncommon(c)["pants"]
	pants_EnchantmentValues["Rare"] = GetEnchatmentLists_Armor_ValuesRare(c)["pants"]
	pants_EnchantmentValues["Epic"] = GetEnchatmentLists_Armor_ValuesEpic(c)["pants"]
	pants_EnchantmentValues["Legend"] = GetEnchatmentLists_Armor_ValuesLegend(c)["pants"]
	pants_EnchantmentValues["Unique"] = GetEnchatmentLists_Armor_ValuesUnique(c)["pants"]

	c.JSON(http.StatusOK, gin.H{
		"listname_uncommon":  pants_EnchantmentName["Uncommon"],
		"listvalue_uncommon": pants_EnchantmentValues["Uncommon"],
		"listname_rare":      pants_EnchantmentName["Rare"],
		"listvalue_rare":     pants_EnchantmentValues["Rare"],
		"listname_epic":      pants_EnchantmentName["Epic"],
		"listvalue_epic":     pants_EnchantmentValues["Epic"],
		"listname_legend":    pants_EnchantmentName["Legend"],
		"listvalue_legend":   pants_EnchantmentValues["Legend"],
		"listname_unique":    pants_EnchantmentName["Unique"],
		"listvalue_unique":   pants_EnchantmentValues["Unique"],
	})

}

func Boots_EnchantmentList_Handler(c *gin.Context) {
	boots_EnchantmentName := map[string][]string{}
	boots_EnchantmentValues := map[string][]float32{}

	boots_EnchantmentName["Uncommon"] = GetEnchatmentLists_Armor_Base(c)["boots"]
	boots_EnchantmentName["Rare"] = GetEnchatmentLists_Armor_TypeRare(c)["boots"]
	boots_EnchantmentName["Epic"] = GetEnchatmentLists_Armor_TypeEpic(c)["boots"]
	boots_EnchantmentName["Legend"] = GetEnchatmentLists_Armor_TypeLegend(c)["boots"]
	boots_EnchantmentName["Unique"] = GetEnchatmentLists_Armor_TypeUnique(c)["boots"]

	boots_EnchantmentValues["Uncommon"] = GetEnchatmentLists_Armor_ValuesUncommon(c)["boots"]
	boots_EnchantmentValues["Rare"] = GetEnchatmentLists_Armor_ValuesRare(c)["boots"]
	boots_EnchantmentValues["Epic"] = GetEnchatmentLists_Armor_ValuesEpic(c)["boots"]
	boots_EnchantmentValues["Legend"] = GetEnchatmentLists_Armor_ValuesLegend(c)["boots"]
	boots_EnchantmentValues["Unique"] = GetEnchatmentLists_Armor_ValuesUnique(c)["boots"]

	c.JSON(http.StatusOK, gin.H{
		"listname_uncommon":  boots_EnchantmentName["Uncommon"],
		"listvalue_uncommon": boots_EnchantmentValues["Uncommon"],
		"listname_rare":      boots_EnchantmentName["Rare"],
		"listvalue_rare":     boots_EnchantmentValues["Rare"],
		"listname_epic":      boots_EnchantmentName["Epic"],
		"listvalue_epic":     boots_EnchantmentValues["Epic"],
		"listname_legend":    boots_EnchantmentName["Legend"],
		"listvalue_legend":   boots_EnchantmentValues["Legend"],
		"listname_unique":    boots_EnchantmentName["Unique"],
		"listvalue_unique":   boots_EnchantmentValues["Unique"],
	})
}

func Cloak_EnchantmentList_Handler(c *gin.Context) {
	cloak_EnchantmentName := map[string][]string{}
	cloak_EnchantmentValues := map[string][]float32{}

	cloak_EnchantmentName["Uncommon"] = GetEnchatmentLists_Armor_Base(c)["cloak"]
	cloak_EnchantmentName["Rare"] = GetEnchatmentLists_Armor_TypeRare(c)["cloak"]
	cloak_EnchantmentName["Epic"] = GetEnchatmentLists_Armor_TypeEpic(c)["cloak"]
	cloak_EnchantmentName["Legend"] = GetEnchatmentLists_Armor_TypeLegend(c)["cloak"]
	cloak_EnchantmentName["Unique"] = GetEnchatmentLists_Armor_TypeUnique(c)["cloak"]

	cloak_EnchantmentValues["Uncommon"] = GetEnchatmentLists_Armor_ValuesUncommon(c)["cloak"]
	cloak_EnchantmentValues["Rare"] = GetEnchatmentLists_Armor_ValuesRare(c)["cloak"]
	cloak_EnchantmentValues["Epic"] = GetEnchatmentLists_Armor_ValuesEpic(c)["cloak"]
	cloak_EnchantmentValues["Legend"] = GetEnchatmentLists_Armor_ValuesLegend(c)["cloak"]
	cloak_EnchantmentValues["Unique"] = GetEnchatmentLists_Armor_ValuesUnique(c)["cloak"]

	c.JSON(http.StatusOK, gin.H{
		"listname_uncommon":  cloak_EnchantmentName["Uncommon"],
		"listvalue_uncommon": cloak_EnchantmentValues["Uncommon"],
		"listname_rare":      cloak_EnchantmentName["Rare"],
		"listvalue_rare":     cloak_EnchantmentValues["Rare"],
		"listname_epic":      cloak_EnchantmentName["Epic"],
		"listvalue_epic":     cloak_EnchantmentValues["Epic"],
		"listname_legend":    cloak_EnchantmentName["Legend"],
		"listvalue_legend":   cloak_EnchantmentValues["Legend"],
		"listname_unique":    cloak_EnchantmentName["Unique"],
		"listvalue_unique":   cloak_EnchantmentValues["Unique"],
	})

}

func Necklace_EnchantmentList_Handler(c *gin.Context) {
	necklace_EnchantmentName := map[string][]string{}
	necklace_EnchantmentValues := map[string][]float32{}

	necklace_EnchantmentName["Uncommon"] = GetEnchatmentLists_Accessory_Base(c)["necklace"]

	necklace_EnchantmentName["Rare"] = GetEnchantmentLists_Accessory_TypeRare(c)["necklace"]
	necklace_EnchantmentName["Epic"] = GetEnchantmentLists_Accessory_TypeEpic(c)["necklace"]
	necklace_EnchantmentName["Legend"] = GetEnchantmentLists_Accessory_TypeLegend(c)["necklace"]
	necklace_EnchantmentName["Unique"] = GetEnchantmentLists_Accessory_TypeUnique(c)["necklace"]

	necklace_EnchantmentValues["Uncommon"] = GetEnchantmentLists_Accessory_ValuesUncommon(c)["necklace"]
	necklace_EnchantmentValues["Rare"] = GetEnchantmentLists_Accessory_ValuesRare(c)["necklace"]
	necklace_EnchantmentValues["Epic"] = GetEnchantmentLists_Accessory_ValuesEpic(c)["necklace"]
	necklace_EnchantmentValues["Legend"] = GetEnchantmentLists_Accessory_ValuesLegend(c)["necklace"]
	necklace_EnchantmentValues["Unique"] = GetEnchantmentLists_Accessory_ValuesUnique(c)["necklace"]

	c.JSON(http.StatusOK, gin.H{
		"listname_uncommon":  necklace_EnchantmentName["Uncommon"],
		"listvalue_uncommon": necklace_EnchantmentValues["Uncommon"],
		"listname_rare":      necklace_EnchantmentName["Rare"],
		"listvalue_rare":     necklace_EnchantmentValues["Rare"],
		"listname_epic":      necklace_EnchantmentName["Epic"],
		"listvalue_epic":     necklace_EnchantmentValues["Epic"],
		"listname_legend":    necklace_EnchantmentName["Legend"],
		"listvalue_legend":   necklace_EnchantmentValues["Legend"],
		"listname_unique":    necklace_EnchantmentName["Unique"],
		"listvalue_unique":   necklace_EnchantmentValues["Unique"],
	})

}

func Ring_EnchantmentList_Handler(c *gin.Context) {
	ring_EnchantmentName := map[string][]string{}
	ring_EnchantmentValues := map[string][]float32{}

	ring_EnchantmentName["Uncommon"] = GetEnchatmentLists_Accessory_Base(c)["ring"]
	ring_EnchantmentName["Rare"] = GetEnchantmentLists_Accessory_TypeRare(c)["ring"]
	ring_EnchantmentName["Epic"] = GetEnchantmentLists_Accessory_TypeEpic(c)["ring"]
	ring_EnchantmentName["Legend"] = GetEnchantmentLists_Accessory_TypeLegend(c)["ring"]
	ring_EnchantmentName["Unique"] = GetEnchantmentLists_Accessory_TypeUnique(c)["ring"]

	ring_EnchantmentValues["Uncommon"] = GetEnchantmentLists_Accessory_ValuesUncommon(c)["ring"]
	ring_EnchantmentValues["Rare"] = GetEnchantmentLists_Accessory_ValuesRare(c)["ring"]
	ring_EnchantmentValues["Epic"] = GetEnchantmentLists_Accessory_ValuesEpic(c)["ring"]
	ring_EnchantmentValues["Legend"] = GetEnchantmentLists_Accessory_ValuesLegend(c)["ring"]
	ring_EnchantmentValues["Unique"] = GetEnchantmentLists_Accessory_ValuesUnique(c)["ring"]

	c.JSON(http.StatusOK, gin.H{
		"listname_uncommon":  ring_EnchantmentName["Uncommon"],
		"listvalue_uncommon": ring_EnchantmentValues["Uncommon"],
		"listname_rare":      ring_EnchantmentName["Rare"],
		"listvalue_rare":     ring_EnchantmentValues["Rare"],
		"listname_epic":      ring_EnchantmentName["Epic"],
		"listvalue_epic":     ring_EnchantmentValues["Epic"],
		"listname_legend":    ring_EnchantmentName["Legend"],
		"listvalue_legend":   ring_EnchantmentValues["Legend"],
		"listname_unique":    ring_EnchantmentName["Unique"],
		"listvalue_unique":   ring_EnchantmentValues["Unique"],
	})

}

func RingTwo_EnchantmentList_Handler(c *gin.Context) {
	ringTwo_EnchantmentName := map[string][]string{}
	ringTwo_EnchantmentValues := map[string][]float32{}

	ringTwo_EnchantmentName["Uncommon"] = GetEnchatmentLists_Accessory_Base(c)["ringtwo"]
	ringTwo_EnchantmentName["Rare"] = GetEnchantmentLists_Accessory_TypeRare(c)["ringtwo"]
	ringTwo_EnchantmentName["Epic"] = GetEnchantmentLists_Accessory_TypeEpic(c)["ringtwo"]
	ringTwo_EnchantmentName["Legend"] = GetEnchantmentLists_Accessory_TypeLegend(c)["ringtwo"]
	ringTwo_EnchantmentName["Unique"] = GetEnchantmentLists_Accessory_TypeUnique(c)["ringtwo"]

	ringTwo_EnchantmentValues["Uncommon"] = GetEnchantmentLists_Accessory_ValuesUncommon(c)["ringtwo"]
	ringTwo_EnchantmentValues["Rare"] = GetEnchantmentLists_Accessory_ValuesRare(c)["ringtwo"]
	ringTwo_EnchantmentValues["Epic"] = GetEnchantmentLists_Accessory_ValuesEpic(c)["ringtwo"]
	ringTwo_EnchantmentValues["Legend"] = GetEnchantmentLists_Accessory_ValuesLegend(c)["ringtwo"]
	ringTwo_EnchantmentValues["Unique"] = GetEnchantmentLists_Accessory_ValuesUnique(c)["ringtwo"]

	c.JSON(http.StatusOK, gin.H{
		"listname_uncommon":  ringTwo_EnchantmentName["Uncommon"],
		"listvalue_uncommon": ringTwo_EnchantmentValues["Uncommon"],
		"listname_rare":      ringTwo_EnchantmentName["Rare"],
		"listvalue_rare":     ringTwo_EnchantmentValues["Rare"],
		"listname_epic":      ringTwo_EnchantmentName["Epic"],
		"listvalue_epic":     ringTwo_EnchantmentValues["Epic"],
		"listname_legend":    ringTwo_EnchantmentName["Legend"],
		"listvalue_legend":   ringTwo_EnchantmentValues["Legend"],
		"listname_unique":    ringTwo_EnchantmentName["Unique"],
		"listvalue_unique":   ringTwo_EnchantmentValues["Unique"],
	})
}

func Pwo_EnchantmentList_Handler(c *gin.Context) {
	primaryWeapon_EnchantmentName := map[string][]string{}
	primaryWeapon_EnchantmentValues := map[string][]float32{}

	primaryWeapon_EnchantmentName["Uncommon"] = GetEnchantmentLists_Weapon_Base(c)["pwo"]
	primaryWeapon_EnchantmentName["Rare"] = GetEnchantmentLists_Weapon_TypeRare(c)["pwo"]
	primaryWeapon_EnchantmentName["Epic"] = GetEnchantmentLists_Weapon_TypeEpic(c)["pwo"]
	primaryWeapon_EnchantmentName["Legend"] = GetEnchantmentLists_Weapon_TypeLegend(c)["pwo"]
	primaryWeapon_EnchantmentName["Unique"] = GetEnchantmentLists_Weapon_TypeUnique(c)["pwo"]

	primaryWeapon_EnchantmentValues["Uncommon"] = GetEnchantmentLists_Weapon_ValuesUncommon(c)["pwo"]
	primaryWeapon_EnchantmentValues["Rare"] = GetEnchantmentLists_Weapon_ValuesRare(c)["pwo"]
	primaryWeapon_EnchantmentValues["Epic"] = GetEnchantmentLists_Weapon_ValuesEpic(c)["pwo"]
	primaryWeapon_EnchantmentValues["Legend"] = GetEnchantmentLists_Weapon_ValuesLegend(c)["pwo"]
	primaryWeapon_EnchantmentValues["Unique"] = GetEnchantmentLists_Weapon_ValuesUnique(c)["pwo"]

	c.JSON(http.StatusOK, gin.H{
		"listname_uncommon":  primaryWeapon_EnchantmentName["Uncommon"],
		"listvalue_uncommon": primaryWeapon_EnchantmentValues["Uncommon"],
		"listname_rare":      primaryWeapon_EnchantmentName["Rare"],
		"listvalue_rare":     primaryWeapon_EnchantmentValues["Rare"],
		"listname_epic":      primaryWeapon_EnchantmentName["Epic"],
		"listvalue_epic":     primaryWeapon_EnchantmentValues["Epic"],
		"listname_legend":    primaryWeapon_EnchantmentName["Legend"],
		"listvalue_legend":   primaryWeapon_EnchantmentValues["Legend"],
		"listname_unique":    primaryWeapon_EnchantmentName["Unique"],
		"listvalue_unique":   primaryWeapon_EnchantmentValues["Unique"],
	})

}

func Pwt_EnchantmentList_Handler(c *gin.Context) {
	primaryOffhandWeapon_EnchantmentName := map[string][]string{}
	primaryOffhandWeapon_EnchantmentValues := map[string][]float32{}

	primaryOffhandWeapon_EnchantmentName["Uncommon"] = GetEnchantmentLists_Weapon_Base(c)["pwt"]
	primaryOffhandWeapon_EnchantmentName["Rare"] = GetEnchantmentLists_Weapon_TypeRare(c)["pwt"]
	primaryOffhandWeapon_EnchantmentName["Epic"] = GetEnchantmentLists_Weapon_TypeEpic(c)["pwt"]
	primaryOffhandWeapon_EnchantmentName["Legend"] = GetEnchantmentLists_Weapon_TypeLegend(c)["pwt"]
	primaryOffhandWeapon_EnchantmentName["Unique"] = GetEnchantmentLists_Weapon_TypeUnique(c)["pwt"]

	primaryOffhandWeapon_EnchantmentValues["Uncommon"] = GetEnchantmentLists_Weapon_ValuesUncommon(c)["pwt"]
	primaryOffhandWeapon_EnchantmentValues["Rare"] = GetEnchantmentLists_Weapon_ValuesRare(c)["pwt"]
	primaryOffhandWeapon_EnchantmentValues["Epic"] = GetEnchantmentLists_Weapon_ValuesEpic(c)["pwt"]
	primaryOffhandWeapon_EnchantmentValues["Legend"] = GetEnchantmentLists_Weapon_ValuesLegend(c)["pwt"]
	primaryOffhandWeapon_EnchantmentValues["Unique"] = GetEnchantmentLists_Weapon_ValuesUnique(c)["pwt"]

	c.JSON(http.StatusOK, gin.H{
		"listname_uncommon":  primaryOffhandWeapon_EnchantmentName["Uncommon"],
		"listvalue_uncommon": primaryOffhandWeapon_EnchantmentValues["Uncommon"],
		"listname_rare":      primaryOffhandWeapon_EnchantmentName["Rare"],
		"listvalue_rare":     primaryOffhandWeapon_EnchantmentValues["Rare"],
		"listname_epic":      primaryOffhandWeapon_EnchantmentName["Epic"],
		"listvalue_epic":     primaryOffhandWeapon_EnchantmentValues["Epic"],
		"listname_legend":    primaryOffhandWeapon_EnchantmentName["Legend"],
		"listvalue_legend":   primaryOffhandWeapon_EnchantmentValues["Legend"],
		"listname_unique":    primaryOffhandWeapon_EnchantmentName["Unique"],
		"listvalue_unique":   primaryOffhandWeapon_EnchantmentValues["Unique"],
	})
}
*/
//////////\\\\\\\\\ ------->   ROUTES  <------- //////////\\\\\\\\\

func setupRoutes(r *gin.Engine) {
	// LISTS ITEMS ENDPOINTS
	r.POST("/helmetlist/", Helmet_List_Handler)
	r.POST("/chestlist/", Chest_List_Handler)
	r.POST("/gloveslist/", Gloves_List_Handler)
	r.POST("/pantslist/", Pants_List_Handler)
	r.POST("/bootslist/", Boots_List_Handler)
	r.POST("/cloaklist/", Cloak_List_Handler)
	r.POST("/pwolist/", Pwo_List_Handler)
	r.POST("/pwtlist/", Pwt_List_Handler)
	r.POST("/necklacelist/", Necklace_List_Handler)
	r.POST("/ringlist/", Ring_List_Handler)

	// RATING LISTS ENDPOINTS
	r.POST("/helmetratinglist/", Helmet_RatingList_Handler)
	r.POST("/chestratinglist/", Chest_RatingList_Handler)
	r.POST("/glovesratinglist/", Gloves_RatingList_Handler)
	r.POST("/pantsratinglist/", Pants_RatingList_Handler)
	r.POST("/bootsratinglist/", Boots_RatingList_Handler)
	r.POST("/cloakratinglist/", Cloak_RatingList_Handler)
	r.POST("/pworatinglist/", Pwo_RatingList_Handler)
	r.POST("/pwtratinglist/", Pwt_RatingList_Handler)

	//ENCHANTMENT LISTS ENDPOINTS
	r.POST("/enchantmentlisthelmet/", Helmet_EnchantmentList_Handler)
	/*
		r.GET("/enchantmentlistchest/", Chest_EnchantmentList_Handler)
		r.GET("/enchantmentlistgloves/", Gloves_EnchantmentList_Handler)
		r.GET("/enchantmentlistpants/", Pants_EnchantmentList_Handler)
		r.GET("/enchantmentlistboots/", Boots_EnchantmentList_Handler)
		r.GET("/enchantmentlistcloak/", Cloak_EnchantmentList_Handler)

		r.GET("/enchantmentlistpwo/", Pwo_EnchantmentList_Handler)
		r.GET("/enchantmentlistpwt/", Pwt_EnchantmentList_Handler)

		r.GET("/enchantmentlistnecklace/", Necklace_EnchantmentList_Handler)
		r.GET("/enchantmentlistring/", Ring_EnchantmentList_Handler)
		r.GET("/enchantmentlistringtwo/", RingTwo_EnchantmentList_Handler)
	*/
	// item display
	r.POST("/itemdisplay/", itemDisplayHandler)

	// Calculate stats
	//r.GET("/charbuilder/:classSelection", updateStatsHandler)

	// json handler
	r.POST("/api/", postHandler)

}
