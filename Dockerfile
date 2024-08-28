FROM golang:1.19

WORKDIR /rewild

COPY go.mod go.sum ./
RUN go mod download

COPY *.go ./
COPY *.json ./

RUN CGO_ENABLED=0 GOOS=linux go build -o /rewild-api

EXPOSE 8080

CMD ["/rewild-api"]