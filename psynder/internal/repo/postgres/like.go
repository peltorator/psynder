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

func (l *likeRepo) GetLikedPsynas(uid domain.AccountId, pg pagination.Info) ([]repo.Psyna, error) {
	var psynaRecords []Psyna
	if err := l.db.Joins("ratings").Where("ratings.liked = ?", true).Find(&psynaRecords).Error; err != nil {
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

func (l *likeRepo) RatePsyna(uid domain.AccountId, pid domain.PsynaId, decision swipe.Decision) error {
	like := Rating{
		AccountId: uid,
		PsynaId:   pid,
		Liked:     l.decisionToDb(decision),
	}
	// TODO: test whether we really want UpdateAll here
	return l.db.Clauses(clause.OnConflict{UpdateAll: true}).Save(&like).Error
}
