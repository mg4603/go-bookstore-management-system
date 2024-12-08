package utils

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
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

type TestStruct struct {
	Name string `json:"name"`
}

type errorReader struct{}

func (e *errorReader) Read(p []byte) (n int, err error) {
	return 0, errors.New("failed to read request body")
}

func TestParseBody(t *testing.T) {
	tests := []struct {
		name            string
		body            io.Reader
		expectedErr     bool
		expectedMessage string
	}{
		{
			name:            "Successful body read and valid JSON",
			body:            bytes.NewBufferString(`{"name":"John"}`),
			expectedErr:     false,
			expectedMessage: "",
		},
		{
			name:            "Failed body read",
			body:            &errorReader{},
			expectedErr:     true,
			expectedMessage: "failed to read request body",
		},
		{
			name:            "Invalid json format",
			body:            bytes.NewBufferString(`{"name":"John",}`),
			expectedErr:     true,
			expectedMessage: "error parsing JSON",
		},
		{
			name:            "Empty body",
			body:            bytes.NewBufferString(``),
			expectedErr:     true,
			expectedMessage: "error parsing JSON",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := httptest.NewRequest("POST", "http://example.com", tt.body)
			assert.NotNil(t, req)

			rec := httptest.NewRecorder()

			var result TestStruct
			err := ParseBody(req, &result)

			if tt.expectedErr {
				assert.Error(t, err)
				assert.Contains(t, err.Error(), tt.expectedMessage)
			} else {
				assert.NoError(t, err)
			}

			assert.Equal(t, http.StatusOK, rec.Code)
		})
	}
}

func TestHandleError(t *testing.T) {
	tests := []struct {
		name           string
		statusCode     int
		errorMessage   string
		expectedStatus int
		expectedBody   string
	}{
		{
			name:           "Client error (400 Bad Request)",
			statusCode:     http.StatusBadRequest,
			errorMessage:   "Invalid Input data",
			expectedStatus: http.StatusBadRequest,
			expectedBody:   "An error occurred. Please try again later.",
		},
		{
			name:           "Server error (500 internal server error)",
			statusCode:     http.StatusInternalServerError,
			errorMessage:   "Database connection failed",
			expectedStatus: http.StatusInternalServerError,
			expectedBody:   "An error occurred. Please try again later.",
		},
		{
			name:           "Server error (500 internal server error)",
			statusCode:     http.StatusInternalServerError,
			errorMessage:   "Database connection failed",
			expectedStatus: http.StatusInternalServerError,
			expectedBody:   "An error occurred. Please try again later.",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			recorder := httptest.NewRecorder()
			HandleError(recorder, tc.statusCode, tc.errorMessage)
			assert.Equal(t, recorder.Code, tc.expectedStatus)

			var response ErrorResponse
			err := json.Unmarshal(recorder.Body.Bytes(), &response)
			assert.NoError(t, err)
			assert.Equal(t, tc.expectedBody, response.Message)
		})
	}

}
