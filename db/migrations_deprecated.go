package db

import (
	"database/sql"

	migrate "github.com/lancer-kit/sql-migrate"
	"github.com/pkg/errors"
)

// DEPRECATED
func directions() map[MigrateDir]migrate.MigrationDirection {
	return map[MigrateDir]migrate.MigrationDirection{
		MigrateUp:   migrate.Up,
		MigrateDown: migrate.Down,
	}
}

// nolint:gochecknoglobals
// DEPRECATED migrations represents all of the schema migration for service
var migrations *migrate.AssetMigrationSource

// DEPRECATED
// SetAssets is a function for injection of precompiled by bindata migrations files.
func SetAssets(assets migrate.AssetMigrationSource) {
	migrations = &assets
}

// DEPRECATED
// Migrate connects to the database and applies migrations.
func Migrate(connStr string, dir MigrateDir) (int, error) {
	if migrations == nil {
		return 0, errors.New("migrations isn't set")
	}
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return 0, errors.Wrap(err, "unable to connect to the database")
	}

	return migrate.Exec(db, "postgres", migrations, directions()[dir])
}
