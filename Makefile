# UpCloud CLI Makefile

GO       = go
CLI      = upctl
MODULE   = $(shell env GO111MODULE=on $(GO) list -m)
DATE    ?= $(shell date +%FT%T%z)
VERSION ?= $(shell git describe --tags --always --dirty --match=v* 2> /dev/null || \
			cat $(CURDIR)/.version 2> /dev/null || echo v0)
PKGS     = $(or $(PKG),$(shell env GO111MODULE=on $(GO) list ./...))
TESTPKGS = $(shell env GO111MODULE=on $(GO) list -f \
			'{{ if or .TestGoFiles .XTestGoFiles }}{{ .ImportPath }}{{ end }}' \
			$(PKGS))

BIN_DIR         = $(CURDIR)/bin
BIN             = $(CLI)
BIN_LINUX       = $(BIN)-$(VERSION)-linux-amd64
BIN_DARWIN      = $(BIN)-$(VERSION)-darwin-amd64
BIN_WINDOWS     = $(BIN)-$(VERSION)-windows-amd64.exe


V = 0
Q = $(if $(filter 1,$V),,@)
M = $(shell printf "\033[34;1m▶\033[0m")

export GO111MODULE=on

.PHONY: build
build: fmt | $(BIN_DIR) ; $(info $(M) building executable for the current target…) @ ## Build program binary for current os/arch
	$Q $(GO) build \
		-tags release \
		-o $(BIN_DIR)/$(BIN) cmd/$(CLI)/main.go

.PHONY: build-all
build-all: build-linux build-darwin build-windows ## Build all targets

.PHONY: build-linux
build-linux: $(BIN_DIR) ; $(info $(M) building executable for Linux x86_64…) @ ## Build program binary for linux x86_64
	$Q GOOS=linux GOARCH=amd64 $(GO) build \
		-tags release \
		-o $(BIN_DIR)/$(BIN_LINUX) cmd/$(CLI)/main.go

.PHONY: build-darwin
build-darwin: $(BIN_DIR) ; $(info $(M) building executable for Darwin x86_64…) @ ## Build program binary for darwin x86_64
	$Q GOOS=darwin GOARCH=amd64 $(GO) build \
		-tags release \
		-o $(BIN_DIR)/$(BIN_DARWIN) cmd/$(CLI)/main.go

.PHONY: build-windows
build-windows: $(BIN_DIR) ; $(info $(M) building executable for Windows x86_64…) @ ## Build program binary for windows x86_64
	$Q GOOS=windows GOARCH=amd64 $(GO) build \
		-tags release \
		-o $(BIN_DIR)/$(BIN_WINDOWS) cmd/$(CLI)/main.go


# Tests

.PHONY: test
test: fmt; $(info $(M) running $(NAME:%=% )tests…) @ ## Run tests
	$Q $(GO) test $(TESTPKGS)

.PHONY: fmt
fmt: ; $(info $(M) running gofmt…) @ ## Run gofmt on all source files
	$Q $(GO) fmt $(PKGS)

# Misc

$(BIN_DIR):
	@mkdir -p $@

.PHONY: clean
clean: ; $(info $(M) cleaning…)	@ ## Cleanup everything
	@rm -rf $(BIN)

.PHONY: help
help:
	@grep -hE '^[ a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | \
		awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-17s\033[0m %s\n", $$1, $$2}'

.PHONY: version
version:
	@echo $(VERSION)
