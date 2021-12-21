package postgres

import (
	"github.com/peltorator/psynder/internal/domain"
	"github.com/peltorator/psynder/internal/pagination"
	"github.com/peltorator/psynder/internal/repo"
	"gorm.io/gorm"
)

type Psyna struct {
	ID          uint64
	Name        string
	Breed       string
	Description string
	PhotoLink   string
}

func psynaIdToDb(pid domain.PsynaId) uint64 {
	return uint64(pid)
}

func psynaIdFromDb(pid uint64) domain.PsynaId {
	return domain.PsynaId(pid)
}

type psynaRepo struct {
	db *gorm.DB
}

func NewPsynaRepo(conn *gorm.DB) *psynaRepo {
	return &psynaRepo{
		db: conn,
	}
}

func (p *psynaRepo) StoreNew(data repo.PsynaData) (domain.PsynaId, error) {
	psyna := Psyna{
		Name:        data.Name,
		Breed:       data.Breed,
		Description: data.Description,
		PhotoLink:   data.PhotoLink,
	}
	err := p.db.Save(&psyna).Error
	return psynaIdFromDb(psyna.ID), err
}

func (p *psynaRepo) LoadSlice(uid domain.AccountId, pg pagination.Info, f domain.PsynaFilter) ([]repo.Psyna, error) {
	var psynaRecords []Psyna
	table := p.db.Table("psynas").Limit(pg.Limit).Offset(pg.Offset)
	if f.Breed != nil && *f.Breed != "" {
		table = table.Where("breed = ?", *f.Breed)
	}
	if f.Shelter != nil || (f.ShelterCity != nil && *f.ShelterCity != "") {
		table = table.
			Joins("JOIN shelter_dogs ON psynas.id = shelter_dogs.psyna_id").
			Joins("JOIN shelter_infos ON shelter_infos.account_id = shelter_dogs.account_id")
		if f.Shelter != nil {
			table = table.Where("shelter_infos.account_id = ?", *f.Shelter)
		}
		if f.ShelterCity != nil && *f.ShelterCity != "" {
			table = table.Where("shelter_infos.city = ?", *f.ShelterCity)
		}
	}
	if err := table.Find(&psynaRecords).Error; err != nil {
		return nil, err
	}

	psynas := make([]repo.Psyna, len(psynaRecords))
	for i, p := range psynaRecords {
		psynas[i] = repo.Psyna{
			Id: psynaIdFromDb(p.ID),
			PsynaData: repo.PsynaData{
				Name:        p.Name,
				Breed:       p.Breed,
				Description: p.Description,
				PhotoLink:   p.PhotoLink,
			},
		}
	}
	return psynas, nil
}
