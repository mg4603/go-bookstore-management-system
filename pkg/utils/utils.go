package utils

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

func SetJSONContentType(n http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		n.ServeHTTP(w, r)
	})
}

func ParseBody(r *http.Request, x interface{}) error {
	body, err := io.ReadAll(r.Body)
	defer r.Body.Close()

	if err != nil {
		return fmt.Errorf("failed to read request body: %w", err)
	}

	if err := json.Unmarshal([]byte(body), x); err != nil {
		return fmt.Errorf("error parsing JSON: %w", err)
	}
	return nil
}

type ErrorResponse struct {
	Message string `json:"message"`
}

func HandleError(w http.ResponseWriter, statusCode int, message string) {
	standardMessage := "An error occurred. Please try again later."

	if statusCode >= 500 {
		log.Printf("Internal Server Error: %s", message)
	} else if statusCode == 404 {
		log.Printf("404 not found: %s", message)
	} else if statusCode == 403 {
		log.Printf("403 forbidden: %s", message)
	} else {
		log.Printf("Status Code: %d;\nError message: %s", statusCode, message)
	}

	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(ErrorResponse{Message: standardMessage})
}
