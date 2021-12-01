package swipeservice

import (
	"github.com/peltorator/psynder/internal/domain"
	"github.com/peltorator/psynder/internal/domain/swipe"
	"github.com/peltorator/psynder/internal/pagination"
	"github.com/peltorator/psynder/internal/repo"
)

type SwipeService struct {
	psyRepo  repo.Psynas
	likeRepo repo.Likes
}

type Args struct {
	Psynas repo.Psynas
	Likes  repo.Likes
}

func New(args Args) *SwipeService {
	return &SwipeService{
		psyRepo:  args.Psynas,
		likeRepo: args.Likes,
	}
}

func (s *SwipeService) BrowsePsynas(uid domain.AccountId, pg pagination.Info, f domain.PsynaFilter) ([]swipe.Psyna, error) {
	psynasStored, err := s.psyRepo.LoadSlice(uid, pg, f)
	if err != nil {
		return nil, err
	}

	return PsynasStoredToSwipe(psynasStored), nil
}

func (s *SwipeService) GetLikedPsynas(uid domain.AccountId, pg pagination.Info) ([]swipe.Psyna, error) {
	psynasStored, err := s.likeRepo.GetLikedPsynas(uid, pg)
	if err != nil {
		return nil, err
	}
	return PsynasStoredToSwipe(psynasStored), nil
}

func (s *SwipeService) RatePsyna(uid domain.AccountId, pid domain.PsynaId, decision swipe.Decision) error {
	return s.likeRepo.RatePsyna(uid, pid, decision)
}

func (s *SwipeService) GetPsynaInfo(pid domain.PsynaId) (swipe.Shelter, error) {
	shelterStored, err := s.likeRepo.GetPsynaInfo(pid)
	if err != nil {
		return swipe.Shelter{}, err
	}
	return shelterStoredToSwipe(shelterStored), nil
}

func (s *SwipeService) GetAllInfo() (swipe.AllInfo, error) {
	allInfo, err := s.likeRepo.GetAllInfo()
	if err != nil {
		return swipe.AllInfo{}, err
	}
	return swipe.AllInfo{
		Users:    allInfo.Users,
		Psynas:   allInfo.Psynas,
		Shelters: allInfo.Shelters,
	}, nil
}

func psynaStoredToSwipe(p repo.Psyna) swipe.Psyna {
	return swipe.Psyna{
		Id: p.Id,
		PsynaData: swipe.PsynaData{
			Name:        p.Name,
			Breed:       p.Breed,
			Description: p.Description,
			PhotoLink:   p.PhotoLink,
		},
	}
}

func shelterStoredToSwipe(p repo.Shelter) swipe.Shelter {
	return swipe.Shelter{
		Id: p.Id,
		ShelterData: swipe.ShelterData{
			City:    p.City,
			Address: p.Address,
			Phone:   p.Phone,
		},
	}
}

func PsynasStoredToSwipe(ps []repo.Psyna) []swipe.Psyna {
	psynas := make([]swipe.Psyna, len(ps))
	for i, psynaStored := range ps {
		psynas[i] = psynaStoredToSwipe(psynaStored)
	}
	return psynas
}
