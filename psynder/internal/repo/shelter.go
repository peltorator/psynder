package repo

import (
	"github.com/peltorator/psynder/internal/domain"
	"github.com/peltorator/psynder/internal/domain/swipe"
	"github.com/peltorator/psynder/internal/pagination"
)

type ShelterData struct {
	City string
	Address string
	Phone string
}

type Shelter struct {
	Id domain.AccountId
	ShelterData
}

type Shelters interface {
	AddInfo(uid domain.AccountId, info domain.ShelterInfo) error
	AddPsyna(uid domain.AccountId, p swipe.PsynaData) (domain.PsynaId, error)
	DeletePsyna(uid domain.AccountId, pid domain.PsynaId) error
	LoadSlice(uid domain.AccountId, pg pagination.Info) ([]Psyna, error)
}
