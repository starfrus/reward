package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"rewardapp/internal/core/config"
	"rewardapp/internal/core/postgres"
	"rewardapp/internal/feature/users/repo"
	"rewardapp/internal/feature/users/service"
	"rewardapp/internal/feature/users/transport"
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

	// 3. Инициализируем слои (Dependency Injection)
	userRepo := repo.NewUserRepository(db)
	userService := service.NewUserService(userRepo)
	userHandler := transport.NewUserHandler(userService)

	// 4. Настраиваем роутер
	mux := http.NewServeMux()

	// Регистрируем эндпоинт для создания пользователя
	// Обратите внимание: метод POST
	mux.HandleFunc("POST /users", userHandler.Create)
	mux.HandleFunc("GET /users/{id}", userHandler.GetUser)

	// 5. Настраиваем HTTP-сервер
	server := &http.Server{
		Addr:    ":8080", // или cfg.HTTP.Port
		Handler: mux,
	}

	// 6. Запускаем сервер в отдельной горутине (чтобы он не блокировал поток)
	go func() {
		log.Printf("Starting server on %s", server.Addr)
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Server error: %v", err)
		}
	}()

	// 7. Graceful Shutdown (Плавная остановка)
	// Ждем сигнал SIGINT или SIGTERM (Ctrl+C)
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutting down server...")

	// Даем серверу 5 секунд, чтобы завершить текущие запросы
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		log.Fatalf("Server forced to shutdown: %v", err)
	}

	log.Println("Server exited properly")

}
