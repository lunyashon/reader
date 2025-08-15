package config

import (
	"flag"
	"os"

	"github.com/joho/godotenv"
)

func InitConfig() Config {
	env := readEnv()

	return Config{
		Env: env,
	}
}

func readFlags() string {

	var env string

	flag.StringVar(&env, "env", "", "path to data file")
	flag.Parse()

	if env == "" {
		env = os.Getenv("CONFIG_ENV")
		if env == "" {
			env = "./configs/config.env"
		}
	}

	return env
}

func readEnv() *Env {
	env := readFlags()

	if err := godotenv.Load(env); err != nil {
		panic(err)
	}

	return &Env{
		RabbitName:              os.Getenv("RABBIT_NAME"),
		RabbitHost:              os.Getenv("RABBIT_HOST"),
		RabbitPort:              os.Getenv("RABBIT_PORT"),
		RabbitPassword:          os.Getenv("RABBIT_PASSWORD"),
		RabbitQueueConfirmEmail: os.Getenv("RABBIT_QUEUE_CONFIRM_EMAIL"),
		RabbitQueueForgotToken:  os.Getenv("RABBIT_QUEUE_FORGOT_TOKEN"),
		LogPath:                 os.Getenv("LOG_PATH"),
		SMTPHost:                os.Getenv("SMTP_HOST"),
		SMTPPort:                os.Getenv("SMTP_PORT"),
		SMTPUsername:            os.Getenv("SMTP_USERNAME"),
		SMTPPassword:            os.Getenv("SMTP_PASSWORD"),
		SMTPFrom:                os.Getenv("SMTP_FROM"),
		MainDomain:              os.Getenv("MAIN_DOMAIN"),
	}
}
