package main

import (
	"fmt"
	"net/smtp"
)

func sendEmail(recipient string, subject string, body string) error {
  // ... configuration details for email server omitted for brevity

  msg := "From: noreply@yourcompany.com\r\n"
  msg += "To: " + recipient + "\r\n"
  msg += "Subject: " + subject + "\r\n"
  msg += "\r\n"
  msg += body + "\r\n"

  err := smtp.SendMail(serverName, auth, msg)
  if err != nil {
    return err
  }
  return nil
}

func main() {
  recipient := "user@example.com"
  subject := "Your Account Update"
  body := "Hello, your account has been updated."

  err := sendEmail(recipient, subject, body)
  if err != nil {
    fmt.Println("Error sending email:", err)
  } else {
    fmt.Println("Email sent successfully!")
  }
}
