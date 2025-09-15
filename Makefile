run:
	go run cmd/main.go
build:
	go build -tags netgo -ldflags '-s -w' -o app
gdocs:
	swag init -g ./cmd/main.go