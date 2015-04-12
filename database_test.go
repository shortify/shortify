package main

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"gopkg.in/gorp.v1"
	"testing"
)

type DatabaseSuite struct {
	suite.Suite
}

func TestDatabaseSuite(t *testing.T) {
	suite.Run(t, new(DatabaseSuite))
}

func (suite *DatabaseSuite) TestDbConnectionDetails() {
	t := suite.T()
	mysql := DbConnectionDetails{"mysql", ""}
	postgres := DbConnectionDetails{"postgres", ""}
	sqlite := DbConnectionDetails{"sqlite", ""}

	assert.IsType(t, gorp.MySQLDialect{}, mysql.Dialect())
	assert.IsType(t, gorp.PostgresDialect{}, postgres.Dialect())
	assert.IsType(t, gorp.SqliteDialect{}, sqlite.Dialect())
}
