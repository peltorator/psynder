package repo

import (
	"psynder/internal/domain/model"
)

type LoadPsynasOptions struct {
	Offset int
	Limit int
	AccountId model.AccountId
}

type LikePsynaOptions struct {
	AccountId model.AccountId
	PsynaId model.PsynaId
}

type SwipeRepo interface {
	LoadPsynasFromRepo(opts LoadPsynasOptions) ([]model.Psyna, error)
	//SaveLastView(account_id model.AccountId, psyna_id model.PsynaId) error
	StoreLikeToRepo(opts LikePsynaOptions) error
	LoadFavoritePsynasFromRepo(id model.AccountId) ([]model.Psyna, error)
}
