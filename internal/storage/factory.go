package storage

import "fmt"

// NewStorage создает хранилище в зависимости от типа в конфигурации
func NewStorage(storageType string) (URLStorage, error) {
	switch storageType {
	case "postgres":
		return NewPostgresStorage()
	case "memory":
		return NewMemoryStorage(), nil
	default:
		return nil, fmt.Errorf("неизвестный тип хранилища: %s", storageType)
	}
} 