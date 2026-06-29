package config

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

type DB struct {
	Pool *pgxpool.Pool
}

func ConnectDB(cfg *Config) (*DB, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	poolCfg, err := pgxpool.ParseConfig(cfg.DSN())
	if err != nil {
		return nil, fmt.Errorf("config: parsing database DSN: %w", err)
	}

	poolCfg.MaxConns = 10
	poolCfg.MinConns = 2
	poolCfg.MaxConnLifetime = time.Hour
	poolCfg.MaxConnIdleTime = 30 * time.Minute

	pool, err := pgxpool.NewWithConfig(ctx, poolCfg)
	if err != nil {
		return nil, fmt.Errorf("config: creating connection pool: %w", err)
	}

	if err := pool.Ping(ctx); err != nil {
		pool.Close()
		return nil, fmt.Errorf("config: pinging database: %w", err)
	}

	log.Printf("config: connected to database %q on %s:%s", cfg.DBName, cfg.DBHost, cfg.DBPort)
	return &DB{Pool: pool}, nil
}

func (db *DB) Close() {
	if db == nil || db.Pool == nil {
		return
	}
	db.Pool.Close()
}
