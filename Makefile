server:
	go run cmd/main.go

swagger:
	swag init -g cmd/main.go --output doc

.PHONY:
	server swagger