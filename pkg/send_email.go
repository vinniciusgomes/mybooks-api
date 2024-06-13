package pkg

import (
	"errors"
	"fmt"
	"os"

	"github.com/resend/resend-go/v2"
)

// SendEmail sends an email using the Resend API.
//
// Parameters:
// - toEmail: a slice of email addresses to send the email to.
// - subject: the subject of the email.
// - body: the body of the email.
//
// Returns:
// - error: an error if there was a problem sending the email.
func SendEmail(toEmail []string, subject string, body string) error {
	const emailFrom = "MyBooks <mybooks@vinniciusgomes.com>"
	const replyTo = "reply@vinniciusgomes.com"

	apiKey := os.Getenv("RESEND_API_KEY")
	if apiKey == "" {
		return errors.New("RESEND_API_KEY not set")
	}

	client := resend.NewClient(apiKey)

	params := &resend.SendEmailRequest{
		From:    emailFrom,
		To:      toEmail,
		Html:    body,
		Subject: subject,
		ReplyTo: replyTo,
	}

	sent, err := client.Emails.Send(params)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	fmt.Println(sent.Id)

	return nil
}
