# Stage 1: Build the Go binary (Updated to Go 1.26 to support modern Templ compilation)
FROM golang:1.26-alpine AS builder
WORKDIR /app

# Install the templ tool binary in the environment
RUN go install github.com/a-h/templ@latest

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN templ generate
RUN CGO_ENABLED=0 GOOS=linux go build -o main .

# Stage 2: Super small deployment alpine package 
FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/main .

EXPOSE 8080
CMD ["./main"]