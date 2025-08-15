# reader — RabbitMQ consumer and SMTP email sender

## THIS PROJECT IS IN DEVELOPMENT

Reader consumes messages from RabbitMQ queues and sends transactional emails (confirm email, reset password) via SMTP. Designed to work alongside an auth service.

## Features
- RabbitMQ consumer with QoS(1) and manual ack/nack
- Two queues: `confirm_email`, `forgot_token`
- SMTP sending with 3 retry attempts and structured JSON logs
- Graceful shutdown on SIGINT/SIGTERM

## Tech stack
- Go 1.21+
- RabbitMQ (`github.com/rabbitmq/amqp091-go`)
- SMTP (`gopkg.in/gomail.v2`)
- Logging: `log/slog` + `lumberjack`

## Requirements
- Go 1.21+
- Running RabbitMQ (local or remote)
- Valid SMTP account (host, port, username, password)

## Configuration

All settings are read from `configs/config.env` (or environment variables):

```ini
# RabbitMQ
RABBIT_NAME=guest
RABBIT_HOST=localhost
RABBIT_PASSWORD=guest
RABBIT_PORT=5672
RABBIT_QUEUE_FORGOT_TOKEN=forgot_token
RABBIT_QUEUE_CONFIRM_EMAIL=confirm_email

# Logs
LOG_PATH=./cmd/app

# SMTP
EMAIL_PROVIDER=smtp
SMTP_HOST=smtp.example.com
SMTP_PORT=465
SMTP_USERNAME=user@example.com
SMTP_PASSWORD=********
SMTP_FROM=user@example.com

# Domain used to build reset links, etc.
MAIN_DOMAIN=https://example.com/password/reset
```

Security note: do not commit real secrets. Keep `.env` files out of VCS and use CI/CD secrets or host-level environment variables.

## Install & run via systemd

You can install and run the service via systemd using `systemd.sh`.

```bash
# 1) Copy the binary and configs to the server (example layout)
sudo mkdir -p /opt/reader
sudo cp ./bin/reader /opt/reader/reader          # your built binary
sudo cp -r ./configs /opt/reader/configs         # your env files

# 2) Create the unit with the helper script (it will ask a few questions)
chmod +x ./systemd.sh
sudo ./systemd.sh
# Provide:
#  - service name (e.g., reader)
#  - working directory (/opt/reader)
#  - exec start binary (/opt/reader/reader)
#  - user (system user to run the service)

# 3) The script does daemon-reload, enable and start
# Check status and tail logs
sudo systemctl status reader | cat
journalctl -u reader -n 200 -f
```

## Queues and payloads

- `confirm_email`
- `forgot_token`

Message payload (JSON):
```json
{
  "email": "user@example.com",
  "token": "opaque-token-or-code"
}
```

## Troubleshooting
- RabbitMQ connection errors: verify `RABBIT_*` and that the broker is reachable (e.g., `nc -z host 5672`)
- SMTP issues: verify `SMTP_*`, port, and TLS requirements of your provider
- Messages not consumed: ensure the publisher routes to `confirm_email` / `forgot_token`
- Logs: check `${LOG_PATH}/app.log` or `journalctl -u reader`

## License
MIT