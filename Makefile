APP_NAME=bin/app1

.PHONY: start
start:
	realize start --server

.PHONY: build
build:
	go build -o $(APP_NAME) cmd/app1/main.go
