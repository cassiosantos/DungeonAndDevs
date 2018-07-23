package main

import (
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/lucasmdrs/DungeonAndDevs/controllers"
	"github.com/lucasmdrs/DungeonAndDevs/db"
)

func main() {
	s := db.Init("mongodb://localhost:27017")
	ctrl := controllers.NewCharController(s)
	router := httprouter.New()

	// Char routes
	router.POST("/char", ctrl.Create)
	router.PATCH("/char/:name", ctrl.Update)
	router.GET("/char/:name", ctrl.Find)
	router.DELETE("/char/:name", ctrl.Delete)
	router.GET("/chars", ctrl.List)

	// //Attr routes
	// router.GET("/attr/:id", Find)
	// router.POST("/attr", Create)
	// router.GET("/attr/:id", Delete)
	// router.POST("/attr/:id", Update)
	//
	// //Item routes
	// router.GET("/item/:id", Find)
	// router.POST("/item", Create)
	// router.GET("/item/:id", Delete)
	// router.POST("/item/:id", Update)
	//
	// //Skill routes
	// router.GET("/skill/:id", Find)
	// router.POST("/skill", Create)
	// router.GET("/skill/:id", Delete)
	// router.POST("/skill/:id", Update)

	log.Fatal(http.ListenAndServe(":8080", router))
}
