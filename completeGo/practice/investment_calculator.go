package main

import (
	"fmt"
	"log"
	"math"
)

func main() {
	const inflationRate = 2.5
	var investmentAmount float64
	var expectedReturnRate = 5.5
	var years float64
	log.Print("Enter Investment Amount : ")
	fmt.Scan(&investmentAmount)

	log.Print("Expected  Return Rate : ")
	fmt.Scan(&expectedReturnRate)

	log.Print("Enter Investment Years : ")
	fmt.Scan(&years)

	var futureValue = investmentAmount * math.Pow(1+expectedReturnRate/100, years)
	futureRealValue := futureValue / math.Pow(1+inflationRate/100,years)
	log.Println(futureValue,futureRealValue)
}