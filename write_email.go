package main

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"bufio"
    "os"
)

func WriteEmail(email string) bool{

	var to_user string 

	fmt.Print("\n\nAlright, please input the email of the user to which you want to send your message:\n\n")

	fmt.Scan(&to_user)

	fmt.Print("\n\nNow, please, send message which you want to send:\n\n")

	scanner := bufio.NewScanner(os.Stdin)
    scanner.Scan()


	db, _ := sql.Open("sqlite3", "db/app.sqlite") // creating connection


	stmt, err := db.Prepare("INSERT INTO messages(from_user, to_user, message) VALUES( ?, ?, ? )")


	if err != nil {
		return false
	}

	if stmt != nil{ 
		defer stmt.Close() // Prepared statements take up server resources and should be closed after use.
	}


	if _, err := stmt.Exec(email, to_user, scanner.Text()); err != nil {
		return false
	}

	return true
	
}