server:
	go generate ./... && go run cmd/main.go

generate:
	go generate ./...

lint:
	golangci-lint run

.PHONY:
	server swagger
