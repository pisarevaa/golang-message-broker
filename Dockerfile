FROM golang:1.24.2-alpine AS builder

WORKDIR /opt/app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 go build -o bin/messagebroker ./cmd

FROM alpine:3.21
RUN addgroup -S messagebroker && adduser -S messagebroker -G messagebroker
USER messagebroker
WORKDIR /app
COPY .env.example .env
COPY --from=builder /opt/app/bin/messagebroker /usr/bin/messagebroker

USER root
# CMD ["messagebroker"]
