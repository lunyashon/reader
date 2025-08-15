package reader

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/lunyashon/reader/internal/lib/waitgroup"
	"github.com/lunyashon/reader/internal/service/email"

	amqp "github.com/rabbitmq/amqp091-go"
)

func (r *Reader) Read(ctx context.Context) error {
	var (
		wg    = waitgroup.InitWg()
		chErr = make(chan error, 2)
	)

	wg.Add("readConfirmEmail")
	go r.ReadConfirmEmail(ctx, wg, chErr)

	wg.Add("readForgotToken")
	go r.ReadForgotToken(ctx, wg, chErr)

	wg.Wait()
	close(chErr)

	for err := range chErr {
		if err != nil {
			r.log.ErrorContext(ctx, "failed to read queue", "error", err)
			return err
		}
	}

	return nil
}

func (r *Reader) ReadConfirmEmail(ctx context.Context, wg *waitgroup.WaitGroup, chErr chan<- error) {

	defer wg.Done("readConfirmEmail")

	queue, err := r.rb.DeclareQueue(
		ctx,
		r.cfg.RabbitQueueConfirmEmail,
	)
	if err != nil {
		r.log.ErrorContext(ctx, "failed to declare queue", "error", err)
		chErr <- err
		return
	}

	if err := r.rb.SetQos(1, 0); err != nil {
		r.log.ErrorContext(ctx, "failed to set qos", "error", err)
		chErr <- err
		return
	}

	deliveries, err := r.rb.Consume(
		ctx,
		queue.Name,
	)
	if err != nil {
		r.log.ErrorContext(ctx, "failed to consume queue", "error", err)
		chErr <- err
		return
	}

	wg.Add("readConfirmEmailInQueue")

	go func() {
		defer wg.Done("readConfirmEmailInQueue")

		for {
			select {
			case <-ctx.Done():
				return
			case delivery, ok := <-deliveries:
				if !ok {
					r.log.InfoContext(ctx, "channel queue is closed")
					return
				}
				r.log.InfoContext(ctx, "received message", "message", string(delivery.Body))
				if err := r.processMessage(ctx, delivery, "confirm_email"); err != nil {
					r.log.ErrorContext(ctx, "failed to process message", "error", err)
					delivery.Nack(false, true)
					continue
				}
				delivery.Ack(false)
			}
		}
	}()
}

func (r *Reader) ReadForgotToken(ctx context.Context, wg *waitgroup.WaitGroup, chErr chan<- error) {

	defer wg.Done("readForgotToken")

	queue, err := r.rb.DeclareQueue(
		ctx,
		r.cfg.RabbitQueueForgotToken,
	)
	if err != nil {
		r.log.ErrorContext(ctx, "failed to declare queue", "error", err)
		chErr <- err
		return
	}

	if err := r.rb.SetQos(1, 0); err != nil {
		r.log.ErrorContext(ctx, "failed to set qos", "error", err)
		chErr <- err
		return
	}

	deliveries, err := r.rb.Consume(
		ctx,
		queue.Name,
	)
	if err != nil {
		r.log.ErrorContext(ctx, "failed to consume queue", "error", err)
		chErr <- err
		return
	}

	wg.Add("readForgotTokenInQueue")

	go func() {
		defer wg.Done("readForgotTokenInQueue")
		for {
			select {
			case <-ctx.Done():
				return
			case delivery, ok := <-deliveries:
				if !ok {
					r.log.InfoContext(ctx, "channel queue is closed")
					return
				}
				if err := r.processMessage(ctx, delivery, "forgot_token"); err != nil {
					r.log.ErrorContext(ctx, "failed to process message", "error", err)
					continue
				}
				delivery.Ack(false)
			}
		}
	}()
}

func (r *Reader) processMessage(
	ctx context.Context,
	delivery amqp.Delivery,
	queue string,
) error {
	select {
	case <-ctx.Done():
		return ctx.Err()
	default:
	}

	var (
		emailData = &email.EmailData{}
		message   string
		title     string
	)

	if err := json.Unmarshal(delivery.Body, emailData); err != nil {
		return err
	}

	switch queue {
	case "confirm_email":
		message = fmt.Sprintf("Confirm your email: %s", emailData.Token)
		title = "Confirm your email"
	case "forgot_token":
		message = getForgotPasswordText(r.cfg.MainDomain, "?key=", emailData.Token)
		title = "Forgot your password"
	}

	r.log.InfoContext(ctx, "sending email", "email", emailData.Email, "title", title, "message", message)
	if err := r.email.SendEmail(ctx, emailData.Email, title, message); err != nil {
		return err
	}

	return nil
}
