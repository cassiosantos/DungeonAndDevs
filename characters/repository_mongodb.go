package characters

import (
	"log"

	"github.com/cassiosantos/DungeonAndDevs/models"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// MongoRepository |
type MongoRepository struct {
	db *mgo.Session
}

// NewMongoRepository |
func NewMongoRepository(session *mgo.Session) *MongoRepository {
	return &MongoRepository{db: session}
}

// AddChar |
func (r *MongoRepository) AddChar(char models.Character) error {
	return r.db.DB("DnDev").C("Characters").Insert(char)
}

// GetAllChars |
func (r *MongoRepository) GetAllChars() ([]models.Character, error) {
	list := []models.Character{}
	err := r.db.DB("DnDev").C("Characters").Find(bson.M{}).All(&list)
	if err != nil {
		log.Panicf("Error searching characters: %s\n", err)
	}
	return list, err
}

// GetCharByName |
func (r *MongoRepository) GetCharByName(name string) (models.Character, error) {
	var c models.Character
	charQuerier := bson.M{"name": name}
	err := r.db.DB("DnDev").C("Characters").Find(charQuerier).One(&c)
	if err != nil && err != mgo.ErrNotFound {
		log.Panicf("Error searching character: %s\n", err)
	}
	return c, err
}

// UpdateChar |
func (r *MongoRepository) UpdateChar(char models.Character) error {
	charQuerier := bson.M{"name": char.Name}
	change := bson.M{"$set": bson.M{
		"name":      char.Name,
		"age":       char.Age,
		"bio":       char.Bio,
		"class":     char.Class,
		"attr":      char.Attr,
		"skills":    char.Skills,
		"inventory": char.Inventory,
	}}
	err := r.db.DB("DnDev").C("Characters").Update(charQuerier, change)
	if err != nil {
		log.Panicf("Error updating character: %s\n", err)
	}
	return nil
}

// DeleteCharByName |
func (r *MongoRepository) DeleteCharByName(name string) error {
	charQuerier := bson.M{"name": name}
	err := r.db.DB("DnDev").C("Characters").Remove(charQuerier)
	if err != nil {
		log.Panicf("Error deleting characters: %s\n", err)
	}
	return nil
}
