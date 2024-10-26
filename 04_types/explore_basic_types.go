package main

import "fmt"

// Function to explore basic data types
func exploreBasicTypes() {
	fmt.Println("\n\nexploreBasicTypes()::Init -------------------------")
	var integer int = 42          // Integer type
	var floatNum float64 = 3.14   // Float type
	var boolean bool = true       // Boolean type
	var str string = "Hello, Go!" // String type

	fmt.Println("Basic Data Types:")
	fmt.Printf("Integer: %d\n", integer)
	fmt.Printf("Float: %.2f\n", floatNum)
	fmt.Printf("Boolean: %t\n", boolean)
	fmt.Printf("String: %s\n", str)
}
