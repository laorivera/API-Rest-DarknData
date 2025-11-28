package routes

import (
	"builder/src/controllers"
	"builder/src/middlewares"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {

	// LISTS ITEMS ENDPOINTS
	r.POST("/helmetlist/", middlewares.Lock(), controllers.Helmet_List_Handler)
	r.POST("/chestlist/", middlewares.Lock(), controllers.Chest_List_Handler)
	r.POST("/gloveslist/", middlewares.Lock(), controllers.Gloves_List_Handler)
	r.POST("/pantslist/", middlewares.Lock(), controllers.Pants_List_Handler)
	r.POST("/bootslist/", middlewares.Lock(), controllers.Boots_List_Handler)
	r.POST("/cloaklist/", middlewares.Lock(), controllers.Cloak_List_Handler)

	r.POST("/pwolist/", middlewares.Lock(), controllers.Pwo_List_Handler)
	r.POST("/pwtlist/", middlewares.Lock(), controllers.Pwt_List_Handler)
	r.POST("/necklacelist/", middlewares.Lock(), controllers.Necklace_List_Handler)
	r.POST("/ringlist/", middlewares.Lock(), controllers.Ring_List_Handler)

	// RATING LISTS ENDPOINTS
	r.POST("/helmetratinglist/", middlewares.Lock(), controllers.Helmet_RatingList_Handler)
	r.POST("/chestratinglist/", middlewares.Lock(), controllers.Chest_RatingList_Handler)
	r.POST("/glovesratinglist/", middlewares.Lock(), controllers.Gloves_RatingList_Handler)
	r.POST("/pantsratinglist/", middlewares.Lock(), controllers.Pants_RatingList_Handler)
	r.POST("/bootsratinglist/", middlewares.Lock(), controllers.Boots_RatingList_Handler)
	r.POST("/cloakratinglist/", middlewares.Lock(), controllers.Cloak_RatingList_Handler)

	r.POST("/pworatinglist/", middlewares.Lock(), controllers.Pwo_RatingList_Handler)
	r.POST("/pwtratinglist/", middlewares.Lock(), controllers.Pwt_RatingList_Handler)

	//ENCHANTMENT LISTS ENDPOINTS
	r.POST("/enchantmentlisthelmet/", middlewares.Lock(), controllers.Helmet_EnchantmentList_Handler)
	r.POST("/enchantmentlistchest/", middlewares.Lock(), controllers.Chest_EnchantmentList_Handler)
	r.POST("/enchantmentlistgloves/", middlewares.Lock(), controllers.Gloves_EnchantmentList_Handler)
	r.POST("/enchantmentlistpants/", middlewares.Lock(), controllers.Pants_EnchantmentList_Handler)
	r.POST("/enchantmentlistboots/", middlewares.Lock(), controllers.Boots_EnchantmentList_Handler)
	r.POST("/enchantmentlistcloak/", middlewares.Lock(), controllers.Cloak_EnchantmentList_Handler)

	r.POST("/enchantmentlistpwo/", middlewares.Lock(), controllers.Pwo_EnchantmentList_Handler)
	r.POST("/enchantmentlistpwt/", middlewares.Lock(), controllers.Pwt_EnchantmentList_Handler)

	r.POST("/enchantmentlistnecklace/", middlewares.Lock(), controllers.Necklace_EnchantmentList_Handler)
	r.POST("/enchantmentlistring/", middlewares.Lock(), controllers.Ring_EnchantmentList_Handler)
	r.POST("/enchantmentlistringtwo/", middlewares.Lock(), controllers.RingTwo_EnchantmentList_Handler)

	// item display
	r.POST("/itemdisplay/", middlewares.Lock(), controllers.ItemDisplayHandler)

	// Calculate stats
	r.POST("/api/", middlewares.Lock(), apiHandler)

}
