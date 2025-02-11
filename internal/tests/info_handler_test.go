package tests

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"url-shortener/internal"
)

func TestInfoHandler(t *testing.T) {
	tests := []struct {
		name         string
		shortURL     string
		setupMock    func(*internal.MockStorage)
		expectedCode int
	}{
		{
			name:     "Existing URL",
			shortURL: "abc123",
			setupMock: func(m *internal.MockStorage) {
				m.SaveURL("https://example.com", "abc123")
			},
			expectedCode: http.StatusOK,
		},
		{
			name:         "Non-existing URL",
			shortURL:     "xyz789",
			setupMock:    func(m *internal.MockStorage) {},
			expectedCode: http.StatusNotFound,
		},
		{
			name:         "Empty URL",
			shortURL:     "",
			setupMock:    func(m *internal.MockStorage) {},
			expectedCode: http.StatusBadRequest,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockStorage := internal.NewMockStorage()
			tt.setupMock(mockStorage)
			service := internal.NewService(mockStorage)
			handler := internal.NewHandler(service)

			req := httptest.NewRequest("GET", "/info/"+tt.shortURL, nil)
			rr := httptest.NewRecorder()

			handler.InfoHandler(rr, req)

			if rr.Code != tt.expectedCode {
				t.Errorf("handler returned wrong status code: got %v want %v",
					rr.Code, tt.expectedCode)
			}
		})
	}
}