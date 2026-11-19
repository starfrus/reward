package main

import (
	"context"
	"log"

	"rewardapp/internal/core/config"
	"rewardapp/internal/core/postgres"
)

func main() {

	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("Failed load config: %v", err)
	}

	db, err := postgres.NewPostgresDB(cfg.DataBaseConfig)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer db.Close(context.Background())
	log.Println("Successfully connected to PostgreSQL")

}
