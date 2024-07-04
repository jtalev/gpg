package handle

import (
	"context"
	"encoding/json"
	"fmt"
	"gpg/portal/internal/localdb"
	"net/http"
	"path/filepath"
)

func encode[T any](w http.ResponseWriter, status int, v T) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	if err := json.NewEncoder(w).Encode(v); err != nil {
		return fmt.Errorf("error encoding json: %w", err)
	}
	return nil
}

func decode[T any](r *http.Request) error {
	var v T
	if err := json.NewDecoder(r.Body).Decode(&v); err != nil {
		return fmt.Errorf("error decoding json: %w", err)
	}
	return nil
}

func ServeIndex(ctx context.Context, db localdb.Db) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r * http.Request){
			path := filepath.Join("..", "..", "web", "pages", "login.html")
			http.ServeFile(w, r, path)
		},
	)
}