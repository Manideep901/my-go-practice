package main

import "log"

var i interface {
}

func found(i interface{}) {
	log.Printf("Type = %T, value = %v\n", i, i)
}

func main() {
	s := "hiii"
	found(s)
}
