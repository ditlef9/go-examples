package main

import (
	"fmt"
)

// viewCart displays the contents of the shopping cart.
func viewCart(cart []Item) {
	// A slice is passed by value here, but it is a reference type in Go.
	// We don't need to modify the original cart, just read its contents.
	for _, item := range cart {
		fmt.Printf("Item: %s, Quantity: %d, Price: %.2f\n", item.Name, item.Quantity, item.Price)
	}
}
