package db

import (
	"github.com/lucasmdrs/DungeonAndDevs/models"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// Init starts a mongo session
func Init(arg string) *mgo.Session {
	session, err := mgo.Dial(arg)
	if err != nil {
		panic(err)
	}
	return session
}

// AddChar |
func AddChar(db *mgo.Session, char models.Character) error {
	return db.DB("DnDev").C("Characters").Insert(char)
}

// GetAllChars |
func GetAllChars(db *mgo.Session) ([]models.Character, error) {
	list := []models.Character{}
	err := db.DB("DnDev").C("Characters").Find(bson.M{}).All(&list)
	if err != nil {
		panic(err)
	}
	return list, err
}

// GetCharByName |
func GetCharByName(db *mgo.Session, name string) (models.Character, error) {
	var c models.Character
	charQuerier := bson.M{"name": name}
	err := db.DB("DnDev").C("Characters").Find(charQuerier).One(&c)
	if err != nil {
		panic(err)
	}
	return c, err
}

// UpdateChar |
func UpdateChar(db *mgo.Session, char models.Character) (models.Character, error) {
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
	err := db.DB("DnDev").C("Characters").Update(charQuerier, change)
	if err != nil {
		panic(err)
	}
	return char, nil
}

// DeleteChar |
func DeleteChar(db *mgo.Session, name string) error {
	charQuerier := bson.M{"name": name}
	err := db.DB("DnDev").C("Characters").Remove(charQuerier)
	if err != nil {
		panic(err)
	}
	return nil
}
