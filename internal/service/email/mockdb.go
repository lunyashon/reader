package email

import (
	"context"
)

type MockEmailSender struct{}

func (m *MockEmailSender) SendEmail(ctx context.Context, to, subject, body string) error {
	return nil
}

func InitMock() *MockEmailSender {
	return &MockEmailSender{}
}
