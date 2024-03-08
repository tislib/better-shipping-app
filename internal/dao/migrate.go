package dao

import (
	"better-shipping-app/internal/config"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
)

type Migrator interface {
	Migrate() error
}

type migrator struct {
	databaseConfig config.DatabaseConfig
}

func (d *migrator) Migrate() error {
	destinationUrl := prepareConnectionStr(d.databaseConfig)

	m, err := migrate.New("file://internal/dao/migrations", destinationUrl)

	if err != nil {
		return err
	}

	err = m.Up()

	if err == nil {
		return nil
	}

	if err.Error() == "no change" {
		return nil
	}

	return err
}

func NewMigrator(databaseConfig config.DatabaseConfig) Migrator {
	return &migrator{databaseConfig: databaseConfig}
}
