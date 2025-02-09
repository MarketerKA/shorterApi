package errors

import "fmt"

// ErrorType определяет тип ошибки
type ErrorType string

const (
	ValidationError ErrorType = "VALIDATION_ERROR"
	DatabaseError  ErrorType = "DATABASE_ERROR"
	NotFoundError  ErrorType = "NOT_FOUND_ERROR"
)

// AppError представляет ошибку в приложении
type AppError struct {
	Type    ErrorType
	Message string
	Err     error
}

func (e *AppError) Error() string {
	if e.Err != nil {
		return fmt.Sprintf("%s: %s (%v)", e.Type, e.Message, e.Err)
	}
	return fmt.Sprintf("%s: %s", e.Type, e.Message)
}

// NewValidationError создает новую ошибку валидации
func NewValidationError(message string) *AppError {
	return &AppError{
		Type:    ValidationError,
		Message: message,
	}
}

// NewDatabaseError создает новую ошибку базы данных
func NewDatabaseError(err error) *AppError {
	return &AppError{
		Type:    DatabaseError,
		Message: "Ошибка базы данных",
		Err:     err,
	}
}

// NewNotFoundError создает новую ошибку "не найдено"
func NewNotFoundError(message string) *AppError {
	return &AppError{
		Type:    NotFoundError,
		Message: message,
	}
} 