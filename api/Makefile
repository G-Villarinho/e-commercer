GREEN=\033[0;32m
YELLOW=\033[0;33m
BLUE=\033[0;34m
NC=\033[0m

### BASE ###
.PHONY: help
help:
	@echo "Use: make [target]"
	@echo ""
	@echo "Tasks:"
	@egrep '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "  %-40s %s\n", $$1, $$2}'
 
.PHONY: setup
setup: ## Install all necessary dependencies
	@echo -e "$(BLUE)Installing dependencies...$(NC)"
	@echo -e "$(YELLOW)Installing gotestfmt...$(NC)"
	@go install github.com/gotesttools/gotestfmt/v2/cmd/gotestfmt@v2.5.0
	@echo -e "$(YELLOW)Installing golangci-lint...$(NC)"
	@go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.64.7
	@echo -e "$(YELLOW)Installing mockery...$(NC)"
	@go install github.com/vektra/mockery/v2@v2.53.2
	@echo -e "$(GREEN)Setup complete!$(NC)"

.PHONY: test
test: ## Run unit test
	@go install github.com/gotesttools/gotestfmt/v2/cmd/gotestfmt@v2.5.0
	@go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.60.7
	@go test -v -coverprofile=rawcover.out -v ./... -json 2>&1 | tee /tmp/gotest.log | gotestfmt -hide successful-tests,empty-packages

.PHONY: docker-run
docker-run: ## Run DB container
	@if docker compose up --build 2>/dev/null; then \
		: ; \
	else \
		echo "Falling back to Docker Compose V1"; \
		docker-compose up --build; \
	fi

.PHONY: docker-down
docker-down: ### Shutdown DB container
	@if docker compose down 2>/dev/null; then \
		: ; \
	else \
		echo "Falling back to Docker Compose V1"; \
		docker-compose down; \
	fi

.PHONY: lint
lint: ## Run lint
	@go install github.com/gotesttools/gotestfmt/v2/cmd/gotestfmt@v2.5.0
	@go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.60.3
	@golangci-lint run

.PHONY: mocks
mocks: ## Generate project mocks
	@mockery


.PHONY: migrations
migrations: ## Run migrations
	@go run migrations/main.go

.PHONY: run
run: ## Run project
	@go run main.go setup.go

### GERANDO PEM ###
.PHONY: generate-key
generate-key: ## Generate private and public key
	@openssl ecparam -name prime256v1 -genkey -noout -out ecdsa_private.pem
	@openssl ec -in ecdsa_private.pem -pubout -out ecdsa_public.pem