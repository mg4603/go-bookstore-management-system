package utils

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestSetJSONContentType(t *testing.T) {
	tests := []struct {
		name           string
		n              http.HandlerFunc
		expectedHeader string
		expectedStatus int
	}{
		{
			name: "Default JSON Content-Type",
			n: func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(http.StatusOK)
			},
			expectedHeader: "application/json",
			expectedStatus: http.StatusOK,
		},
		{
			name: "Overwrite existing Content-Type header",
			n: func(w http.ResponseWriter, r *http.Request) {
				w.Header().Set("Content-Type", "text/plain")
				w.WriteHeader(http.StatusOK)
			},
			expectedHeader: "text/plain",
			expectedStatus: http.StatusOK,
		},
		{
			name: "Ensure next handler is called",
			n: func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(http.StatusTeapot)
			},
			expectedHeader: "application/json",
			expectedStatus: http.StatusTeapot,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			wrappedHandler := SetJSONContentType(http.HandlerFunc(tt.n))

			req := httptest.NewRequest("GET", "/", nil)
			rec := httptest.NewRecorder()

			wrappedHandler.ServeHTTP(rec, req)

			if rec.Code != tt.expectedStatus {
				t.Errorf("Status code = %v; want %v", rec.Code, tt.expectedStatus)
			}

			if got := rec.Header().Get("Content-Type"); got != tt.expectedHeader {
				t.Errorf("Content-Type = %v, want %v", got, tt.expectedHeader)
			}
		})
	}
}
