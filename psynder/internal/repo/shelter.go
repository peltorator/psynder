package repo

import (
	"github.com/peltorator/psynder/internal/domain"
	"github.com/peltorator/psynder/internal/domain/swipe"
	"github.com/peltorator/psynder/internal/pagination"
)

type Shelters interface {
	AddInfo(uid domain.AccountId, info domain.ShelterInfo) error
	AddPsyna(uid domain.AccountId, p swipe.PsynaData) (domain.PsynaId, error)
	LoadSlice(uid domain.AccountId, pg pagination.Info) ([]Psyna, error)
}
