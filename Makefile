server:
	swag init -v3.1 -g cmd/main.go --output doc && go run cmd/main.go

swagger:
	swag init -v3.1 -g cmd/main.go --output doc

lint:
	golangci-lint run

.PHONY:
	server swagger
