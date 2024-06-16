package database

import (
	"database/sql"
	"fmt"

	"github.com/geovanymds/balance/internal/infra/config"
	_ "github.com/go-sql-driver/mysql"
)

func InitDb() (*sql.DB, error) {
	config := config.NewDbConnectionConfig()

	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", config.User, config.Password, config.Host, config.Port, config.DbName))

	if err != nil {
		panic(err)
	}

	return db, nil
}
