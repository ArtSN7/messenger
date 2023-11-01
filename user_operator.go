package main

import (
	"database/sql"
	"log"
	_ "github.com/mattn/go-sqlite3"
)


func CheckPassword(email, password string) bool{
	db, err := sql.Open("sqlite3", "db/app.sqlite") // creating connection
	checkErr(err) // checking if there are any errors

	rows, _ := db.Query("SELECT password FROM users WHERE mail=?", email) // getting password

    var pas string

	for rows.Next() {
		err = rows.Scan(&pas) // getting password from db in string format

		if err != nil{
			return false
		}

	}

	if pas == password{ // if passwords match , then approve and return true
		return true
	}

	return false

}

func CreateAccount(email, password string) bool{

	db, _ := sql.Open("sqlite3", "db/app.sqlite") // creating connection


	stmt, err := db.Prepare("INSERT INTO users(mail, password) VALUES( ?, ? )")


	if err != nil {
		return false
	}


	if stmt != nil{ 
		defer stmt.Close() // Prepared statements take up server resources and should be closed after use.
	}

	

	if _, err := stmt.Exec(email, password); err != nil {
		return false
	}


	return true
}



func checkErr(err error) { // checking if there is any mistakes while oppening db
	if err != nil {
		log.Fatal(err)
	}
}