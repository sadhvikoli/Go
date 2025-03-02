package main

import (
	"fmt"
	"mymodule/practice"
)

func main() {

	// 	No parentheses () around conditions.
	// Curly braces {} are mandatory, even for one-line blocks.
	age := 18

	if age >= 18 {
		fmt.Println("You are eligible to vote!")
	} else {
		fmt.Println("You are not eligible to vote.")
	}

	// If-Else with Multiple Conditions
	num := 10

	if num > 0 {
		fmt.Println("Positive number")
	} else if num < 0 {
		fmt.Println("Negative number")
	} else {
		fmt.Println("Zero")
	}

	// Short Variable Declaration in If Statement
	if num := 42; num%2 == 0 {
		fmt.Println(num, "is even")
	} else {
		fmt.Println(num, "is odd")
	}
	// num := 42 is declared inside the if statement and is not accessible outside.

	// Basic Switch Example
	day := 3

	switch day {
	case 1:
		fmt.Println("Monday")
	case 2:
		fmt.Println("Tuesday")
	case 3:
		fmt.Println("Wednesday")
	default:
		fmt.Println("Invalid day")
	}
	// Switch Without break
	// Unlike other languages, Go does not require break, as cases exit automatically.

	// Switch with Multiple Cases
	grade := "A"

	switch grade {
	case "A", "B":
		fmt.Println("Excellent!")
	case "C", "D":
		fmt.Println("Good, but can improve.")
	default:
		fmt.Println("Fail.")
	}

	// Switch with fallthrough (Forcing Next Case Execution)
	num1 := 2

	switch num1 {
	case 1:
		fmt.Println("One")
	case 2:
		fmt.Println("Two")
		fallthrough
	case 3:
		fmt.Println("Three") // Will execute due to fallthrough
	default:
		fmt.Println("Other number")
	}
	// fallthrough forces the next case to execute.

	// Basic for Loop
	for i := 1; i <= 5; i++ {
		fmt.Println("Iteration:", i)
	}
	// i++ is used for incrementing.

	// For Loop Without Initialization and Post Statement
	i := 1
	for i <= 5 { // Acts like a while loop
		fmt.Println("Count:", i)
		i++
	}
	// This behaves like a while loop.

	// Infinite Loop (for {})
	// for {
	//     fmt.Println("This is an infinite loop!")
	// }
	// Use break to stop an infinite loop.

	//  Using break and continue
	for i := 1; i <= 5; i++ {
		if i == 3 {
			continue // Skips 3
		}
		fmt.Println(i)
		if i == 4 {
			break // Exits loop when i == 4
		}
	}
	// 	continue → Skips current iteration.
	// break → Exits loop immediately.

	// Using range with Arrays & Slices
	nums := []int{10, 20, 30, 40, 50}

	for index, value := range nums {
		fmt.Printf("Index: %d, Value: %d\n", index, value)
	}
	// range is used for iterating over slices, arrays, and maps.

	practice.ExeputeTasks()
}
