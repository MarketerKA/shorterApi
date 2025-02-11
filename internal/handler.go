// HTTP обработчики

package internal

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"url-shortener/internal/errors"
)

type Handler struct {
	service *Service
}

type URL struct {
	OriginalURL string `json:"original_url,omitempty"`
	ShortURL    string `json:"short_url,omitempty"`
}

func NewHandler(service *Service) *Handler {
	return &Handler{service: service}
}

// CreateHandler обрабатывает POST-запросы для создания коротких URL
func (h *Handler) CreateHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusOK)
		return
	}

	if r.Method != http.MethodPost {
		http.Error(w, "Разрешены только POST запросы", http.StatusMethodNotAllowed)
		return
	}

	var url URL
	if err := json.NewDecoder(r.Body).Decode(&url); err != nil {
		h.sendError(w, errors.NewValidationError("Некорректный формат запроса"))
		return
	}

	shortURL, err := h.service.CreateShortURL(url.OriginalURL)
	if err != nil {
		h.handleError(w, err)
		return
	}

	url.ShortURL = shortURL
	h.sendJSON(w, url)
}

// GetHandler обрабатывает GET-запросы для получения оригинального URL по короткой ссылке
func (h *Handler) GetHandler(w http.ResponseWriter, r *http.Request) {
	// Добавляем CORS заголовки
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type")

	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusOK)
		return
	}

	if r.Method != http.MethodGet {
		http.Error(w, "Разрешены только GET запросы", http.StatusMethodNotAllowed)
		return
	}

	// Если это корневой путь, отправляем index.html
	if r.URL.Path == "/" {
		http.ServeFile(w, r, "static/test.html")
		return
	}

	// Игнорируем запросы к favicon.ico
	if r.URL.Path == "/favicon.ico" {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	shortURL := r.URL.Path[1:]
	if shortURL == "" {
		http.Error(w, "Короткая ссылка не указана", http.StatusBadRequest)
		return
	}

	originalURL, err := h.service.GetOriginalURL(shortURL)
	if err == sql.ErrNoRows {
		http.Error(w, "Ссылка не найдена", http.StatusNotFound)
		return
	}
	if err != nil {
		http.Error(w, "Ошибка базы данных", http.StatusInternalServerError)
		return
	}

	// Возвращаем редирект вместо JSON
	http.Redirect(w, r, originalURL, http.StatusFound)
}

// InfoHandler обрабатывает GET-запросы для получения информации о ссылке
func (h *Handler) InfoHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type")

	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusOK)
		return
	}

	if r.Method != http.MethodGet {
		http.Error(w, "Разрешены только GET запросы", http.StatusMethodNotAllowed)
		return
	}

	shortURL := r.URL.Path[len("/info/"):]
	if shortURL == "" {
		http.Error(w, "Короткая ссылка не указана", http.StatusBadRequest)
		return
	}

	originalURL, err := h.service.GetOriginalURL(shortURL)
	if err != nil {
		h.handleError(w, err)
		return
	}

	h.sendJSON(w, URL{
		OriginalURL: originalURL,
		ShortURL:    shortURL,
	})
}

func (h *Handler) handleError(w http.ResponseWriter, err error) {
	if appErr, ok := err.(*errors.AppError); ok {
		switch appErr.Type {
		case errors.ValidationError:
			h.sendError(w, appErr, http.StatusBadRequest)
		case errors.NotFoundError:
			h.sendError(w, appErr, http.StatusNotFound)
		case errors.DatabaseError:
			h.sendError(w, appErr, http.StatusInternalServerError)
		default:
			h.sendError(w, appErr, http.StatusInternalServerError)
		}
		return
	}
	h.sendError(w, errors.NewDatabaseError(err), http.StatusInternalServerError)
}

func (h *Handler) sendError(w http.ResponseWriter, err error, status ...int) {
	statusCode := http.StatusInternalServerError
	if len(status) > 0 {
		statusCode = status[0]
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(map[string]string{
		"error": err.Error(),
	})
}

func (h *Handler) sendJSON(w http.ResponseWriter, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}
