package sessions

import (
	"net/http"
	"github.com/gorilla/sessions"
)

//Store the cookie store which is going to store session data in the cookie
var Store = sessions.NewCookieStore([]byte("VHUgr4RBwjMvvfhm"))
var session *sessions.Session

func init() {
	//set expiry in second
	Store.MaxAge(60 * 10) //10 minutes
}
//IsLoggedIn will check if the user has an active session and return True
func IsLoggedIn(r *http.Request) bool {

	session, err := Store.Get(r, "byod_session")

	if err == nil && (session.Values["loggedin"] == "true") {
		return true
	}
	return false
}

//GetCurrentUserName returns the username of the logged in user
func GetCurrentUserName(r *http.Request) string {
	session, err := Store.Get(r, "byod_session")
	if err == nil {
		return session.Values["username"].(string)
	}
	return ""
}
