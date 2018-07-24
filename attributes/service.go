package attributes

import "github.com/lucasmdrs/DungeonAndDevs/models"

//AttrService |
type AttrService struct {
	repo Repository
}

// NewService |
func NewService(r Repository) *AttrService {
	return &AttrService{repo: r}
}

// IsValid |
func (s *AttrService) IsValid(attr models.Attribute) (bool, string) {
	if attr.Name == "" {
		return false, "The attribute needs a name!"
	}
	if s.AttrExists(attr.Name) {
		return false, "An attribute with this name was already created"
	}
	return true, ""
}

// AttrExists |
func (s *AttrService) AttrExists(name string) bool {
	_, err := s.repo.GetAttrByName(name)
	return err == nil
}

// AddAttr |
func (s *AttrService) AddAttr(attr models.Attribute) error {
	return s.repo.AddAttr(attr)
}

// UpdateAttr |
func (s *AttrService) UpdateAttr(attr models.Attribute) error {
	return s.repo.UpdateAttr(attr)
}

// GetAttrByName |
func (s *AttrService) GetAttrByName(name string) (models.Attribute, error) {
	return s.repo.GetAttrByName(name)
}

// GetAllAttrs |
func (s *AttrService) GetAllAttrs() ([]models.Attribute, error) {
	return s.repo.GetAllAttrs()
}

// DeleteAttrByName |
func (s *AttrService) DeleteAttrByName(name string) error {
	return s.repo.DeleteAttrByName(name)
}
