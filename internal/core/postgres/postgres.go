package postgres

import (
	"context"
	"fmt"
	"rewardapp/internal/core/config"

	"github.com/jackc/pgx/v5"
)

func NewPostgresDB(cfg config.DataBaseConfig) (*pgx.Conn, error) {
	dsn := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=%s",
		cfg.User,
		cfg.Password,
		cfg.Host,
		cfg.Port,
		cfg.Name,
		cfg.SSLMode,
	)
	conn, err := pgx.Connect(context.Background(), dsn)
	if err != nil {
		return nil, fmt.Errorf("failed connect to database: %w", err)
	}

	if err := conn.Ping(context.Background()); err != nil {
		return nil, fmt.Errorf("failed ping to database: %w", err)
	}

	return conn, nil
}
