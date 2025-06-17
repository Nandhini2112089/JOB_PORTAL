# FROM golang:1.24

# WORKDIR /app

# COPY . .

# RUN go mod tidy
# RUN go build -o main .

# CMD ["./main"]


FROM golang:1.24-alpine AS builderAdd commentMore actions

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o main .



FROM alpine:latest
WORKDIR /root/


COPY --from=builder /app/main .
COPY ./config/config.yml .

EXPOSE 8080
CMD ["./main"]
