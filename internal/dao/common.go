package dao

import (
	"better-shipping-app/internal/config"
	"context"
	"database/sql"
	"fmt"
	"net/url"
)

func NewDb(config config.DatabaseConfig) (*sql.DB, error) {
	connectionStr := prepareConnectionStr(config)

	db, err := sql.Open("postgres", connectionStr)
	if err != nil {
		return nil, err
	}

	if err := db.PingContext(context.TODO()); err != nil {
		return nil, err
	}

	return db, nil
}

func prepareConnectionStr(config config.DatabaseConfig) string {
	return fmt.Sprintf("postgresql://%s:%s@%s:%d/%s?sslmode=disable&search_path=%s", config.Username, url.QueryEscape(config.Password), config.Host, config.Port, config.Database, config.Schema)
}
