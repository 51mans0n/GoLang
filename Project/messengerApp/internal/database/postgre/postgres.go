package database

import (
	"database/sql"
	"fmt"
	"messengerApp/cmd/config"

	_ "github.com/lib/pq"
)

// Database represents the database connection
type Database struct {
	db *sql.DB
}

// DB returns the underlying *sql.DB instance
func (d *Database) DB() *sql.DB {
	return d.db
}

func (d *Database) Login(username, password string) (string, error) {
	// Implement login logic here
	return "", nil
}

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

	fmt.Println("Connected to database")

	return &Database{db: db}, nil
}
