package validator

import (
	"net/url"
	"strings"
	"url-shortener/internal/errors"
)

// ValidateURL проверяет корректность URL
func ValidateURL(urlStr string) error {
	// Проверка на пустую строку
	if strings.TrimSpace(urlStr) == "" {
		return errors.NewValidationError("URL не может быть пустым")
	}

	// Проверка на корректность URL
	u, err := url.Parse(urlStr)
	if err != nil {
		return errors.NewValidationError("Некорректный формат URL")
	}

	// Проверка наличия схемы и хоста
	if u.Scheme == "" || u.Host == "" {
		return errors.NewValidationError("URL должен содержать протокол и домен")
	}

	// Проверка протокола
	if u.Scheme != "http" && u.Scheme != "https" {
		return errors.NewValidationError("URL должен использовать протокол HTTP или HTTPS")
	}

	return nil
}

// ValidateShortURL проверяет корректность короткого URL
func ValidateShortURL(shortURL string) error {
	if strings.TrimSpace(shortURL) == "" {
		return errors.NewValidationError("Короткий URL не может быть пустым")
	}

	if len(shortURL) > 50 {
		return errors.NewValidationError("Короткий URL слишком длинный")
	}

	return nil
} 