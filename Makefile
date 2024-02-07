export PATH := $(shell go env GOPATH)/bin:$(PATH)

BUILD_DIR = app/cmd/build
BINARY_APP = binary_app
BUILD_SRC = app/cmd/main.go

build:
	go build -o ./$(BUILD_DIR)/$(BINARY_APP) $(BUILD_SRC)

run: swag-init build
	./$(BUILD_DIR)/$(BINARY_APP)

swag-init:
	swag init -g ./app/cmd/main.go -o ./app/docks

migrate:
	migrate create -ext sql -dir migrations -seq schema

migrate-up:
	migrate -path migrations -database "postgres://postgres:asyl12345.@localhost:5432/postgres?sslmode=disable" up

migrate-down:
	migrate -path migrations -database "postgres://postgres:asyl12345.@localhost:5432/postgres?sslmode=disable" down

migrate-dirty:
	migrate -path migrations -database "postgres://postgres:asyl12345.@localhost:5432/postgres?sslmode=disable" force 1
