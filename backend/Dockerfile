FROM golang:1.23.4

WORKDIR /app

RUN go install github.com/air-verse/air@latest

COPY . .
EXPOSE 3000
RUN go mod tidy
