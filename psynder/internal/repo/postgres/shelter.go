package postgres

import (
	"github.com/peltorator/psynder/internal/domain"
	"github.com/peltorator/psynder/internal/domain/swipe"
	"github.com/peltorator/psynder/internal/pagination"
	"github.com/peltorator/psynder/internal/repo"
	"github.com/pkg/errors"
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

type ShelterInfo struct {
	AccountId uint64
	City      string
	Address   string
	Phone     string
}

func shelterIdFromDb(pid uint64) domain.AccountId {
	return domain.AccountId(pid)
}

func (p *shelterRepo) AddInfo(uid domain.AccountId, info domain.ShelterInfo) error {
	shelterInfo := ShelterInfo{
		AccountId: accountIdToDb(info.AccountId),
		City:      info.City,
		Address:   info.Address,
		Phone:     info.Phone,
	}

	return p.db.Create(&shelterInfo).Error
}

func (p *shelterRepo) AddPsyna(uid domain.AccountId, pd swipe.PsynaData) (domain.PsynaId, error) {
	psyna := Psyna{
		Name:        pd.Name,
		Breed:       pd.Breed,
		Description: pd.Description,
		PhotoLink:   pd.PhotoLink,
	}
	err := p.db.Create(&psyna).Error
	return psynaIdFromDb(psyna.ID), err
}

type ShelterPsyna struct {
	AccountId uint64
	PsynaId   uint64
}

func (p *shelterRepo) DeletePsyna(uid domain.AccountId, pid domain.PsynaId) error {
	var r []ShelterPsyna
	if err := p.db.Where("psyna_id = ?", pid).
		Where("account_id = ?", uid).
		Find(&r).Error; err != nil {
		return err
	}
	if len(r) == 0 {
		return errors.Errorf("Can't delete another's psyna")
	}
	var psyna Psyna
	p.db.First(&psyna, "id = ?", pid)
	return p.db.Delete(&psyna).Error
}

func (p *shelterRepo) LoadSlice(uid domain.AccountId, pg pagination.Info) ([]repo.Psyna, error) {
	var psynaRecords []Psyna
	if err := p.db.Limit(pg.Limit).Offset(pg.Offset).
		Joins("JOIN shelter_dogs ON id = shelter_dogs.psyna_id").
		Where("shelter_dogs.account_id = ?", uid).
		Find(&psynaRecords).Error; err != nil {
		return nil, err
	}
	return psynasFromDb(psynaRecords), nil
}

func (p *shelterRepo) GetPsynaLikes(pid domain.PsynaId) (int64, error) {
	var r int64
	var psynaRecords []Psyna
	err := p.db.Table("ratings").Where( "psyna_id = ? AND liked = ?", pid, true).Find(&psynaRecords).Error
	if err != nil {
		return 0, err
	}
	r = int64(len(psynaRecords))
	return r, nil
}
