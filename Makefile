server:
	go generate ./... && go run cmd/main.go

generate:
	go generate ./...

lint:
	golangci-lint run

start:
	docker-compose up -d && air

.PHONY:
	server swagger start
	