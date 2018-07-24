package items

import "github.com/lucasmdrs/DungeonAndDevs/models"

//ItemService |
type ItemService struct {
	repo Repository
}

// NewService |
func NewService(r Repository) *ItemService {
	return &ItemService{repo: r}
}

// IsValid |
func (s *ItemService) IsValid(item models.Item) (bool, string) {
	if item.Name == "" {
		return false, "The item needs a name!"
	}
	if s.ItemExists(item.Name) {
		return false, "An item with this name was already created"
	}
	return true, ""
}

// ItemExists |
func (s *ItemService) ItemExists(name string) bool {
	_, err := s.repo.GetItemByName(name)
	return err == nil
}

// AddItem |
func (s *ItemService) AddItem(item models.Item) error {
	return s.repo.AddItem(item)
}

// UpdateItem |
func (s *ItemService) UpdateItem(item models.Item) error {
	return s.repo.UpdateItem(item)
}

// GetItemByName |
func (s *ItemService) GetItemByName(name string) (models.Item, error) {
	return s.repo.GetItemByName(name)
}

// GetAllItems |
func (s *ItemService) GetAllItems() ([]models.Item, error) {
	return s.repo.GetAllItems()
}

// DeleteItemByName |
func (s *ItemService) DeleteItemByName(name string) error {
	return s.repo.DeleteItemByName(name)
}
