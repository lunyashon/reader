package rabbit

import (
	"context"
	"log/slog"
	"time"

	"github.com/lunyashon/reader/internal/lib/config"

	amqp "github.com/rabbitmq/amqp091-go"
)

type RabbitService struct {
	Rabbit  RabbitProvider
	Connect ConnectProvider
}

type Rabbit struct {
	log   *slog.Logger
	cfg   *config.Env
	conn  *amqp.Connection
	ch    *amqp.Channel
	rbcfg *RabbitConfig
}

type ConnectProvider interface {
	Connect() error
	Channel() error
	CloseConnection() error
	CloseChannel() error
	IsConnected() error
}

type RabbitProvider interface {
	DeclareQueue(
		ctx context.Context,
		queue string,
	) (amqp.Queue, error)
	Consume(
		ctx context.Context,
		queue string,
	) (<-chan amqp.Delivery, error)
	SetQos(
		prefetchCount int,
		prefetchSize int,
	) error
}

type RabbitConfig struct {
	MaxRetries int           `yaml:"max_retries" env-default:"5"`
	RetryDelay time.Duration `yaml:"retry_delay" env-default:"1s"`
}
