package main

import "fmt"

// Function to explore slices
func exploreSlices() {
	fmt.Println("\n\nexploreSlices()::Init -------------------------")

	// A slice is a dynamically-sized, flexible view into the elements of an array.
	// Unlike arrays, which have a fixed size, slices can grow and shrink as needed.
	// Slices are more commonly used in Go because of their flexibility.

	// Slices in Go are similar to lists in Python and ArrayList in Java.
	// In Python, lists are dynamic arrays that can hold elements of different types.
	// In Java, ArrayList provides a resizable array implementation that allows for dynamic resizing.

	// Creating a slice of strings
	slice := []string{"apple", "banana", "cherry"}

	fmt.Println("Slices:")
	fmt.Println("Slice:", slice)

	// Appending to a slice
	slice = append(slice, "date") // Slices can grow in size dynamically
	fmt.Println("After appending:", slice)

	// Removing an item from a slice (not a built-in function, but we can demonstrate how to do it)
	// For example, to remove "banana":
	slice = append(slice[:1], slice[2:]...) // Keep elements before "banana" and after "banana"
	fmt.Println("After removing 'banana':", slice)
}
