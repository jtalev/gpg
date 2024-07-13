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
			auth := session.Values["is_authenticated"].(bool)
			if !auth {
				log.Printf("unauthorised user: %v", err)
				http.Error(w, "unauthorised user", http.StatusInternalServerError)
				
			} else {
				log.Println("serving dashboard")
				http.ServeFile(w, r, "../../web/pages/dashboard.html")
				tmpl := template.Must(template.ParseFiles("../../web/pages/dashboard.html"))
				if err := tmpl.Execute(w, r.Body); err != nil {
					http.Error(w, "error executing template", http.StatusInternalServerError)
				}
				log.Println("dashboard served")
			}
		},
	)
}
