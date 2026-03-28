FROM golang:1.25-alpine AS builder

WORKDIR /app

RUN apk add --no-cache git

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o example_app example/main.go

FROM alpine:latest
WORKDIR /root/
COPY --from=builder /app/example_app .
CMD ["./example_app"]