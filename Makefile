install:
	go install -v ./cmd/...

build:
	go build -o ./bin/ ./cmd/...

test:
	go test -race ./...
fix: ## Fix lint violations
	golangci-lint run --fix
	gofmt -s -w .
	goimports -local "github.com/profclems/compozify" -w $$(find . -type f -name '*.go' -not -path "*/vendor/*")

.PHONY: fix