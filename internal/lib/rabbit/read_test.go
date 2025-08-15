package rabbit

import (
	"context"
	"testing"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestDeclareQueue(t *testing.T) {
	mockDB := InitMock()
	expectedQueue := amqp.Queue{Name: "test-queue"}
	mockDB.Rabbit.On("DeclareQueue", mock.Anything, "test-queue").Return(expectedQueue, nil)

	queue, err := mockDB.Rabbit.DeclareQueue(context.Background(), "test-queue")

	assert.NoError(t, err)
	assert.Equal(t, expectedQueue, queue)
	mockDB.Rabbit.AssertExpectations(t)
}

func TestConsume(t *testing.T) {
	mockDB := InitMock()
	mockDB.Rabbit.On("Consume", mock.Anything, "test-queue").Return(nil, nil)

	tests := []struct {
		name     string
		queue    string
		delivery string
		wantErr  error
	}{
		{
			name:     "test-queue",
			queue:    "test-queue",
			delivery: "email",
			wantErr:  nil,
		},
		{
			name:     "test-queue",
			queue:    "test-queue",
			delivery: "sms",
			wantErr:  nil,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {

			mockDB.Rabbit.On("Consume", mock.Anything, test.queue).Return(mockDB.Rabbit.deleviry, nil)

			deliveries, err := mockDB.Rabbit.Consume(context.Background(), test.queue)
			t.Log("deliveries", deliveries)
			t.Log("err", err)
			select {
			case <-deliveries:
				t.Log("delivery received")
			case <-time.After(1 * time.Second):
				t.Fatal("timeout")
			}
		})
	}
	mockDB.Rabbit.AssertExpectations(t)
}

func TestSetQos(t *testing.T) {
	mockDB := InitMock()
	mockDB.Rabbit.On("SetQos", mock.Anything, mock.Anything).Return(nil)

	err := mockDB.Rabbit.SetQos(10, 0)

	assert.NoError(t, err)
	mockDB.Rabbit.AssertExpectations(t)
}
