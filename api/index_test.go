package handler

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHandler(t *testing.T) {
	for _, path := range []string{"/", "/static/main.mjs", "/config/config.json.template"} {
		t.Run(path, func(t *testing.T) {
			response := httptest.NewRecorder()
			Handler(response, httptest.NewRequest(http.MethodGet, path, nil))
			if response.Code != http.StatusOK {
				t.Fatalf("GET %s: got status %d, want %d", path, response.Code, http.StatusOK)
			}
		})
	}
}
