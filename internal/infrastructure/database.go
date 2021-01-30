package infrastructure

import (
	"database/sql"
	"fmt"

	"github.com/spf13/viper"
)

type Database struct {
	Conn *sql.DB
}

func NewDatabaseConnection() Database {
	databaseInfo := fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s",
		viper.GetString("DB_USERNAME"),
		viper.GetString("DB_PASSWORD"),
		viper.GetString("DB_HOST"),
		viper.GetInt("DB_PORT"),
		viper.GetString("DB_DATABASE"),
	)

	db, err := sql.Open(viper.GetString("DB_DRIVER"), databaseInfo)
	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	return Database{Conn: db}
}
