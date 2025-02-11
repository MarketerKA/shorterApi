package internal

import "database/sql"

// MockStorage реализует интерфейс Storage для тестирования
type MockStorage struct {
    urls map[string]string
}

func NewMockStorage() *MockStorage {
    return &MockStorage{
        urls: make(map[string]string),
    }
}

func (m *MockStorage) CreateTable() error {
    return nil
}

func (m *MockStorage) SaveURL(originalURL, shortURL string) error {
    m.urls[shortURL] = originalURL
    return nil
}

func (m *MockStorage) GetOriginalURL(shortURL string) (string, error) {
    if shortURL == "" {
        return "", sql.ErrNoRows
    }
    if url, ok := m.urls[shortURL]; ok {
        return url, nil
    }
    return "", sql.ErrNoRows
}

func (m *MockStorage) GetExistingShortURL(originalURL string) (string, error) {
    for short, orig := range m.urls {
        if orig == originalURL {
            return short, nil
        }
    }
    return "", sql.ErrNoRows
}

func (m *MockStorage) Close() error {
    return nil
}