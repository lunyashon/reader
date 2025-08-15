package config

import "time"

type Env struct {
	RabbitName              string `env:"RABBIT_NAME" env-default:"reader"`
	RabbitHost              string `env:"RABBIT_HOST" env-default:"localhost"`
	RabbitPort              string `env:"RABBIT_PORT" env-default:"5672"`
	RabbitPassword          string `env:"RABBIT_PASSWORD" env-default:"guest"`
	RabbitQueueConfirmEmail string `env:"RABBIT_QUEUE_CONFIRM_EMAIL" env-default:"confirm_email"`
	RabbitQueueForgotToken  string `env:"RABBIT_QUEUE_FORGOT_TOKEN" env-default:"forgot_token"`
	LogPath                 string `env:"LOG_PATH" env-default:"logs"`
	SMTPHost                string `env:"SMTP_HOST" env-default:""`
	SMTPPort                string `env:"SMTP_PORT" env-default:""`
	SMTPUsername            string `env:"SMTP_USERNAME" env-default:""`
	SMTPPassword            string `env:"SMTP_PASSWORD" env-default:""`
	SMTPFrom                string `env:"SMTP_FROM" env-default:""`
	MainDomain              string `env:"MAIN_DOMAIN" env-default:""`
}

type Config struct {
	Env *Env
}

type ConfigRabbit struct {
	MaxRetries int           `env:"RABBIT_MAX_RETRIES" env-default:"5"`
	RetryDelay time.Duration `env:"RABBIT_RETRY_DELAY" env-default:"1s"`
}
