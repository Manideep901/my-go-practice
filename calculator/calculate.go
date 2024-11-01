package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"
	"unicode"
)

func main() {
	TakeInput()
	// (8*9+7-8*(7/8))
}

func TakeInput() {
	log.Println("Select the option to Calculate:\n1. Addition\n2. Subtraction\n3. Multiplication\n4. Division\n5. Expression")
	var selectedOperation string
	fmt.Scanln(&selectedOperation)
	switch selectedOperation {
	case "1":
		log.Println("Adition Selected\n enter first digit :")
		var firstDigit float64
		fmt.Scanln(&firstDigit)
		log.Println("enter second digit :")
		var secondDigit float64
		fmt.Scanln(&secondDigit)
		addition := firstDigit + secondDigit

		log.Printf("Addition of %v and %v is: %v\n", firstDigit, secondDigit, addition)
	case "2":
		log.Println("Subtraction Selected\n enter first digit :")
		var firstDigit float64
		fmt.Scanln(&firstDigit)
		log.Println("enter second digit :")
		var secondDigit float64
		fmt.Scanln(&secondDigit)
		addition := firstDigit - secondDigit

		log.Printf("Subtraction of %v and %v is: %v\n", firstDigit, secondDigit, addition)
	case "3":
		log.Println("Multiplication Selected\n enter first digit :")
		var firstDigit float64
		fmt.Scanln(&firstDigit)
		log.Println("enter second digit :")
		var secondDigit float64
		fmt.Scanln(&secondDigit)
		addition := firstDigit * secondDigit

		log.Printf("Multiplication of %v and %v is: %v\n", firstDigit, secondDigit, addition)
	case "4":
		log.Println("Divison Selected\n enter first digit :")
		var firstDigit float64
		fmt.Scanln(&firstDigit)
		log.Println("enter second digit :")
		var secondDigit float64
		fmt.Scanln(&secondDigit)
		addition := firstDigit / secondDigit

		log.Printf("Divition of %v and %v is: %v\n", firstDigit, secondDigit, addition)
	case "5":
		log.Println("Expression Selected\n enter Expression to calculate :")
		var expression string
		fmt.Scanln(&expression)
		SolveExpression(expression)
	}
}

func SolveExpression(exp string) {
	if ValidateExpression(exp) {
		log.Println("entered expresion is valid", exp)
		result, _ := Calculate(exp)
		log.Printf("Result of the expression '%s' is: %.2f\n", exp, result)
	} else {
		log.Println("Entered expresion is invalid\n Enter valid expression\n enter Expression to calculate :")
		var expression string
		fmt.Scanln(&expression)
		SolveExpression(expression)
	}

}

func Calculate(exp string) (float64, error) {
	var nums []float64
	var ops []rune
	exp = strings.ReplaceAll(exp, " ", "")
	applyOperator := func() error {
		if len(nums) < 2 || len(ops) == 0 {
			return fmt.Errorf("not enough values to apply operator")
		}
		b := nums[len(nums)-1]
		a := nums[len(nums)-2]
		operator := ops[len(ops)-1]

		// Remove the last numbers and operator
		nums = nums[:len(nums)-2]
		ops = ops[:len(ops)-1]

		var result float64
		switch operator {
		case '+':
			result = a + b
		case '-':
			result = a - b
		case '*':
			result = a * b
		case '/':
			if b == 0 {
				return fmt.Errorf("division by zero")
			}
			result = a / b
		}
		nums = append(nums, result)
		return nil
	}

	for i := 0; i < len(exp); i++ {
		char := exp[i]
		if unicode.IsDigit(rune(char)) {
			// If the character is a digit, extract the full number
			start := i
			for i < len(exp) && (unicode.IsDigit(rune(exp[i])) || exp[i] == '.') {
				i++
			}
			num, err := strconv.ParseFloat(exp[start:i], 64)
			if err != nil {
				return 0, err
			}
			nums = append(nums, num)
			i-- // Adjust because of the loop increment
		} else if char == '(' {
			ops = append(ops, rune(char))
		} else if char == ')' {
			for len(ops) > 0 && ops[len(ops)-1] != '(' {
				if err := applyOperator(); err != nil {
					return 0, err
				}
			}
			ops = ops[:len(ops)-1] // Remove the '(' from stack
		} else if char == '+' || char == '-' || char == '*' || char == '/' {
			for len(ops) > 0 && precedence(rune(char)) <= precedence(ops[len(ops)-1]) {
				if err := applyOperator(); err != nil {
					return 0, err
				}
			}
			ops = append(ops, rune(char))
		}
	}

	// Apply remaining operators
	for len(ops) > 0 {
		if err := applyOperator(); err != nil {
			return 0, err
		}
	}

	if len(nums) != 1 {
		return 0, fmt.Errorf("invalid expression")
	}

	return nums[0], nil
}

// Determine precedence of operators
func precedence(op rune) int {
	switch op {
	case '+', '-':
		return 1
	case '*', '/':
		return 2
	}
	return 0
}
