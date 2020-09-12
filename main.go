package main

import (
	"log"
	"net/http"

	"github.com/cassiosantos/DungeonAndDevs/attributes"
	"github.com/cassiosantos/DungeonAndDevs/characters"
	"github.com/cassiosantos/DungeonAndDevs/db"
	"github.com/cassiosantos/DungeonAndDevs/items"
	"github.com/cassiosantos/DungeonAndDevs/skills"
	"github.com/julienschmidt/httprouter"
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
