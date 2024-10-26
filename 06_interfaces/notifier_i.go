package main

// Notifier is an interface for sending notifications
type NotifierI interface {
	Notify(message string) // Method signature for notifying
}

// Function to send notifications using the Notifier interface
func NotifyAll(notifiers []NotifierI, message string) {
	for _, notifier := range notifiers {
		notifier.Notify(message) // Call the Notify method of each notifier
	}
}
