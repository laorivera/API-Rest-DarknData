package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

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
	fmt.Println(Sm.variable, "2")
	//fmt.Println(Sm.totalrating, "2")
	//fmt.Println(Sm.totalspeed, "3")
	//fmt.Println(resultother, "4")

	totally := Calculate(Sm.base, Sm.totalrating, Sm.totalspeed, resultother, Sm.variable)
	total := totally.AddEnchant(resultother)

	Sm.WeaponDamageCalcx(selection, total)

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
