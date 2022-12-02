PROJECT_NAME := "gitconfswap"
PKG := "github.com/einzigartigerName/$(PROJECT_NAME)"
PKG_LIST := $(shell go list ${PKG}/... | grep -v /vendor/)
GO_FILES := $(shell find . -name '*.go' | grep -v /vendor/ | grep -v _test.go)
PINFO :=

.PHONY: all dep lint vet build clean

all: build

dep: ## Get the dependencies
	@go mod download

lint: ## Lint Golang files
	@golangci-lint -c .golangci.yml run

build: dep ## Build the binary file
	@go build -o out/gitconfswap $(PKG)/cmd/gitconfswap

clean: ## Remove previous build
	@rm -f $(PROJECT_NAME)/build

help: ## Display this help screen
	@grep -h -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'
