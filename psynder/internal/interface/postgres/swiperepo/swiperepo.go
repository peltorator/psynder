package swiperepo

import (
	"errors"
	"github.com/peltorator/psynder/internal/domain/model"
	"github.com/peltorator/psynder/internal/domain/repo"
	"github.com/peltorator/psynder/internal/interface/postgres"
)

type Psyna struct {
	postgres.Model
	Name        string
	Description string
	PhotoLink   string
}

type Like struct {
	AccountId uint
	PsynaId   uint
}

type Postgres struct {
	conn *postgres.Connection
}

func New(conn *postgres.Connection) *Postgres {
	return &Postgres{conn: conn}
}

func (p *Postgres) LoadPsynasFromRepo(opts repo.LoadPsynasOptions) ([]model.Psyna, error) {
	var ps []Psyna
	r := p.conn.Db.Limit(opts.Limit).Offset(opts.Offset).Find(&ps)
	var psynas []model.Psyna
	for _, psyna := range ps {
		psynas = append(psynas, model.Psyna{
			Id:          model.PsynaId(psyna.ID),
			Name:        psyna.Name,
			Description: psyna.Description,
			PhotoLink:   psyna.PhotoLink})
	}
	return psynas, r.Error
}

func (p *Postgres) StoreLikeToRepo(opts repo.LikePsynaOptions) error {
	like := Like{
		AccountId: uint(opts.AccountId),
		PsynaId:   uint(opts.PsynaId),
	}
	r := p.conn.Db.Create(&like)
	if r.Error != nil {
		// TODO: ???
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

func (p *Postgres) LoadFavoritePsynasFromRepo(id model.AccountId) ([]model.Psyna, error) {
	//rows, err := p.conn.Query(queryGetFavoritePsynas, id)
	//if err != nil {
	//	return []model.Psyna{}, err
	//}
	//var psynas []model.Psyna
	//for rows.Next() {
	//	p := new(model.Psyna)
	//	err = rows.Scan(&p.Id, &p.Name, &p.Description, &p.PhotoLink)
	//	if err != nil {
	//		return []model.Psyna{}, err
	//	}
	//	psynas = append(psynas, *p)
	//}
	//return psynas, nil
	var ps []Psyna

	r := p.conn.Db.Joins("likes").Find(&ps)
	// TODO
	var psynas []model.Psyna
	for _, psyna := range ps {
		psynas = append(psynas, model.Psyna{
			model.PsynaId(psyna.ID),
			psyna.Name,
			psyna.Description,
			psyna.PhotoLink})
	}
	return psynas, r.Error
}
