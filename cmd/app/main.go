package main

import (
	"context"
	"log"

	// Импортируем наши внутренние пакеты
	"rewardapp/internal/core/config"
	"rewardapp/internal/core/postgres"
)

func main() {
	// 1. Загружаем конфигурацию
	// (Предполагается, что у вас есть функция config.Load(),
	// которая читает .env или yaml и возвращает структуру)
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("Failed load config: %v", err)
	}

	// 2. Подключаемся к базе данных
	db, err := postgres.NewPostgresDB(cfg.DatabaseConfig)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer db.Close(context.Background())
	log.Println("Successfully connected to PostgreSQL")

}
