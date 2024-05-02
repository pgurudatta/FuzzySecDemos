package main

import (
	"fmt"
	"net/smtp"
	"os"
)

// sendEmail sends an email with the given subject and body to the specified recipient
func sendEmail(subject, body, recipient string) error {
	// Simulated SMTP server configuration
	smtpServer := "smtp.example.com"
	smtpPort := "587"
	senderEmail := "sender@example.com"
	senderPassword := "secretpassword123" // Simulated sender's password

	// Construct the email message
	message := fmt.Sprintf("From: %s\r\nTo: %s\r\nSubject: %s\r\n\r\n%s", senderEmail, recipient, subject, body)

	// Send the email using SMTP
	auth := smtp.PlainAuth("", senderEmail, senderPassword, smtpServer)
	err := smtp.SendMail(smtpServer+":"+smtpPort, auth, senderEmail, []string{recipient}, []byte(message))
	if err != nil {
		return err
	}
	return nil
}

func main() {
	// Simulated user input (subject and body of the email)
	subject := "Important Announcement"
	body := "Please click on the following link to claim your prize: https://example.com/claim?token=123456"

	// Simulated recipient email address
	recipient := "recipient@example.com"

	// Send the email
	err := sendEmail(subject, body, recipient)
	if err != nil {
		fmt.Println("Error sending email:", err)
		os.Exit(1)
	}

	fmt.Println("Email sent successfully")
}
