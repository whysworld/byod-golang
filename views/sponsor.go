package views

import (
	"fmt"
	"log"
	"html/template"
	"net/http"
	"whysworld.net/byod/types"
	"whysworld.net/byod/sessions"
	"whysworld.net/byod/db"
	"github.com/gorilla/mux"
)

func replaceQueryParams(r *http.Request, key string, value string) (string){
	q := r.URL.Query()
	q.Set(key, value)
	r.URL.RawQuery = q.Encode()
	redirectURI := fmt.Sprintf("%s", r.URL)
	return redirectURI
}

func loadSponsorLoginPage() (*types.SponsorPage, error) {
	title := "Sponsor Portal"
	subTitle := "LOGIN"
	welcomeTitle := ""
	welcomeMessage := "Login here to manage your sponsored guest users"
	content := ""
	return &types.SponsorPage{Title: title, SubTitle: subTitle, WelcomeTitle: welcomeTitle, WelcomeMessage: welcomeMessage, Content: content, Information: DefaultInfo, Status: "", TimeLeft: ""}, nil
}

func loadSponsorUsersPage(status string) (*types.SponsorPage, error) {
	title := "Sponsor Portal"
	subTitle := "GUEST USERS"
	welcomeTitle := ""
	welcomeMessage := ""
	content := ""
	if status == "waiting"{
		info := []types.GuestInfo{
			{Name: "John Doe", Email: "John.doe@gmail1.com", Company: "ABC Corp", SponsorEmail: "", Option1: "Here for Interview", Option2: "Representing self", Status: "Waiting"},
			{Name: "John Doe1", Email: "John.doe@gmail2.com", Company: "ABC Corp", SponsorEmail: "", Option1: "Here for Interview", Option2: "Representing self", Status: "Waiting"},
		}
		return &types.SponsorPage{Title: title, SubTitle: subTitle, WelcomeTitle: welcomeTitle, WelcomeMessage: welcomeMessage, Content: content, Information: info, TimeLeft: "2 hr 3m remaining", Status: "waiting"}, nil
	} else if status == "admitted"{
		info := []types.GuestInfo{
			{Name: "John Doe2", Email: "John.doe@gmail3.com", Company: "ABC Corp", SponsorEmail: "", Option1: "Here for Interview", Option2: "Representing self", Status: "Admitted", CreatedAt: "9/14/2021 04:00 PM"},
			{Name: "John Doe3", Email: "John.doe@gmail4.com", Company: "ABC Corp", SponsorEmail: "", Option1: "Here for Interview", Option2: "Representing self", Status: "Admitted", CreatedAt: "9/14/2021 04:00 PM"},
		}
		return &types.SponsorPage{Title: title, SubTitle: subTitle, WelcomeTitle: welcomeTitle, WelcomeMessage: welcomeMessage, Content: content, Information: info, TimeLeft: "2 hr 3m remaining", Status: "admitted"}, nil
	} else {
		info := []types.GuestInfo{
			{Name: "John Doe", Email: "John.doe@gmail1.com", Company: "ABC Corp", SponsorEmail: "", Option1: "Here for Interview", Option2: "Representing self", Status: "Waiting"},
			{Name: "John Doe1", Email: "John.doe@gmail2.com", Company: "ABC Corp", SponsorEmail: "", Option1: "Here for Interview", Option2: "Representing self", Status: "Waiting"},
			{Name: "John Doe2", Email: "John.doe@gmail3.com", Company: "ABC Corp", SponsorEmail: "", Option1: "Here for Interview", Option2: "Representing self", Status: "Admitted", CreatedAt: "9/14/2021 04:00 PM"},
			{Name: "John Doe3", Email: "John.doe@gmail4.com", Company: "ABC Corp", SponsorEmail: "", Option1: "Here for Interview", Option2: "Representing self", Status: "Admitted", CreatedAt: "9/14/2021 04:00 PM"},
		}
		return &types.SponsorPage{Title: title, SubTitle: subTitle, WelcomeTitle: welcomeTitle, WelcomeMessage: welcomeMessage, Content: content, Information: info, TimeLeft: "2 hr 3m remaining", Status: "all"}, nil
	}

}

func loadSponsorAddUserPage() (*types.SponsorPage, error) {
	title := "Sponsor Portal"
	subTitle := "Add a user"
	welcomeTitle := ""
	welcomeMessage := ""
	content := ""
	info := []types.GuestInfo{
		{Name: "John Doe", Email: "John.doe@gmail.com", Company: "ABC Corp", SponsorEmail: "", Option1: "Here for Interview", Option2: "Representing self"},
		{Name: "John Doe1", Email: "John.doe@gmail1.com", Company: "ABC Corp", SponsorEmail: "", Option1: "Here for Interview", Option2: "Representing self"},
	}
	return &types.SponsorPage{Title: title, SubTitle: subTitle, WelcomeTitle: welcomeTitle, WelcomeMessage: welcomeMessage, Content: content, Information: info, Status: "Admitted", TimeLeft: "2 hr 3m remaining"}, nil
}

func SponsorLoginPageHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
    portal_id := vars["portal_id"]
	log.Print("login: ", portal_id)
	if sessions.IsLoggedIn(r) {
		redirectURI := fmt.Sprintf("/guestportal/%s/users", portal_id)
		http.Redirect(w, r, redirectURI, http.StatusFound)
		return
	}
	session, _ := sessions.Store.Get(r, "byod_session")
	p, err := loadSponsorLoginPage()
	if err != nil {
		return
	}
	switch r.Method {
	case "GET":
		templateName := fmt.Sprintf("%s-login", portal_id)
		renderSponsorTemplate(w, templateName, p)
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
			session.Values["oid"] = portal_id
			session.Save(r, w)
			log.Print("user ", username, " is authenticated")
			templateName := fmt.Sprintf("/guestportal/%s/users?accept=true", portal_id)
			http.Redirect(w, r, templateName, 302)
			return
		}
		errorMessage := "Invalid username or password."
		log.Print(errorMessage)
		templateName := fmt.Sprintf("%s-login", portal_id)
		renderSponsorTemplate(w, templateName, &types.SponsorPage{Title: "LOGIN", WelcomeMessage: "Login here to manage your sponsored guest users", Information: DefaultInfo, ErrorMessage: errorMessage})
	default:
		// oid := uuid.NewV4().String()
		// redirectURI := fmt.Sprintf("/%s/sponsor/login", oid)
		// http.Redirect(w, r, redirectURI, http.StatusUnauthorized)
		
		renderSponsorTemplate(w, "registration-decline", p)
		return
	}
}

func SponsorUsersPageHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
    portal_id := vars["portal_id"]
	p, err := loadSponsorUsersPage("all")
	if err != nil {
		return	
	}
	switch r.Method {
		case "GET":
			templateName := fmt.Sprintf("%s-users", portal_id)
			renderSponsorTemplate(w, templateName, p)
		case "POST":
			r.ParseForm()
			action := r.Form.Get("action")
			status := r.Form.Get("status")
			selectedItems := r.PostForm["selected"]
			//get selected emails
			log.Print("action: ", selectedItems)
			log.Print("status: ", status)
			if status == "waiting" {
				p, _ := loadSponsorUsersPage("waiting")
				templateName := fmt.Sprintf("%s-users", portal_id)
				renderSponsorTemplate(w, templateName, p)
				return
			} else if status == "admitted"{
				p, _ := loadSponsorUsersPage("admitted")
				templateName := fmt.Sprintf("%s-users", portal_id)
				renderSponsorTemplate(w, templateName, p)
				return
			} else {
				p, _ := loadSponsorUsersPage("all")
				templateName := fmt.Sprintf("%s-users", portal_id)
				renderSponsorTemplate(w, templateName, p)
				return
			}
			if action == "logout" {
				redirectURI := fmt.Sprintf("/guestportal/%s/logout?loggedout=true", portal_id)
				http.Redirect(w, r, redirectURI, http.StatusFound)
				return
			} else if action == "add-user"{
				redirectURI := fmt.Sprintf("/guestportal/%s/adduser?action=add-user", portal_id)
				http.Redirect(w, r, redirectURI, http.StatusFound)
				return
			} else if action == "admit_selected"{
				redirectURI := replaceQueryParams(r, "action", "admit_selected")
				http.Redirect(w, r, redirectURI, http.StatusFound)
				return
			} else if action == "admit_all"{
				redirectURI := replaceQueryParams(r, "action", "admit_all")
				http.Redirect(w, r, redirectURI, http.StatusFound)
				return
			} else if action == "extend_time"{
				redirectURI := replaceQueryParams(r, "action", "extend_time")
				http.Redirect(w, r, redirectURI, http.StatusFound)
				return
			} else if action == "remove_selected"{
				redirectURI := replaceQueryParams(r, "action", "remove_selected")
				http.Redirect(w, r, redirectURI, http.StatusFound)
				return
			}
		default: 
			renderSponsorTemplate(w, "sponsor-users", p)
			return
	}
}

func SponsorAddUserPageHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
    portal_id := vars["portal_id"]
	p, err := loadSponsorAddUserPage()
	print(p)
	if err != nil {
		return	
	}
	switch r.Method {
	case "GET":
		templateName := fmt.Sprintf("%s-adduser", portal_id)
		renderSponsorTemplate(w, templateName, p)
	case "POST":
		r.ParseForm()
		action := r.Form.Get("action")
		if action == "save" {
			name := r.Form.Get("name")
			email := r.Form.Get("email")
			company := r.Form.Get("company")
			option1 := r.Form.Get("option1")
			option2 := r.Form.Get("option2")

			log.Print("username: ", name)
			log.Print("email: ", email)
			log.Print("company: ", company)
			log.Print("option2: ", option2)
			log.Print("option1: ", option1)

			templateName := fmt.Sprintf("%s-adduser", portal_id)
			renderSponsorTemplate(w, templateName, p)
			return
		}
	default: 
		renderSponsorTemplate(w, "sponsor-adduser", p)
	}
}

var sponsorTemplate = template.Must(template.ParseFiles("templates/sponsor/sponsor-adduser.html", "templates/sponsor/sponsor-login.html", "templates/sponsor/sponsor-users.html"))

func renderSponsorTemplate(w http.ResponseWriter, tmpl string, p *types.SponsorPage) {
	err := sponsorTemplate.ExecuteTemplate(w, tmpl+".html", p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
