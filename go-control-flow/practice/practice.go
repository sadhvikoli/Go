package practice

import "fmt"

func ExeputeTasks() {
	// Write a program that checks if a number is positive, negative, or zero using if-else.

	a := 2
	if a > 0 {
		fmt.Println("IS A POSITIVE INTEGER")
	} else if a == 0 {
		fmt.Println("IS A ZERO")
	} else {
		fmt.Println("IS A NEGATIVE INTEGER")
	}

	// Use a switch statement to print the name of the day based on a number (1-7).

	switch a {
	case 1:
		fmt.Println("Monday")
	case 2:
		fmt.Println("Tuesday")
	case 3:
		fmt.Println("Wednesday")
	case 4:
		fmt.Println("Thursday")
	case 5:
		fmt.Println("Friday")
	case 6:
		fmt.Println("Saturday")
	case 7:
		fmt.Println("Sunday")
	default:
		fmt.Println("Invalid day")
	}

	// Write a for loop to print numbers from 10 to 1 in reverse order.
	for i := 10; i > 0; i-- {
		fmt.Println(i)
	}

	//Create a multiplication table program using a loop.
	for i := 0; i <= 10; i++ {
		fmt.Println(i * 2)
	}

	//Iterate over an array of names using range and print them.
	nums := []string{"I", "love", "to", "learn", "go"}

	for _, value := range nums { // Fixed range usage
		fmt.Println(value)
	}

}
