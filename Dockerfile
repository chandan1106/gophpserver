FROM golang:1.18-alpine AS builder

WORKDIR /app
COPY . .
RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux go build -o gophpserver ./cmd/gophpserver

FROM alpine:latest

RUN apk --no-cache add php

WORKDIR /app
COPY --from=builder /app/gophpserver /app/
COPY public /app/public

EXPOSE 8080

CMD ["./gophpserver", "-host", "0.0.0.0", "-port", "8080", "-docroot", "./public"]