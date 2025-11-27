package routes

import (
	"builder/src/controllers"

	"github.com/gin-gonic/gin"
)

// SetupRoutes configures all routes on the gin router using controllers
func SetupRoutes(r *gin.Engine) {
	// LISTS ITEMS ENDPOINTS
	r.POST("/helmetlist/", controllers.Helmet_List_Handler)
	r.POST("/chestlist/", controllers.Chest_List_Handler)
	r.POST("/gloveslist/", controllers.Gloves_List_Handler)
	r.POST("/pantslist/", controllers.Pants_List_Handler)
	r.POST("/bootslist/", controllers.Boots_List_Handler)
	r.POST("/cloaklist/", controllers.Cloak_List_Handler)
	r.POST("/pwolist/", controllers.Pwo_List_Handler)
	r.POST("/pwtlist/", controllers.Pwt_List_Handler)
	r.POST("/necklacelist/", controllers.Necklace_List_Handler)
	r.POST("/ringlist/", controllers.Ring_List_Handler)

	// RATING LISTS ENDPOINTS
	r.POST("/helmetratinglist/", controllers.Helmet_RatingList_Handler)
	r.POST("/chestratinglist/", controllers.Chest_RatingList_Handler)
	r.POST("/glovesratinglist/", controllers.Gloves_RatingList_Handler)
	r.POST("/pantsratinglist/", controllers.Pants_RatingList_Handler)
	r.POST("/bootsratinglist/", controllers.Boots_RatingList_Handler)
	r.POST("/cloakratinglist/", controllers.Cloak_RatingList_Handler)
	r.POST("/pworatinglist/", controllers.Pwo_RatingList_Handler)
	r.POST("/pwtratinglist/", controllers.Pwt_RatingList_Handler)

	//ENCHANTMENT LISTS ENDPOINTS
	r.POST("/enchantmentlisthelmet/", Helmet_EnchantmentList_Handler)
	r.POST("/enchantmentlistchest/", Chest_EnchantmentList_Handler)
	r.POST("/enchantmentlistgloves/", Gloves_EnchantmentList_Handler)
	r.POST("/enchantmentlistpants/", Pants_EnchantmentList_Handler)
	r.POST("/enchantmentlistboots/", Boots_EnchantmentList_Handler)
	r.POST("/enchantmentlistcloak/", Cloak_EnchantmentList_Handler)

	r.POST("/enchantmentlistpwo/", Pwo_EnchantmentList_Handler)
	r.POST("/enchantmentlistpwt/", Pwt_EnchantmentList_Handler)

	r.POST("/enchantmentlistnecklace/", Necklace_EnchantmentList_Handler)
	r.POST("/enchantmentlistring/", Ring_EnchantmentList_Handler)
	r.POST("/enchantmentlistringtwo/", RingTwo_EnchantmentList_Handler)

	// item display
	r.POST("/itemdisplay/", controllers.itemDisplayHandler)

	// Calculate stats
	r.POST("/api/", postHandler)

}
