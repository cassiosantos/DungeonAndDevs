package characters

import "github.com/lucasmdrs/DungeonAndDevs/models"

//CharacterService |
type CharacterService struct {
	repo Repository
}

// NewService |
func NewService(r Repository) *CharacterService {
	return &CharacterService{repo: r}
}

// IsValid |
func (s *CharacterService) IsValid(char models.Character) (bool, string) {
	if char.Name == "" {
		return false, "The character needs a name!"
	}
	if s.CharExists(char.Name) {
		return false, "An character with this name was already created"
	}
	return true, ""
}

// CharExists |
func (s *CharacterService) CharExists(name string) bool {
	_, err := s.repo.GetCharByName(name)
	return err == nil
}

// AddChar |
func (s *CharacterService) AddChar(char models.Character) error {
	return s.repo.AddChar(char)
}

// UpdateChar |
func (s *CharacterService) UpdateChar(char models.Character) error {
	return s.repo.UpdateChar(char)
}

// GetCharByName |
func (s *CharacterService) GetCharByName(name string) (models.Character, error) {
	return s.repo.GetCharByName(name)
}

// GetAllChars |
func (s *CharacterService) GetAllChars() ([]models.Character, error) {
	return s.repo.GetAllChars()
}

// DeleteCharByName |
func (s *CharacterService) DeleteCharByName(name string) error {
	return s.repo.DeleteCharByName(name)
}
