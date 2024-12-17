package main

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHandleCalculate(t *testing.T) {

	testServer := httptest.NewServer(http.HandlerFunc(handleCalculate))
	defer testServer.Close()

	tests := []struct {
		name           string
		requestBody    string
		expectedStatus int
		expectedBody   string
	}{
		{
			name:           "Valid expression",
			requestBody:    `{"expression": "2+2*2"}`,
			expectedStatus: http.StatusOK,
			expectedBody:   `{"result":"6.00"}`,
		},
		{
			name:           "Invalid JSON",
			requestBody:    `{"expression": "2+2*2"`, // некорректный JSON
			expectedStatus: http.StatusBadRequest,
			expectedBody:   `{"error":"Invalid JSON format"}`,
		},
		{
			name:           "Invalid expression",
			requestBody:    `{"expression": "2+abc"}`,
			expectedStatus: http.StatusUnprocessableEntity,
			expectedBody:   `{"error":"Expression is not valid"}`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Создаём HTTP-запрос
			req, err := http.NewRequest("POST", testServer.URL+"/api/v1/calculate", bytes.NewBuffer([]byte(tt.requestBody)))
			if err != nil {
				t.Fatalf("could not create request: %v", err)
			}
			req.Header.Set("Content-Type", "application/json")

			// Отправляем запрос
			resp, err := http.DefaultClient.Do(req)
			if err != nil {
				t.Fatalf("could not send request: %v", err)
			}
			defer resp.Body.Close()

			// Проверка статуса
			if resp.StatusCode != tt.expectedStatus {
				t.Errorf("expected status %d, got %d", tt.expectedStatus, resp.StatusCode)
			}

			// Проверка тела ответа
			buf := new(bytes.Buffer)
			buf.ReadFrom(resp.Body)
			if buf.String() != tt.expectedBody {
				t.Errorf("expected body %q, got %q", tt.expectedBody, buf.String())
			}
		})
	}
}
