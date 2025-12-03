FROM golang:1.25-alpine AS builder
WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY shared ./shared
COPY services/api-gateway ./services/api-gateway

RUN go build -o /api-gateway ./services/api-gateway

FROM alpine
WORKDIR /app
COPY --from=builder /api-gateway .
ENTRYPOINT ["./api-gateway"]