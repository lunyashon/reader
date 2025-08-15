package rabbit

import (
	"fmt"
	"log/slog"
	"sync"

	"github.com/lunyashon/reader/internal/lib/config"

	amqp "github.com/rabbitmq/amqp091-go"
)

func InitRabbit(
	log *slog.Logger,
	cfg *config.Env,
	rbcfg config.ConfigRabbit,
) *RabbitService {

	var (
		once     sync.Once
		provider *Rabbit
		err      error
	)

	once.Do(func() {
		provider = &Rabbit{
			log: log,
			cfg: cfg,
			rbcfg: &RabbitConfig{
				MaxRetries: rbcfg.MaxRetries,
				RetryDelay: rbcfg.RetryDelay,
			},
		}
		if err = provider.Connect(); err != nil {
			panic(err)
		}
		if err = provider.Channel(); err != nil {
			panic(err)
		}
	})
	return &RabbitService{
		Rabbit:  provider,
		Connect: provider,
	}
}

func (r *Rabbit) Connect() error {
	conn, err := amqp.Dial(
		fmt.Sprintf(
			"amqp://%s:%s@%s:%s/",
			r.cfg.RabbitName,
			r.cfg.RabbitPassword,
			r.cfg.RabbitHost,
			r.cfg.RabbitPort,
		),
	)
	if err != nil {
		r.log.Error("failed to dial", "error", err)
		return err
	}
	r.conn = conn
	return nil
}

func (r *Rabbit) Channel() error {
	ch, err := r.conn.Channel()
	if err != nil {
		r.log.Error("failed to open a channel", "error", err)
		return err
	}
	r.ch = ch
	return nil
}

func (r *Rabbit) CloseConnection() error {
	return r.conn.Close()
}

func (r *Rabbit) CloseChannel() error {
	return r.ch.Close()
}

func (r *Rabbit) IsConnected() error {
	if r.conn == nil || r.conn.IsClosed() {
		return fmt.Errorf("connection is not established")
	}
	if r.ch == nil || r.ch.IsClosed() {
		return fmt.Errorf("channel is not established")
	}
	return nil
}
