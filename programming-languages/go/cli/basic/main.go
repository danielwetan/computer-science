package main

import (
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"
)

// Print usage documentation
func printUsage() {
	fmt.Println(`Usage: <number1> <operator> <number2>

A basic calculator that performs operations on two numbers.

Positional arguments:
  number1          The first number (integer or float)
  operator         The operation to perform: +, -, *, /
  number2          The second number (integer or float)

Examples:
  ./calculator 10 + 20          # Performs addition
  ./calculator 10 * 20          # Performs multiplication
  ./calculator 10 / 0           # Outputs an error for division by zero
  ./calculator 20.5 - 5.1       # Performs subtraction`)
}

// Function that performs calculation, now easier to test
func calculate(op1, op2 float64, operator string) (float64, error) {
	switch operator {
	case "+":
		return op1 + op2, nil
	case "-":
		return op1 - op2, nil
	case "*":
		return op1 * op2, nil
	case "/":
		if op2 == 0 {
			return 0, errors.New("division by zero")
		}
		return op1 / op2, nil
	default:
		return 0, fmt.Errorf("unsupported operator '%s'", operator)
	}
}

func main() {
	if len(os.Args) != 4 {
		printUsage()
		os.Exit(1)
	}

	op1, err := strconv.ParseFloat(os.Args[1], 64)
	if err != nil {
		log.Fatalf("Error: Invalid first number '%s'. Please provide a valid number.", os.Args[1])
	}

	operator := os.Args[2] // Operator is the second argument

	op2, err := strconv.ParseFloat(os.Args[3], 64)
	if err != nil {
		log.Fatalf("Error: Invalid second number '%s'. Please provide a valid number.", os.Args[3])
	}

	result, err := calculate(op1, op2, operator)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%.2f %s %.2f = %.2f\n", op1, operator, op2, result)
}
