ifneq (,$(wildcard ./.env))
    include .env
    export
endif

.PHONY: generate
generate:
	@echo running all go:generate
	@go generate ./...


.PHONY: configure
configure:
	@echo installing dependency...
	@go mod download
	@echo "creating .env file from example (will not overwrite)"
	@cp -n .env.example .env

.PHONY: run-local
run-local:
	@go run ./cmd/oapi-server

.PHONY: lint
lint:
	golangci-lint run -v

.PHONY: build
build: generate
	@echo compile the app
	@CGO_ENABLED=0 go build -v \
		-o oapi-server ./cmd/oapi-server
	@echo compiled to oapi-server

.PHONY: test
test: generate
	@echo run all go test
	@go test -coverprofile cover.out ./...

.PHONY: cover
cover: test
	@go tool cover -html cover.out