FROM golang:1.19

WORKDIR /rewild

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /rewild-api ./cmd/server.go

EXPOSE 8080

ENTRYPOINT ["/rewild-api"]