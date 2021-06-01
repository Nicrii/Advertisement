package configuration

import (
	"os"
)

type Configuration struct {
	Host     string
	Port     string
	User     string
	Password string
	Dbname   string
}

func (config *Configuration) ReadConfiguration() {
	config.Dbname = os.Getenv("POSTGRES_DB")
	config.Password = os.Getenv("POSTGRES_PASSWORD")
	config.User = os.Getenv("POSTGRES_USER")
	config.Host = "localhost"
	config.Port = "5432"

	dbHost := os.Getenv("DB_HOST")

	if len(dbHost) > 0 {
		config.Host = dbHost
	}

	dbPort := os.Getenv("DB_PORT")

	if len(dbPort) > 0 {
		config.Port = dbPort
	}

	//file, err := os.Open("./config.json")
	//
	//if err != nil {
	//	return err
	//}
	//defer file.Close()
	//decoder := json.NewDecoder(file)
	//err = decoder.Decode(&config)
	//
	//if err != nil {
	//	panic(err)
	//	return err
	//}
	//return err
}
