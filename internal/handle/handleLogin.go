package handle

import (
	"context"
	"gpg/portal/internal/localdb"
	"html/template"
	"net/http"
)

func ServeIndex(ctx context.Context, db localdb.Db) http.Handler {
	users := db.GetUsers()
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			tmpl := template.Must(template.ParseFiles("../../web/pages/login.html"))
			tmpl.Execute(w, users)
		},
	)
}

// todo: generate session cookie
func createSessionCookie() {

}

type Login struct {
	Username string
	Password string
}

// todo: create func to take user entered password, hash password and compare to db
