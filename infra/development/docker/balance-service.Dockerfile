FROM golang:1.25-alpine AS builder
WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY shared ./shared
COPY services/balance-service ./services/balance-service

RUN go build -o /balance-service ./services/balance-service/cmd

FROM alpine
WORKDIR /app
COPY --from=builder /balance-service .
ENTRYPOINT ["./balance-service"]