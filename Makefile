.DEFAULT_GOAL = help

help:
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) \
		| sort \
		| awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-20s\033[0m %s\n", $$1, $$2}'

run: build ## Запустить приложение
	@./bin/main
build: ## Сборать проект
	@go build -o bin/main cmd/app/main.go
clean: ## Удалить артефакты сборки
	rm -rf bin
fmt: ## Оформатировать файлы .go

lint: ## Запустить линтер

test: ## Запустить тесты

