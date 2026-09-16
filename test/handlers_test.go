package test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/YOUR_USERNAME/go-mini/src/handlers"
)

func TestHealthHandler(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/health", nil)
	w := httptest.NewRecorder()

	handlers.HealthHandler(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("expected status 200, got %d", w.Code)
	}

	expected := `{"message":"I'm alive!","status":200,"data":null}` + "\n"
	if w.Body.String() != expected {
		t.Errorf("unexpected body: got %q, want %q", w.Body.String(), expected)
	}
}

func TestHomeHandler(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	w := httptest.NewRecorder()

	handlers.HomeHandler(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("expected status 200, got %d", w.Code)
	}

	expected := "Hello, World!\n"
	if w.Body.String() != expected {
		t.Errorf("unexpected body: got %q, want %q", w.Body.String(), expected)
	}
}
