package practice

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

// User struct for JSON handling
type User struct {
	Name  string `json:"name"`
	Age   int    `json:"age"`
	Email string `json:"email"`
}

func Practice() {
	// Task 1: Create a file and write user input
	writeToFile("user_input.txt")

	// Task 2: Read and print file content
	readFromFile("user_input.txt")

	// Task 3: Create a JSON file and read it back
	createAndReadJSON("users.json")

	// Task 4: Modify an existing JSON file by appending data
	modifyJSON("users.json", User{"Charlie", 28, "charlie@example.com"})

	// Task 5: Convert struct to JSON, save it, and read it back
	convertStructToJSON("user.json", User{"David", 35, "david@example.com"})
}

// ✅ Task 1: Write user input to a file
func writeToFile(filename string) {
	file, err := os.Create(filename)
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer file.Close()

	fmt.Println("Enter text (press Ctrl+D / Ctrl+Z to stop):")
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		_, err := file.WriteString(scanner.Text() + "\n")
		if err != nil {
			fmt.Println("Error writing to file:", err)
			return
		}
	}
	fmt.Println("Data saved to", filename)
}

// ✅ Task 2: Read and print a file’s content
func readFromFile(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	fmt.Println("\nReading", filename, "content:")
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}

// ✅ Task 3: Create a JSON file storing a list of users and read the data
func createAndReadJSON(filename string) {
	users := []User{
		{"Alice", 25, "alice@example.com"},
		{"Bob", 30, "bob@example.com"},
	}

	data, err := json.MarshalIndent(users, "", "  ")
	if err != nil {
		fmt.Println("Error encoding JSON:", err)
		return
	}

	err = ioutil.WriteFile(filename, data, 0644)
	if err != nil {
		fmt.Println("Error writing JSON file:", err)
		return
	}
	fmt.Println("\nJSON file created successfully:", filename)

	// Read JSON file
	file, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("Error reading JSON file:", err)
		return
	}

	var loadedUsers []User
	json.Unmarshal(file, &loadedUsers)
	fmt.Println("Users from JSON file:", loadedUsers)
}

// ✅ Task 4: Modify an existing JSON file by appending data
func modifyJSON(filename string, newUser User) {
	file, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("Error reading JSON file:", err)
		return
	}

	var users []User
	json.Unmarshal(file, &users)

	users = append(users, newUser)

	updatedData, err := json.MarshalIndent(users, "", "  ")
	if err != nil {
		fmt.Println("Error encoding JSON:", err)
		return
	}

	err = ioutil.WriteFile(filename, updatedData, 0644)
	if err != nil {
		fmt.Println("Error writing to JSON file:", err)
		return
	}

	fmt.Println("New user added to", filename)
}

// ✅ Task 5: Convert struct to JSON, save it, and read it back
func convertStructToJSON(filename string, user User) {
	data, err := json.MarshalIndent(user, "", "  ")
	if err != nil {
		fmt.Println("Error encoding JSON:", err)
		return
	}

	err = ioutil.WriteFile(filename, data, 0644)
	if err != nil {
		fmt.Println("Error writing JSON file:", err)
		return
	}
	fmt.Println("\nStruct saved as JSON:", filename)

	// Read and decode JSON back into struct
	file, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("Error reading JSON file:", err)
		return
	}

	var loadedUser User
	json.Unmarshal(file, &loadedUser)
	fmt.Println("Read from JSON file:", loadedUser)
}
