package data

import (
	"database/sql"
	"fmt"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	"github.com/jinzhu/gorm"
	"github.com/pkg/errors"
)

type IConnection interface{
	// Migrate will apply the migrations found under `path` on `branch` at `github.com/[repo]`.
	Migrate(repo, path, branch string) error

	// Ping is used to test the connection with the underlying database.
	Ping() error

	// Raw returns the underlying database connection.
	Raw() *sql.DB
}

// Connection is the default implementation of the IConnection interface utilizing Gorm.
type Connection struct {
	db *gorm.DB
}

func NewConnection(connectionString string) (Connection, error) {
	db, err := gorm.Open("postgres", connectionString)
	if err != nil {
		return Connection{}, errors.Errorf("failed to create new connection: %v", err)
	}

	return Connection{db: db}, nil
}

func (c Connection) Migrate(repo, path, branch string) error {
	drv, err := postgres.WithInstance(c.Raw(), &postgres.Config{})
	if err != nil {
		return errors.WithStack(err)
	}

	source := fmt.Sprintf("github://:@%s/%s#%s", repo, path, branch)

	m, err := migrate.NewWithDatabaseInstance(source, "postgres", drv)
	if err != nil {
		return errors.Errorf("failed to create a new migration: %v", err)
	}

	err = m.Up()
	if err != nil {
		if err.Error() == "no change" {
			return err
		}
		return errors.Errorf("migration failed: %v", err)
	}

	return nil
}

func (c Connection) Ping() error {
	return c.Raw().Ping()
}

func (c Connection) Raw() *sql.DB {
	return c.db.DB()
}