
FROM golang:1.21.0-alpine3.17 AS builder
WORKDIR /app
COPY . .
RUN go mod download
RUN go build -o main .
FROM alpine:3.18.3
WORKDIR /app
COPY --from=builder /app/main .
COPY .env .env
EXPOSE 8080
CMD ["./main"]
