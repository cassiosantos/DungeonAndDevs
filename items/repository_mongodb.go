package items

import (
	"log"

	"github.com/lucasmdrs/DungeonAndDevs/models"
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

// AddItem |
func (r *MongoRepository) AddItem(item models.Item) error {
	return r.db.DB("DnDev").C("Items").Insert(item)
}

// GetAllItems |
func (r *MongoRepository) GetAllItems() ([]models.Item, error) {
	list := []models.Item{}
	err := r.db.DB("DnDev").C("Items").Find(bson.M{}).All(&list)
	if err != nil {
		log.Panicf("Error searching item: %s\n", err)
	}
	return list, err
}

// GetItemByName |
func (r *MongoRepository) GetItemByName(name string) (models.Item, error) {
	var c models.Item
	itemQuerier := bson.M{"name": name}
	err := r.db.DB("DnDev").C("Items").Find(itemQuerier).One(&c)
	if err != nil && err != mgo.ErrNotFound {
		log.Panicf("Error searching item: %s\n", err)
	}
	return c, err
}

// UpdateItem |
func (r *MongoRepository) UpdateItem(item models.Item) error {
	itemQuerier := bson.M{"name": item.Name}
	change := bson.M{"$set": bson.M{
		"name":   item.Name,
		"type":   item.Type,
		"value":  item.Value,
		"weight": item.Weight,
	}}
	err := r.db.DB("DnDev").C("Items").Update(itemQuerier, change)
	if err != nil {
		log.Panicf("Error updating item: %s\n", err)
	}
	return nil
}

// DeleteItemByName |
func (r *MongoRepository) DeleteItemByName(name string) error {
	itemQuerier := bson.M{"name": name}
	err := r.db.DB("DnDev").C("Items").Remove(itemQuerier)
	if err != nil {
		log.Panicf("Error deleting item: %s\n", err)
	}
	return nil
}
