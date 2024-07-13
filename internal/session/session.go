package session

import (
	"encoding/base64"
	"encoding/json"
	"gpg/portal/internal/user"
	"log"
	"net/http"
	"time"
)

type Session struct {
	ID        	int
	SessionId 	string
	UserId		int
	CreatedAt 	time.Time
	ModifiedAt	time.Time
}

// todo: generate session cookie
func CreateSessionCookie(user user.User) http.Cookie {
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
		Name:   "user_session",
		Value:  encodedValue,
		Secure: true,
	}
	return c
}