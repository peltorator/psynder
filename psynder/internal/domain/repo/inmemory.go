package repo

import (
	"github.com/peltorator/psynder/internal/domain/model"
	"sync"
)

type InMemoryAccountModel struct {
	Id model.AccountId
	Email string
	PasswordHash model.PasswordHash
}

type InMemoryAccountRepo struct {
	idToAccount map[model.AccountId]*InMemoryAccountModel
	emailToAccount map[string]*InMemoryAccountModel
	nextId uint64
	mu *sync.RWMutex
}

func NewInMemoryAccountRepo() *InMemoryAccountRepo {
	return &InMemoryAccountRepo{
		idToAccount:    make(map[model.AccountId]*InMemoryAccountModel),
		emailToAccount: make(map[string]*InMemoryAccountModel),
		mu:             &sync.RWMutex{},
	}
}

func (r *InMemoryAccountRepo) StoreAccountToRepo(opts CreateAccountOptions) (model.AccountId, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	r.nextId++
	acc := &InMemoryAccountModel{
		Id:           model.AccountId(r.nextId),
		Email:        opts.Email,
		PasswordHash: opts.PasswordHash,
	}
	r.idToAccount[acc.Id] = acc
	r.emailToAccount[acc.Email] = acc
	return acc.Id, nil
}

func (r *InMemoryAccountRepo) LoadIdByEmailFromRepo(email string) (model.AccountId, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	return r.emailToAccount[email].Id, nil
}

func (r *InMemoryAccountRepo) LoadPasswordHashByIdFromRepo(id model.AccountId) (model.PasswordHash, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	return r.idToAccount[id].PasswordHash, nil
}





type Views struct {
	accountId model.AccountId
	psynaId model.PsynaId
	liked bool
}

type InMemorySwipeRepo struct {
	idToPsyna map[model.PsynaId]*model.Psyna
	views []Views
	mu *sync.RWMutex
}

func NewInMemorySwipeRepo() *InMemorySwipeRepo {
	return &InMemorySwipeRepo{
		idToPsyna: make(map[model.PsynaId]*model.Psyna),
		views: []Views{},
		mu:             &sync.RWMutex{},
	}
}

func (r *InMemorySwipeRepo) LoadPsynasFromRepo(opts LoadPsynasOptions) ([]model.Psyna, error) {
	return [] model.Psyna{}, nil
}

//func (r *InMemorySwipeRepo) SaveLastView(id model.AccountId, psynas_id []model.PsynaId) error {
//	return nil
//}

func (r *InMemorySwipeRepo) StoreLikeToRepo(opts LikePsynaOptions) error {
	return nil
}

func (r *InMemorySwipeRepo) LoadFavoritePsynasFromRepo(id model.AccountId) ([]model.Psyna, error) {
	return [] model.Psyna{}, nil
}