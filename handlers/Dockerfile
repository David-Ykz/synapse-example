FROM golang:1.25-alpine AS builder

WORKDIR /app

RUN apk add --no-cache git

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o handler_server handlers/cmd/main.go

FROM alpine:latest
WORKDIR /root/
COPY --from=builder /app/handler_server .
CMD ["./handler_server"]