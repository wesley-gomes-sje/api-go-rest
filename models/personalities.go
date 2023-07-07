package models

type Personalitie struct {
	Id    int    `json:"id"`
	Name  string `json:"name"`
	About string `json:"about"`
}

var Personalities []Personalitie
