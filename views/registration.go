package views

import (
	"html/template"
	"net/http"
	"whysworld.net/byod/types"
	"github.com/gorilla/mux"
	"log"
	"fmt"
)

func loadRegistrationHomePage() (*types.RegistrationPage, error) {
	title := "Acceptable Use Policy"
	welcomeTitle := "BYOD Welcome"
	welcomeMessage := "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Aliquam sed enim sollicitudin, venenatis massa sit amet, placerat neque. Proin blandit arcu."
	content := "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Aliquam sed enim sollicitudin, venenatis massa sit amet, placerat neque. Proin blandit arcu at metus bibendum, a fringilla purus feugiat. Mauris nec nulla auctor, porta enim et, eleifend magna. Phasellus blandit quam ut ante posuere tempus. Sed condimentum elementum libero, sit amet semper sapien tempus eget. Duis egestas vehicula semper. Nulla eleifend, felis tincidunt porttitor porta, quam lacus rhoncus ipsum.Lorem ipsum dolor sit amet, consectetur adipiscing elit. Aliquam sed enim sollicitudin, venenatis massa sit amet, placerat neque. Proin blandit arcu at metus bibendum, a fringilla purus feugiat. Mauris nec nulla auctor, porta enim et, eleifend magna. Phasellus blandit quam ut ante posuere tempus. Sed condimentum elementum libero, sit amet semper sapien tempus eget. Duis egestas vehicula semper. Nulla eleifend, felis tincidunt porttitor porta, quam lacus rhoncus ipsum.Lorem ipsum dolor sit amet, consectetur adipiscing elit. Aliquam sed enim sollicitudin, venenatis massa sit amet, placerat neque. Proin blandit arcu at metus bibendum, a fringilla purus feugiat. Mauris nec nulla auctor, porta enim et, eleifend magna. Phasellus blandit quam ut ante posuere tempus. Sed condimentum elementum libero, sit amet semper sapien tempus eget. Duis egestas vehicula semper. Nulla eleifend, felis tincidunt porttitor porta, quam lacus rhoncus ipsum"

	return &types.RegistrationPage{Title: title, WelcomeTitle: welcomeTitle, WelcomeMessage: welcomeMessage, Content: content, Information: DefaultInfo}, nil
}

func loadRegistrationInfoPage() (*types.RegistrationPage, error) {
	title := "Guest Information"
	welcomeTitle := ""
	welcomeMessage := ""
	content := ""
	return &types.RegistrationPage{Title: title, WelcomeTitle: welcomeTitle, WelcomeMessage: welcomeMessage, Content: content, Information: DefaultInfo}, nil
}

func loadRegistrationAcceptPage() (*types.RegistrationPage, error) {
	title := "Success"
	welcomeTitle := ""
	welcomeMessage := "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Aliquam sed enim sollicitudin, venenatis massa sit amet, placerat neque. Proin blandit arcu."
	content := ""
	return &types.RegistrationPage{Title: title, WelcomeTitle: welcomeTitle, WelcomeMessage: welcomeMessage, Content: content, Information: DefaultInfo}, nil
}

func loadRegistrationLoginPage() (*types.RegistrationPage, error) {
	title := "Success"
	welcomeTitle := "Install"
	welcomeMessage := "Your password is in your email. If Company policy requires your sponsor to approve your access, email will be sent to you .. After that approval"
	content := ""
	return &types.RegistrationPage{Title: title, WelcomeTitle: welcomeTitle, WelcomeMessage: welcomeMessage, Content: content, Information: DefaultInfo}, nil
}

func loadRegistrationDeclinePage() (*types.RegistrationPage, error) {
	title := "Decline"
	welcomeTitle := "Install"
	welcomeMessage := "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Aliquam sed enim sollicitudin, venenatis massa sit amet, placerat neque. Proin blandit arcu."
	content := ""
	return &types.RegistrationPage{Title: title, WelcomeTitle: welcomeTitle, WelcomeMessage: welcomeMessage, Content: content, Information: DefaultInfo}, nil
}
func RegistratoinHomePageHandler(w http.ResponseWriter, r *http.Request) {
	//get params (portal_id)
	vars := mux.Vars(r)
	portal_id, ok := vars["portal_id"]
	if !ok {
		log.Print("portal_id is missing in parameters")
	}
	log.Print(`portal_id := `, portal_id)

	p, err := loadRegistrationHomePage()
	print(p)
	if err != nil {
		return
	}
	switch r.Method {
		case "GET":
			templateName := fmt.Sprintf("%s-home", portal_id)
			renderRegistrationTemplate(w, templateName, p)
		case "POST":
			r.ParseForm()
			action := r.Form.Get("action")
			if action == "accept" {
				redirectURI := fmt.Sprintf("/guestportal/%s/info?accept=true", portal_id)
				http.Redirect(w, r, redirectURI, http.StatusFound)
				return
			} else {
				redirectURI := fmt.Sprintf("/guestportal/%s/decline?accept=false", portal_id)
				http.Redirect(w, r, redirectURI, http.StatusFound)
				return
			}
		default: 
			renderRegistrationTemplate(w, "registration-home", p)
			return
	}
}

func RegistratoinInfoPageHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	portal_id, ok := vars["portal_id"]
	if !ok {
		log.Print("portal_id is missing in parameters")
	}
	log.Print(`portal_id := `, portal_id)

	p, err := loadRegistrationInfoPage()
	print(p)
	if err != nil {
		return
	}
	switch r.Method {
		case "GET":
			templateName := fmt.Sprintf("%s-information", portal_id)
			renderRegistrationTemplate(w, templateName, p)
		case "POST":
			r.ParseForm()
			action := r.Form.Get("action")
			if action == "register" {
				redirectURI := fmt.Sprintf("/guestportal/%s/login?accept=true", portal_id)
				http.Redirect(w, r, redirectURI, http.StatusFound)
				return
			} else {
				redirectURI := fmt.Sprintf("/guestportal/%s/login?accept=false", portal_id)
				http.Redirect(w, r, redirectURI, http.StatusFound)
				return
			}
		default: 
			renderRegistrationTemplate(w, "registration-information", p)
			return
	}
}

func RegistratoinLoginPageHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	portal_id, ok := vars["portal_id"]
	if !ok {
		log.Print("portal_id is missing in parameters")
	}
	log.Print(`portal_id := `, portal_id)

	p, err := loadRegistrationLoginPage()
	print(p)
	if err != nil {
		return
	}
	switch r.Method {
		case "GET":
			templateName := fmt.Sprintf("%s-login", portal_id)
			renderRegistrationTemplate(w, templateName, p)
		case "POST":
			r.ParseForm()
			action := r.Form.Get("action")
			if action == "success" {
				redirectURI := fmt.Sprintf("/guestportal/%s/accept?accept=true", portal_id)
				http.Redirect(w, r, redirectURI, http.StatusFound)
				return
			} else {
				redirectURI := fmt.Sprintf("/guestportal/%s/info?action=signup", portal_id)
				http.Redirect(w, r, redirectURI, http.StatusFound)
				return
			}
		default: 
			renderRegistrationTemplate(w, "registration-login", p)
			return
	}
}

func RegistratoinAcceptPageHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	portal_id, ok := vars["portal_id"]
	if !ok {
		log.Print("portal_id is missing in parameters")
	}
	log.Print(`portal_id := `, portal_id)

	p, err := loadRegistrationAcceptPage()
	print(p)
	if err != nil {
		return
	}
	switch r.Method {
		case "GET":
			templateName := fmt.Sprintf("%s-accept", portal_id)
			renderRegistrationTemplate(w, templateName, p)
		default: 
			renderRegistrationTemplate(w, "registration-accept", p)
			return
	}
}

func RegistratoinDeclinePageHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	portal_id, ok := vars["portal_id"]
	if !ok {
		log.Print("portal_id is missing in parameters")
	}
	log.Print(`portal_id := `, portal_id)

	p, err := loadRegistrationDeclinePage()
	print(p)
	if err != nil {
		return
	}
	switch r.Method {
		case "GET":
			templateName := fmt.Sprintf("%s-decline", portal_id)
			renderRegistrationTemplate(w, templateName, p)
		case "POST":
			r.ParseForm()
			action := r.Form.Get("action")
			if action == "back" {
				redirectURI := fmt.Sprintf("/guestportal/%s/home?refresh=true", portal_id)
				http.Redirect(w, r, redirectURI, http.StatusFound)
				return
			} else {
				redirectURI := fmt.Sprintf("/guestportal/%s/decline?accept=false", portal_id)
				http.Redirect(w, r, redirectURI, http.StatusFound)
				return
			}
		default: 
			renderRegistrationTemplate(w, "registration-decline", p)
			return
	}
}

var registrationTemplate = template.Must(template.ParseFiles("templates/registration/registration-decline.html", "templates/registration/registration-home.html", "templates/registration/registration-information.html", "templates/registration/registration-accept.html","templates/registration/registration-login.html"))

func renderRegistrationTemplate(w http.ResponseWriter, tmpl string, p *types.RegistrationPage) {
	err := registrationTemplate.ExecuteTemplate(w, tmpl+".html", p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}