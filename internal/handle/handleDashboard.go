package handle

import (
	"context"
	"gpg/portal/internal/localdb"
	"html/template"
	"log"
	"net/http"
)

func ServeDashboard(ctx context.Context, db localdb.Db) http.Handler {
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
			tmpl := template.Must(template.ParseFiles("../../web/pages/dashboard.html"))
			if err := tmpl.Execute(w, nil); err != nil {
				log.Printf("error executing template: %v", err)
				http.Error(w, "error executing template", http.StatusInternalServerError)
			}
			log.Println("dashboard served")
		},
	)
}