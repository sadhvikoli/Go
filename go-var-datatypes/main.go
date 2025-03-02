package main

import (
	"fmt"
	"mymodule/practice"
)

// Import your custom package
func main() {

	// 	var is used for explicit type declaration.
	// Each variable must have a type (string, int, float64, bool).
	var name string = "GoLang"
	var age int = 25
	var pi float64 = 3.14
	var isGoFun bool = true

	fmt.Println("Name:", name)
	fmt.Println("Age:", age)
	fmt.Println("Pi:", pi)
	fmt.Println("Is Go Fun?", isGoFun)

	// 	:= allows implicit type inference.
	// It can only be used inside functions, not at the package level.
	myName := "GoLang"
	myAge := 25
	myPi := 3.14
	isMyGoFun := true

	fmt.Println(myName, myAge, myPi, isMyGoFun)

	// 	const values must be assigned at compile time.
	// They cannot be reassigned later in the code.
	const gravity = 9.8
	const constpi float64 = 3.14159

	fmt.Println("Gravity:", gravity)
	fmt.Println("Pi:", constpi)
	// 	const pi = 3.14
	//  pi = 3.1415 // ❌ Cannot assign a new value to a constant

	// datatypes
	var i int = 42
	var f float64 = 3.1415
	var b bool = true
	var s string = "Hello, Go!"

	fmt.Printf("Type of i: %T, Value: %v\n", i, i)
	fmt.Printf("Type of f: %T, Value: %v\n", f, f)
	fmt.Printf("Type of b: %T, Value: %v\n", b, b)
	fmt.Printf("Type of s: %T, Value: %v\n", s, s)

	// 	Type Conversion in Go
	// Go does not support implicit type conversion. You must explicitly convert types.

	var num int = 42
	var floatNum float64 = float64(num) // Explicit conversion

	fmt.Println("Integer:", num)
	fmt.Println("Converted to Float:", floatNum)
	// 	var i int = 10
	// var f float64 = i // ❌ Error: cannot use i (type int) as type float64

	// 	Zero Values in Go
	// If a variable is declared but not initialized, it gets a zero value.
	var y int     // Default: 0
	var z float64 // Default: 0.0
	var x bool    // Default: false
	var d string  // Default: ""

	fmt.Println("Zero values:")
	fmt.Println(y, z, x, d)

	practice.ExecuteTasks()
}
