package reader

import (
	"context"
	"log/slog"
	"os"
	"testing"

	"cmd/app/internal/lib/config"
	"cmd/app/internal/lib/rabbit"
	"cmd/app/internal/service/email"

	"github.com/stretchr/testify/assert"
)

func TestConfirmEmail(t *testing.T) {
	cfg := &config.Env{
		SMTPHost:     "test-smtp-host",
		SMTPPort:     "test-smtp-port",
		SMTPUsername: "test-smtp-username",
		SMTPPassword: "test-smtp-password",
	}
	log := slog.New(slog.NewTextHandler(os.Stdout, nil))

	rb := &rabbit.RabbitService{
		Rabbit:  &rabbit.MockRabbit{},
		Connect: &rabbit.MockConnect{},
	}
	mockEmail := email.InitMock()

	reader := NewReader(cfg, log, rb.Rabbit, mockEmail)

	tests := []struct {
		name string
		want error
	}{
		{name: "test-1", want: nil},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {

			rb.Rabbit.Consume(context.Background(), "test-queue-confirm-email")
			err := reader.Read(context.Background())
			assert.Equal(t, test.want, err)
		})
	}
}
