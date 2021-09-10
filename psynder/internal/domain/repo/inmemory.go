package repo

import (
	"psynder/internal/domain/model"
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

func (r *InMemoryAccountRepo) CreateAccount(opts CreateAccountOptions) (model.AccountId, error) {
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

func (r *InMemoryAccountRepo) GetIdByEmail(email string) (model.AccountId, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	return r.emailToAccount[email].Id, nil
}

func (r *InMemoryAccountRepo) GetPasswordHashById(id model.AccountId) (model.PasswordHash, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	return r.idToAccount[id].PasswordHash, nil
}