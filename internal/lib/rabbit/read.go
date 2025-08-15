package rabbit

import (
	"context"

	amqp "github.com/rabbitmq/amqp091-go"
)

func (r *Rabbit) DeclareQueue(
	ctx context.Context,
	queue string,
) (amqp.Queue, error) {
	return r.ch.QueueDeclare(
		queue,
		true,
		false,
		false,
		false,
		amqp.Table{},
	)
}

func (r *Rabbit) Consume(
	ctx context.Context,
	queue string,
) (<-chan amqp.Delivery, error) {
	return r.ch.Consume(
		queue,
		"",
		false,
		false,
		false,
		false,
		nil,
	)
}

func (r *Rabbit) SetQos(prefetchCount int, prefetchSize int) error {
	return r.ch.Qos(
		prefetchCount,
		prefetchSize,
		false,
	)
}
