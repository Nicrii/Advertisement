package ads_db

import (
	"database/sql"
	"fmt"
	"github.com/Nicrii/Advertisement/configuration"
	_ "github.com/lib/pq"
)

var (
	Client *sql.DB
)

func init() {
	var err error
	config := configuration.Configuration{}
	config.ReadConfiguration()
	dbConnectionString := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		config.Host, config.Port, config.User, config.Password, config.Dbname,
	)

	Client, err = sql.Open("postgres", dbConnectionString)
	if err != nil {
		panic(err)
	}
	err = Client.Ping()
	if err != nil {
		panic(err)
	}
}
