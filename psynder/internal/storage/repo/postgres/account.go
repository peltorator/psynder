package postgres

import (
	"github.com/peltorator/psynder/internal/domain"
	"github.com/peltorator/psynder/internal/storage"
	"github.com/peltorator/psynder/internal/storage/repo"
	"gorm.io/gorm"
)

type Account struct {
	ID uint64
	storage.LoginCredentials
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

func (r *accountRepo) StoreNew(data storage.AccountData) (domain.AccountId, error) {
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

func (r *accountRepo) LoadByEmail(email string) (storage.Account, error) {
	var acc Account
	if err := r.db.First(&acc, "email = ?", email).Error; err != nil {
		return storage.Account{}, err
	}

	return storage.Account{
		Id: accountIdFromDb(acc.ID),
		AccountData: storage.AccountData{
			LoginCredentials: acc.LoginCredentials,
			Kind:             domain.AccountKindFromString(acc.Kind),
		},
	}, nil
}
