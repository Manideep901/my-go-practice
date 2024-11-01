package main

import (
	"io"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	content := "Hello mani good moring"
	fileName, err := os.Create("./Hello.txt")

	if err != nil {
		panic(err)
	}

	length, err := io.WriteString(fileName, content)
	if err != nil {
		panic(err)
	}
	log.Println("lenth of the file is: ", length)
	defer fileName.Close()
	ReadFile("./Hello.txt")
}

func ReadFile(filename string) {
	dataByte, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}

	log.Println("text data in the file is: ", string(dataByte))
}
