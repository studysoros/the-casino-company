FROM golang:1.25-alpine AS builder
WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY shared ./shared
COPY services/cashier-service ./services/cashier-service

RUN go build -o /cashier-service ./services/cashier-service/cmd

FROM alpine
WORKDIR /app
COPY --from=builder /cashier-service .
ENTRYPOINT ["./cashier-service"]