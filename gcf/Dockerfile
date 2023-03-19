FROM golang:1.20-alpine

WORKDIR /app

COPY go.mod go.sum ./
RUN go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.51.1
RUN go mod download

COPY .. .

RUN go build -o main ./cmd/updater

CMD ["./main"]
