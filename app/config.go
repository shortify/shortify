package app

import (
	"code.google.com/p/gcfg"
	"log"
	"os"
	"strings"
)

type appConfig struct {
	Database struct {
		Provider   string
		DataSource string
	}
	Settings struct {
		Alphabet string
		Port     string
	}
}

func (self *appConfig) sanitize() {
	self.Database.Provider = valueOrEnvVar(self.Database.Provider)
	self.Database.DataSource = valueOrEnvVar(self.Database.DataSource)
	self.Settings.Alphabet = valueOrEnvVar(self.Settings.Alphabet)
	self.Settings.Port = valueOrEnvVar(self.Settings.Port)
}

func Configure(configFile string) bool {
	cfg, err := loadConfigFromFile(configFile)

	if err != nil {
		log.Fatalf("Could not load config file from %s", configFile)
		return false
	}

	shortifyDb = newDatabase(cfg.Database.Provider, cfg.Database.DataSource)
	shortifyEncoder = encoder{cfg.Settings.Alphabet}
	shortifyPort = cfg.Settings.Port
	return true
}

func loadConfigFromString(configString string) (appConfig, error) {
	cfg := new(appConfig)
	err := gcfg.ReadStringInto(cfg, configString)

	if cfg != nil {
		cfg.sanitize()
	}

	return *cfg, err
}

func loadConfigFromFile(filePath string) (appConfig, error) {
	cfg := new(appConfig)
	err := gcfg.ReadFileInto(cfg, filePath)

	if cfg != nil {
		cfg.sanitize()
	}

	return *cfg, err
}

func valueOrEnvVar(value string) string {
	if strings.HasPrefix(value, "$") {
		return os.Getenv(strings.TrimPrefix(value, "$"))
	}

	return value
}
