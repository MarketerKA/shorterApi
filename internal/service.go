// Бизнес-логика и генерация ссылок

package internal

import (
	"crypto/rand"
	"math/big"
)

const chars = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789_"

// Service предоставляет бизнес-логику для работы с URL
type Service struct {
	storage *Storage
}

// NewService создает новый экземпляр сервиса
func NewService(storage *Storage) *Service {
	return &Service{storage: storage}
}

// CreateShortURL создает короткий URL для оригинального URL
// Если оригинальный URL уже существует, возвращает существующий короткий URL
func (s *Service) CreateShortURL(originalURL string) (string, error) {
	// Проверяем, существует ли уже такой URL
	if existingURL, err := s.storage.GetExistingShortURL(originalURL); err == nil {
		return existingURL, nil
	}

	shortURL := generateShortURL()
	err := s.storage.SaveURL(originalURL, shortURL)
	if err != nil {
		return "", err
	}

	return shortURL, nil
}

// GetOriginalURL возвращает оригинальный URL по короткому URL
func (s *Service) GetOriginalURL(shortURL string) (string, error) {
	return s.storage.GetOriginalURL(shortURL)
}

// generateShortURL генерирует случайную строку для короткого URL
func generateShortURL() string {
	result := make([]byte, 10)
	for i := 0; i < 10; i++ {
		n, _ := rand.Int(rand.Reader, big.NewInt(int64(len(chars))))
		result[i] = chars[n.Int64()]
	}
	return string(result)
}
