package main

import "log"

func main() {
	age := 32

	var agePointer *int
	agePointer = &age
	log.Println("Age: ",*agePointer)
	adultYears := getAdultYears(age)
	log.Println(adultYears)
}

func getAdultYears(age int) int {
	return age - 18
}