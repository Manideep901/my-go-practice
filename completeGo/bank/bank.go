package main

import (
	"bank/fileops"
	"fmt"
	"log"
)
const  accountBlanaceFile = "balance.txt"
func main() {
	var accountBalance, err = fileops.GetFloatFromFile(accountBlanaceFile)

	if err != nil {
		fmt.Println("ERROR")
		fmt.Println(err)
		fmt.Println("---------")
		// panic("Can't continue, sorry.")
	}

	log.Println("Welcome to Mani Bank")

	for {
		presentOptions()
		var choice int
		log.Println("Your Choice: ")
		fmt.Scan(&choice)

		if choice == 1 {
			log.Println("Your current balance is: ", accountBalance)
		} else if choice == 2 {
			log.Println("Enter Your deposit: ")
			var despositeAmount float64
			fmt.Scan(&despositeAmount)
			if despositeAmount <= 0 {
				log.Println("Invalid amount, Must be greater than 0")
				// return
				continue
			}
			accountBalance += despositeAmount;
			fileops.WriteFloatToFile(accountBalance,accountBlanaceFile)
			log.Println("Balance updated! New Amount: ", accountBalance)
		} else if choice == 3 {
			log.Println("Enter Your withdrawal: ")
			var withdrawalAmount float64
			fmt.Scan(&withdrawalAmount)
			if withdrawalAmount <= 0 {
				log.Println("Invalid amount, Must be greater than 0")
				// return
				continue
			}
			if withdrawalAmount > accountBalance {
				log.Println("Insufficient balance")
				return
			}
			accountBalance -= withdrawalAmount;
			fileops.WriteFloatToFile(accountBalance,accountBlanaceFile)
			log.Println("Balance updated! New Amount: ", accountBalance)
		} else {
			log.Println("Good bye!")
			break
		}
	}

	log.Println("Thanks for choosing our bank")

}

