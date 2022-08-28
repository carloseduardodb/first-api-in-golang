package database

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql" // connection driver to mysql
)

func Connect() (*sql.DB, error) {
	stringConnection := "root:RootPassword@/Backoffice?charset=utf8&parseTime=True&loc=Local"
	db, err := sql.Open("mysql", stringConnection)
	if err != nil {
		return nil, err
	}
	if err := db.Ping(); err != nil {
		return nil, err
	}
	return db, nil
}
