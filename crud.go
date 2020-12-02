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

	// close database after all work is done
	defer db.Close()

	PingDB(db)

	// INSERT INTO DB
	// prepare
	// stmt, e := db.Prepare("insert into info values (?, ?, ?)")
	// ErrorCheck(e)

	// //execute
	// res, e := stmt.Exec(2, "Harsh", "Jaipur")
	// ErrorCheck(e)

	// id, e := res.LastInsertId()
	// ErrorCheck(e)

	// fmt.Println("Insert id", id)

	//Update db
	// stmt, e := db.Prepare("update info set city=? where id=?")
	// ErrorCheck(e)

	// // execute
	// res, e := stmt.Exec("Kotputli", 2)
	// ErrorCheck(e)

	// a, e := res.RowsAffected()
	// ErrorCheck(e)

	// fmt.Println(a)

	// // query all data
	rows, e := db.Query("select * from info")
	ErrorCheck(e)

	var post = Post{}

	for rows.Next() {
		e = rows.Scan(&post.Id, &post.Name, &post.City)
		ErrorCheck(e)
		fmt.Println(post)
	}

	// delete data
	// stmt, e := db.Prepare("delete from posts where id=?")
	// ErrorCheck(e)

	// // delete 5th post
	// res, e := stmt.Exec("5")
	// ErrorCheck(e)

	// // affected rows
	// a, e := res.RowsAffected()
	// ErrorCheck(e)

	// fmt.Println(a) // 1
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
