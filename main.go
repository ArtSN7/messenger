package main

import (
	"fmt"
	_ "github.com/mattn/go-sqlite3"
)



func main() {

	var approve bool

	var scan int
	

	response := log_reg()

    // log in or sign up 
	if response == "try again"{ // first attempt

		response = log_reg()

		if response == "try again"{ // second attempt

			fmt.Println("Your activity is strange for us. Bye!")

			approve = false // if user was not able to log in or sign up, then we refuse his other attempts to log in

		}
	}

    // if user was approved by the syste,
	if approve == true{

		fmt.Print("\n\n------------------------------\n\n")
		fmt.Printf("\n\nHello, %s .\n\nWhat do you want to do?\n\n1 - write email\n2 - check your inbox\n3 - leave app\n\n", response)

		

		for{

			fmt.Scan(&scan)

			switch scan{

			case 3: break

			case 1:
				write email 

			case 2:
				check inbox

			}

			fmt.Print("\n\nDONE! \n\nWhat do you want to do next?\n\n1 - write email\n2 - check your inbox\n3 - leave app\n\n")
			
		}

	}

}

