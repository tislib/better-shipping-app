package dao

import (
	"better-shipping-app/internal/config"
	"database/sql"
	_ "github.com/lib/pq"
)

type DbShell interface {
	getDb() *sql.DB
}

type dbShell struct {
	db *sql.DB
}

func (d dbShell) getDb() *sql.DB {
	return d.db
}

func NewDbShell(config config.DatabaseConfig) (DbShell, error) {
	db, err := NewDb(config)
	if err != nil {
		return nil, err
	}

	return &dbShell{
		db: db,
	}, nil
}
