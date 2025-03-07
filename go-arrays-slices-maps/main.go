package main

import (
	"fmt"
	"mymodule/practice"
)

func main() {

	numbers := [3]int{10, 20, 30}
	numbers[1] = 50      // Change the second element
	fmt.Println(numbers) // [10 50 30]

	// Iterating Over an Array
	for _, num := range numbers {
		fmt.Print(num)
	}

	// Slice declaration
	nums := []int{10, 20, 30}
	fmt.Println(nums) // [10 20 30]

	nums = append(nums, 4, 5) // Adding elements
	fmt.Println(nums)         // [1 2 3 4 5]

	subSlice := nums[1:4] // Slices from index 1 to 3 (excluding 4)
	fmt.Println(subSlice) // [20 30 4]

	original := []int{1, 2, 3, 4}
	copySlice := make([]int, len(original))
	copy(copySlice, original) // Copy elements
	fmt.Println(copySlice)    // [1 2 3 4]
	// 	make([]int, size) → Creates an empty slice of given size.
	// copy(destination, source) → Copies elements.

	// student := make(map[string]int) // Key: string, Value: int
	// fmt.Println(student) // map[]
	// make(map[keyType]valueType) → Creates an empty map.

	student := map[string]int{"Alice": 25, "Bob": 30}
	fmt.Println(student["Alice"]) // 25

	student["Charlie"] = 22 // Adding new entry
	fmt.Println(student)

	age, exists := student["Bob"]
	if exists {
		fmt.Println("Bob's age:", age)
	} else {
		fmt.Println("Bob not found")
	}
	// value, exists := map[key] → Returns value and a boolean indicating key presence.

	delete(student, "Alice") // Deletes "Alice"
	fmt.Println(student)     // map[Bob:30]

	for key, value := range student {
		fmt.Println("Name:", key, "Age:", value)
	}

	practice.ExeputeTasks()
}
