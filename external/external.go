package main

import (
	"fmt"
	"log"

	"github.com/Knetic/govaluate"
)

func SolveExpression(expression string) (float64, error) {
	// Create a new expression from the input string
	expr, err := govaluate.NewEvaluableExpression(expression)
	if err != nil {
		return 0, err
	}

	// Evaluate the expression
	result, err := expr.Evaluate(nil) // No variables to pass, just evaluate the expression
	if err != nil {
		return 0, err
	}

	// Assert the result to float64
	return result.(float64), nil
}

func TakeInput() {
	log.Println("Enter a mathematical expression to evaluate (e.g., 3 + 4 * 2):")
	var expression string
	fmt.Scanln(&expression)

	// Solve the expression
	result, err := SolveExpression(expression)
	if err != nil {
		log.Println("Error evaluating expression:", err)
		return
	}

	log.Printf("Result of the expression '%s' is: %f\n", expression, result)
}

func main() {
	// Call the function to take input and solve the expression
	TakeInput()
}
