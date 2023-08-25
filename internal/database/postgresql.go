package database

import (
	"github.com/golang-migrate/migrate/v4"
	"github.com/hypertonyc/rpc-incrementor-service/internal/config"

	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
)

func Migrate(config config.AppConfig) error {
	var err error

	m, err := migrate.New(config.PgMigrationsPath, config.PgConUrl)
	if err != nil {
		return err
	}

	err = m.Up()
	if err != nil && err != migrate.ErrNoChange {
		return err
	}

	return nil
}
