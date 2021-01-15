package controllers

import (
	"database/sql"
	"github.com/stretchr/testify/mock"
)

type DbMock struct {
	mock.Mock
	MigrationSource string
}

func (d *DbMock) Migrate(repo, path, branch string) error {
	args := d.Called(repo, path, branch)
	return args.Error(0)
}

func (d *DbMock) Ping() error {
	args := d.Called()
	return args.Error(0)
}

func (d *DbMock) Raw() *sql.DB {
	args := d.Called()
	return args.Get(0).(*sql.DB)
}

