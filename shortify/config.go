package shortify

import (
	"code.google.com/p/gcfg"
	"log"
)

var shortifyConfig appConfig

type appConfig struct {
	Database struct {
		Provider   string
		DataSource string
	}
	Settings struct {
		Alphabet string
	}
}

func Configure(configFile string) bool {
	cfg, err := loadConfigFromFile(configFile)

	if err != nil {
		log.Fatalf("Could not load config file from %s", configFile)
		return false
	}

	shortifyConfig = cfg
	shortifyEncoder = encoder{cfg.Settings.Alphabet}
	return true
}

func loadConfigFromString(configString string) (appConfig, error) {
	cfg := new(appConfig)
	err := gcfg.ReadStringInto(cfg, configString)
	return *cfg, err
}

func loadConfigFromFile(filePath string) (appConfig, error) {
	cfg := new(appConfig)
	err := gcfg.ReadFileInto(cfg, filePath)
	return *cfg, err
}
