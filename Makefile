export CGO_ENABLED ?= 0

CGO_CPPFLAGS ?= ${CPPFLAGS}
export CGO_CPPFLAGS
CGO_CFLAGS ?= ${CFLAGS}
export CGO_CFLAGS
CGO_LDFLAGS ?= $(filter -g -L% -l% -O%,${LDFLAGS})
export CGO_LDFLAGS

DATE_FMT = +%Y-%m-%d
ifdef SOURCE_DATE_EPOCH
    BUILD_DATE ?= $(shell date -u -d "@$(SOURCE_DATE_EPOCH)" "$(DATE_FMT)" 2>/dev/null || date -u -r "$(SOURCE_DATE_EPOCH)" "$(DATE_FMT)" 2>/dev/null || date -u "$(DATE_FMT)")
else
    BUILD_DATE ?= $(shell date "$(DATE_FMT)")
endif

GO_LDFLAGS := -X github.com/profclems/compozify/internal/version.BuildDate=$(BUILD_DATE) $(GO_LDFLAGS)
ifdef COMPOZIFY_VERSION
	GO_LDFLAGS := $(GO_LDFLAGS) -X github.com/profclems/compozify/internal/version.Version=$(COMPOZIFY_VERSION)
endif

EXE =
ifeq ($(shell go env GOOS),windows)
EXE = .exe
endif

install:
	go install -v -trimpath -ldflags "$(GO_LDFLAGS)" ./cmd/compozify

.PHONY: build
build:
	go build -trimpath -ldflags "$(GO_LDFLAGS)" -o ./bin/compozify ./cmd/compozify

.PHONY: test
test:
	go test -race ./...

.PHONY: fix
fix: ## Fix lint violations
	golangci-lint run --fix
	gofmt -s -w .
	goimports -local "github.com/profclems/compozify" -w $$(find . -type f -name '*.go' -not -path "*/vendor/*")

.PHONY: completions
completions: build
	mkdir -p ./share/bash-completion/completions ./share/fish/vendor_completions.d ./share/zsh/site-functions
	bin/compozify$(EXE) completion bash > ./share/bash-completion/completions/compozify
	bin/compozify$(EXE) completion fish > ./share/fish/vendor_completions.d/compozify.fish
	bin/compozify$(EXE) completion zsh > ./share/zsh/site-functions/_compozify

.PHONY: manpage
manpage: ## Generate manual pages
	go run -trimpath -ldflags "$(GO_LDFLAGS)" ./cmd/gen-docs/main.go --manpage --path ./share/man/man1
