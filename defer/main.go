package main

import "log"

func main() {
	defer log.Println("World")
	log.Println("hello")
	Defer()
}

func Defer() {
	for i := 0; i < 5; i++ {
		defer log.Println(i)
	}
}
