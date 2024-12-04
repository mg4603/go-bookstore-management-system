package utils

import "net/http"

func SetJSONContentType(n http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		n.ServeHTTP(w, r)
	})
}
