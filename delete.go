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
	stmt, e := db.Prepare("delete from posts where id=?")
	ErrorCheck(e)

	// delete 5th post
	res, e := stmt.Exec("5")
	ErrorCheck(e)

	// affected rows
	a, e := res.RowsAffected()
	ErrorCheck(e)

	fmt.Println(a)
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
