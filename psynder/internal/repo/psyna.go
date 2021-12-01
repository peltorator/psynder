package repo

import (
	"github.com/peltorator/psynder/internal/domain"
	"github.com/peltorator/psynder/internal/pagination"
)

type PsynaData struct {
	Name        string
	Breed       *string
	Description string
	PhotoLink   string
}

type Psyna struct {
	Id domain.PsynaId
	PsynaData
}

type Psynas interface {
	StoreNew(data PsynaData) (domain.PsynaId, error)
	LoadSlice(uid domain.AccountId, pg pagination.Info, f domain.PsynaFilter) ([]Psyna, error)
}
