package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

type Post struct {
	Id   int
	Name string
	City string
}

func main() {
	dbUser := "manish"
	dbpass := "manish"
	dbName := "student"
	db, e := sql.Open("mysql", dbUser+":"+dbpass+"@/"+dbName)
	ErrorCheck(e)
	defer db.Close()

	PingDB(db)
	stmt, e := db.Prepare("update info set city=? where id=?")
	ErrorCheck(e)

	res, e := stmt.Exec("Kotputli", 2)
	ErrorCheck(e)
	fmt.Print("data updated")
}

func ErrorCheck(err error) {
	if err != nil {
		panic(err.Error())
	}
}

func PingDB(db *sql.DB) {
	err := db.Ping()
	ErrorCheck(err)
}
