package swipe

import "github.com/peltorator/psynder/internal/domain"

type PsynaData struct {
	Name string `json:"name"`
	Description string `json:"description"`
	PhotoLink string `json:"photoLink"`
}

type Psyna struct {
	Id domain.PsynaId `json:"id"`
	PsynaData
}
