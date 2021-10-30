package storage

import "github.com/peltorator/psynder/internal/domain"

type PsynaData struct {
	Name string
	Description string
	PhotoLink string
}

type Psyna struct {
	Id domain.PsynaId
	PsynaData
}