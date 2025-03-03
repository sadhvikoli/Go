package main

import (
	"fmt"
	"mymodule/practice"
)

// type StructName struct {} → Declares a struct.
// Fields are variables inside the struct.
type Person struct {
	name string
	age  int
}

// Struct with Methods
func (p Person) greet() {
	fmt.Println("Hello, my name is", p.name)
}

// Pass Struct by Value (Does Not Modify Original)
func updateAge(p Person) {
	p.age = 35
}

// Pass Struct by Reference (Modifies Original)
func updateAge1(p *Person) {
	p.age = 35
}

// Embedding One Struct into Another
type Employee struct {
	name   string
	salary int
}

type Manager struct {
	Employee
	teamSize int
}

//	Declaring and Implementing an Interface
//
// Interface
type Shape interface {
	area() float64
}

// Struct implementing interface
type Circle struct {
	radius float64
}

func (c Circle) area() float64 {
	return 3.14 * c.radius * c.radius
}

// Multiple Structs Implementing the Same Interface
type Rectangle struct {
	width, height float64
}

func (r Rectangle) area() float64 {
	return r.width * r.height
}

func printArea(s Shape) {
	fmt.Println("Area:", s.area())
}

// Empty Interface (interface{}) - Accepts Any Type
func printAnything(val interface{}) {
	fmt.Println(val)
}

func main() {
	var p1 Person
	p1.name = "Alice"
	p1.age = 25

	fmt.Println(p1) // {Alice 25}

	// Initializing a Struct with Values
	p2 := Person{"Bob", 30}
	fmt.Println(p2)

	p3 := Person{name: "Charlie", age: 22}
	fmt.Println(p3)
	// Order-based → Person{"Bob", 30}
	// Key-value-based → Person{name: "Charlie", age: 22}

	p := Person{"David", 28}
	p.greet()
	//(p Person) greet() → greet is a method of the struct.

	updateAge(p)       // p is passed by value (copy)
	fmt.Println(p.age) // 30 (unchanged)

	updateAge1(&p)     // Pass by reference
	fmt.Println(p.age) // 35 (updated)
	// *Person → Pointer receiver modifies the original struct.

	m := Manager{Employee{"Alice", 50000}, 10}
	fmt.Println(m.name, m.salary, m.teamSize)
	// 	Manager inherits from Employee.
	// Fields of Employee are directly accessible in Manager.

	var s Shape = Circle{5}
	fmt.Println("Area:", s.area())
	// 	No implements keyword needed → If a struct has the methods of an interface, it automatically implements it.
	// var s Shape = Circle{5} → Polymorphism in action.

	r := Rectangle{10, 5}
	c := Circle{7}

	printArea(r)
	printArea(c)
	// 	Different structs (Rectangle, Circle) implement the same interface.
	// printArea(s Shape) → Works with any Shape.

	printAnything(42)
	printAnything("Hello, Go!")
	printAnything([]int{1, 2, 3})
	// interface{} can hold any data type.

	practice.ExecuteTasks()

}
