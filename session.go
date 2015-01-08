// Parlay Session Model
// A session requires an owner, typically a user and implements a User struct to handle the incoming data.

package repor

import (
	"errors"
	"fmt"
	rdb "github.com/dancannon/gorethink"
	"github.com/gorilla/securecookie"
	"github.com/gorilla/sessions"
	"log"
	"net/http"
	"time"
)

// Session Struct.  May use this as the Go Class type. Belongs to a User.
// type Session interface {
// 	Create()
// }

type SessionItem struct {
	Username  string
	CreatedAt time.Time
	UpdatedAt time.Time
}

// Create creates a new user session.
func CreateSession(w http.ResponseWriter, r *http.Request, userName string) {
	rdbSession := GetConn()
	key := securecookie.GenerateRandomKey(32)
	// encryptKey := securecookie.GenerateRandomKey(32)

	// Creating a session, creates a new session and stores it in the DB for the given user.
	var store = sessions.NewCookieStore(
		[]byte(key),
	) //CHANGE THE SESSION NAME: Export to Env Variables

	// Either going to create or update the existing session named parlay
	gs, err := store.Get(r, string(key))
	if err != nil {
		errors.New("Error getting the session.")
		return
	}
	gs.Options = &sessions.Options{
		Path:     "/",
		MaxAge:   86400 * 14, //two weeks
		HttpOnly: true,       // for dev, need to change to https only in prod, maybe through an ENV VARIABLE
	}
	err = gs.Save(r, w)
	session := SessionItem{
		Username:  userName,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	_, err = rdb.Table("sessions").Insert(session).RunWrite(rdbSession)
	if err != nil {
		log.Println(err)
	}
	fmt.Println(session)
}

// func (s *SessionItem) Get(w http.ResponseWriter, r *http.Request, p *User) {
// 	session, err := store.Get(r, "parlay")
// 	if err != nil {
// 		errors.New("Error creating or getting the session.")
// 		return
// 	}
// }

// IsActive returns true if a user has an active session.
// func (s *SessionItem) IsActive(userOrkey string) bool {
//   s.
// }
