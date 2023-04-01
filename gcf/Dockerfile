FROM golang:1.20-alpine

WORKDIR /app

COPY go.mod go.sum ./
RUN go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.51.1
RUN go mod download

COPY . .

RUN go build -o notice-latest-program-version main.go

CMD ["./notice-latest-program-version"]
