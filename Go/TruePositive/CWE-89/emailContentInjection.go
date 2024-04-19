package main

import (
	"fmt"
	"log"
	"net/smtp"
)

func main() {
	to := "recipient@example.com"
	from := "sender@example.com"
	subject := "Important Notification"
	body := "Hello, please click on the following link to reset your password: https://example.com/reset?token=abc123"

	// Send email using SendMail function
	err := SendMail(to, from, subject, body)
	if err != nil {
		log.Fatalf("Failed to send email: %v", err)
	}
	fmt.Println("Email sent successfully!")
}

// SendMail sends an email using SMTP
func SendMail(to, from, subject, body string) error {
	// SMTP server configuration (example only - replace with actual SMTP server details)
	smtpServer := "smtp.example.com"
	smtpPort := "587"
	smtpUsername := "username"
	smtpPassword := "password"

	// Compose the email message
	message := fmt.Sprintf("To: %s\r\n", to) +
		fmt.Sprintf("From: %s\r\n", from) +
		fmt.Sprintf("Subject: %s\r\n", subject) +
		"\r\n" +
		body

	// Connect to the SMTP server
	auth := smtp.PlainAuth("", smtpUsername, smtpPassword, smtpServer)
	err := smtp.SendMail(smtpServer+":"+smtpPort, auth, from, []string{to}, []byte(message))
	if err != nil {
		return err
	}
	return nil
}
