package main

import "fmt"

// Define a struct type for a Person
type Person struct {
	Name string
	Age  int
}

// Function to explore structs
func exploreStructs() {
	fmt.Println("\n\nexploreStructs()::Init -------------------------")
	// Creating an instance of Person struct
	p := Person{Name: "Alice", Age: 30}

	fmt.Println("Structs:")
	fmt.Printf("Person: Name=%s, Age=%d\n", p.Name, p.Age)
}
