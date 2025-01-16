FROM golang:1.23.4-alpine3.21 AS builder

RUN apk add --no-cache git

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -ldflags="-s -w" -o myapp

FROM alpine:3.21

WORKDIR /app

RUN apk add --no-cache ca-certificates

COPY --from=builder /app/myapp .

COPY .env .

RUN chmod 644 .env

EXPOSE 1523

CMD ["./myapp"]
