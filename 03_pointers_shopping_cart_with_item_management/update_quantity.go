package main

import (
	"fmt"
)

// updateQuantity updates the quantity of a specific item in the shopping cart.
func updateQuantity(cart *[]Item, itemName string, newQuantity int) {
	// Pointer to the cart slice is used here to modify the original cart.
	for i := range *cart {
		if (*cart)[i].Name == itemName {
			(*cart)[i].Quantity = newQuantity // Update quantity directly in the original slice.
			fmt.Printf("Updated %s quantity to %d.\n", itemName, newQuantity)
			return
		}
	}
	fmt.Println("Item not found in cart.")
}
