package utils

import (
	"encoding/json"
	"fmt"
	"io"
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
