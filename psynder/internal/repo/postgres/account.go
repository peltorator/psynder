package postgres

import (
	"github.com/peltorator/psynder/internal/domain"
	"github.com/peltorator/psynder/internal/repo"
	"gorm.io/gorm"
)

type Account struct {
	ID uint64
	repo.LoginCredentials
	Kind string `sql:"type:account_kind"`
}

func accountIdToDb(id domain.AccountId) uint64 {
	return uint64(id)
}

func accountIdFromDb(id uint64) domain.AccountId {
	return domain.AccountId(id)
}

type accountRepo struct {
	db *gorm.DB
}

func NewAccountRepo(conn *gorm.DB) *accountRepo {
	return &accountRepo{
		db: conn,
	}
}

func (r *accountRepo) StoreNew(data repo.AccountData) (domain.AccountId, error) {
	acc := Account{
		LoginCredentials: data.LoginCredentials,
		Kind:             data.Kind.String(),
	}
	var sameEmailCount int64
	if err := r.db.Model(&acc).Where("email = ?", data.Email).Count(&sameEmailCount).Error; err != nil {
		return 0, err
	}
	if sameEmailCount != 0 {
		return 0, repo.AccountStoreError{Kind: repo.AccountStoreErrorDuplicate}
	}
	err := r.db.Create(&acc).Error
	return accountIdFromDb(acc.ID), err
}

func (r *accountRepo) LoadByEmail(email string) (repo.Account, error) {
	var acc Account
	if err := r.db.First(&acc, "email = ?", email).Error; err != nil {
		return repo.Account{}, err
	}

	return repo.Account{
		Id: accountIdFromDb(acc.ID),
		AccountData: repo.AccountData{
			LoginCredentials: acc.LoginCredentials,
			Kind:             domain.AccountKindFromString(acc.Kind),
		},
	}, nil
}
