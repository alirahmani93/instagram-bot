FROM golang:1.20-alpine

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod tidy
RUN go mod download
COPY . .
RUN go build -o main ./cmd/server/main.go
CMD ["./main"]
