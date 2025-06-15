# FROM golang:1.24

# WORKDIR /app

# COPY . .

# RUN go mod tidy
# RUN go build -o main .

# CMD ["./main"]

# -------- build stage --------
FROM golang:1.24-alpine AS builder

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o main ./cmd/server

# -------- runtime stage ------
FROM alpine:latest
WORKDIR /root/

# copy binary + config.yml
COPY --from=builder /app/main .
COPY config.yml .

EXPOSE 8080
CMD ["./main"]
