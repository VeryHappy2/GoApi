FROM golang:1.23.3-alpine AS build
COPY . /src
WORKDIR /src
COPY go.mod go.sum ./
RUN go mod download
RUN go build -o cmd/main cmd/main.go
CMD ["./cmd/main"]