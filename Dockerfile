# Build stage
FROM golang:1.21-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o accelerometer-client-test .

# Final stage
FROM alpine:latest

WORKDIR /root/

COPY --from=builder /app/accelerometer-client-test .
COPY config.properties .

CMD ["./accelerometer-client-test"]
