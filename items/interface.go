package items

import (
	"github.com/cassiosantos/DungeonAndDevs/models"
)

// Read |
type Read interface {
	GetAllItems() ([]models.Item, error)
	GetItemByName(name string) (models.Item, error)
}

// Write |
type Write interface {
	AddItem(item models.Item) error
	UpdateItem(item models.Item) error
	DeleteItemByName(name string) error
}

// Repository |
type Repository interface {
	Read
	Write
}

// Service |
type Service interface {
	IsValid(Item models.Item) bool
	Read
	Write
}
