include .env
export

export PROJECT_ROOT=$(shell pwd)

run: build ## Запустить приложение
	@./bin/main

build: ## Сборать проект
	@go build -o bin/main cmd/app/main.go

clean: ## Удалить артефакты сборки
	@read -p "Удалить артефакты сборки? Опасность утери данных. [y/N]: " ans; \
	if [ "$$ans" = "y" ]; then \
		rm -rf bin out/pgdata && \
		echo "Артефакты сборки удалены"; \
	else \
		echo "Удаление артефактов сборки отменена"; \
	fi

migrate-create: ## Создать миграцию
	@if [ -z "$(seq)" ]; then \
		echo "Не указан параметр seq. Пример: make migrate-create seq=init"; \
		exit 1; \
	fi; \
	docker compose run migrate-postgres
	migrate create \
	-ext sql \
	-dir ./migrations \
	-seq "$(seq)"

migrate-up: ## Поднять миграции все
	docker compose up -d migrate

db-up:
	docker compose up -d postgres

db-down:
	docker compose down postgres

db-logs:
	docker compose logs -f

migrate-create-local: ## Создать миграцию локально
	@if [ -z "$(seq)" ]; then \
		echo "Не указан параметр seq. Пример: make migrate-create-local seq=init"; \
		exit 1; \
	fi; \
	migrate create \
	-ext sql \
	-dir ./migrations \
	-seq "$(seq)"



	

