package services

import (
	"bytes"
	"fmt"
	"html/template"
	"net/smtp"
	"os"
)

type NotificationService struct{}

func NewNotificationService() *NotificationService {
	return &NotificationService{}
}

func (s *NotificationService) SendNotification(to, subject, message string) error {
	from := os.Getenv("EMAIL_USER")
	password := os.Getenv("EMAIL_PASS")

	if from == "" || password == "" {
		return fmt.Errorf("email credentials not set in environment variables")
	}

	auth := smtp.PlainAuth("", from, password, "smtp.gmail.com")

	// Load the HTML email template
	tmpl, err := template.ParseFiles("internal/services/notification_template.txt")
	if err != nil {
		return err
	}

	var body bytes.Buffer
	tmpl.Execute(&body, map[string]string{"Message": message})

	msg := []byte(fmt.Sprintf(
		"To: %s\r\nSubject: %s\r\n\r\n%s",
		to, subject, body.String(),
	))

	err = smtp.SendMail("smtp.gmail.com:587", auth, from, []string{to}, msg)
	if err != nil {
		return fmt.Errorf("failed to send email: %v", err)
	}

	fmt.Printf("📧 Email successfully sent to %s\n", to)
	return nil
}
