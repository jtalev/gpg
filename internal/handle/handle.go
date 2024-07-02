package handle

import (
	"context"
	"encoding/json"
	"fmt"
	"gpg/portal/internal/user"
	"net/http"
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

func HandleDoubleInt(u user.User) http.Handler {
	s := "Hello, " + u.Username
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			encode(w, http.StatusOK, u)
			json.NewEncoder(w).Encode(s)
		},
	)
}

func ServeIndex(ctx context.Context) http.Handler {
	u := ctx.Value("user_name")
	return http.HandlerFunc(
		func(w http.ResponseWriter, r * http.Request){
			encode(w, http.StatusOK, u)
		},
	)
}