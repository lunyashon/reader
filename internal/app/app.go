package app

import (
	"context"
	"log/slog"

	"github.com/lunyashon/reader/internal/lib/config"
	"github.com/lunyashon/reader/internal/lib/rabbit"
	"github.com/lunyashon/reader/internal/service/email"
	"github.com/lunyashon/reader/internal/service/reader"
)

func Launch(
	ctx context.Context,
	rb *rabbit.RabbitService,
	cfg *config.Env,
	log *slog.Logger,
	mail email.EmailInterface,
) error {
	rdr := reader.NewReader(cfg, log, rb.Rabbit, mail)

	return rdr.Read(ctx)
}
