package shelterservice

import (
	"github.com/peltorator/psynder/internal/domain"
	"github.com/peltorator/psynder/internal/domain/swipe"
	"github.com/peltorator/psynder/internal/pagination"
	"github.com/peltorator/psynder/internal/repo"
	"github.com/peltorator/psynder/internal/serviceimpl/swipeservice"
)

type ShelterService struct {
	shelterRepo repo.Shelters
}

func New(s repo.Shelters) *ShelterService {
	return &ShelterService{
		shelterRepo: s,
	}
}

func (s *ShelterService) AddInfo(uid domain.AccountId, info domain.ShelterInfo) error {
	return s.shelterRepo.AddInfo(uid, info)
}

func (s *ShelterService) AddPsyna(uid domain.AccountId, p swipe.PsynaData) (domain.PsynaId, error) {
	return s.shelterRepo.AddPsyna(uid, p)
}

func (s *ShelterService) DeletePsyna(uid domain.AccountId, pid domain.PsynaId) error {
	return s.shelterRepo.DeletePsyna(uid, pid)
}

func (s *ShelterService) GetMyPsynas(uid domain.AccountId, pg pagination.Info) ([]swipe.Psyna, error) {
	psynasStored, err := s.shelterRepo.LoadSlice(uid, pg)
	if err != nil {
		return nil, err
	}
	return swipeservice.PsynasStoredToSwipe(psynasStored), nil
}

func (s *ShelterService) GetPsynaLikes(pid domain.PsynaId) (int64, error) {
	r, err := s.shelterRepo.GetPsynaLikes(pid)
	if err != nil {
		return 0, err
	}
	return r, nil
}