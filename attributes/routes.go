package attributes

import (
	"github.com/julienschmidt/httprouter"
)

//RouterAttr |
func RouterAttr(attrRepo Repository, router *httprouter.Router) {

	attrService := NewService(attrRepo)
	attrController := NewAttrController(attrService)

	router.POST("/attr", attrController.Create)
	router.PATCH("/attr", attrController.Update)
	router.GET("/attr/:name", attrController.Find)
	router.DELETE("/attr/:name", attrController.Delete)
	router.GET("/attrs", attrController.List)

}
