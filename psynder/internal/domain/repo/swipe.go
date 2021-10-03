package repo

import (
	"psynder/internal/domain/model"
)

type LoadPsynasOptions struct {
	Count int
	AccountId model.AccountId
}

type LikePsynaOptions struct {
	AccountId model.AccountId
	PsynaId model.PsynaId
}

type SwipeRepo interface {
	LoadPsynas(opts LoadPsynasOptions) ([]model.Psyna, error)
	SaveLastView(account_id model.AccountId, psyna_id model.PsynaId) error
	LikePsyna(opts LikePsynaOptions) error
	GetFavoritePsynas(id model.AccountId) ([]model.Psyna, error)
}
