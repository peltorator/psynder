package repo

import (
	"github.com/peltorator/psynder/internal/domain/model"
)

type CreateAccountOptions struct {
	Email string
	PasswordHash model.PasswordHash
}

type AccountRepo interface {
	StoreAccountToRepo(opts CreateAccountOptions) (model.AccountId, error)
	LoadIdByEmailFromRepo(email string) (model.AccountId, error)
	LoadPasswordHashByIdFromRepo(id model.AccountId) (model.PasswordHash, error)
	//GetAccountById(id string) (model.Account, error)
	//GetAccountByEmail(email string) (model.Account, error)
}