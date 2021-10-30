package storage

import "github.com/peltorator/psynder/internal/domain"

type AccountData struct {
	LoginCredentials
	Kind domain.AccountKind
}

type Account struct {
	Id domain.AccountId
	AccountData
}