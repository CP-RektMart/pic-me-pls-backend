server:
	go run cmd/main.go

swagger:
	swag init -g cmd/main.go --output doc

lint:
	golangci-lint run

.PHONY:
	server swagger
