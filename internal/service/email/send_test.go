package email

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSendEmail(t *testing.T) {
	mail := InitMock()

	tests := []struct {
		name string
		want error
	}{
		{
			name: "test-1",
			want: nil,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			err := mail.SendEmail(context.Background(), "test@test.com", "test", "test")
			assert.NoError(t, err)
		})
	}
}
