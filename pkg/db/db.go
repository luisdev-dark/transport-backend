package db

import (
	"context"
	"errors"
	"os"
	"sync"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

var (
	pool     *pgxpool.Pool
	poolOnce sync.Once
	poolErr  error
)

func GetPool(ctx context.Context) (*pgxpool.Pool, error) {
	poolOnce.Do(func() {
		databaseURL := os.Getenv("DATABASE_URL")
		if databaseURL == "" {
			poolErr = errors.New("DATABASE_URL environment variable is not set")
			return
		}

		config, err := pgxpool.ParseConfig(databaseURL)
		if err != nil {
			poolErr = err
			return
		}

		config.MaxConns = 5
		config.MinConns = 1
		config.MaxConnLifetime = 30 * time.Minute
		config.MaxConnIdleTime = 5 * time.Minute

		pingCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
		defer cancel()

		pool, err = pgxpool.NewWithConfig(pingCtx, config)
		if err != nil {
			poolErr = err
			return
		}

		if err := pool.Ping(pingCtx); err != nil {
			poolErr = err
			pool = nil
			return
		}
	})

	return pool, poolErr
}
