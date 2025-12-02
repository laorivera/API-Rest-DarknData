package routes

import (
	"builder/src/controllers"
	"builder/src/models"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func apiHandler(c *gin.Context) {
	var selection models.Selection

	if err := c.BindJSON(&selection); err != nil {
		fmt.Println("JSON Bind Error:", err)
		return
	}
	//fmt.Println(selection)

	var Sm = controllers.NewStatsManager()

	Sm.ItemsBaseStats(selection)
	Sm.BaseItemCalc(selection)
	Sm.TotalSpeedRating(selection)
	Sm.TotalArmorRating(selection)

	Sm.Base = controllers.ProcessBaseEnchantments(selection).AddStats(Sm.Base)

	resultother := controllers.ProcessOtherEnchantments(selection).AddEnchant(Sm.Variable)

	totally := controllers.Calculate(Sm.Base, Sm.Totalrating, Sm.Totalspeed, resultother, Sm.Variable)
	total := totally.AddEnchant(resultother)

	Sm.WeaponDamageCalcx(selection, total)

	c.JSON(http.StatusOK, gin.H{

		"stats":               Sm.Base,
		"computedstats":       total,
		"computedstatsweapon": Sm.Weapon,
	})

}
