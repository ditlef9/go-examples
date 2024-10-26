package main

import (
	"fmt"
)

func main() {
	var cart []Item // A slice to hold items in the shopping cart (not a pointer since the slice itself is a reference type).

	for {
		fmt.Println("\n--- Grocery Shopping Cart ---")
		viewCart(cart)
		viewTotal(cart)
		fmt.Println("\n--- Menu ---")
		fmt.Println("1. Add item to cart")
		fmt.Println("2. Update item quantity")
		fmt.Println("3. View cart")
		fmt.Println("4. Exit")
		fmt.Print("Choose an option: ")

		var choice int
		fmt.Scan(&choice)

		switch choice {
		case 1:
			// Add item to cart
			var name string
			var quantity int
			var price float64

			fmt.Print("Enter item name: ")
			fmt.Scan(&name)
			fmt.Print("Enter quantity: ")
			fmt.Scan(&quantity)
			fmt.Print("Enter price: ")
			fmt.Scan(&price)

			// Add item to cart (using a value type)
			// 		Using value types here because we're creating new instances of Item
			// 		and appending them to the slice. There’s no need to modify existing data.
			// Alternative method to use same cart:
			// 		Use a pointer to cart when appending new items
			// 		This is a more explicit way to show that we're potentially modifying the slice.
			// 		addItem(&cart, Item{Name: name, Quantity: quantity, Price: price})
			cart = append(cart, Item{Name: name, Quantity: quantity, Price: price})
			fmt.Println("Item added to cart.")

		case 2:
			// Update item quantity
			var name string
			var newQuantity int

			fmt.Print("Enter item name to update: ")
			fmt.Scan(&name)
			fmt.Print("Enter new quantity: ")
			fmt.Scan(&newQuantity)

			// Update quantity using a pointer to the cart
			// Using a pointer here allows us to modify the original slice directly.
			updateQuantity(&cart, name, newQuantity)

		case 3:
			// View cart
			fmt.Println("\n--- Items in Cart ---")
			// Not using a pointer here because we are only reading the data.
			// The viewCart function will not modify the cart; it simply displays the contents.
			viewCart(cart)
			viewTotal(cart)

		case 4:
			// Exit the program
			fmt.Println("Exiting the program.")
			return

		default:
			fmt.Println("Invalid choice. Please choose a valid option.")
		}
	}
}
