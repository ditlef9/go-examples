package main

import "fmt"

// SMSNotifier is a struct that represents an SMS notification.
type SMSNotifier struct {
	PhoneNumber string // Phone number to send notifications to
}

// Notify method for SMSNotifier to satisfy the Notifier interface.
// This method sends a notification by formatting the phone number and message.
func (s SMSNotifier) Notify(message string) {
	fmt.Printf("Sending SMS to %s: %s\n", s.PhoneNumber, message)
}
