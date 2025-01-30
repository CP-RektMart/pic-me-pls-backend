FROM golang:1.23.5-alpine3.21

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . ./

RUN CGO_ENABLED=0 GOOS=linux go build -o app cmd/main.go

CMD ["./app"]
