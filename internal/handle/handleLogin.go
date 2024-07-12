package handle

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"gpg/portal/internal/localdb"
	"gpg/portal/internal/user"
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

func ServeDashboard(ctx context.Context, db localdb.Db) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			log.Println("serving dashboard")
			http.ServeFile(w, r, "../../web/pages/dashboard.html")
			log.Println("dashboard served")
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
			var html string
			if !result.IsValid {
				html = fmt.Sprintf(`<p class="err" id="err">*%s</p>`, result.Msg)
				if _, err := w.Write([]byte(html)); err != nil {
					log.Printf("error writing response: %v", err)
					http.Error(w, "error writing response", http.StatusInternalServerError)
					return
				}
			}
			if result.IsValid {
				user, _ := db.Ur.GetUserByUsername(username)
				cookie := createSessionCookie(user)
				if err := cookie.Valid(); err != nil {
					log.Printf("cookie invalid: %v", err)
					return
				}
				http.SetCookie(w, &cookie)
				log.Println("cookie sent to client")
				log.Println("user logged in successfully")
				w.Header().Set("HX-Redirect", "/dashboard")
			}
			log.Println("login validation request handled")
		},
	)
}

// todo: generate session cookie
func createSessionCookie(user user.User) http.Cookie {
	data := map[string]interface{}{
		"SessionId":        "secret_session_id",
		"UserId":           user.Id,
		"IsAuthentication": true,
	}
	value, err := json.Marshal(data)
	if err != nil {
		log.Printf("error marshalling json: %v", err)
	}
	encodedValue := base64.StdEncoding.EncodeToString(value)
	c := http.Cookie{
		Name:   "session",
		Value:  encodedValue,
		Secure: true,
	}
	return c
}
