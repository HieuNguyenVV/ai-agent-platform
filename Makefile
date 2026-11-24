.PHONY: help swagger migrate-diff migrate-apply migrate-status migrate-hash migrate-lint migrate-validate migrate-docker

ATLAS ?= atlas
COMPOSE ?= docker compose
ENV ?= local
NAME ?=
SWAG ?= swag

help: 
	@echo "Usage: make <target>"
	@echo "Targets:"
	@echo "  swagger - Generate Swagger documentation"
	@echo "  migrate-diff - Generate a new migration (requires name=...)"
	@echo "  migrate-apply - Apply pending migrations locally"
	@echo "  migrate-status - Show applied / pending migrations"
	@echo "  migrate-hash - Update atlas.sum checksums"
	@echo "  migrate-lint - Lint migrations against current DB"
	@echo "  migrate-validate - Validate migration files and atlas.sum"
	@echo "  migrate-docker - Apply migrations with the migrator container"

swagger: ## Generate Swagger documentation
	${SWAG} init -g cmd/server/main.go -o docs

migrate-diff: ## Generate a new migration (requires name=...)
ifndef NAME
	$(error Usage: make migrate-diff name=<migration_name>)
endif
	$(ATLAS) migrate diff $(NAME) --env $(ENV)

migrate-apply: ## Apply pending migrations locally
	$(ATLAS) migrate apply --env $(ENV)

migrate-status: ## Show applied / pending migrations
	$(ATLAS) migrate status --env $(ENV)

migrate-hash: ## Update atlas.sum checksums
	$(ATLAS) migrate hash --env $(ENV)

migrate-lint: ## Lint migrations against current DB
	$(ATLAS) migrate lint --env $(ENV) --latest 1

migrate-validate: ## Validate migration files and atlas.sum
	$(ATLAS) migrate validate --env $(ENV)

migrate-docker: ## Apply migrations with the migrator container
	$(COMPOSE) up --build --abort-on-container-exit migrator