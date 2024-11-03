package main

import (
	"fmt"
	"log"
	"time"
)

type user struct{
	 firstName string
	 lastName string
	 birthdate string
	 createdAt time.Time
}

func main() {
	userFirstName := getUserData("Please enter your first name: ")
	userLastName := getUserData("Please enter your last name: ")
	userBirthdate := getUserData("Please enter your birthdate (MM/DD/YYYY): ")

	var appUser user

	appUser = user{
		firstName: userFirstName,
		lastName: userLastName,
		birthdate: userBirthdate,
		createdAt: time.Now(),
	}

	// ... do something awesome with that gathered data!
	appUser.outputUserDetails()
	
}
func (u user)outputUserDetails(){
	log.Println(u.firstName,u.lastName,u.birthdate)
}

func getUserData(promptText string) string {
	log.Print(promptText)
	var value string
	fmt.Scan(&value)
	return value
}