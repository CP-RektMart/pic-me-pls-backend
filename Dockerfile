FROM golang:1.23.5-alpine3.21

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . ./

RUN go install github.com/swaggo/swag/v2/cmd/swag@v2.0.0-rc3
RUN swag init -v3.1 -g cmd/main.go --output doc
RUN CGO_ENABLED=0 GOOS=linux go build -o app cmd/main.go

CMD ["./app"]
