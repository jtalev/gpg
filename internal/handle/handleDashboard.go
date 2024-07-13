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
			log.Println("serving dashboard")
			http.ServeFile(w, r, "../../web/pages/dashboard.html")
			tmpl := template.Must(template.ParseFiles("../../web/pages/dashboard.html"))
			if err := tmpl.Execute(w, r.Body); err != nil {
				http.Error(w, "error executing template", http.StatusInternalServerError)
			}
			log.Println("dashboard served")
		},
	)
}