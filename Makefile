HASGOCILINT := $(shell which golangci-lint 2> /dev/null)


ifdef HASGOCILINT
    GOLINT=golangci-lint
else
    GOLINT=bin/golangci-lint
endif

# Dependency versions
GOLANGCI_VERSION = 1.53.3

install:
	go install -v github.com/profclems/compozify/cmd

build:
	go build -o ./bin/ ./cmd

test:
	go test -race ./...
fix: ## Fix lint violations
	gofmt -s -w .
	goimports -w $$(find . -type f -name '*.go' -not -path "*/vendor/*")

.PHONY: fix