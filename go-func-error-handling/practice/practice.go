package practice

import (
	"errors"
	"fmt"
)

// Write a function to calculate the factorial of a number using recursion.
func factorial(num int) int {
	if num == 0 || num == 1 {
		return 1
	}
	return num * factorial(num-1) // Recursive call
}

// Create a function that returns the sum and average of an array of numbers.
func avgSum() (int, float64) { // Return sum (int) and average (float64)
	nums := []int{1, 2, 3, 4, 5}
	res := 0

	for i := 0; i < len(nums); i++ { // Fixed loop condition
		res += nums[i]
	}

	avg := float64(res) / float64(len(nums)) // Ensure floating-point division
	return res, avg
}

// Function to handle division by zero
func divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("division by zero is not allowed")
	}
	return a / b, nil
}

// Function with panic and recover to prevent crashes
func safeDivide(a, b float64) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered from panic:", r)
		}
	}()
	if b == 0 {
		panic("cannot divide by zero!")
	}
	fmt.Println("Result:", a/b)
}

// Variadic function to find the maximum number
func findMax(nums ...int) int {
	if len(nums) == 0 {
		panic("no numbers provided")
	}
	maxNum := nums[0]
	for _, num := range nums {
		if num > maxNum {
			maxNum = num
		}
	}
	return maxNum
}

func ExecuteTasks() {

	num := 5
	fmt.Println("Factorial of", num, "is:", factorial(num))
	sum, average := avgSum()
	fmt.Println("Sum:", sum)
	fmt.Println("Average:", average)
	// Example 1: Handling division by zero
	result, err := divide(10, 0)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Division Result:", result)
	}

	// Example 2: Using panic and recover
	fmt.Println("Attempting safe division:")
	safeDivide(10, 0)

	// Example 3: Finding the maximum number
	fmt.Println("Max number:", findMax(4, 2, 8, 1, 5))

}
