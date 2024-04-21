##### Commands Start #####
help: ## Shows this help
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z_\-\.]+:.*?## / {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}' $(MAKEFILE_LIST)

run: ## Run the application
	go run src/cmd/main.go

lint: ## Run linter
	golangci-lint run -v  ./... --timeout=2m
