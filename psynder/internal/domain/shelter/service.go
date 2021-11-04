package shelter

import (
	"github.com/peltorator/psynder/internal/domain"
	"github.com/peltorator/psynder/internal/domain/swipe"
	"github.com/peltorator/psynder/internal/pagination"
)

type Service interface {
	AddInfo(uid domain.AccountId, info domain.ShelterInfo) error
	AddPsyna(uid domain.AccountId, p swipe.PsynaData) (domain.PsynaId, error)
	GetMyPsynas(uid domain.AccountId, pg pagination.Info) ([]swipe.Psyna, error)
}
