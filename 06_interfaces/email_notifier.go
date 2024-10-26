package main

import "fmt"

// EmailNotifier is a struct that represents an email notification.
// It contains an EmailAddress field, which holds the recipient's email address.
type EmailNotifier struct {
	EmailAddress string // Email address to send notifications to
}

// Send method for EmailNotifier to satisfy the Notifier interface.
// The 'e' parameter is a receiver of type EmailNotifier, allowing access to its fields.
// This method takes a string 'message' as an argument and formats it for email notification.
// It prints the message to the console, simulating the action of sending an email.
// The function name is Send, and it is a method of the EmailNotifier struct.
// It adheres to the Notifier interface's method signature, enabling polymorphism.
func (e EmailNotifier) Notify(message string) {
	fmt.Printf("Sending Email to %s: %s\n", e.EmailAddress, message)
}
