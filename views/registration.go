package views

import (
	"html/template"
	"net/http"

	"whysworld.net/byod/types"
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
	title := "Login"
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

func RegistratoinHomePageHandler(w http.ResponseWriter, r *http.Request) {
	p, err := loadRegistrationHomePage()
	print(p)
	if err != nil {
		http.Redirect(w, r, "/registration/home/", http.StatusFound)
		return
	}
	renderRegistrationTemplate(w, "registration-home", p)
}

func RegistratoinInfoPageHandler(w http.ResponseWriter, r *http.Request) {
	p, err := loadRegistrationInfoPage()
	print(p)
	if err != nil {
		http.Redirect(w, r, "/registration/info/", http.StatusFound)
		return
	}
	renderRegistrationTemplate(w, "registration-information", p)
}

func RegistratoinLoginPageHandler(w http.ResponseWriter, r *http.Request) {
	p, err := loadRegistrationLoginPage()
	print(p)
	if err != nil {
		http.Redirect(w, r, "/registration/login/", http.StatusFound)
		return
	}
	renderRegistrationTemplate(w, "registration-login", p)
}

func RegistratoinAcceptPageHandler(w http.ResponseWriter, r *http.Request) {
	p, err := loadRegistrationLoginPage()
	print(p)
	if err != nil {
		http.Redirect(w, r, "/registration/accept/", http.StatusFound)
		return
	}
	renderRegistrationTemplate(w, "registration-accept", p)
}

var registrationTemplate = template.Must(template.ParseFiles("templates/registration/registration-home.html", "templates/registration/registration-information.html", "templates/registration/registration-accept.html","templates/registration/registration-login.html"))

func renderRegistrationTemplate(w http.ResponseWriter, tmpl string, p *types.RegistrationPage) {
	err := registrationTemplate.ExecuteTemplate(w, tmpl+".html", p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}