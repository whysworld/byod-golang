package main

import (
	"log"
	"net/http"
	"github.com/gorilla/mux"
	"whysworld.net/byod/views"
)

func main() {
	// myRouter := mux.NewRouter().StrictSlash(true)
	myRouter := NewRouter()
	//HotSpot Portal endpoints
	myRouter.HandleFunc("/hotspot/home/", views.MakeHandler(views.HotSpotHomePageHandler))
	myRouter.HandleFunc("/hotspot/accept/", views.MakeHandler(views.HotSpotAcceptPageHandler))
	myRouter.HandleFunc("/hotspot/decline/", views.MakeHandler(views.HotSpotDeclinePageHandler))

	//Self Registration endpoints
	myRouter.HandleFunc("/registration/home/", views.MakeHandler(views.RegistratoinHomePageHandler))
	myRouter.HandleFunc("/registration/info/", views.MakeHandler(views.RegistratoinInfoPageHandler))
	myRouter.HandleFunc("/registration/login/", views.MakeHandler(views.RegistratoinLoginPageHandler))
	myRouter.HandleFunc("/registration/accept/", views.MakeHandler(views.RegistratoinAcceptPageHandler))

	//Sponsor Portal endpoints
	myRouter.HandleFunc("/{id}/sponsor/login/",(views.SponsorLoginPageHandler))
	myRouter.HandleFunc("/sponsor/users/", views.MakeHandler(views.RequiresLogin(views.SponsorUsersPageHandler)))
	//logout
	myRouter.HandleFunc("/sponsor/logout/", views.RequiresLogin(views.LogoutFunc))

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