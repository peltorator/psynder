package swiperepo

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

const queryLoadPsynas = `
	SELECT psynas.id, psynas.name, psynas.description, psynas.photoLink
	FROM psynas
	WHERE psynas.id > (select coalesce(max(lastView.psynaId), -1) FROM lastView WHERE lastView.accountId = $1)
	ORDER BY psynas.id
	LIMIT $2;
`

func (p *Postgres) LoadPsynas(opts repo.LoadPsynasOptions) ([]model.Psyna, error) {
	rows, err := p.conn.Query(queryLoadPsynas, opts.AccountId, opts.Count)
	if err != nil {
		return []model.Psyna{}, nil
	}
	var psynas []model.Psyna
	var count = 0
	for rows.Next() {
		count += 1
		r := new(model.Psyna)
		err = rows.Scan(&r.Id, &r.Name, &r.Description, &r.PhotoLink)
		if err != nil {
			return []model.Psyna{}, nil
		}
		psynas = append(psynas, *r)
	}
	if count == 0 {
		return []model.Psyna{}, nil
	}
	return psynas, nil
}

const querySaveLastView = `
	INSERT INTO lastView(accountId, psynaId)
	VALUES ($1, $2)
	ON CONFLICT (accountId) DO UPDATE SET psynaId=EXCLUDED.psynaId;
`

func (p *Postgres) SaveLastView(account_id model.AccountId, psyna_id model.PsynaId) error {
	row := p.conn.QueryRow(querySaveLastView, account_id, psyna_id)
	err := row.Err()
	if err != nil {
		return err
	}
	return nil
}

const queryLikePsyna = `
	INSERT INTO likes(accountId, psynaId)
	VALUES ($1, $2)
`

func (p *Postgres) LikePsyna(opts repo.LikePsynaOptions) error {
	row := p.conn.QueryRow(queryLikePsyna, opts.AccountId, opts.PsynaId)
	err := row.Err()
	if err != nil {
		return errors.New("Account or psyna doesn't exist")
	}
	return nil
}

const queryGetFavoritePsynas = `
	SELECT psynaId, name, description, photoLink
	FROM likes
	INNER JOIN psynas p ON likes.psynaId = p.id
	WHERE likes.accountId = $1;
`

func (p *Postgres) GetFavoritePsynas(id model.AccountId) ([]model.Psyna, error) {
	rows, err := p.conn.Query(queryGetFavoritePsynas, id)
	if err != nil {
		return []model.Psyna{}, err
	}
	var psynas []model.Psyna
	for rows.Next() {
		p := new(model.Psyna)
		err = rows.Scan(&p.Id, &p.Name, &p.Description, &p.PhotoLink)
		if err != nil {
			return []model.Psyna{}, err
		}
		psynas = append(psynas, *p)
	}
	return psynas, nil
}