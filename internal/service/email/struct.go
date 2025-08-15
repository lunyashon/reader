package email

import (
	"context"
	"log/slog"
	"strconv"

	"github.com/lunyashon/reader/internal/lib/config"
)

type SMTPEmail struct {
	host     string
	port     int
	username string
	password string
	log      *slog.Logger
}

type EmailData struct {
	Email string `json:"email"`
	Token string `json:"token"`
}

type EmailInterface interface {
	SendEmail(ctx context.Context, to, subject, body string) error
}

func NewSMTPEmail(cfg *config.Env, log *slog.Logger) *SMTPEmail {
	port, err := strconv.Atoi(cfg.SMTPPort)
	if err != nil {
		log.Error("failed to convert SMTP port to int", "error", err)
		port = 465
	}
	return &SMTPEmail{
		host:     cfg.SMTPHost,
		port:     port,
		username: cfg.SMTPUsername,
		password: cfg.SMTPPassword,
		log:      log,
	}
}
