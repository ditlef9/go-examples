package main

import "fmt"

// CalculateTotal computes the total cost of items in the shopping cart.
func viewTotal(cart []Item) {
	total := 0.0 // Initialize total cost
	for _, item := range cart {
		total += float64(item.Quantity) * item.Price // Calculate total cost
	}
	fmt.Printf("Total: %.2f\n", total)
}
