package main

import (
	"log"
	"net/http"
	"github.com/gorilla/mux"
	"whysworld.net/byod/views"
)
//guestportal homepage handler
func GuestPortalHomePageHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	portal_id, ok := vars["portal_id"]
	if !ok {
		log.Print("portal_id is missing in parameters")
	}
	switch portal_id{
		case "hotspot":
			views.HotSpotHomePageHandler(w, r)
			return
		case "registration":
			views.RegistratoinHomePageHandler(w, r)
			return
		case "sponsor":
			return
		default:
			http.NotFound(w, r)
			return
	}
	http.NotFound(w, r)
}
//guestportal accept page handler
func GuestPortalAcceptPageHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	portal_id, ok := vars["portal_id"]
	if !ok {
		log.Print("portal_id is missing in parameters")
	}
	switch portal_id{
		case "hotspot":
			views.HotSpotAcceptPageHandler(w, r)
			return
		case "registration":
			views.RegistratoinAcceptPageHandler(w, r)
			return
		case "sponsor":
			return
		default:
			http.NotFound(w, r)
			return
	}
	http.NotFound(w, r)
}
//guestportal decline page handler
func GuestPortalDeclinePageHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	portal_id, ok := vars["portal_id"]
	if !ok {
		log.Print("portal_id is missing in parameters")
	}
	switch portal_id{
		case "hotspot":
			views.HotSpotDeclinePageHandler(w, r)
			return
		case "registration":
			views.RegistratoinDeclinePageHandler(w, r)
			return
		case "sponsor":
			return
		default:
			http.NotFound(w, r)
			return
	}
	http.NotFound(w, r)
}
//guestportal information page handler
func GuestPortalRegistrationPageHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	portal_id, ok := vars["portal_id"]
	if !ok {
		log.Print("portal_id is missing in parameters")
	}
	switch portal_id{
		case "registration":
			views.RegistratoinInfoPageHandler(w, r)
			return
		default:
			http.NotFound(w, r)
			return
	}
	http.NotFound(w, r)
}
//guestportal login page handler
func GuestPortalLoginPageHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	portal_id, ok := vars["portal_id"]
	if !ok {
		log.Print("portal_id is missing in parameters")
	}
	switch portal_id{
		case "registration":
			views.RegistratoinLoginPageHandler(w, r)
			return
		case "sponsor":
			views.SponsorLoginPageHandler(w, r)
			return
		default:
			http.NotFound(w, r)
			return
	}
	http.NotFound(w, r)
}
//guestportal uesrs page handler
func GuestPortalUsersPageHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	portal_id, ok := vars["portal_id"]
	if !ok {
		log.Print("portal_id is missing in parameters")
	}
	switch portal_id{
		case "sponsor":
			views.SponsorUsersPageHandler(w, r)
			return
		default:
			http.NotFound(w, r)
			return
	}
	http.NotFound(w, r)
}

func main() {
	myRouter := NewRouter()

	//Guest Portal endpoints
	myRouter.HandleFunc("/guestportal/{portal_id}/home/", GuestPortalHomePageHandler)
	myRouter.HandleFunc("/guestportal/{portal_id}/accept/", GuestPortalAcceptPageHandler)
	myRouter.HandleFunc("/guestportal/{portal_id}/decline/", GuestPortalDeclinePageHandler)
	myRouter.HandleFunc("/guestportal/{portal_id}/info/", GuestPortalRegistrationPageHandler)
	myRouter.HandleFunc("/guestportal/{portal_id}/login/", GuestPortalLoginPageHandler)
	myRouter.HandleFunc("/guestportal/{portal_id}/users/", GuestPortalUsersPageHandler)
	myRouter.HandleFunc("/guestportal/{portal_id}/logout/", views.LogoutFunc)

	log.Fatal(http.ListenAndServe(":8080", myRouter))
}

func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)

	// Choose the folder to serve
	staticDir := "/static/"

	// Create the route
	router.
		PathPrefix(staticDir).
		Handler(http.StripPrefix(staticDir, http.FileServer(http.Dir("."+staticDir))))

	return router
}