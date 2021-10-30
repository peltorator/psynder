package swipeservice

import (
	"github.com/peltorator/psynder/internal/domain"
	"github.com/peltorator/psynder/internal/domain/swipe"
	"github.com/peltorator/psynder/internal/pagination"
	"github.com/peltorator/psynder/internal/storage/repo"
)

type swipeService struct {
	psyRepo repo.Psynas
}

func New(psynasRepo repo.Psynas) *swipeService {
	return &swipeService{
		psyRepo: psynasRepo,
	}
}

func (s *swipeService) BrowsePsynas(uid domain.AccountId, pg pagination.Info) ([]swipe.Psyna, error) {
	psynasStored, err := s.psyRepo.LoadSlice(uid, pg)
	if err != nil {
		return nil, err
	}

	psynas := make([]swipe.Psyna, len(psynasStored))
	for i, p := range psynasStored {
		psynas[i] = swipe.Psyna{
			Id:          p.Id,
			Name:        p.Name,
			Description: p.Description,
			PhotoLink:   p.PhotoLink,
		}
	}
	return psynas, nil
}

func (s *swipeService) GetLikedPsynas(uid domain.AccountId, pg pagination.Info) ([]swipe.Psyna, error) {
	panic("implement me")
}

func (s *swipeService) RatePsyna(uid domain.AccountId, args struct {
	Pid      domain.PsynaId
	Decision swipe.Decision
}) error {
	panic("implement me")
}


