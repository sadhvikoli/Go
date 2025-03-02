package main

import (
	"errors"
	"fmt"
	"mymodule/practice"
)

// Functions are defined with func functionName() {}.
// Call the function using functionName().
func greet() {
	fmt.Println("Hello, Go Developer!")
}

// Function with Parameters
// name string → Function expects a string argument.
func greetUser(name string) {
	fmt.Println("Hello,", name)
}

// Function with Multiple Parameters
// a int, b int → Parameters with data types.
// int → Return type of the function
func add(a int, b int) int {
	return a + b
}

// Functions with Multiple Return Values
// func divide(a, b int) (int, int) → Function returns two values.
// q, r := divide(10, 3) → Multiple assignments when calling the function.
func divide(a, b int) (int, int) {
	quotient := a / b
	remainder := a % b
	return quotient, remainder
}

// Function Accepting Multiple Arguments
// numbers ...int → Accepts any number of integer arguments.
func summation(numbers ...int) int {
	total := 0
	for _, num := range numbers {
		total += num
	}
	return total
}

// Basic Error Handling Example
func divideError(a, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("division by zero is not allowed")
	}
	return a / b, nil
}

// Example of Custom Error Handling
func withdraw(balance, amount int) (int, error) {
	if amount > balance {
		return balance, fmt.Errorf("insufficient balance: requested %d but available %d", amount, balance)
	}
	return balance - amount, nil
}

// Recover - Handle Panic and Continue Execution
func safeFunction() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered from Panic:", r)
		}
	}()

	panic("Oops! Something went wrong.")
}

func main() {
	greet() // Calling the function
	greetUser("Alice")
	sum := add(5, 10)
	fmt.Println("Sum:", sum)
	q, r := divide(10, 3)
	fmt.Println("Quotient:", q, "Remainder:", r)
	fmt.Println("Sum:", summation(1, 2, 3, 4, 5))

	// Anonymous Function Assigned to a Variable
	multiply := func(a, b int) int {
		return a * b
	}

	result := multiply(5, 6)
	fmt.Println("Multiplication:", result)
	// The function is assigned to a variable (multiply).

	// Immediately Invoked Function (IIFE)
	func(msg string) {
		fmt.Println(msg)
	}("Hello, Go!") // Function invoked immediately

	result, err := divideError(10, 0)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Result:", result)
	}
	// 	errors.New("message") → Creates an error.
	// return 0, nil → nil means no error.
	// Always check if err != nil.

	newBalance, err := withdraw(500, 700)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("New Balance:", newBalance)
	}
	// fmt.Errorf() allows formatted error messages.

	fmt.Println("Starting...")
	safeFunction()
	fmt.Println("Execution continued after panic.")
	// recover() catches panic and prevents a crash.
	// defer ensures execution before function exits.

	practice.ExecuteTasks()

	// Panic - Stops Program Execution Immediately
	fmt.Println("Before Panic")
	panic("Something went wrong!")
	// fmt.Println("After Panic") // ❌ This will not execute
	// panic("message") stops execution immediately.

}
