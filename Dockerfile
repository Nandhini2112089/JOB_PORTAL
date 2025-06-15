# Build stage
FROM golang:1.24-alpine AS builder

WORKDIR /app
COPY . .

RUN go mod tidy
RUN go build -o myapp main.go

# Runtime stage
FROM alpine:latest

WORKDIR /root/
COPY --from=builder /app/myapp .

CMD ["./myapp"]
