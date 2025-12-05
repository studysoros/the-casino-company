FROM golang:1.25-alpine AS builder
WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY shared ./shared
COPY services/betting-service ./services/betting-service

RUN go build -o /betting-service ./services/betting-service/cmd

FROM alpine
WORKDIR /app
COPY --from=builder /betting-service .
ENTRYPOINT ["./betting-service"]