// Работа с базой данных

package internal

import (
	"database/sql"
	"fmt"
	"os"
)

type Storage struct {
	db *sql.DB
}

func NewStorage() (*Storage, error) {
	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_SSLMODE"))

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, fmt.Errorf("ошибка подключения к БД: %w", err)
	}

	return &Storage{db: db}, nil
}

func (s *Storage) CreateTable() error {
	query := `
        CREATE TABLE IF NOT EXISTS urls (
            id SERIAL PRIMARY KEY,
            original_url TEXT NOT NULL,
            short_url TEXT NOT NULL UNIQUE,
            created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
        )
    `
	_, err := s.db.Exec(query)
	return err
}

func (s *Storage) SaveURL(originalURL, shortURL string) error {
	_, err := s.db.Exec("INSERT INTO urls (original_url, short_url) VALUES ($1, $2)",
		originalURL, shortURL)
	return err
}

func (s *Storage) GetOriginalURL(shortURL string) (string, error) {
	var originalURL string
	err := s.db.QueryRow("SELECT original_url FROM urls WHERE short_url = $1", shortURL).Scan(&originalURL)
	return originalURL, err
}

func (s *Storage) GetExistingShortURL(originalURL string) (string, error) {
	var shortURL string
	err := s.db.QueryRow("SELECT short_url FROM urls WHERE original_url = $1", originalURL).Scan(&shortURL)
	return shortURL, err
}

func (s *Storage) Close() error {
	return s.db.Close()
}
