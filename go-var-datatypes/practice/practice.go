package practice

import (
	"fmt"
)

// ExecuteTasks is a function that performs multiple tasks
func ExecuteTasks() {
	// Task 1: Declare variables using var and :=
	var name string = "GoLang"
	age := 10
	fmt.Println("Task 1:", name, age)

	// Task 2: Declare constants and print them
	const (
		pi           = 3.141592653589793
		gravity      = 9.81
		speedOfLight = 299792458
	)
	fmt.Println("Task 2:", pi, gravity, speedOfLight)

	// Task 3: Convert an integer to a float
	integerValue := 42
	floatValue := float64(integerValue)
	fmt.Println("Task 3:", floatValue)

	// Task 4: Declare variables without assigning values
	var num int
	var str string
	var boolean bool
	fmt.Println("Task 4:", num, str, boolean)
}
