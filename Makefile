ifneq (,$(wildcard ./.env))
    include .env
    export
endif

__API_GEN_OUTPUT = "./internal/api/api.gen.go"
__API_GEN_CFG = "./api/generate-server.config.yaml"
__API_GEN_MODEL_OUTPUT = "./internal/entity/model.gen.go"
__API_GEN_MODEL_CFG = "./api/generate-model.config.yaml"
__API_GEN_DEFINITION = "./api/api.yaml"

.PHONY: generate
generate: ${__API_GEN_MODEL_OUTPUT} ${__API_GEN_OUTPUT}

${__API_GEN_OUTPUT}:
	go run github.com/oapi-codegen/oapi-codegen/v2/cmd/oapi-codegen -config ${__API_GEN_CFG} -o ${__API_GEN_OUTPUT} ${__API_GEN_DEFINITION}

${__API_GEN_MODEL_OUTPUT}:
	go run github.com/oapi-codegen/oapi-codegen/v2/cmd/oapi-codegen -config ${__API_GEN_MODEL_CFG} -o ${__API_GEN_MODEL_OUTPUT} ${__API_GEN_DEFINITION}

.PHONY: configure
configure:
	@echo installing dependency...
	@go mod download
	@echo install go mock
	@go install go.uber.org/mock/mockgen@latest
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

.PHONY: __go-gen
__go-gen:
	go generate ./...

.PHONY: test
test: generate __go-gen
	@echo run all go test
	@go test -coverprofile cover.out ./...

.PHONY: cover
cover: test
	@go tool cover -html cover.out