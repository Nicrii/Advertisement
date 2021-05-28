package configuration

import (
	"encoding/json"
	"os"
)

type Configuration struct {
	Host     string
	Port     string
	User     string
	Password string
	Dbname   string
}

func (config *Configuration) ReadConfiguration() error {
	file, err := os.Open("./config.json")

	if err != nil {
		return err
	}
	defer file.Close()
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&config)

	if err != nil {
		panic(err)
		return err
	}
	return err
}
