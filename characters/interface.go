package characters

import (
	"github.com/cassiosantos/DungeonAndDevs/models"
)

// Read |
type Read interface {
	GetAllChars() ([]models.Character, error)
	GetCharByName(name string) (models.Character, error)
}

// Write |
type Write interface {
	AddChar(char models.Character) error
	UpdateChar(char models.Character) error
	DeleteCharByName(name string) error
}

// Repository |
type Repository interface {
	Read
	Write
}

// Service |
type Service interface {
	IsValid(char models.Character) bool
	Read
	Write
}
