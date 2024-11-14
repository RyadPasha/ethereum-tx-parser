package notification

import "log"

// SendNotification is a method to simulate sending a notification.
// Currently, it does nothing but can be extended later.
func SendNotification(message string, recipient string) {
	log.Printf("Notification to %s: %s", recipient, message)
}
