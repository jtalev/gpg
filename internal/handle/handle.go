package handle

import (
	"context"
	"encoding/json"
	"fmt"
	"gpg/portal/internal/localdb"
	"html/template"
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

func ServeIndex(ctx context.Context, db localdb.Db) http.Handler {
	users := db.GetUsers()
	return http.HandlerFunc(
		func(w http.ResponseWriter, r * http.Request){
			tmpl := template.Must(template.ParseFiles("../../web/pages/login.html"))
			tmpl.Execute(w, users)
		},
	)
}