package db

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

func DbConnection() *sql.DB {
	dbDriver := "mysql"
	dbUser := "root"
	dbPass := "password"
	dbName := "bonder"
	db, err := sql.Open(dbDriver, dbUser+":"+dbPass+"@/"+dbName)
	if err != nil {
		panic(err.Error())
	}
	return db
}
