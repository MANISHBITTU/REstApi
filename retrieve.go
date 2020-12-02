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
	rows, e := db.Query("select * from info")
	ErrorCheck(e)

	var post = Post{}

	for rows.Next() {
		e = rows.Scan(&post.Id, &post.Name, &post.City)
		ErrorCheck(e)
		fmt.Println(post)
	}
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
