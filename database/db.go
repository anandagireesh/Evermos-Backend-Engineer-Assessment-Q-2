package database

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

var Db *sql.DB

//DbConnection -
func DbConnection() {
	var err error
	Db, err = sql.Open("mysql", "root:@tcp(localhost:3306)/evermos_q2?parseTime=true")

	if err != nil {
		panic(err.Error())
	}

	fmt.Println("OK")

}

//GetConnection -
func GetConnection() *sql.DB {
	return Db

}
