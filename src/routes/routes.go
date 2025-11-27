package routes

import (
	"builder/src/controllers"
	middleware "builder/src/middlewares"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {

	// LISTS ITEMS ENDPOINTS
	r.POST("/helmetlist/", middleware.Lock(), controllers.Helmet_List_Handler)
	r.POST("/chestlist/", middleware.Lock(), controllers.Chest_List_Handler)
	r.POST("/gloveslist/", middleware.Lock(), controllers.Gloves_List_Handler)
	r.POST("/pantslist/", middleware.Lock(), controllers.Pants_List_Handler)
	r.POST("/bootslist/", middleware.Lock(), controllers.Boots_List_Handler)
	r.POST("/cloaklist/", middleware.Lock(), controllers.Cloak_List_Handler)

	r.POST("/pwolist/", middleware.Lock(), controllers.Pwo_List_Handler)
	r.POST("/pwtlist/", middleware.Lock(), controllers.Pwt_List_Handler)
	r.POST("/necklacelist/", middleware.Lock(), controllers.Necklace_List_Handler)
	r.POST("/ringlist/", middleware.Lock(), controllers.Ring_List_Handler)

	// RATING LISTS ENDPOINTS
	r.POST("/helmetratinglist/", middleware.Lock(), controllers.Helmet_RatingList_Handler)
	r.POST("/chestratinglist/", middleware.Lock(), controllers.Chest_RatingList_Handler)
	r.POST("/glovesratinglist/", middleware.Lock(), controllers.Gloves_RatingList_Handler)
	r.POST("/pantsratinglist/", middleware.Lock(), controllers.Pants_RatingList_Handler)
	r.POST("/bootsratinglist/", middleware.Lock(), controllers.Boots_RatingList_Handler)
	r.POST("/cloakratinglist/", middleware.Lock(), controllers.Cloak_RatingList_Handler)

	r.POST("/pworatinglist/", middleware.Lock(), controllers.Pwo_RatingList_Handler)
	r.POST("/pwtratinglist/", middleware.Lock(), controllers.Pwt_RatingList_Handler)

	//ENCHANTMENT LISTS ENDPOINTS
	r.POST("/enchantmentlisthelmet/", middleware.Lock(), controllers.Helmet_EnchantmentList_Handler)
	r.POST("/enchantmentlistchest/", middleware.Lock(), controllers.Chest_EnchantmentList_Handler)
	r.POST("/enchantmentlistgloves/", middleware.Lock(), controllers.Gloves_EnchantmentList_Handler)
	r.POST("/enchantmentlistpants/", middleware.Lock(), controllers.Pants_EnchantmentList_Handler)
	r.POST("/enchantmentlistboots/", middleware.Lock(), controllers.Boots_EnchantmentList_Handler)
	r.POST("/enchantmentlistcloak/", middleware.Lock(), controllers.Cloak_EnchantmentList_Handler)

	r.POST("/enchantmentlistpwo/", middleware.Lock(), controllers.Pwo_EnchantmentList_Handler)
	r.POST("/enchantmentlistpwt/", middleware.Lock(), controllers.Pwt_EnchantmentList_Handler)

	r.POST("/enchantmentlistnecklace/", middleware.Lock(), controllers.Necklace_EnchantmentList_Handler)
	r.POST("/enchantmentlistring/", middleware.Lock(), controllers.Ring_EnchantmentList_Handler)
	r.POST("/enchantmentlistringtwo/", middleware.Lock(), controllers.RingTwo_EnchantmentList_Handler)

	// item display
	r.POST("/itemdisplay/", middleware.Lock(), controllers.ItemDisplayHandler)

	// Calculate stats
	r.POST("/api/", middleware.Lock(), apiHandler)

}
