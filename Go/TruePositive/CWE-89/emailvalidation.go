package main

import (
	"fmt"
	"net/smtp"
)

func sendEmail(to, subject, content string) error {
	// User input directly used in email content
	msg := fmt.Sprintf("Subject: %s\n\n%s", subject, content)

	// ... (code to send email using msg)

	return nil
}

func main() {
	userEmail := "user@example.com"
	emailContent := "This is some content with <script>alert('XSS!')</script>"
	sendEmail(userEmail, "Test Email", emailContent)
}
