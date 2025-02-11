package tests

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"url-shortener/internal"
)

func TestCreateHandler(t *testing.T) {
	tests := []struct {
		name           string
		inputURL      string
		expectedCode  int
		expectedError bool
	}{
		{
			name:           "Valid URL",
			inputURL:      "https://example.com",
			expectedCode:  http.StatusOK,
			expectedError: false,
		},
		{
			name:           "Invalid URL",
			inputURL:      "not-a-url",
			expectedCode:  http.StatusBadRequest,
			expectedError: true,
		},
		{
			name:           "Empty URL",
			inputURL:      "",
			expectedCode:  http.StatusBadRequest,
			expectedError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockStorage := internal.NewMockStorage()
			service := internal.NewService(mockStorage)
			handler := internal.NewHandler(service)

			url := internal.URL{OriginalURL: tt.inputURL}
			body, _ := json.Marshal(url)
			req := httptest.NewRequest("POST", "/create", bytes.NewBuffer(body))
			req.Header.Set("Content-Type", "application/json")

			rr := httptest.NewRecorder()
			handler.CreateHandler(rr, req)

			if rr.Code != tt.expectedCode {
				t.Errorf("handler returned wrong status code: got %v want %v",
					rr.Code, tt.expectedCode)
			}

			if !tt.expectedError {
				var response internal.URL
				err := json.NewDecoder(rr.Body).Decode(&response)
				if err != nil {
					t.Fatalf("Could not decode response: %v", err)
				}
				if response.ShortURL == "" {
					t.Error("Expected short URL in response, got empty string")
				}
			}
		})
	}
}