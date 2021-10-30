package swipe

import "github.com/peltorator/psynder/internal/domain"

type Psyna struct {
	Id domain.PsynaId `json:"id"`
	Name string `json:"name"`
	Description string `json:"description"`
	PhotoLink string `json:"photoLink"`
}
