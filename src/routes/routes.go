package routes

import (
	"builder/src/controllers"
	"builder/src/middlewares"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {

	// LISTS ITEMS ENDPOINTS
	r.POST("/helmetlist/", middlewares.Auth(), controllers.Helmet_List_Handler)
	r.POST("/chestlist/", middlewares.Auth(), controllers.Chest_List_Handler)
	r.POST("/gloveslist/", middlewares.Auth(), controllers.Gloves_List_Handler)
	r.POST("/pantslist/", middlewares.Auth(), controllers.Pants_List_Handler)
	r.POST("/bootslist/", middlewares.Auth(), controllers.Boots_List_Handler)
	r.POST("/cloaklist/", middlewares.Auth(), controllers.Cloak_List_Handler)

	r.POST("/pwolist/", middlewares.Auth(), controllers.Pwo_List_Handler)
	r.POST("/pwtlist/", middlewares.Auth(), controllers.Pwt_List_Handler)
	r.POST("/necklacelist/", middlewares.Auth(), controllers.Necklace_List_Handler)
	r.POST("/ringlist/", middlewares.Auth(), controllers.Ring_List_Handler)

	// RATING LISTS ENDPOINTS
	r.POST("/helmetratinglist/", middlewares.Auth(), controllers.Helmet_RatingList_Handler)
	r.POST("/chestratinglist/", middlewares.Auth(), controllers.Chest_RatingList_Handler)
	r.POST("/glovesratinglist/", middlewares.Auth(), controllers.Gloves_RatingList_Handler)
	r.POST("/pantsratinglist/", middlewares.Auth(), controllers.Pants_RatingList_Handler)
	r.POST("/bootsratinglist/", middlewares.Auth(), controllers.Boots_RatingList_Handler)
	r.POST("/cloakratinglist/", middlewares.Auth(), controllers.Cloak_RatingList_Handler)

	r.POST("/pworatinglist/", middlewares.Auth(), controllers.Pwo_RatingList_Handler)
	r.POST("/pwtratinglist/", middlewares.Auth(), controllers.Pwt_RatingList_Handler)

	//ENCHANTMENT LISTS ENDPOINTS
	r.POST("/enchantmentlisthelmet/", middlewares.Auth(), controllers.Helmet_EnchantmentList_Handler)
	r.POST("/enchantmentlistchest/", middlewares.Auth(), controllers.Chest_EnchantmentList_Handler)
	r.POST("/enchantmentlistgloves/", middlewares.Auth(), controllers.Gloves_EnchantmentList_Handler)
	r.POST("/enchantmentlistpants/", middlewares.Auth(), controllers.Pants_EnchantmentList_Handler)
	r.POST("/enchantmentlistboots/", middlewares.Auth(), controllers.Boots_EnchantmentList_Handler)
	r.POST("/enchantmentlistcloak/", middlewares.Auth(), controllers.Cloak_EnchantmentList_Handler)

	r.POST("/enchantmentlistpwo/", middlewares.Auth(), controllers.Pwo_EnchantmentList_Handler)
	r.POST("/enchantmentlistpwt/", middlewares.Auth(), controllers.Pwt_EnchantmentList_Handler)

	r.POST("/enchantmentlistnecklace/", middlewares.Auth(), controllers.Necklace_EnchantmentList_Handler)
	r.POST("/enchantmentlistring/", middlewares.Auth(), controllers.Ring_EnchantmentList_Handler)
	r.POST("/enchantmentlistringtwo/", middlewares.Auth(), controllers.RingTwo_EnchantmentList_Handler)

	// item display
	r.POST("/itemdisplay/", middlewares.Auth(), controllers.ItemDisplayHandler)

	// Calculate stats
	r.POST("/api/", middlewares.Auth(), apiHandler)

}
