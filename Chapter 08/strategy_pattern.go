package main

import "fmt"

// EmailSender defines the strategy function type for sending emails
type EmailSender func(subject, body, recipient string) error

// MockEmailSender simulates sending an email using a mock service
func MockEmailSender(subject, body, recipient string) error {
    fmt.Printf("Sending email using Mock service\nTo: %s\nSubject: %s\nBody: %s\n", recipient, subject, body)
    return nil // Simulate success
}

// SMTPEmailSender simulates sending an email using an SMTP server
func SMTPEmailSender(subject, body, recipient string) error {
    fmt.Printf("Sending email using SMTP server\nTo: %s\nSubject: %s\nBody: %s\n", recipient, subject, body)
    return nil // Simulate success
}

// SendEmail uses the provided email sender strategy to send an email
func SendEmail(sender EmailSender, subject, body, recipient string) error {
    return sender(subject, body, recipient)
}

func main() {
    recipient := "user@example.com"
    subject := "Hello World"
    body := "This is a test email."

    // Dynamically choose the email sending strategy
    var currentSender EmailSender

    // Example scenario: use the mock sender for development environment
    environment := "development" // This could be set based on external configuration

    if environment == "development" {
        currentSender = MockEmailSender
    } else {
        currentSender = SMTPEmailSender
    }

    // Use the selected email sender
    if err := SendEmail(currentSender, subject, body, recipient); err != nil {
        fmt.Println("Error sending email:", err)
    }
}
