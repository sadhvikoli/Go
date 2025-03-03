package practice

import (
	"fmt"
	"reflect"
)

// Create a struct Car with fields brand, model, year and a method details().
type Car struct {
	brand string
	model string
	year  int
}

func (c Car) details() {
	fmt.Println("Car details:", c)
}

// Task 2: Implement a struct BankAccount with deposit() and withdraw() methods
type BankAccount struct {
	balance float64
}

func (b *BankAccount) Deposit(amount float64) {
	b.balance += amount
	fmt.Printf("Deposited: %.2f, New Balance: %.2f\n", amount, b.balance)
}

func (b *BankAccount) Withdraw(amount float64) {
	if amount > b.balance {
		fmt.Println("Insufficient balance!")
		return
	}
	b.balance -= amount
	fmt.Printf("Withdrawn: %.2f, New Balance: %.2f\n", amount, b.balance)
}

// Task 3: Create a struct Employee, embed it inside Manager, and print details
type Employee struct {
	Name   string
	Age    int
	Salary float64
}

type Manager struct {
	Employee
	Department string
}

func (m Manager) PrintDetails() {
	fmt.Printf("Manager Name: %s, Age: %d, Salary: %.2f, Department: %s\n",
		m.Name, m.Age, m.Salary, m.Department)
}

// Task 4: Define an interface Animal with a method speak(). Implement Dog and Cat structs.
type Animal interface {
	Speak() string
}

type Dog struct{}
type Cat struct{}

func (d Dog) Speak() string {
	return "Woof!"
}

func (c Cat) Speak() string {
	return "Meow!"
}

// Task 5: Function that takes an empty interface (interface{}) and prints the type
func printType(value interface{}) {
	fmt.Printf("Type: %s, Value: %v\n", reflect.TypeOf(value), value)
}

func ExecuteTasks() {

	c := Car{brand: "BMW", model: "F13", year: 2030}
	c.details()

	// Task 2: BankAccount operations
	account := BankAccount{balance: 1000}
	account.Deposit(500)
	account.Withdraw(300)
	account.Withdraw(1500) // Should print insufficient balance

	// Task 3: Manager details
	manager := Manager{
		Employee:   Employee{Name: "Alice", Age: 35, Salary: 75000},
		Department: "IT",
	}
	manager.PrintDetails()

	// Task 4: Animal interface implementation
	var animals []Animal = []Animal{Dog{}, Cat{}}
	for _, animal := range animals {
		fmt.Println(animal.Speak())
	}

	// Task 5: Print type of different values
	printType(42)
	printType(3.14)
	printType("Hello, Go!")
	printType(true)
	printType([]int{1, 2, 3})
}
