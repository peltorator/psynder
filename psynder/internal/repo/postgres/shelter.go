package postgres

import (
	"github.com/peltorator/psynder/internal/domain"
	"github.com/peltorator/psynder/internal/domain/swipe"
	"github.com/peltorator/psynder/internal/pagination"
	"github.com/peltorator/psynder/internal/repo"
	"gorm.io/gorm"
)

type shelterRepo struct {
	db *gorm.DB
}

func NewShelterRepo(conn *gorm.DB) *shelterRepo {
	return &shelterRepo{
		db: conn,
	}
}

func (p *shelterRepo) AddInfo(uid domain.AccountId, info domain.ShelterInfo) error {
	// TODO
	return nil
}

func (p *shelterRepo) AddPsyna(uid domain.AccountId, pd swipe.PsynaData) (domain.PsynaId, error) {
	// TODO
	return domain.PsynaId(0), nil
}

func (p *shelterRepo) LoadSlice(uid domain.AccountId, pg pagination.Info) ([]repo.Psyna, error) {
	// TODO
	return []repo.Psyna{}, nil
}
