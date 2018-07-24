package attributes

import (
	"github.com/lucasmdrs/DungeonAndDevs/models"
)

// Read |
type Read interface {
	GetAllAttrs() ([]models.Attribute, error)
	GetAttrByName(name string) (models.Attribute, error)
}

// Write |
type Write interface {
	AddAttr(attr models.Attribute) error
	UpdateAttr(attr models.Attribute) error
	DeleteAttrByName(name string) error
}

// Repository |
type Repository interface {
	Read
	Write
}

// Service |
type Service interface {
	IsValid(Attr models.Attribute) bool
	Read
	Write
}
