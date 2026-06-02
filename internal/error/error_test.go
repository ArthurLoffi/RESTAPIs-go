package error

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
)

func init() {
	gin.SetMode(gin.TestMode)
}

func TestNew(t *testing.T) {
	tests := []struct {
		name string
		status int
		message string
	}{
		{"bad request", http.StatusBadRequest, "invalid request body"},
		{"not found", http.StatusNotFound, "User not found"},
		{"internal error", http.StatusInternalServerError, "something went wrong"},
		{"unauthorized", http.StatusUnauthorized, "Token not found"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := New(tt.status, tt.message)
			if err == nil {
				t.Fatal("expected non-nil AppError")
			}
			if err.Status != tt.status {
				t.Errorf("Status = %d, want %d", err.Status, tt.status)
			}
			if err.Message != tt.message {
				t.Errorf("Message = %q, want %q", err.Message, tt.message)
			}
		})
	}
}

func TestAppError_Error(t *testing.T) {
	err := New(http.StatusNotFound, "User not found")
	if err.Error() != "User not found" {
		t.Errorf("Error() = %q, want %q", err.Error(), "User not found")
	}
}

func TestRespond(t *testing.T) {
	tests := []struct {
		name string
		status int
		message string
		expectedStatus int
	}{
		{"400 bad request", http.StatusBadRequest, "invalid body", http.StatusBadRequest},
		{"404 not found", http.StatusNotFound, "not found", http.StatusNotFound},
		{"500 internal error", http.StatusInternalServerError, "server error", http.StatusInternalServerError},
		{"401 unauthorized", http.StatusUnauthorized, "unauthorized", http.StatusUnauthorized},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := httptest.NewRecorder()
			c, _ := gin.CreateTestContext(w)

			appErr := New(tt.status, tt.message)
			Respond(c, appErr)

			if w.Code != tt.expectedStatus {
				t.Errorf("status code = %d, want %d", w.Code, tt.expectedStatus)
			}

			body := w.Body.String()
			if body == "" {
				t.Error("expected non-empty response body")
			}
		})
	}
}