package postgres

import (
	"github.com/peltorator/psynder/internal/domain"
	"github.com/peltorator/psynder/internal/domain/swipe"
	"github.com/peltorator/psynder/internal/pagination"
	"github.com/peltorator/psynder/internal/repo"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type Rating struct {
	AccountId domain.AccountId
	PsynaId   domain.PsynaId
	Liked     bool
}

func (l *likeRepo) decisionToDb(decision swipe.Decision) bool {
	switch decision {
	case swipe.DecisionLike:
		return true
	case swipe.DecisionSkip:
		return false
	default:
		l.logger.DPanicf("Unknown decision: %v", decision)
		return false
	}
}

type likeRepo struct {
	db     *gorm.DB
	logger *zap.SugaredLogger
}

func NewLikeRepo(conn *gorm.DB) *likeRepo {
	return &likeRepo{db: conn}
}

func psynasFromDb(psynaRecords []Psyna) []repo.Psyna {
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
	return psynas
}

func (l *likeRepo) GetLikedPsynas(uid domain.AccountId, pg pagination.Info) ([]repo.Psyna, error) {
	var psynaRecords []Psyna
	if err := l.db.Limit(pg.Limit).Offset(pg.Offset).
		Joins("JOIN ratings ON id = ratings.psyna_id").
		Where("ratings.account_id = ?", uid).
		Where("ratings.liked = ?", true).
		Find(&psynaRecords).Error; err != nil {
		return nil, err
	}

	return psynasFromDb(psynaRecords), nil
}

func (l *likeRepo) RatePsyna(uid domain.AccountId, pid domain.PsynaId, decision swipe.Decision) error {
	like := Rating{
		AccountId: uid,
		PsynaId:   pid,
		Liked:     l.decisionToDb(decision),
	}
	// TODO: test whether we really want UpdateAll here
	return l.db.Clauses(clause.OnConflict{
		UpdateAll:    true,
		OnConstraint: "pk_ratings",
	}).Create(&like).Error
}

func shelterFromDb(shelterRecord ShelterInfo) repo.Shelter {
	var shelter = repo.Shelter{
		Id: shelterIdFromDb(shelterRecord.AccountId),
		ShelterData: repo.ShelterData{
			City:    shelterRecord.City,
			Address: shelterRecord.Address,
			Phone:   shelterRecord.Phone,
		},
	}

	return shelter
}

func (l *likeRepo) GetPsynaInfo(pid domain.PsynaId) (repo.Shelter, error) {
	var shelterRecord ShelterInfo
	if err := l.db.Table("shelter_infos").
		Joins("JOIN shelter_dogs ON shelter_infos.account_id = shelter_dogs.account_id").
		Where("shelter_dogs.psyna_id = ?", pid).
		Find(&shelterRecord).Error; err != nil {
		return repo.Shelter{}, err
	}

	return shelterFromDb(shelterRecord), nil
}

func (l *likeRepo) GetAllInfo() (repo.AllInfo, error) {
	var psynaRecords []Psyna
	err1 := l.db.Table("psynas").
		Find(&psynaRecords).Error
	if err1 != nil {
		return repo.AllInfo{}, err1
	}

	var users []Account
	err2 := l.db.Table("accounts").Where("kind = ?", "person").Find(&users).Error
	if err2 != nil {
		return repo.AllInfo{}, err2
	}

	var shelters []Account
	err3 := l.db.Table("accounts").Where("kind = ?", "shelter").Find(&shelters).Error
	if err3 != nil {
		return repo.AllInfo{}, err3
	}

	return repo.AllInfo{
		Users:    int64(len(users)),
		Shelters: int64(len(shelters)),
		Psynas:   int64(len(psynaRecords)),
	}, nil
}
