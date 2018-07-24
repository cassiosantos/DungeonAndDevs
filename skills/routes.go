package skills

import (
	"github.com/julienschmidt/httprouter"
)

//RouterSkill |
func RouterSkill(skillRepo Repository, router *httprouter.Router) {

	skillService := NewService(skillRepo)
	skillController := NewSkillController(skillService)

	router.POST("/skill", skillController.Create)
	router.PATCH("/skill", skillController.Update)
	router.GET("/skill/:name", skillController.Find)
	router.DELETE("/skill/:name", skillController.Delete)
	router.GET("/skills", skillController.List)

}
