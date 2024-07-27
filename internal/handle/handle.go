package handle

import (
	"context"
	"encoding/json"
	"fmt"
	"gpg/portal/internal/database"
	"html/template"
	"log"
	"net/http"
	"path/filepath"
)

type Cookie struct {
	Username		string
	IsAdmin			bool
}
var cookie = Cookie{}

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

func ServeAdmin() http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			path := filepath.Join("..", "..", "web", "pages", "admin.html")
			tmpl := template.Must(template.ParseFiles(path))
			if err := tmpl.Execute(w, nil); err != nil {
				log.Printf("error executing template: %v", err)
				http.Error(w, "error executing template", http.StatusInternalServerError)
			}
		},
	)
}

func ServeDashboard(ctx context.Context, db *database.Db) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			session, err := store.Get(r, "user_session")
			if err != nil {
				log.Printf("error getting session: %v", err)
				http.Error(w, "error getting session", http.StatusInternalServerError)
				return
			}
			auth, ok := session.Values["is_authenticated"].(bool)
			if !ok || !auth {
				log.Printf("unauthorized user")
				http.Error(w, "unauthorized user", http.StatusUnauthorized)
				return
			}
			log.Println("serving dashboard")

			username := session.Values["username"].(string)
			isAdmin := session.Values["is_admin"].(bool)
			cookie.Username = username
			cookie.IsAdmin = isAdmin

			path := filepath.Join("..", "..", "web", "pages", "dashboard.html")
			navPath := filepath.Join("..", "..", "web", "tmpl", "nav.html")

			tmpl := template.Must(template.ParseFiles(path, navPath))
			if err := tmpl.Execute(w, cookie); err != nil {
				log.Printf("error executing template: %v", err)
				http.Error(w, "error executing template", http.StatusInternalServerError)
			}
			log.Println("dashboard served")
		},
	)
}