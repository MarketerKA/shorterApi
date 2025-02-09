// HTTP обработчики

package internal

import (
	"database/sql"
	"encoding/json"
	"net/http"
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
		http.Error(w, "Ошибка чтения запроса", http.StatusBadRequest)
		return
	}

	shortURL, err := h.service.CreateShortURL(url.OriginalURL)
	if err != nil {
		http.Error(w, "Ошибка сохранения в базу данных", http.StatusInternalServerError)
		return
	}

	url.ShortURL = shortURL
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(url)
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

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(URL{
		OriginalURL: originalURL,
	})
}
