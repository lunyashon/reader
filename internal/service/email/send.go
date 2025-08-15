package email

import (
	"context"
	"errors"
	"time"

	"gopkg.in/gomail.v2"
)

func (e *SMTPEmail) SendEmail(ctx context.Context, to, subject, body string) error {
	m := gomail.NewMessage()
	m.SetHeader("From", e.username)
	m.SetHeader("To", to)
	m.SetHeader("Subject", subject)
	m.SetBody("text/html", body)

	d := gomail.NewDialer(e.host, e.port, e.username, e.password)

	for i := 0; i < 3; i++ {
		if err := d.DialAndSend(m); err != nil {
			e.log.Error("failed to send email from %s to %s", "from", e.username, "to", to, "error", err)
			time.Sleep(time.Second * 2)
			continue
		}

		return nil
	}

	return errors.New("failed to send email after 3 attempts")
}
