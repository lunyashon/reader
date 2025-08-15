package rabbit

import (
	"context"

	amqp "github.com/rabbitmq/amqp091-go"
	"github.com/stretchr/testify/mock"
)

type MockRabbit struct {
	mock.Mock
	deleviry chan amqp.Delivery
}

func (m *MockRabbit) DeclareQueue(ctx context.Context, queue string) (amqp.Queue, error) {
	return m.Called(ctx, queue).Get(0).(amqp.Queue), m.Called(ctx, queue).Error(1)
}

func (m *MockRabbit) Consume(ctx context.Context, queue string) (<-chan amqp.Delivery, error) {
	m.deleviry = make(chan amqp.Delivery, 1)
	m.deleviry <- amqp.Delivery{
		Body: []byte("test-message"),
	}
	close(m.deleviry)
	return m.deleviry, m.Called(ctx, queue).Error(1)
}

func (m *MockRabbit) SetQos(prefetchCount int, prefetchSize int) error {
	return m.Called(prefetchCount, prefetchSize).Error(0)
}

type MockConnect struct {
	mock.Mock
}

func (m *MockConnect) Connect() error {
	return m.Called().Error(0)
}

func (m *MockConnect) Channel() error {
	return m.Called().Error(0)
}

func (m *MockConnect) CloseConnection() error {
	return m.Called().Error(0)
}

func (m *MockConnect) CloseChannel() error {
	return m.Called().Error(0)
}

func (m *MockConnect) IsConnected() error {
	return m.Called().Error(0)
}

type MockDB struct {
	Rabbit  *MockRabbit
	Connect *MockConnect
}

func InitMock() *MockDB {
	return &MockDB{
		Rabbit:  new(MockRabbit),
		Connect: new(MockConnect),
	}
}
