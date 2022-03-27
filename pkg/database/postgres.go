package database

import (
	"fmt"
	"github.com/jmoiron/sqlx"
)

func NewPostgresDB(cfg *Config) (*sqlx.DB, error) {
	return sqlx.Connect("postgres", fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
		cfg.host, cfg.port, cfg.username, cfg.dbName, cfg.password, cfg.sslMode))
}
