package config

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func GetConnection() (*sql.DB, error) {
	var dsn = fmt.Sprintf("%s:%s@/%s", "root", "", "go_graphql")
	//db,err := sql.Open("mysql","user:pass@tcp(xxx.xxx.xxx.xxx:PORT)/namedb")
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		panic(err.Error())
	}
	err = db.Ping()
	if err != nil {
		panic(err.Error())
	}
	return db, err
}
