package main

import (
	"context"
	"log/slog"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/lunyashon/reader/internal/app"
	"github.com/lunyashon/reader/internal/lib/config"
	"github.com/lunyashon/reader/internal/lib/logger"
	"github.com/lunyashon/reader/internal/lib/rabbit"
	"github.com/lunyashon/reader/internal/lib/waitgroup"
	"github.com/lunyashon/reader/internal/service/email"
)

func main() {

	cfg := config.InitConfig()
	log := logger.InitLog(cfg.Env.LogPath)

	mail := email.NewSMTPEmail(cfg.Env, log)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)

	rb := rabbit.InitRabbit(
		log,
		cfg.Env,
		config.ConfigRabbit{
			MaxRetries: 5,
			RetryDelay: 5 * time.Second,
		},
	)

	wg := waitgroup.InitWg()
	wg.Add("main")

	go func(wg *waitgroup.WaitGroup) {
		defer wg.Done("main")
		log.Info("launch app")
		if err := app.Launch(ctx, rb, cfg.Env, log, mail); err != nil {
			log.Error("failed to launch app", "error", err)
			cancel()
		}
	}(wg)

	select {
	case <-sigChan:
		log.Info("received signal to shutdown")
		cancel()
	case <-ctx.Done():
		log.Info("context done")
	}

	shdCtx, shdCancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer shdCancel()

	wg.Wait()

	shutdown(shdCtx, rb, log)

}

func shutdown(ctx context.Context, rb *rabbit.RabbitService, log *slog.Logger) {
	closeWithContext(ctx, rb.Connect.CloseChannel, log)
	closeWithContext(ctx, rb.Connect.CloseConnection, log)
	log.Info("app is shutdown")
}

func closeWithContext(ctx context.Context, fn func() error, log *slog.Logger) {
	done := make(chan error, 1)
	go func() {
		done <- fn()
	}()

	select {
	case err := <-done:
		if err != nil {
			log.Error("failed to close", "error", err)
		}
	case <-ctx.Done():
		log.Info("context done")
	}
}
