package models

//Character |
type Character struct {
	Name      string      `json:"name"`
	Age       int         `json:"age"`
	Bio       string      `json:"bio"`
	Race      string      `json:"race"`
	Class     string      `json:"class"`
	Attr      []Attribute `json:"attributes"`
	Skills    []Skill     `json:"skills"`
	Inventory []Item      `json:"inventory"`
}
