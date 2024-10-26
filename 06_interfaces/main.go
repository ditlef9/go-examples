package main

// Main function
func main() {
	// Create instances of different notifiers
	emailNotifier := EmailNotifier{EmailAddress: "user@example.com"}
	smsNotifier := SMSNotifier{PhoneNumber: "+1234567890"}
	pushNotifier := PushNotifier{AppID: "com.example.app"}

	// Create a slice of Notifier interfaces
	notifiers := []NotifierI{emailNotifier, smsNotifier, pushNotifier}

	// Send a notification to all notifiers
	NotifyAll(notifiers, "This is a test notification!")
}
