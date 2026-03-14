BIN := bin/action

.PHONY: build run test clean help

build: ## Build the action binary
	go build -o $(BIN) ./cmd/action

run: build ## Build and run the action locally
	GITHUB_TOKEN=$(GITHUB_TOKEN) WORKING_DIRECTORY=$(or $(WORKING_DIRECTORY),.) ./$(BIN)

test: ## Run unit tests
	go test ./...

clean: ## Remove build artifacts
	rm -rf bin/

help: ## Show this help
	@grep -E '^[a-zA-Z_-]+:.*##' $(MAKEFILE_LIST) | awk 'BEGIN {FS=":.*##"}; {printf "  %-12s %s\n", $$1, $$2}'
