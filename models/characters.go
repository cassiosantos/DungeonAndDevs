package models

//Character |
type Character struct {
	Name      string
	Age       int
	Bio       string
	Race      string
	Class     string
	Attr      []Attribute
	Skills    []Skill
	Inventory map[Item]int
}
