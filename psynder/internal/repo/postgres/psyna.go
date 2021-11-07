package postgres

import (
	"github.com/peltorator/psynder/internal/domain"
	"github.com/peltorator/psynder/internal/pagination"
	"github.com/peltorator/psynder/internal/repo"
	"gorm.io/gorm"
)

type Psyna struct {
	ID uint64
	Name string
	Description string
	PhotoLink string
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
		Description: data.Description,
		PhotoLink:   data.PhotoLink,
	}
	err := p.db.Save(&psyna).Error
	return psynaIdFromDb(psyna.ID), err
}

func (p *psynaRepo) LoadSlice(uid domain.AccountId, pg pagination.Info) ([]repo.Psyna, error) {
	var psynaRecords []Psyna
	if err := p.db.Limit(pg.Limit).Offset(pg.Offset).Find(&psynaRecords).Error; err != nil {
		return nil, err
	}

	psynas := make([]repo.Psyna, len(psynaRecords))
	for i, p := range psynaRecords {
		psynas[i] = repo.Psyna{
			Id: psynaIdFromDb(p.ID),
			PsynaData: repo.PsynaData{
				Name:        p.Name,
				Description: p.Description,
				PhotoLink:   p.PhotoLink,
			},
		}
	}
	return psynas, nil
}



