
FROM golang:alpine
WORKDIR /server
COPY ./ /server

RUN go mod download

ENTRYPOINT go run cmd/rewild-it/server.go