package accountrepo

import (
	"errors"
	"psynder/internal/domain/model"
	"psynder/internal/domain/repo"
	"psynder/internal/interface/postgres"
)

type Account struct {
	postgres.Model
	Email        string
	PasswordHash []byte
}

type Postgres struct {
	conn *postgres.Connection
}

func New(conn *postgres.Connection) *Postgres {
	return &Postgres{conn: conn}
}

func (p *Postgres) StoreAccountToRepo(opts repo.CreateAccountOptions) (model.AccountId, error) {
	acc := Account{
		Email:        opts.Email,
		PasswordHash: opts.PasswordHash,
	}
	r := p.conn.Db.Create(&acc)
	if r.Error != nil {
		// TODO: ???
		return model.AccountId(acc.ID), errors.New("already exists")
	}
	return model.AccountId(acc.ID), nil
}

func (p *Postgres) LoadIdByEmailFromRepo(email string) (model.AccountId, error) {
	var acc Account
	r := p.conn.Db.First(&acc, "email = ?", email)
	return model.AccountId(acc.ID), r.Error
}

func (p *Postgres) LoadPasswordHashByIdFromRepo(id model.AccountId) (model.PasswordHash, error) {
	var acc Account
	r := p.conn.Db.First(&acc, id)
	// TODO
	return acc.PasswordHash, r.Error
}
