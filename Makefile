
.PHONY: default
default: dev

.PHONY: dev
dev:
	@go run main.go

.PHONY: env
env:
	@cp .env.example .env

# Update deps
.PHONY: update-deps
update-deps:
	@go get -u ./...
	@go mod tidy

# Help
.PHONY: help
help:
	@echo "Makefile for project"
	@echo ""
	@echo "Usage:"
	@echo "  make [target] [parameter=value]"
	@echo ""
	@echo "Targets:"
	@echo "  default           				: Run dev"
	@echo "  update-deps       				: Update all depedency"
	@echo "  env       						: Generated env"
	@echo "  help              				: Display this help information"