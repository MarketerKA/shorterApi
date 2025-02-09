package storage

// URLStorage определяет интерфейс для хранения URL
type URLStorage interface {
	// CreateTable инициализирует хранилище
	CreateTable() error
	// SaveURL сохраняет пару оригинальный/короткий URL
	SaveURL(originalURL, shortURL string) error
	// GetOriginalURL получает оригинальный URL по короткому
	GetOriginalURL(shortURL string) (string, error)
	// GetExistingShortURL проверяет существование оригинального URL
	GetExistingShortURL(originalURL string) (string, error)
	// Close освобождает ресурсы хранилища
	Close() error
} 