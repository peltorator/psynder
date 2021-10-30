package postgres

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Connection struct {
	Db *gorm.DB
}

func New(dsn string) (*Connection, error) {
	db, err := gorm.Open(postgres.Open(dsn))
	if err != nil {
		return nil, err
	}
	return &Connection{Db: db}, nil
}

type Model struct {
	ID uint `gorm:"primarykey"`
}
