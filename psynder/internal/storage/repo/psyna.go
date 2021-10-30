package repo

import (
	"github.com/peltorator/psynder/internal/domain"
	"github.com/peltorator/psynder/internal/pagination"
	"github.com/peltorator/psynder/internal/storage"
)

type Psynas interface {
	StoreNew(data storage.PsynaData) (domain.PsynaId, error)
	LoadSlice(uid domain.AccountId, pg pagination.Info) ([]storage.Psyna, error)
}
