package main

import (
	"log"
	"strconv"
)

func main() {
	var a string = "7"
	b := 7
	aInt, _ := strconv.Atoi(a)
	c := aInt + b
	log.Print("toatal value: ", c)
}
