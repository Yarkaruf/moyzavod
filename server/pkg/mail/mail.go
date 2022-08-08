package mail

import (
	"net/smtp"
)

/*
 moyzavod
SMTP Username:
AKIA25OEZTUSTRAIQ3OE
SMTP Password:
BBXoBXfuCMC9C3MjoMmP5WTI46QSAjtNYDDLKGu0yKif

email-smtp.eu-central-1.amazonaws.com


SMTP endpoint
email-smtp.eu-central-1.amazonaws.com

STARTTLS Port
25, 587 or 2587

Transport Layer Security (TLS)
Required

TLS Wrapper Port
465 or 2465


*/

// SendMail ...
func SendMail(recipient string, body string) error {
	// Connect to the server, authenticate, set the sender and recipient,
	// and send the email all in one step.
	to := []string{recipient}
	msg := []byte("To: " + recipient + "\r\n" +
		"Subject: МойЗавод\r\n" +
		"\r\n" + body + "\r\n")
	return smtp.SendMail("localhost:25", nil, "noreply@moyzavod.com", to, msg)
}
