package main

import (
	"database/sql"

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
	stmt, e := db.Prepare("insert into info values (?, ?, ?)")
	ErrorCheck(e)

	//execute
	res, e := stmt.Exec(2, "Harsh", "Jaipur")
	ErrorCheck(e)
	fmt.println("data inserted")
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
