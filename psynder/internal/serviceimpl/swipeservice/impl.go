package swipeservice

import (
	"github.com/peltorator/psynder/internal/domain"
	"github.com/peltorator/psynder/internal/domain/swipe"
	"github.com/peltorator/psynder/internal/pagination"
	"github.com/peltorator/psynder/internal/repo"
)

type swipeService struct {
	psyRepo repo.Psynas
	likeRepo repo.Likes
}

type Args struct {
	Psynas repo.Psynas
	Likes repo.Likes
}

func New(args Args) *swipeService {
	return &swipeService{
		psyRepo: args.Psynas,
		likeRepo: args.Likes,
	}
}

func (s *swipeService) BrowsePsynas(uid domain.AccountId, pg pagination.Info) ([]swipe.Psyna, error) {
	psynasStored, err := s.psyRepo.LoadSlice(uid, pg)
	if err != nil {
		return nil, err
	}

	return psynasStoredToSwipe(psynasStored), nil
}

func (s *swipeService) GetLikedPsynas(uid domain.AccountId, pg pagination.Info) ([]swipe.Psyna, error) {
	psynasStored, err := s.likeRepo.GetLikedPsynas(uid, pg)
	if err != nil {
		return nil, err
	}
	return psynasStoredToSwipe(psynasStored), nil
}

func (s *swipeService) RatePsyna(uid domain.AccountId, pid domain.PsynaId, decision swipe.Decision) error {
	return s.likeRepo.RatePsyna(uid, pid, decision)
}

func psynaStoredToSwipe(p repo.Psyna) swipe.Psyna {
	return swipe.Psyna{
		Id:          p.Id,
		Name:        p.Name,
		Description: p.Description,
		PhotoLink:   p.PhotoLink,
	}
}

func psynasStoredToSwipe(ps []repo.Psyna) []swipe.Psyna {
	psynas := make([]swipe.Psyna, len(ps))
	for i, psynaStored := range ps {
		psynas[i] = psynaStoredToSwipe(psynaStored)
	}
	return psynas
}
