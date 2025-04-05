package notifications

import (
	"bytes"
	"context"
	"fmt"
	"log"
	"text/template"

	"github.com/g-villarinho/xp-life-api/clients"
	"github.com/g-villarinho/xp-life-api/models"
	"github.com/g-villarinho/xp-life-api/pkgs"
)

type EmailNotification interface {
	SendVerificationEmail(ctx context.Context, email, code string) error
}

type emailNotification struct {
	di *pkgs.Di
	sc clients.SMTPClient
}

func NewEmailNotification(di *pkgs.Di) (EmailNotification, error) {
	sc, err := pkgs.Invoke[clients.SMTPClient](di)
	if err != nil {
		return nil, fmt.Errorf("invoke clients.SMTPClient: %w", err)
	}

	return &emailNotification{
		di: di,
		sc: sc,
	}, nil
}

func (e *emailNotification) SendVerificationEmail(ctx context.Context, email string, code string) error {
	tmpl, err := template.ParseFiles("notifications/templates/verification-email.html")
	if err != nil {
		log.Fatalf("parse template: %v", err)
	}

	var htmlBuffer bytes.Buffer
	data := models.VerificationEmailData{
		Code: code,
	}

	if err := tmpl.Execute(&htmlBuffer, data); err != nil {
		return fmt.Errorf("execute template: %w", err)
	}

	if err := e.sc.SendEmail(ctx, email, "XP Life - Verification Email", htmlBuffer.String()); err != nil {
		return fmt.Errorf("send email: %w", err)
	}

	return nil
}
