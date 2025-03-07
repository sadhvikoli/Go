package practice

import "fmt"

func ExeputeTasks() {

	//  Create an array of 5 integers and print each element using a loop.
	numbers := [5]int{23, 24, 25, 26, 27}

	for _, num := range numbers {
		fmt.Print(num)
	}

	//  Create a slice, append elements to it, and print its length and capacity.
	names := []string{}

	names = append(names, "sadhvi", "sterash")
	fmt.Println(len(names))

	// Slice an array into two parts and print both.
	subSlice1 := numbers[0:3]
	subSlice2 := numbers[3:5]
	fmt.Println("First half:", subSlice1)
	fmt.Println("Second half:", subSlice2)

	// Create a map of student names and their scores, add a new student, and delete one.
	student := map[string]int{"Alice": 25, "Bob": 30}
	delete(student, "Alice")
	fmt.Println(student)

	num := score(student)
	fmt.Print(num)

}

// Write a function to return the highest score from a map of student marks.
func score(m map[string]int) int {
	maxScore := 0
	for _, num := range m {
		if num > maxScore {
			maxScore = num
		}
	}
	return maxScore
}
