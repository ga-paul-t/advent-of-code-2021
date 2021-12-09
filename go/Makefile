BINDIR     := $(CURDIR)/bin
BINNAME    ?= aoc-2021

GOBIN = $(shell go env GOBIN)
ifeq ($(GOBIN),)
GOBIN = $(shell go env GOPATH)/bin
endif

.PHONY: all
all: clean lint test build

.PHONY: build
build: $(BINDIR)/$(BINNAME)

$(BINDIR)/$(BINNAME): $(SRC)
	go build -o '$(BINDIR)/$(BINNAME)' ./cmd/aoc2021

.PHONY: test
test:
	go test -race -vet=off -p 1 -covermode=atomic -coverprofile=coverage.out ./...

.PHONY: lint
lint:
	golangci-lint run --timeout 5m0s

.PHONY: clean
clean:
	@rm -rf '$(BINDIR)'