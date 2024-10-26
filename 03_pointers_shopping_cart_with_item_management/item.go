package main

// Item struct represents a grocery item with a name, quantity, and price.
type Item struct {
	Name     string  // Name of the item (not a pointer because we want to hold the value directly)
	Quantity int     // Quantity of the item (not a pointer as it's a simple integer value)
	Price    float64 // Price of the item (not a pointer since we only need the value)
}
