.DEFAULT_GOAL := help
include .env
export

export PROJECT_ROOT=$(shell pwd)
MIGRATE_URL = postgres://$(DATABASE_USER):$(DATABASE_PASSWORD)@postgres:5432/$(DATABASE_NAME)?sslmode=disable

help:
	@awk 'BEGIN {FS = ":.*## "} \
	     /^[a-zA-Z_-]+:.*## / {printf "  \033[36m%-20s\033[0m %s\n", $$1, $$2}' \
	     $(MAKEFILE_LIST)

run: build ## Запустить приложение
	@$(PROJECT_ROOT)/bin/main

build: ## Сборка проект
	@go build -o bin/main cmd/app/main.go

clean: ## Удалить артефакты сборки
	@read -p "Удалить артефакты сборки? Опасность утери данных. [y/N]: " ans; \
	if [ "$$ans" = "y" ]; then \
		rm -rf bin && \
		echo "Артефакты сборки удалены"; \
	else \
		echo "Удаление артефактов сборки отменена"; \
	fi

migrate-create: ## Создать миграцию локально: make migrate-create seq=init
	@if [ -z "$(seq)" ]; then \
		echo "Не указан параметр seq. Пример: make migrate-create seq=init"; \
		exit 1; \
	fi; \
	migrate create -ext sql -dir ./migrations -seq "$(seq)"

migrate-up: ## Применить одну миграцию
	docker compose run --rm migrate -path /migrations -database "$(MIGRATE_URL)" up 1

migrate-down: ## Откатить одну миграцию
	docker compose run --rm migrate -path /migrations -database "$(MIGRATE_URL)" down 1

postgres-up: ## Поднять БД
	docker compose up -d postgres

postgres-down: ## Остановить БД
	docker compose stop postgres

.PHONY: help run build clean migrate-create migrate-up postgres-up postgres-down