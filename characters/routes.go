package characters

import (
	"github.com/julienschmidt/httprouter"
)

//RouterCharacters |
func RouterCharacters(charRepo Repository, router *httprouter.Router) {

	charService := NewService(charRepo)
	charController := NewCharacterController(charService)

	router.POST("/char", charController.Create)
	router.PATCH("/char", charController.Update)
	router.GET("/char/:name", charController.Find)
	router.DELETE("/char/:name", charController.Delete)
	router.GET("/chars", charController.List)

}
