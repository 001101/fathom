package api

import (
	"context"
	"encoding/json"
	"net/http"
	"os"

	"github.com/gorilla/sessions"
	"github.com/usefathom/fathom/pkg/datastore"
	"golang.org/x/crypto/bcrypt"
)

type key int

const (
	userKey key = 0
)

type login struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

var store = sessions.NewCookieStore([]byte(os.Getenv("FATHOM_SECRET")))

// URL: POST /api/session
var LoginHandler = HandlerFunc(func(w http.ResponseWriter, r *http.Request) error {

	// check login creds
	var l login
	json.NewDecoder(r.Body).Decode(&l)

	u, err := datastore.GetUserByEmail(l.Email)

	// compare pwd
	if err != nil || bcrypt.CompareHashAndPassword([]byte(u.HashedPassword), []byte(l.Password)) != nil {
		w.WriteHeader(http.StatusUnauthorized)
		return respond(w, envelope{Error: "invalid_credentials"})
	}

	session, _ := store.Get(r, "auth")
	session.Values["user_id"] = u.ID
	err = session.Save(r, w)
	if err != nil {
		return err
	}

	return respond(w, envelope{Data: true})
})

// URL: DELETE /api/session
var LogoutHandler = HandlerFunc(func(w http.ResponseWriter, r *http.Request) error {
	session, _ := store.Get(r, "auth")
	if !session.IsNew {
		session.Options.MaxAge = -1
		err := session.Save(r, w)
		if err != nil {
			return err
		}
	}

	return respond(w, envelope{Data: true})
})

/* middleware */
func Authorize(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		session, _ := store.Get(r, "auth")
		userID, ok := session.Values["user_id"]

		if !ok {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		// find user
		u, err := datastore.GetUser(userID.(int64))
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		ctx := context.WithValue(r.Context(), userKey, u)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
