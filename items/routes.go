package items

import (
	"github.com/julienschmidt/httprouter"
)

//RouterItems |
func RouterItems(itemRepo Repository, router *httprouter.Router) {

	itemService := NewService(itemRepo)
	itemController := NewItemController(itemService)

	router.POST("/item", itemController.Create)
	router.PATCH("/item", itemController.Update)
	router.GET("/item/:name", itemController.Find)
	router.DELETE("/item/:name", itemController.Delete)
	router.GET("/items", itemController.List)

}
