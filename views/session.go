package views

import (
	"fmt"
	"net/http"
	"github.com/gorilla/mux"
	"whysworld.net/byod/sessions"
	"log"
)

//RequiresLogin is a middleware which will be used for each httpHandler to check if there is any active session
func RequiresLogin(handler func(w http.ResponseWriter, r *http.Request)) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		if !sessions.IsLoggedIn(r) {
			vars := mux.Vars(r)
			portal_id, ok := vars["portal_id"]
			if !ok {
				log.Print("portal_id is missing in parameters")
			}
			redirectURI := fmt.Sprintf("/guestportal/%s/login", portal_id)
			http.Redirect(w, r, redirectURI, 302)
			return
		}
		handler(w, r)
	}
}

//LogoutFunc Implements the logout functionality. WIll delete the session information from the cookie store
func LogoutFunc(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	portal_id := vars["portal_id"]
	session, err := sessions.Store.Get(r, "byod_session")
	if err == nil { //If there is no error, then remove session
		if session.Values["loggedin"] != "false" {
			session.Values["loggedin"] = "false"
			session.Save(r, w)
		}
	}
	// oid := uuid.NewV4().String()
	redirectURI := fmt.Sprintf("/guestportal/%s/login", portal_id)
	http.Redirect(w, r, redirectURI, 302) //redirect to login irrespective of error or not
}
