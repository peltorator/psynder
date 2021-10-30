package usecases

import (
	"psynder/internal/domain/model"
	"psynder/internal/domain/repo"
)



type SwipeUseCases interface {
	GetPsynas(opts repo.LoadPsynasOptions) ([]model.Psyna, error)
	LikePsyna(opts repo.LikePsynaOptions) error
	GetFavoritePsynas(id model.AccountId) ([]model.Psyna, error)
}

type SwipeUseCasesImpl struct {
	SwipeRepo repo.SwipeRepo
}

func NewSwipeUseCases(swipeRepo repo.SwipeRepo) *SwipeUseCasesImpl {
	return &SwipeUseCasesImpl{
		SwipeRepo: swipeRepo,
	}
}

func (u *SwipeUseCasesImpl) GetPsynas(opts repo.LoadPsynasOptions) ([]model.Psyna, error) {
	psynas, err := u.SwipeRepo.LoadPsynasFromRepo(opts)
	return psynas, err
}

func (u *SwipeUseCasesImpl) LikePsyna(opts repo.LikePsynaOptions) error {
	err := u.SwipeRepo.StoreLikeToRepo(opts)
	return err
}

func (u *SwipeUseCasesImpl) GetFavoritePsynas(id model.AccountId) ([]model.Psyna, error) {
	psynas, err := u.SwipeRepo.LoadFavoritePsynasFromRepo(id)
	return psynas, err
}