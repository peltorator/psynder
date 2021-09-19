package repo

import (
	"psynder/internal/domain/model"
)

type CreateAccountOptions struct {
	Email string
	PasswordHash model.PasswordHash
}

type AccountRepo interface {
	CreateAccount(opts CreateAccountOptions) (model.AccountId, error)
	GetIdByEmail(email string) (model.AccountId, error)
	GetPasswordHashById(id model.AccountId) (model.PasswordHash, error)
	//GetAccountById(id string) (model.Account, error)
	//GetAccountByEmail(email string) (model.Account, error)
}