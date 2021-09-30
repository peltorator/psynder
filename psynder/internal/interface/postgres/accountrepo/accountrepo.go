package accountrepo

import (
	"database/sql"
	"errors"
	"psynder/internal/domain/model"
	"psynder/internal/domain/repo"
)

type Postgres struct {
	conn *sql.DB
}

func New(conn *sql.DB) *Postgres {
	return &Postgres{conn: conn}
}

const queryCreateAccount = `
	INSERT INTO accounts(
		email,
		password_hash
	) VALUES ($1, $2)
	ON CONFLICT DO NOTHING
	RETURNING id
`

func (p *Postgres) CreateAccount(opts repo.CreateAccountOptions) (model.AccountId, error) {
	var a model.AccountId
	row := p.conn.QueryRow(queryCreateAccount, opts.Email, opts.PasswordHash)
	err := row.Scan(&a)
	if err != nil {
		if err == sql.ErrNoRows {
			// TODO: АНТОХА А НУ КА БЫСТРО НАПИШИ ОШИБКИ
			return a, errors.New("Already exist")
		}
		return a, err
	}
	return a, nil
}

const queryGetIdByEmail = `
	SELECT
		id
	FROM accounts 
	WHERE email = $1
`

func (p *Postgres) GetIdByEmail(email string) (model.AccountId, error) {
	var a model.AccountId
	row := p.conn.QueryRow(queryGetIdByEmail, email)
	err := row.Scan(&a)
	if err != nil {
		if err == sql.ErrNoRows {
			// TODO: АНТОХА А НУ КА БЫСТРО НАПИШИ ОШИБКИ
			return a, errors.New("Email doesn't exist")
		}
		return a, err
	}
	return a, nil
}

const queryGetPasswordHashById = `
	SELECT
		password_hash
	FROM accounts 
	WHERE id = $1
`

func (p *Postgres) GetPasswordHashById(id model.AccountId) (model.PasswordHash, error) {
	var a model.PasswordHash
	row := p.conn.QueryRow(queryGetPasswordHashById, id)
	err := row.Scan(&a)
	if err != nil {
		if err == sql.ErrNoRows {
			// TODO: АНТОХА А НУ КА БЫСТРО НАПИШИ ОШИБКИ
			return a, errors.New("Id doesn't exist")
		}
		return a, err
	}
	return a, nil
}
