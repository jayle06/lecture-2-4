package database

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)


func Connection() (db *sql.DB){
	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:8080)/crawl")
	if err != nil {
		panic(err.Error())
	}
	return db
}
