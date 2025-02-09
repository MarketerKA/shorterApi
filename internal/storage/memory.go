package storage

import (
	"database/sql"
	"sync"
)

// MemoryStorage реализует хранение URL в памяти
type MemoryStorage struct {
	urls     map[string]string // shortURL -> originalURL
	origURLs map[string]string // originalURL -> shortURL
	mu       sync.RWMutex
}

// NewMemoryStorage создает новое in-memory хранилище
func NewMemoryStorage() *MemoryStorage {
	return &MemoryStorage{
		urls:     make(map[string]string),
		origURLs: make(map[string]string),
	}
}

func (m *MemoryStorage) CreateTable() error {
	// Для in-memory хранилища ничего создавать не нужно
	return nil
}

func (m *MemoryStorage) SaveURL(originalURL, shortURL string) error {
	m.mu.Lock()
	defer m.mu.Unlock()

	m.urls[shortURL] = originalURL
	m.origURLs[originalURL] = shortURL
	return nil
}

func (m *MemoryStorage) GetOriginalURL(shortURL string) (string, error) {
	m.mu.RLock()
	defer m.mu.RUnlock()

	if originalURL, ok := m.urls[shortURL]; ok {
		return originalURL, nil
	}
	return "", sql.ErrNoRows
}

func (m *MemoryStorage) GetExistingShortURL(originalURL string) (string, error) {
	m.mu.RLock()
	defer m.mu.RUnlock()

	if shortURL, ok := m.origURLs[originalURL]; ok {
		return shortURL, nil
	}
	return "", sql.ErrNoRows
}

func (m *MemoryStorage) Close() error {
	// Для in-memory хранилища ничего закрывать не нужно
	return nil
} 