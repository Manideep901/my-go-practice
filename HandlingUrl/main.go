package main

import (
	"log"
	"net/url"
)

const myUrl string = "https://lco.dev:3000/learn?coursename=reactjs&paymentid=54646"

func main() {
	log.Println(myUrl)
	result,_ := url.Parse(myUrl);
	log.Println("my url ",result.Scheme)
	log.Println("my url ",result.Host)
	log.Println("my url ",result.Path)
	log.Println("my url ",result.Port())
	log.Println("my url ",result.RawQuery)
	qparams := result.Query()
	log.Printf("The type of query params are: %T\n",qparams)
	log.Println(qparams["coursname"])
	for _,val:= range qparams{
		log.Println("params is: ",val)
	}
	partsofUrl := &url.URL{
		Scheme: "https",
		Host: "lco.dev",
		Path: "/turcss",
		RawPath: "user=manideep",
	}
	anotherUrl := partsofUrl.String()
	log.Println(anotherUrl)
}