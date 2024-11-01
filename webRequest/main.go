package main

import (
	"io"
	"log"
	"net/http"
)

const url = "https://www.google.com"

func main() {
	response, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	log.Printf("response is of type : %T\n", response)
	defer response.Body.Close()
	byteCode, err := io.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}
	content := string(byteCode)
	log.Println("response from the url", content)
}
