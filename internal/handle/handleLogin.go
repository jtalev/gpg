package handle

import (
	"context"
	"fmt"
	"gpg/portal/internal/localdb"
	"gpg/portal/internal/validation"
	"html/template"
	"log"
	"net/http"
)

func ServeLogin(ctx context.Context, db localdb.Db) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			log.Println("serving login page")
			tmpl := template.Must(template.ParseFiles("../../web/pages/login.html"))
			if err := tmpl.Execute(w, nil); err != nil {
				http.Error(w, "error executing template", http.StatusInternalServerError)
			}
		},
	)
}

func HandleValidateLogin(ctx context.Context, db localdb.Db) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			log.Println("handling login validation request")
			if err := r.ParseForm(); err != nil {
				log.Printf("error parsing form: %v", err)
				http.Error(w, "error parsing form", http.StatusBadRequest)
				return
			}
			username := r.FormValue("username")
			password := r.FormValue("password")

			result := validation.ValidateLogin(db.Ur, username, password)
			// if result.Msg == "" {
			// 	handle redirect to dashboard
			// }
			var html string
			if result.Msg != "" {
				html = fmt.Sprintf(`<p class="err" id="err">*%s</p>`, result.Msg)
			}
			if _, err := w.Write([]byte(html)); err != nil {
				log.Printf("error writing response: %v", err)
				http.Error(w, "error writing response", http.StatusInternalServerError)
				return
			}

			log.Println("login validation request handled")
		},
	)
}

// todo: generate session cookie
func createSessionCookie() {

}

// todo: create func to take user entered password, hash password and compare to db
