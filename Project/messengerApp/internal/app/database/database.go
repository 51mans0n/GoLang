package database

import (
	"database/sql"
	"fmt"
	"messengerApp/cmd/config"
)

// Database обертывает *sql.DB для удобства работы с БД
type Database struct {
	db *sql.DB
}

// InitDatabase инициализирует подключение к БД и возвращает *Database
func InitDatabase(cfg *config.DatabaseConfig) (*Database, error) {
	dbURI := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		cfg.Host, cfg.Port, cfg.User, cfg.Password, cfg.DBName)

	db, err := sql.Open("postgres", dbURI)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return &Database{db: db}, nil
}

// DB возвращает внутренний *sql.DB из Database
func (d *Database) DB() *sql.DB {
	return d.db
}
