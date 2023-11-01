package main

import (
	"fmt"
	_ "github.com/mattn/go-sqlite3"
)



func main() {

	var response_reg int
	var email string
	var password string

	fmt.Print("Hello, what do you want to do.\n\nFunctions:\n\n1 - log in into existing account\n2 - create a new account\n\n")
	fmt.Scan(&response_reg)

	if response_reg == 1{

		// getting data
		fmt.Print("\n\nTo log in into your account, please, input your email:\n\n")
		fmt.Scan(&email)
		fmt.Print("\n\nNow, please, input your password:\n\n")
		fmt.Scan(&password)


        // checking if email and password match 
		if CheckPassword(email, password) == true{
			fmt.Println("\n\nAccess approved.")

		}else{fmt.Println("\n\nAccess denied. Check your data.")}


	}else{

		// getting data
		fmt.Print("\n\nTo create a new account, please, input your email:\n\n")
		fmt.Scan(&email)
		fmt.Print("\n\nNow, please, input your password:\n\n")
		fmt.Scan(&password)

		if CreateAccount(email, password) == true{ // creating account
			fmt.Println("\n\nAccess approved.")
		}else{
			fmt.Println("\n\nSomething went wrong. Restart application.")
		}

	}


}

