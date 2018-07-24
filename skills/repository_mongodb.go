package skills

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

// AddSkill |
func (r *MongoRepository) AddSkill(skill models.Skill) error {
	return r.db.DB("DnDev").C("Skills").Insert(skill)
}

// GetAllSkills |
func (r *MongoRepository) GetAllSkills() ([]models.Skill, error) {
	list := []models.Skill{}
	err := r.db.DB("DnDev").C("Skills").Find(bson.M{}).All(&list)
	if err != nil {
		log.Panicf("Error searching skill: %s\n", err)
	}
	return list, err
}

// GetSkillByName |
func (r *MongoRepository) GetSkillByName(name string) (models.Skill, error) {
	var c models.Skill
	skillQuerier := bson.M{"name": name}
	err := r.db.DB("DnDev").C("Skills").Find(skillQuerier).One(&c)
	if err != nil && err != mgo.ErrNotFound {
		log.Panicf("Error searching skill: %s\n", err)
	}
	return c, err
}

// UpdateSkill |
func (r *MongoRepository) UpdateSkill(skill models.Skill) error {
	skillQuerier := bson.M{"name": skill.Name}
	change := bson.M{"$set": bson.M{
		"name":        skill.Name,
		"type":        skill.Type,
		"description": skill.Description,
		"power":       skill.Power,
	}}
	err := r.db.DB("DnDev").C("Skills").Update(skillQuerier, change)
	if err != nil {
		log.Panicf("Error updating skill: %s\n", err)
	}
	return nil
}

// DeleteSkillByName |
func (r *MongoRepository) DeleteSkillByName(name string) error {
	skillQuerier := bson.M{"name": name}
	err := r.db.DB("DnDev").C("Skills").Remove(skillQuerier)
	if err != nil {
		log.Panicf("Error deleting skill: %s\n", err)
	}
	return nil
}
