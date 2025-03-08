package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"mymodule/practice"
	"os"
)

func main() {
	data := "Hello, Go File Handling!"

	err := os.WriteFile("example.txt", []byte(data), 0644)
	if err != nil {
		fmt.Println("Error writing to file:", err)
		return
	}

	fmt.Println("File written successfully!")
	// 	os.WriteFile(filename, data, permissions) → Writes data to a file.
	// 0644 → File permission (read & write for owner, read-only for others).

	file, err := os.Create("example2.txt")
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer file.Close()

	lines := []string{"Go is awesome!", "File handling is easy!", "Enjoy coding!"}

	for _, line := range lines {
		_, err := file.WriteString(line + "\n")
		if err != nil {
			fmt.Println("Error writing line:", err)
			return
		}
	}

	fmt.Println("Lines written successfully!")
	// 	os.Create("filename") → Creates a new file (or overwrites if it exists).
	// file.WriteString() → Writes a string to a file.
	// defer file.Close() → Ensures file closes properly.

	dta, err := os.ReadFile("example.txt")
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	fmt.Println("File Content:\n", string(dta))
	// os.ReadFile(filename) → Reads the entire file into memory.

	fle, err := os.Open("example2.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer fle.Close()

	scanner := bufio.NewScanner(fle)
	for scanner.Scan() {
		fmt.Println(scanner.Text()) // Print each line
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}
	// 	os.Open("filename") → Opens a file for reading.
	// bufio.NewScanner(file) → Reads line by line efficiently.

	fil, err := os.OpenFile("example.txt", os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer fil.Close()

	_, err = fil.WriteString("\nNew Line Appended!")
	if err != nil {
		fmt.Println("Error appending to file:", err)
		return
	}

	fmt.Println("Data appended successfully!")
	// 	os.OpenFile(filename, mode, permissions) → Opens file in append mode.
	// Modes: os.O_APPEND (append), os.O_WRONLY (write only).

	p := Person{Name: "Alice", Age: 25, Email: "alice@example.com"}

	jsonData, err := json.MarshalIndent(p, "", "  ")
	if err != nil {
		fmt.Println("Error encoding JSON:", err)
		return
	}

	fmt.Println(string(jsonData))
	// 	json.Marshal() → Converts struct to JSON.
	// json.MarshalIndent() → Pretty-prints JSON.
	// json:"fieldname" → Maps struct fields to JSON keys.

	// jsonData := `{"name":"Bob","age":30,"email":"bob@example.com"}`

	// var p Person
	// err := json.Unmarshal([]byte(jsonData), &p)
	// if err != nil {
	//     fmt.Println("Error decoding JSON:", err)
	//     return
	// }

	// fmt.Println("Decoded Struct:", p)

	// p := Person{Name: "Charlie", Age: 35, Email: "charlie@example.com"}

	// file, err := os.Create("person.json")
	// if err != nil {
	//     fmt.Println("Error creating file:", err)
	//     return
	// }
	// defer file.Close()

	// encoder := json.NewEncoder(file)
	// encoder.SetIndent("", "  ")
	// err = encoder.Encode(p)
	// if err != nil {
	//     fmt.Println("Error encoding JSON to file:", err)
	//     return
	// }

	// fmt.Println("JSON saved to file!")
	// json.NewEncoder(file).Encode(structVar) → Writes JSON to a file.

	// ile, err := os.ReadFile("person.json")
	// if err != nil {
	//     fmt.Println("Error reading file:", err)
	//     return
	// }

	// var p Person
	// err = json.Unmarshal(ile, &p)
	// if err != nil {
	//     fmt.Println("Error decoding JSON:", err)
	//     return
	// }

	// fmt.Println("Data from file:", p)
	// Read JSON file, convert to struct & print data.

	practice.Practice()

}

type Person struct {
	Name  string `json:"name"`
	Age   int    `json:"age"`
	Email string `json:"email"`
}
