package db

import (
	"github.com/jmoiron/sqlx"
	migrate "github.com/lancer-kit/sql-migrate"
	"github.com/pkg/errors"
)

// MigrateDir represents a direction in which to perform schema migrations.
type MigrateDir string

const (
	// MigrateUp causes migrations to be run in the "up" direction.
	MigrateUp MigrateDir = "up"
	// MigrateDown causes migrations to be run in the "down" direction.
	MigrateDown MigrateDir = "down"
)

// MigrateSet applies given set of migrations to SQL database.
func MigrateSet(connStr, driver string, dir MigrateDir, set Migrations) (int, error) {
	applier, err := NewExecutor(connStr, driver)
	if err != nil {
		return 0, err
	}
	return applier.SetMigrations(set).Migrate(dir)
}

// MigrateMultiSet applies the given list of sets of migrations to the SQL database.
// Some cases for the multiset:
//	— separate sets of migrations for different database schemas;
//	— separate sets of migrations for different purposes (ex. one for tables, another for functions & procedures).
// Each set bounded to own table with migration history.
// When applying multiset UP, executor begins from the first set to last.
// When applying multiset DOWN, executor begins from the last set to first.
func MigrateMultiSet(connStr, driver string, dir MigrateDir, sets ...Migrations) (int, error) {
	switch dir {
	case MigrateUp:
		return migrateMultiSetUp(connStr, driver, sets...)
	case MigrateDown:
		return migrateMultiSetDown(connStr, driver, sets...)

	}
	return 0, nil
}

func migrateMultiSetUp(connStr, driver string, sets ...Migrations) (int, error) {
	var total int

	applier, err := NewExecutor(connStr, driver)
	if err != nil {
		return 0, err
	}

	for _, migrations := range sets {
		count, err := applier.SetMigrations(migrations).Migrate(MigrateUp)
		if err != nil {
			return 0, err
		}
		total += count
	}

	return total, nil
}

func migrateMultiSetDown(connStr, driver string, sets ...Migrations) (int, error) {
	var total int

	applier, err := NewExecutor(connStr, driver)
	if err != nil {
		return 0, err
	}

	for i := len(sets) - 1; i >= 0; i-- {
		count, err := applier.SetMigrations(sets[i]).Migrate(MigrateDown)
		if err != nil {
			return 0, err
		}
		total += count
	}

	return total, nil
}

// Migrations is a configuration of migration set.
type Migrations struct {
	// Table name of the table used to store migration info.
	Table string
	// Schema that the migration table be referenced.
	Schema string
	// EnablePatchMode enables patch mode for migrations
	// Now it requires a new migration name format: 0001_00_name.sql and new table structure for save migrations
	EnablePatchMode bool
	// IgnoreUnknown skips the check to see if there is a migration
	// ran in the database that is not in MigrationSource.
	IgnoreUnknown bool
	// Assets are sql-migrate.MigrationSource assets. Ex.:
	// 	- migrate.HttpFileSystemMigrationSource
	// 	- migrate.FileMigrationSource
	// 	- migrate.AssetMigrationSource
	// 	- migrate.PackrMigrationSource
	Assets migrate.MigrationSource
}

// MigrationsExecutor is a helper that initializes database connection and applies migrations to the database.
type MigrationsExecutor struct {
	Migrations

	connStr string
	driver  string
	db      *sqlx.DB
}

// NewExecutor returns new MigrationsExecutor.
func NewExecutor(connStr, driver string) (*MigrationsExecutor, error) {
	if driver == "" {
		driver = "postgres"
	}
	conn, err := sqlx.Open(driver, connStr)
	if err != nil {
		return nil, errors.Wrap(err, "unable to connect to the database")
	}

	return &MigrationsExecutor{connStr: connStr, db: conn, driver: driver}, nil
}

// SetMigrations sets Migrations for executor.
func (executor *MigrationsExecutor) SetMigrations(set Migrations) *MigrationsExecutor {
	executor.Migrations = set
	return executor
}

// Migrate connects to the database and applies migrations.
func (executor *MigrationsExecutor) Migrate(dir MigrateDir) (int, error) {
	if executor.Assets == nil {
		return 0, errors.New("migrations isn't set")
	}

	if executor.db == nil {
		var err error
		executor.db, err = sqlx.Open(executor.driver, executor.connStr)
		if err != nil {
			return 0, errors.Wrap(err, "unable to connect to the database")
		}

	}

	applier := migrate.MigrationSet{
		TableName:       executor.Table,
		SchemaName:      executor.Schema,
		EnablePatchMode: executor.EnablePatchMode,
		IgnoreUnknown:   false,
	}

	return applier.Exec(executor.db.DB, executor.driver, executor.Assets, executor.directions()[dir])
}

func (executor *MigrationsExecutor) directions() map[MigrateDir]migrate.MigrationDirection {
	return map[MigrateDir]migrate.MigrationDirection{
		MigrateUp:   migrate.Up,
		MigrateDown: migrate.Down,
	}
}
