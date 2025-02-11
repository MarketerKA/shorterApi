package tests

import (
	"database/sql"
	"testing"
	"url-shortener/internal"
)

func TestMockStorage(t *testing.T) {
	storage := internal.NewMockStorage()

	t.Run("Save and Get URL", func(t *testing.T) {
		originalURL := "https://example.com"
		shortURL := "abc123"

		err := storage.SaveURL(originalURL, shortURL)
		if err != nil {
			t.Errorf("SaveURL failed: %v", err)
		}

		got, err := storage.GetOriginalURL(shortURL)
		if err != nil {
			t.Errorf("GetOriginalURL failed: %v", err)
		}
		if got != originalURL {
			t.Errorf("GetOriginalURL = %v, want %v", got, originalURL)
		}
	})

	t.Run("Get Non-existing URL", func(t *testing.T) {
		_, err := storage.GetOriginalURL("nonexistent")
		if err != sql.ErrNoRows {
			t.Errorf("Expected sql.ErrNoRows, got %v", err)
		}
	})

	t.Run("Get Existing Short URL", func(t *testing.T) {
		originalURL := "https://example.com"
		shortURL := "abc123"

		storage.SaveURL(originalURL, shortURL)

		got, err := storage.GetExistingShortURL(originalURL)
		if err != nil {
			t.Errorf("GetExistingShortURL failed: %v", err)
		}
		if got != shortURL {
			t.Errorf("GetExistingShortURL = %v, want %v", got, shortURL)
		}
	})
}