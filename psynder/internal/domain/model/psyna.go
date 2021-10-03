package model

type PsynaId uint64

type Psyna struct {
	Id PsynaId `json:"id"`
	Name string `json:"name"`
	Description string `json:"description"`
	PhotoLink string `json:"photo_link"`
}

