package main

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"fmt"
)

func CheckInbox(email string) bool {

	db, err := sql.Open("sqlite3", "db/app.sqlite") // creating connection

	rows, err := db.Query("SELECT message FROM messages WHERE to_user=?", email) // getting password

	if err != nil{
		return false
	}

    var text string

	for rows.Next() {
		err = rows.Scan(&text) // getting password from db in string format

		fmt.Printf("\n\n%s\n\n", text)

		if err != nil{
			return false
		}

	}

	return true
	
}