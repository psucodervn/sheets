package database

import (
	"api/config"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/rs/zerolog/log"
	"os"
)

func NewPostgresGorm(cfg config.PostgresConfig) (*gorm.DB, error) {
	dsn := fmt.Sprintf(
		"host=%s port=%v user=%s dbname=%s password=%s sslmode=%s",
		cfg.Host, cfg.Port, cfg.User, cfg.DB, cfg.Password, cfg.SSLMode,
	)
	db, err := gorm.Open("postgres", dsn)
	if err != nil {
		return nil, err
	}
	if cfg.Debug {
		db = db.Debug()
	}
	return db, nil
}

func MustNewPostgresGorm(cfg config.PostgresConfig) *gorm.DB {
	db, err := NewPostgresGorm(cfg)
	if err != nil {
		log.Err(err).Msg("connect postgres failed")
		os.Exit(1)
	}
	return db
}
