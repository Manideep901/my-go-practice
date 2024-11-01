package main

import (
	"strings"
)

func ValidateExpression(exp string) bool {
	exp = strings.ReplaceAll(exp, " ", "")
	return areParenthesesBalanced(exp)
}
func areParenthesesBalanced(expression string) bool {
	count := 0
	for _, char := range expression {
		if char == '(' {
			count++
		} else if char == ')' {
			count--
			if count < 0 {
				return false
			}
		}
	}
	return count == 0
}

// Check if the character is an operator
func isOperator(char string) bool {
	return char == "+" || char == "-" || char == "*" || char == "/"
}
