package models

//Item |
type Item struct {
	Name   string `json:"name"`
	Type   string `json:"type"`
	Weight int    `json:"weight"`
	Value  int    `json:"value"`
}
