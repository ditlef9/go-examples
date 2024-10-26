package main

import "fmt"

// PushNotifier is a struct that represents a push notification.
type PushNotifier struct {
	AppID string // App ID to send push notifications to
}

// Notify method for PushNotifier to satisfy the Notifier interface.
// This method sends a notification by formatting the app ID and message.
func (p PushNotifier) Notify(message string) {
	fmt.Printf("Sending Push Notification to App %s: %s\n", p.AppID, message)
}
