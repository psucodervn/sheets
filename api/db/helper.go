package db

import (
	"database/sql"
	"fmt"
	"os"
	"time"

	"github.com/go-pg/pg/v9"
	_ "github.com/lib/pq"
	"github.com/rs/zerolog/log"

	"api/internal/config"
)

func ConnectDB(cfg config.PostgresConfig) *sql.DB {
	dsn := fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=%s", cfg.User, cfg.Password, cfg.Host, cfg.Port, cfg.DB, cfg.SSLMode)
	conn, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Err(err).Msg("connect postgres failed")
		os.Exit(1)
	}
	if err := conn.Ping(); err != nil {
		log.Err(err).Msg("ping postgres failed")
		os.Exit(1)
	}
	return conn
}

func ConnectGoPGDB(cfg config.PostgresConfig) *pg.DB {
	db := pg.Connect(&pg.Options{
		Network:               "tcp",
		Addr:                  fmt.Sprintf("%s:%d", cfg.Host, cfg.Port),
		Dialer:                nil,
		User:                  cfg.User,
		Password:              cfg.Password,
		Database:              cfg.DB,
		ApplicationName:       "sheets",
		TLSConfig:             nil,
		DialTimeout:           10 * time.Second,
		ReadTimeout:           10 * time.Second,
		WriteTimeout:          10 * time.Second,
		OnConnect:             nil,
		MaxRetries:            0,
		RetryStatementTimeout: false,
		MinRetryBackoff:       0,
		MaxRetryBackoff:       0,
		PoolSize:              0,
		MinIdleConns:          0,
		MaxConnAge:            0,
		PoolTimeout:           0,
		IdleTimeout:           0,
		IdleCheckFrequency:    0,
	})
	return db
}
