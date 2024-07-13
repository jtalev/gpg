package handle

import (
	"context"
	"gpg/portal/internal/database"
	"gpg/portal/internal/session"
	"html/template"
	"log"
	"net/http"
)

var cookie = session.Cookie{}

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

			tmpl := template.Must(template.ParseFiles("../../web/pages/dashboard.html"))
			if err := tmpl.Execute(w, cookie); err != nil {
				log.Printf("error executing template: %v", err)
				http.Error(w, "error executing template", http.StatusInternalServerError)
			}
			log.Println("dashboard served")
		},
	)
}