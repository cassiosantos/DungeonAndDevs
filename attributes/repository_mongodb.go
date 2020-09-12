package attributes

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

// AddAttr |
func (r *MongoRepository) AddAttr(attr models.Attribute) error {
	return r.db.DB("DnDev").C("Attrs").Insert(attr)
}

// GetAllAttrs |
func (r *MongoRepository) GetAllAttrs() ([]models.Attribute, error) {
	list := []models.Attribute{}
	err := r.db.DB("DnDev").C("Attrs").Find(bson.M{}).All(&list)
	if err != nil {
		log.Panicf("Error searching attr: %s\n", err)
	}
	return list, err
}

// GetAttrByName |
func (r *MongoRepository) GetAttrByName(name string) (models.Attribute, error) {
	var c models.Attribute
	attrQuerier := bson.M{"name": name}
	err := r.db.DB("DnDev").C("Attrs").Find(attrQuerier).One(&c)
	if err != nil && err != mgo.ErrNotFound {
		log.Panicf("Error searching attr: %s\n", err)
	}
	return c, err
}

// UpdateAttr |
func (r *MongoRepository) UpdateAttr(attr models.Attribute) error {
	attrQuerier := bson.M{"name": attr.Name}
	change := bson.M{"$set": bson.M{
		"name":  attr.Name,
		"value": attr.Value,
	}}
	err := r.db.DB("DnDev").C("Attrs").Update(attrQuerier, change)
	if err != nil {
		log.Panicf("Error updating attr: %s\n", err)
	}
	return nil
}

// DeleteAttrByName |
func (r *MongoRepository) DeleteAttrByName(name string) error {
	attrQuerier := bson.M{"name": name}
	err := r.db.DB("DnDev").C("Attrs").Remove(attrQuerier)
	if err != nil {
		log.Panicf("Error deleting attr: %s\n", err)
	}
	return nil
}
