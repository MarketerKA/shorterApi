// Бизнес-логика и генерация ссылок

package internal

import (
	"crypto/rand"
	"database/sql"
	"encoding/base64"
	"net/url"
	"url-shortener/internal/errors"
	"url-shortener/internal/storage"
	"url-shortener/internal/validator"
)

const chars = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789_"

// Service предоставляет бизнес-логику для работы с URL
type Service struct {
	storage storage.URLStorage
}

// NewService создает новый экземпляр сервиса
func NewService(storage storage.URLStorage) *Service {
	return &Service{storage: storage}
}

// CreateShortURL создает короткий URL для оригинального URL
// Если оригинальный URL уже существует, возвращает существующий короткий URL
func (s *Service) CreateShortURL(originalURL string) (string, error) {
	// Валидация URL
	if _, err := url.ParseRequestURI(originalURL); err != nil {
		return "", errors.NewValidationError("Некорректный URL")
	}

	// Проверяем, существует ли уже короткая ссылка для этого URL
	if shortURL, err := s.storage.GetExistingShortURL(originalURL); err == nil {
		return shortURL, nil
	}

	// Генерируем короткий URL
	shortURL, err := s.generateShortURL()
	if err != nil {
		return "", errors.NewDatabaseError(err)
	}

	// Сохраняем в базу данных
	if err := s.storage.SaveURL(originalURL, shortURL); err != nil {
		return "", errors.NewDatabaseError(err)
	}

	return shortURL, nil
}

// GetOriginalURL возвращает оригинальный URL по короткому URL
func (s *Service) GetOriginalURL(shortURL string) (string, error) {
	if err := validator.ValidateShortURL(shortURL); err != nil {
		return "", err
	}

	originalURL, err := s.storage.GetOriginalURL(shortURL)
	if err == sql.ErrNoRows {
		return "", errors.NewNotFoundError("Ссылка не найдена")
	}
	if err != nil {
		return "", errors.NewDatabaseError(err)
	}

	return originalURL, nil
}

// generateShortURL генерирует случайную строку для короткого URL
func (s *Service) generateShortURL() (string, error) {
	b := make([]byte, 6)
	_, err := rand.Read(b)
	if err != nil {
		return "", err
	}
	return base64.URLEncoding.EncodeToString(b)[:6], nil
}
