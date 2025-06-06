FROM golang:1.24-alpine AS builder

WORKDIR /app

COPY chat/go.mod chat/go.sum ./
RUN go mod download
RUN go install github.com/pressly/goose/v3/cmd/goose@latest

COPY . .

# Собираем сервис
WORKDIR /app/chat/cmd/chat
RUN go build -o chat_service

FROM alpine:latest
WORKDIR /root/

COPY --from=builder /go/bin/goose /usr/local/bin/goose
COPY --from=builder /app/chat/cmd/chat/chat_service .
COPY --from=builder /app/chat/config/config.yaml ./config/config.yaml
COPY --from=builder /app/chat/db/migrations /root/db/migrations

EXPOSE 8080

CMD goose -dir /root/db/migrations postgres "postgresql://${APP_POSTGRES_USER}:${APP_POSTGRES_PASSWORD}@${APP_POSTGRES_ADDRESS}/${APP_POSTGRES_NAME}?sslmode=disable" up && ./chat_service

