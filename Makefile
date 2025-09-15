run:
	go run cmd/main.go
build:
	go build -o app ./cmd/main.go
gdocs:
	swag init -g ./cmd/main.go