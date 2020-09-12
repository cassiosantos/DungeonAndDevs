package skills

import "github.com/cassiosantos/DungeonAndDevs/models"

//SkillService |
type SkillService struct {
	repo Repository
}

// NewService |
func NewService(r Repository) *SkillService {
	return &SkillService{repo: r}
}

// IsValid |
func (s *SkillService) IsValid(skill models.Skill) (bool, string) {
	if skill.Name == "" {
		return false, "The skillibute needs a name!"
	}
	if s.SkillExists(skill.Name) {
		return false, "An skillibute with this name was already created"
	}
	return true, ""
}

// SkillExists |
func (s *SkillService) SkillExists(name string) bool {
	_, err := s.repo.GetSkillByName(name)
	return err == nil
}

// AddSkill |
func (s *SkillService) AddSkill(skill models.Skill) error {
	return s.repo.AddSkill(skill)
}

// UpdateSkill |
func (s *SkillService) UpdateSkill(skill models.Skill) error {
	return s.repo.UpdateSkill(skill)
}

// GetSkillByName |
func (s *SkillService) GetSkillByName(name string) (models.Skill, error) {
	return s.repo.GetSkillByName(name)
}

// GetAllSkills |
func (s *SkillService) GetAllSkills() ([]models.Skill, error) {
	return s.repo.GetAllSkills()
}

// DeleteSkillByName |
func (s *SkillService) DeleteSkillByName(name string) error {
	return s.repo.DeleteSkillByName(name)
}
