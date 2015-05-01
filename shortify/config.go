package shortify

import "code.google.com/p/gcfg"

type config struct {
	Database struct {
		Provider   string
		DataSource string
	}
	Settings struct {
		Alphabet string
	}
}

func loadConfigFromString(configString string) (config, error) {
	config := new(config)
	err := gcfg.ReadStringInto(config, configString)
	return *config, err
}

func loadConfigFromFile(filePath string) (config, error) {
	config := new(config)
	err := gcfg.ReadFileInto(config, filePath)
	return *config, err
}
