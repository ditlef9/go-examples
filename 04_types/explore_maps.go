package main

import "fmt"

// Function to explore maps
func exploreMaps() {
	fmt.Println("\n\nexploreMaps()::Init -------------------------")

	// A map is an unordered collection of key-value pairs.
	// In Go, maps are built-in data types that allow you to associate values with keys for efficient lookups.
	// The keys must be unique, and they are used to access the corresponding values.

	// Maps in Go are similar to dictionaries in Python and HashMaps in Java.
	// In Python, dictionaries are mutable, unordered collections of key-value pairs.
	// In Java, HashMap allows for storing key-value pairs where keys are unique, and it provides constant-time performance for basic operations (like add, remove, contains, and get).

	// Creating a map of string keys to int values
	m := make(map[string]int)
	m["apple"] = 1
	m["banana"] = 2
	m["cherry"] = 3

	fmt.Println("Maps:")
	fmt.Println("Map:", m)

	// Accessing a value by its key
	fmt.Printf("Value for 'banana': %d\n", m["banana"])

	// Checking if a key exists in the map
	if value, exists := m["banana"]; exists {
		fmt.Printf("Key 'banana' exists with value: %d\n", value)
	} else {
		fmt.Println("Key 'banana' does not exist.")
	}

	// Attempting to access a non-existent key
	if value, exists := m["grape"]; exists {
		fmt.Printf("Key 'grape' exists with value: %d\n", value)
	} else {
		fmt.Println("Key 'grape' does not exist.")
	}
}
