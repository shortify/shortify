package shortify

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"testing"
)

type ConfigSuite struct {
	suite.Suite
}

func TestConfigSuite(t *testing.T) {
	suite.Run(t, new(ConfigSuite))
}

func (suite *ConfigSuite) TestLoadConfigFromString() {
	t := suite.T()
	configString := `
	[database]
	provider = sqlite3
	dataSource = database.bin

	[settings]
	alphabet = abcdefg
	`

	cfg, err := loadConfigFromString(configString)
	assert.Nil(t, err)
	assert.Equal(t, "sqlite3", cfg.Database.Provider)
	assert.Equal(t, "database.bin", cfg.Database.DataSource)
	assert.Equal(t, "abcdefg", cfg.Settings.Alphabet)
}

func (suite *ConfigSuite) TestLoadConfigFromFile() {
	t := suite.T()
	cfg, err := loadConfigFromFile("../examples/sqlite3.gcfg")

	assert.Nil(t, err)
	assert.Equal(t, "sqlite3", cfg.Database.Provider)
	assert.Equal(t, "/tmp/shortify_sqlite3_sample.bin", cfg.Database.DataSource)
	assert.Equal(t, "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ", cfg.Settings.Alphabet)
}
