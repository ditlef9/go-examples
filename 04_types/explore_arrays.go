package main

import "fmt"

// Function to explore arrays
func exploreArrays() {
	fmt.Println("\n\nexploreArrays()::Init -------------------------")

	// An array is a fixed-size sequence of elements of the same type.
	// In Go, arrays have a defined length that cannot be changed after creation.
	// They are useful when you know the number of elements ahead of time.

	// Arrays in Go are similar to arrays in Java and lists in Python (to some extent).
	// In Java, arrays are fixed in size and can hold elements of a single data type.
	// In Python, lists are dynamic arrays that can hold elements of different types and have a variable size.

	// Creating an array of integers with a fixed size of 5
	var arr [5]int
	arr[0] = 1
	arr[1] = 2
	arr[2] = 3
	arr[3] = 4
	arr[4] = 5

	fmt.Println("Arrays:")
	fmt.Println("Array:", arr)

	// Accessing an element by index
	fmt.Printf("Element at index 2: %d\n", arr[2]) // Access the third element (0-based index)

	// Looping through the array
	fmt.Println("Looping through the array:")
	for i := 0; i < len(arr); i++ {
		fmt.Printf("Element at index %d: %d\n", i, arr[i])
	}
}
