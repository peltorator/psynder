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

func (p *shelterRepo) checkShelterAccountKind(uid domain.AccountId) error {
	var acc Account
	if err := p.db.First(&acc, "id = ?", uid).Error; err != nil {
		return err
	}
	if domain.AccountKindFromString(acc.Kind) != domain.AccountKindShelter {
		return errors.Errorf("Can't use shelter api for %s accound kind", domain.AccountKindFromString(acc.Kind))
	}
	return nil
}

func (p *shelterRepo) AddInfo(uid domain.AccountId, info domain.ShelterInfo) error {
	if err := p.checkShelterAccountKind(uid); err != nil {
		return err
	}
	shelterInfo := ShelterInfo{
		AccountId: accountIdToDb(uid),
		City:      info.City,
		Address:   info.Address,
		Phone:     info.Phone,
	}
	return p.db.Table("shelter_info").Create(&shelterInfo).Error
}

type ShelterPsynas struct {
	AccountId uint64
	PsynaId   uint64
}

func (p *shelterRepo) AddPsyna(uid domain.AccountId, pd swipe.PsynaData) (domain.PsynaId, error) {
	if err := p.checkShelterAccountKind(uid); err != nil {
		return domain.PsynaId(0), err
	}
	psyna := Psyna{
		Name:        pd.Name,
		Breed:       pd.Breed,
		Description: pd.Description,
		PhotoLink:   pd.PhotoLink,
	}
	err := p.db.Create(&psyna).Error
	if err != nil {
		return domain.PsynaId(0), err
	}
	pid := psynaIdFromDb(psyna.ID)
	shelter_psyna := ShelterPsynas{
		AccountId: uint64(uid),
		PsynaId:   uint64(pid),
	}
	err = p.db.Table("shelter_dogs").Create(&shelter_psyna).Error
	return pid, err
}

type ShelterPsyna struct {
	AccountId uint64
	PsynaId   uint64
}

func (p *shelterRepo) DeletePsyna(uid domain.AccountId, pid domain.PsynaId) error {
	var r []ShelterPsyna
	if err := p.db.Table("shelter_dogs").Where("psyna_id = ?", pid).
		Where("account_id = ?", uid).
		Find(&r).Error; err != nil {
		return err
	}
	if len(r) == 0 {
		return errors.Errorf("Can't delete another's psyna")
	}
	shelter_psyna := ShelterPsyna{
		AccountId: uint64(uid),
		PsynaId:   uint64(pid),
	}
	if err := p.db.Table("shelter_dogs").Where("psyna_id = ?", pid).
		Where("account_id = ?", uid).Delete(&shelter_psyna).Error; err != nil {
		return err
	}
	var psyna Psyna
	p.db.First(&psyna, "id = ?", pid)
	return p.db.Delete(&psyna).Error
}

func (p *shelterRepo) LoadSlice(uid domain.AccountId, pg pagination.Info) ([]repo.Psyna, error) {
	var psynaRecords []Psyna
	if err := p.db.Table("psynas").Limit(pg.Limit).Offset(pg.Offset).
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
	err := p.db.Table("ratings").Where("psyna_id = ? AND liked = ?", pid, true).Find(&psynaRecords).Error
	if err != nil {
		return 0, err
	}
	r = int64(len(psynaRecords))
	return r, nil
}
