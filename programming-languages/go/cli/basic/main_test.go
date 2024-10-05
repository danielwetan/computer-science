package main

import (
	"testing"
)

// Test addition
func TestCalculate_Addition(t *testing.T) {
	result, err := calculate(10, 20, "+")
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}
	expected := 30.0
	if result != expected {
		t.Errorf("Expected %.2f, but got %.2f", expected, result)
	}
}

// Test subtraction
func TestCalculate_Subtraction(t *testing.T) {
	result, err := calculate(15, 5, "-")
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}
	expected := 10.0
	if result != expected {
		t.Errorf("Expected %.2f, but got %.2f", expected, result)
	}
}

// Test multiplication
func TestCalculate_Multiplication(t *testing.T) {
	result, err := calculate(10, 5, "*")
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}
	expected := 50.0
	if result != expected {
		t.Errorf("Expected %.2f, but got %.2f", expected, result)
	}
}

// Test division
func TestCalculate_Division(t *testing.T) {
	result, err := calculate(10, 2, "/")
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}
	expected := 5.0
	if result != expected {
		t.Errorf("Expected %.2f, but got %.2f", expected, result)
	}
}

// Test division by zero
func TestCalculate_DivisionByZero(t *testing.T) {
	_, err := calculate(10, 0, "/")
	if err == nil {
		t.Fatalf("Expected error, but got none")
	}
	expectedError := "division by zero"
	if err.Error() != expectedError {
		t.Errorf("Expected error '%s', but got '%s'", expectedError, err.Error())
	}
}

// Test unsupported operator
func TestCalculate_UnsupportedOperator(t *testing.T) {
	_, err := calculate(10, 5, "%")
	if err == nil {
		t.Fatalf("Expected error, but got none")
	}
	expectedError := "unsupported operator '%'"
	if err.Error() != expectedError {
		t.Errorf("Expected error '%s', but got '%s'", expectedError, err.Error())
	}
}
