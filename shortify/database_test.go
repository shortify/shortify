package shortify

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

func (suite *DatabaseSuite) TestdbConnectionInfo() {
	t := suite.T()
	mysql := dbConnectionInfo{"mysql", ""}
	postgres := dbConnectionInfo{"postgres", ""}
	sqlite := dbConnectionInfo{"sqlite", ""}

	assert.IsType(t, gorp.MySQLDialect{}, mysql.Dialect())
	assert.IsType(t, gorp.PostgresDialect{}, postgres.Dialect())
	assert.IsType(t, gorp.SqliteDialect{}, sqlite.Dialect())
}
