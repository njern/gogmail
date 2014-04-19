package gogmail

import (
	"fmt"
	"net/smtp"
	"strings"
)

const (
	SERVER = "smtp.gmail.com"
	PORT   = 587
)

// Gmail encapsulates an authenticated GMail user.
type GMail struct {
	username string
	auth     *smtp.Auth
}

// GmailConnection creates a GMail object, which encapsulates the connection to GMail's server(s).
func GmailConnection(username, password string) *GMail {
	auth := smtp.PlainAuth("", username, password, SERVER)

	return &GMail{username, &auth}
}

// SendMail() sends a mail via GMail - assuming you have valid credentials and a working internet connection.
func (gmail *GMail) SendMail(recipients []string, subject, body string, isHTML bool) error {
	// Remove any newlines before the body...
	sender := strings.Replace(gmail.username, "\n", "", -1)
	recipientString := strings.Replace(strings.Join(recipients, ", "), "\n", "", -1)
	subject = strings.Replace(subject, "\n", "", -1)

	var message string
	if isHTML {
		mime := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"
		message = fmt.Sprintf("From: %s\nTo: %s\nSubject: %s\n\n%s%s", sender, recipientString, subject, mime, body)
	} else {
		message = fmt.Sprintf("From: %s\nTo: %s\nSubject: %s\n\n%s", sender, recipientString, subject, body)
	}

	err := smtp.SendMail(fmt.Sprintf("%s:%d", SERVER, PORT),
		*gmail.auth,
		gmail.username,
		recipients,
		[]byte(message))

	if err != nil {
		return err
	}

	return nil
}
