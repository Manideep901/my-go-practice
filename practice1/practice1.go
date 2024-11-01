package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Function to evaluate expressions
func eval(expr string) float64 {
	// Remove spaces
	expr = strings.ReplaceAll(expr, " ", "")

	// Process parentheses first
	for strings.Contains(expr, "(") {
		open := strings.LastIndex(expr, "(")
		close := strings.Index(expr[open:], ")") + open
		subExpr := expr[open+1 : close]
		value := eval(subExpr)
		expr = expr[:open] + fmt.Sprintf("%f", value) + expr[close+1:]
	}

	// Handle multiplication and division before addition and subtraction
	return evalAddSub(evalMulDiv(expr))
}

// Helper function for multiplication and division
func evalMulDiv(expr string) string {
	var result []string
	num := ""
	// var prevOp rune = '+'
	for i, c := range expr {
		if c == '*' || c == '/' {
			val1, _ := strconv.ParseFloat(num, 64)
			num = ""
			i++
			// Get the next number
			for i < len(expr) && (expr[i] >= '0' && expr[i] <= '9' || expr[i] == '.') {
				num += string(expr[i])
				i++
			}
			val2, _ := strconv.ParseFloat(num, 64)
			// Perform multiplication or division
			if c == '*' {
				val1 = val1 * val2
			} else {
				val1 = val1 / val2
			}
			// Replace the result of multiplication/division in the expression
			result = append(result, fmt.Sprintf("%f", val1))
			if i < len(expr) {
				result = append(result, string(expr[i]))
			}
			num = ""
		} else {
			num += string(c)
		}
		if i == len(expr)-1 && num != "" {
			result = append(result, num)
		}
	}
	return strings.Join(result, "")
}

// Helper function for addition and subtraction
func evalAddSub(expr string) float64 {
	var result float64
	num := ""
	operation := '+'
	for _, c := range expr {
		if c == '+' || c == '-' {
			if num != "" {
				val, _ := strconv.ParseFloat(num, 64)
				if operation == '+' {
					result += val
				} else {
					result -= val
				}
				num = ""
			}
			operation = c
		} else {
			num += string(c)
		}
	}
	if num != "" {
		val, _ := strconv.ParseFloat(num, 64)
		if operation == '+' {
			result += val
		} else {
			result -= val
		}
	}
	return result
}

func main() {
	// Get input from the user
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter a mathematical expression: ")
	expression, _ := reader.ReadString('\n')

	// Trim the input to remove newline characters
	expression = strings.TrimSpace(expression)

	// Evaluate the expression
	result := eval(expression)

	// Print the result
	fmt.Printf("The result of the expression '%s' is: %f\n", expression, result)
}
