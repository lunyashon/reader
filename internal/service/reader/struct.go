package reader

import (
	"log/slog"

	"github.com/lunyashon/reader/internal/lib/config"
	"github.com/lunyashon/reader/internal/lib/rabbit"
	"github.com/lunyashon/reader/internal/service/email"
)

type Reader struct {
	cfg   *config.Env
	log   *slog.Logger
	rb    rabbit.RabbitProvider
	email email.EmailInterface
}

func NewReader(
	cfg *config.Env,
	log *slog.Logger,
	rb rabbit.RabbitProvider,
	mail email.EmailInterface,
) *Reader {
	return &Reader{
		cfg:   cfg,
		log:   log,
		rb:    rb,
		email: mail,
	}
}
