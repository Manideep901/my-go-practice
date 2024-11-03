package main

import (
	"fmt"
	"log"
)

func main() {

	revenue := getUserInput("Enter the revenue: ")
	expenses := getUserInput("Enter the expemses: ")
	taxRate := getUserInput("Enter the tax rated: ")

	ratio, ebt, profit := calculateFinancials(revenue, expenses, taxRate)

	log.Printf("%.1f\n",ebt)
	log.Printf("%.3f\n",profit)
	log.Printf("%.2f\n",ratio)

}

func getUserInput(infoText string) float64 {
	var userInput float64
	log.Print(infoText)
	fmt.Scan(&userInput)
	return userInput
}

func calculateFinancials(revenue, expenses, taxRate float64) (float64, float64, float64) {
	ebt := revenue - expenses
	profit := ebt * (1 - taxRate/100)
	ratio := ebt / profit
	return ratio, ebt, profit
}
