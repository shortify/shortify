package app

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"os"
	"testing"
)

type ConfigSuite struct {
	suite.Suite
}

func TestConfigSuite(t *testing.T) {
	suite.Run(t, new(ConfigSuite))
}

func (suite *ConfigSuite) SetupSuite() {
	os.Setenv("TEST_ENV_VAR", "sqlite3")
}

func (suite *ConfigSuite) TearDownSuite() {
	os.Unsetenv("TEST_ENV_VAR")
}

func (suite *ConfigSuite) TestLoadConfigFromString() {
	t := suite.T()
	configString := `
	[database]
	provider = sqlite3
	dataSource = database.bin

	[settings]
	alphabet = abcdefg
	port = 8080
	`

	cfg, err := loadConfigFromString(configString)
	assert.Nil(t, err)
	assert.Equal(t, "sqlite3", cfg.Database.Provider)
	assert.Equal(t, "database.bin", cfg.Database.DataSource)
	assert.Equal(t, "abcdefg", cfg.Settings.Alphabet)
	assert.Equal(t, "8080", cfg.Settings.Port)
}

func (suite *ConfigSuite) TestLoadConfigFromFile() {
	t := suite.T()
	cfg, err := loadConfigFromFile("../examples/sqlite3.gcfg")

	assert.Nil(t, err)
	assert.Equal(t, "sqlite3", cfg.Database.Provider)
	assert.Equal(t, "/tmp/shortify_sqlite3_sample.bin", cfg.Database.DataSource)
	assert.Equal(t, "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ", cfg.Settings.Alphabet)
}

func (suite *ConfigSuite) TestEnvVarSupport() {
	t := suite.T()

	configString := `
	[database]
	provider = $TEST_ENV_VAR
	dataSource = database.bin

	[settings]
	alphabet = abcdefg
	port = 8080
	`

	cfg, err := loadConfigFromString(configString)
	assert.Nil(t, err)
	assert.Equal(t, "sqlite3", cfg.Database.Provider)
}
