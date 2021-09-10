package views

import (
	"fmt"
	"log"
	"html/template"
	"net/http"

	"whysworld.net/byod/types"
	"whysworld.net/byod/sessions"
	"whysworld.net/byod/db"
	"github.com/satori/go.uuid"
	"github.com/gorilla/mux"
)

func loadSponsorLoginPage() (*types.SponsorPage, error) {
	title := "LOGIN"
	welcomeTitle := ""
	welcomeMessage := "Login here to manage your sponsored guest users"
	content := ""
	return &types.SponsorPage{Title: title, WelcomeTitle: welcomeTitle, WelcomeMessage: welcomeMessage, Content: content, Information: DefaultInfo, Status: "", TimeLeft: ""}, nil
}

func loadSponsorUsersPage() (*types.SponsorPage, error) {
	title := "GUEST USERS"
	welcomeTitle := ""
	welcomeMessage := ""
	content := ""
	info := types.GuestInfo{Name: "John Doe", Email: "John.doe@gmail.com", Company: "ABC Corp", SponsorEmail: "", Option1: "Here for Interview", Option2: "Representing self"}
	return &types.SponsorPage{Title: title, WelcomeTitle: welcomeTitle, WelcomeMessage: welcomeMessage, Content: content, Information: info, Status: "Admitted", TimeLeft: "2 hr 3m remaining"}, nil
}



func SponsorLoginPageHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
    key := vars["id"]
	log.Print("PATH: ", r.URL.Path)
	log.Print("Key: ", key)
	if sessions.IsLoggedIn(r) {
		http.Redirect(w, r, "/sponsor/users", 302)
		return
	}
	session, _ := sessions.Store.Get(r, "byod_session")
	p, err := loadSponsorLoginPage()
	if err != nil {
		http.Redirect(w, r, "sponsor/users", 302)
		return
	}
	switch r.Method {
	case "GET":
		renderSponsorTemplate(w, "sponsor-login", p)
	case "POST":
		log.Print("Inside POST")
		r.ParseForm()
		username := r.Form.Get("username")
		password := r.Form.Get("password")
		log.Print("username: ", username)
		log.Print("password: ", password)
		if (username != "" && password != "") && db.ValidUser(username, password) {
			session.Values["loggedin"] = "true"
			session.Values["username"] = username
			session.Values["oid"] = key
			session.Save(r, w)
			log.Print("user ", username, " is authenticated")
			http.Redirect(w, r, "/sponsor/users", 302)
			return
		}
		errorMessage := "Invalid username or password."
		log.Print(errorMessage)
		renderSponsorTemplate(w, "sponsor-login", &types.SponsorPage{Title: "LOGIN", WelcomeMessage: "Login here to manage your sponsored guest users", Information: DefaultInfo, ErrorMessage: errorMessage})
	default:
		oid := uuid.NewV4().String()
		redirectURI := fmt.Sprintf("/%s/sponsor/login", oid)
		http.Redirect(w, r, redirectURI, http.StatusUnauthorized)
		// http.Redirect(w, r, "/sponsor/login/", http.StatusUnauthorized)
	}
}

func SponsorUsersPageHandler(w http.ResponseWriter, r *http.Request) {
	p, err := loadSponsorUsersPage()
	print(p)
	if err != nil {

		http.Redirect(w, r, "/sponsor/users", http.StatusFound)
		return	
	}
	renderSponsorTemplate(w, "sponsor-users", p)
}

var sponsorTemplate = template.Must(template.ParseFiles("templates/sponsor/sponsor-login.html", "templates/sponsor/sponsor-users.html"))

func renderSponsorTemplate(w http.ResponseWriter, tmpl string, p *types.SponsorPage) {
	err := sponsorTemplate.ExecuteTemplate(w, tmpl+".html", p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
