# enoch-luminaries — containerized toolchain (no host Go/Node required).
# All builds/tests run in throwaway containers.

GO_IMAGE := golang:1.22
MODULE   := github.com/enoch/luminaries
# Mount the parent so the collector can read ../enoch/.local manifest at runtime.
BTC_ROOT := $(abspath ..)

.PHONY: help
help: ## Show this help
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | \
		awk 'BEGIN{FS=":.*?## "}{printf "  \033[36m%-14s\033[0m %s\n", $$1, $$2}'

.PHONY: build
build: ## Compile the collector in a container
	docker run --rm -v $(CURDIR):/src -w /src $(GO_IMAGE) go build ./...

.PHONY: vet
vet: ## go vet in a container
	docker run --rm -v $(CURDIR):/src -w /src $(GO_IMAGE) go vet ./...

.PHONY: test
test: ## Run tests in a container
	docker run --rm -v $(CURDIR):/src -w /src $(GO_IMAGE) go test ./...

.PHONY: fmt
fmt: ## gofmt the tree in a container
	docker run --rm -v $(CURDIR):/src -w /src $(GO_IMAGE) gofmt -w .

.PHONY: collector
collector: ## Run the collector against a local regtest federation
	docker run --rm -p 8090:8090 \
		-v $(CURDIR):/src -v $(BTC_ROOT)/enoch/.local:/enoch/.local:ro \
		-w /src $(GO_IMAGE) \
		go run ./cmd/collector --manifest=/enoch/.local/federation/federation_manifest.json --listen=:8090

.PHONY: orrery
orrery: ## Serve the Orrery web UI (placeholder until web/ is scaffolded)
	@echo "web/ not yet scaffolded — see docs/roadmap.md Phase 1"
