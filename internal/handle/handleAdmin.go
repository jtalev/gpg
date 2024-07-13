package handle

import (
	"html/template"
	"log"
	"net/http"
	"path/filepath"
)

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