.PHONY: default run build tests docs clean
# Variables
APP_NAME=mybooks-api

# Tasks
default: run

run: 
  @go run /cmd/api/main.go
build:
  @go build -o $(APP_NAME) /cmd/api/main.go
tests:
  @go test ./...
docs:
  @swag init
clean: 
  @rm -f $(APP_NAME)
  @rm -rf ./docs