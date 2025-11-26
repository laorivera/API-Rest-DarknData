package main

import (
	"builder/src/models"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func postHandler(c *gin.Context) {
	var selection models.Selection

	if err := c.BindJSON(&selection); err != nil {
		fmt.Println("JSON Bind Error:", err)
		return
	}
	fmt.Println(selection)

	var Sm = NewStatsManager()

	Sm.ItemsBaseStats(selection)
	Sm.BaseItemCalc(selection)
	Sm.TotalSpeedRating(selection)
	Sm.TotalArmorRating(selection)

	Sm.base = ProcessBaseEnchantments(selection).AddStats(Sm.base)

	resultother := ProcessOtherEnchantments(selection).AddEnchant(Sm.variable)

	totally := Calculate(Sm.base, Sm.totalrating, Sm.totalspeed, resultother, Sm.variable)
	total := totally.AddEnchant(resultother)

	Sm.WeaponDamageCalcx(selection, total)

	c.JSON(http.StatusOK, gin.H{

		"stats":               Sm.base,
		"computedstats":       total,
		"computedstatsweapon": Sm.weapon,
	})

}
