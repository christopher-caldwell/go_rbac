install:
	@go mod tidy

gen-openapi:
	@go run github.com/ogen-go/ogen/cmd/ogen@latest \
		--target ./internal/server \
		-package server \
		--config ./openapi/.ogen.yml \
		--clean \
		./openapi/schema.yml

build: gen-openapi
	@go build -o bin/main cmd/api/main.go

run: build
	@./bin/main

run-dev:
	@air