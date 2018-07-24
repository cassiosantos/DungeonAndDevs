package main

import (
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/lucasmdrs/DungeonAndDevs/attributes"
	"github.com/lucasmdrs/DungeonAndDevs/characters"
	"github.com/lucasmdrs/DungeonAndDevs/db"
	"github.com/lucasmdrs/DungeonAndDevs/items"
	"github.com/lucasmdrs/DungeonAndDevs/skills"
)

func main() {

	s := db.InitMongo("mongodb://localhost:27017")
	defer func() {
		if s != nil {
			s.Close()
		}
	}()

	router := httprouter.New()
	charRepo := characters.NewMongoRepository(s)
	itemRepo := items.NewMongoRepository(s)
	attrRepo := attributes.NewMongoRepository(s)
	skillRepo := skills.NewMongoRepository(s)

	characters.RouterCharacters(charRepo, router)
	items.RouterItems(itemRepo, router)
	attributes.RouterAttr(attrRepo, router)
	skills.RouterSkill(skillRepo, router)

	log.Fatal(http.ListenAndServe(":8080", router))
}
