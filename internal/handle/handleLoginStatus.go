package handle

import (
	"context"
	"fmt"
	"gpg/portal/internal/database"
	"gpg/portal/internal/validation"
	"html/template"
	"log"
	"net/http"

	"github.com/gorilla/sessions"
)

var (
	key           = []byte("secret_key")
	encryptionKey = []byte("another_secfasdfsdfdsfsdfsdfsdfy")
	store         = sessions.NewCookieStore(key, encryptionKey)
)

func ServeLogin(ctx context.Context, db *database.Db) http.Handler {
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
				log.Println("serving login page")
				tmpl := template.Must(template.ParseFiles("../../web/pages/login.html"))
				if err := tmpl.Execute(w, nil); err != nil {
					http.Error(w, "error executing template", http.StatusInternalServerError)
				}
				return
			}

			log.Println("user authenticated")
			http.Redirect(w, r, "/dashboard", http.StatusSeeOther)
		},
	)
}

func HandleValidateLogin(ctx context.Context, db *database.Db) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			log.Println("handling login validation request")
			if err := r.ParseForm(); err != nil {
				log.Printf("HandleValidateLogin: error parsing form: %v", err)
				http.Error(w, "error parsing form", http.StatusBadRequest)
				return
			}
			username := r.FormValue("username")
			password := r.FormValue("password")

			result := validation.ValidateLogin(db.Ur, username, password)
			log.Println(result)
			var html string
			if !result.IsValid {
				html = fmt.Sprintf(`<p class="err" id="err">*%s</p>`, result.Msg)
				if _, err := w.Write([]byte(html)); err != nil {
					log.Printf("HandleValidateLogin: error writing response: %v", err)
					http.Error(w, "error writing response", http.StatusInternalServerError)
					return
				}
			} else {
				session, err := store.Get(r, "user_session")
				if err != nil {
					log.Printf("HandleValidateLogin: error getting session: %v", err)
					http.Error(w, "error getting session", http.StatusInternalServerError)
					return
				}
				u, err := db.Ur.GetUserByUsername(username)
				if err != nil {
					log.Printf("error getting user: %v", err)
					http.Error(w, "error getting user", http.StatusBadRequest)
					return
				}
				session.Values["is_authenticated"] = true
				session.Values["username"] = username
				if u.IsAdmin {
					session.Values["role"] = "admin"
				} else {
					session.Values["role"] = "user"
				}
				err = session.Save(r, w)
				if err != nil {
					log.Printf("HandleValidateLogin: error saving session: %v", err)
					http.Error(w, "error saving session", http.StatusInternalServerError)
					return
				}
				log.Println("cookie sent to client")
				log.Println("user logged in successfully")
				w.Header().Set("HX-Redirect", "/dashboard")
			}
			log.Println("login validation request handled")
		},
	)
}

func HandleLogout() http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			log.Println("logging out")
			session, err := store.Get(r, "user_session")
			if err != nil {
				log.Printf("error getting session: %v", err)
				http.Error(w, "error getting session", http.StatusInternalServerError)
				return
			}
			session.Values["is_authenticated"] = false
			err = session.Save(r, w)
			if err != nil {
				log.Printf("error saving session: %v", err)
				http.Error(w, "error saving session", http.StatusInternalServerError)
				return
			}
			http.Redirect(w, r, "/", http.StatusFound)
			log.Println("user logged out")
		},
	)
}
