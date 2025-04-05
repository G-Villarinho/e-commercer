package clients

import (
	"context"

	"github.com/g-villarinho/flash-buy-api/config"
	"github.com/g-villarinho/flash-buy-api/pkgs"
	"gopkg.in/gomail.v2"
)

type SMTPClient interface {
	SendEmail(ctx context.Context, to, subject, content string) error
}

type smptClient struct {
	di     *pkgs.Di
	dialer *gomail.Dialer
	from   string
}

func NewSMTPClient(di *pkgs.Di) (SMTPClient, error) {
	return &smptClient{
		di:     di,
		dialer: gomail.NewDialer(config.Env.SMTP.Host, config.Env.SMTP.Port, config.Env.SMTP.User, config.Env.SMTP.Password),
		from:   config.Env.SMTP.User,
	}, nil
}

func (s *smptClient) SendEmail(ctx context.Context, to string, subject string, content string) error {
	msg := gomail.NewMessage()
	msg.SetHeader("From", s.from)
	msg.SetHeader("To", to)
	msg.SetHeader("Subject", subject)
	msg.SetBody("text/html", content)

	if err := s.dialer.DialAndSend(msg); err != nil {
		return err
	}

	return nil
}
