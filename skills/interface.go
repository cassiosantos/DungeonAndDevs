package skills

import (
	"github.com/cassiosantos/DungeonAndDevs/models"
)

// Read |
type Read interface {
	GetAllSkills() ([]models.Skill, error)
	GetSkillByName(name string) (models.Skill, error)
}

// Write |
type Write interface {
	AddSkill(skill models.Skill) error
	UpdateSkill(skill models.Skill) error
	DeleteSkillByName(name string) error
}

// Repository |
type Repository interface {
	Read
	Write
}

// Service |
type Service interface {
	IsValid(Skill models.Skill) bool
	Read
	Write
}
