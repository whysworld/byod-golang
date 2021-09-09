package main

import (
	"html/template"
	// "io/ioutil"
	"log"
	"net/http"
	"regexp"
)

type HotSpotPage struct {
	Title          string
	WelcomeTitle   string
	WelcomeMessage string
	Content        string
}

type GuestInfo struct {
	Name         string
	Email        string
	Company      string
	SponsorEmail string
	Option1      string
	Option2      string
}

type RegistrationPage struct {
	Title          string
	WelcomeTitle   string
	WelcomeMessage string
	Content        string
	Information    GuestInfo
}

type SponsorPage struct {
	Title          string
	WelcomeTitle   string
	WelcomeMessage string
	Content        string
	Information    GuestInfo
	Status         string
	TimeLeft       string
}

var defaultInfo = GuestInfo{Name: "", Email: "", Company: "", SponsorEmail: "", Option1: "Other", Option2: "Other"}

func loadHotSpotHomePage() (*HotSpotPage, error) {
	title := "Acceptable Use Policy"
	welcomeTitle := "BYOD Welcome"
	welcomeMessage := "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Aliquam sed enim sollicitudin, venenatis massa sit amet, placerat neque. Proin blandit arcu."
	content := "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Aliquam sed enim sollicitudin, venenatis massa sit amet, placerat neque. Proin blandit arcu at metus bibendum, a fringilla purus feugiat. Mauris nec nulla auctor, porta enim et, eleifend magna. Phasellus blandit quam ut ante posuere tempus. Sed condimentum elementum libero, sit amet semper sapien tempus eget. Duis egestas vehicula semper. Nulla eleifend, felis tincidunt porttitor porta, quam lacus rhoncus ipsum.Lorem ipsum dolor sit amet, consectetur adipiscing elit. Aliquam sed enim sollicitudin, venenatis massa sit amet, placerat neque. Proin blandit arcu at metus bibendum, a fringilla purus feugiat. Mauris nec nulla auctor, porta enim et, eleifend magna. Phasellus blandit quam ut ante posuere tempus. Sed condimentum elementum libero, sit amet semper sapien tempus eget. Duis egestas vehicula semper. Nulla eleifend, felis tincidunt porttitor porta, quam lacus rhoncus ipsum.Lorem ipsum dolor sit amet, consectetur adipiscing elit. Aliquam sed enim sollicitudin, venenatis massa sit amet, placerat neque. Proin blandit arcu at metus bibendum, a fringilla purus feugiat. Mauris nec nulla auctor, porta enim et, eleifend magna. Phasellus blandit quam ut ante posuere tempus. Sed condimentum elementum libero, sit amet semper sapien tempus eget. Duis egestas vehicula semper. Nulla eleifend, felis tincidunt porttitor porta, quam lacus rhoncus ipsum"

	return &HotSpotPage{Title: title, WelcomeTitle: welcomeTitle, WelcomeMessage: welcomeMessage, Content: content}, nil
}

func loadHotSpotAcceptPage() (*HotSpotPage, error) {
	title := "Success"
	welcomeTitle := "Install"
	welcomeMessage := "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Aliquam sed enim sollicitudin, venenatis massa sit amet, placerat neque. Proin blandit arcu."
	content := ""
	return &HotSpotPage{Title: title, WelcomeTitle: welcomeTitle, WelcomeMessage: welcomeMessage, Content: content}, nil
}

func loadHotSpotDeclinePage() (*HotSpotPage, error) {
	title := "DECLINED"
	welcomeTitle := "Install"
	welcomeMessage := "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Aliquam sed enim sollicitudin, venenatis massa sit amet, placerat neque. Proin blandit arcu."
	content := ""
	return &HotSpotPage{Title: title, WelcomeTitle: welcomeTitle, WelcomeMessage: welcomeMessage, Content: content}, nil
}

func loadRegistrationHomePage() (*RegistrationPage, error) {
	title := "Acceptable Use Policy"
	welcomeTitle := "BYOD Welcome"
	welcomeMessage := "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Aliquam sed enim sollicitudin, venenatis massa sit amet, placerat neque. Proin blandit arcu."
	content := "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Aliquam sed enim sollicitudin, venenatis massa sit amet, placerat neque. Proin blandit arcu at metus bibendum, a fringilla purus feugiat. Mauris nec nulla auctor, porta enim et, eleifend magna. Phasellus blandit quam ut ante posuere tempus. Sed condimentum elementum libero, sit amet semper sapien tempus eget. Duis egestas vehicula semper. Nulla eleifend, felis tincidunt porttitor porta, quam lacus rhoncus ipsum.Lorem ipsum dolor sit amet, consectetur adipiscing elit. Aliquam sed enim sollicitudin, venenatis massa sit amet, placerat neque. Proin blandit arcu at metus bibendum, a fringilla purus feugiat. Mauris nec nulla auctor, porta enim et, eleifend magna. Phasellus blandit quam ut ante posuere tempus. Sed condimentum elementum libero, sit amet semper sapien tempus eget. Duis egestas vehicula semper. Nulla eleifend, felis tincidunt porttitor porta, quam lacus rhoncus ipsum.Lorem ipsum dolor sit amet, consectetur adipiscing elit. Aliquam sed enim sollicitudin, venenatis massa sit amet, placerat neque. Proin blandit arcu at metus bibendum, a fringilla purus feugiat. Mauris nec nulla auctor, porta enim et, eleifend magna. Phasellus blandit quam ut ante posuere tempus. Sed condimentum elementum libero, sit amet semper sapien tempus eget. Duis egestas vehicula semper. Nulla eleifend, felis tincidunt porttitor porta, quam lacus rhoncus ipsum"

	return &RegistrationPage{Title: title, WelcomeTitle: welcomeTitle, WelcomeMessage: welcomeMessage, Content: content, Information: defaultInfo}, nil
}

func loadRegistrationInfoPage() (*RegistrationPage, error) {
	title := "Guest Information"
	welcomeTitle := ""
	welcomeMessage := ""
	content := ""
	return &RegistrationPage{Title: title, WelcomeTitle: welcomeTitle, WelcomeMessage: welcomeMessage, Content: content, Information: defaultInfo}, nil
}

func loadRegistrationAcceptPage() (*RegistrationPage, error) {
	title := "Login"
	welcomeTitle := ""
	welcomeMessage := "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Aliquam sed enim sollicitudin, venenatis massa sit amet, placerat neque. Proin blandit arcu."
	content := ""
	return &RegistrationPage{Title: title, WelcomeTitle: welcomeTitle, WelcomeMessage: welcomeMessage, Content: content, Information: defaultInfo}, nil
}

func loadRegistrationLoginPage() (*RegistrationPage, error) {
	title := "Success"
	welcomeTitle := "Install"
	welcomeMessage := "Your password is in your email. If Company policy requires your sponsor to approve your access, email will be sent to you .. After that approval"
	content := ""
	return &RegistrationPage{Title: title, WelcomeTitle: welcomeTitle, WelcomeMessage: welcomeMessage, Content: content, Information: defaultInfo}, nil
}

func loadSponsorLoginPage() (*SponsorPage, error) {
	title := "LOGIN"
	welcomeTitle := ""
	welcomeMessage := "Login here to manage your sponsored guest users"
	content := ""
	return &SponsorPage{Title: title, WelcomeTitle: welcomeTitle, WelcomeMessage: welcomeMessage, Content: content, Information: defaultInfo, Status: "", TimeLeft: ""}, nil
}

func loadSponsorUsersPage() (*SponsorPage, error) {
	title := "GUEST USERS"
	welcomeTitle := ""
	welcomeMessage := ""
	content := ""
	info := GuestInfo{Name: "John Doe", Email: "John.doe@gmail.com", Company: "ABC Corp", SponsorEmail: "", Option1: "Here for Interview", Option2: "Representing self"}
	return &SponsorPage{Title: title, WelcomeTitle: welcomeTitle, WelcomeMessage: welcomeMessage, Content: content, Information: info, Status: "Admitted", TimeLeft: "2 hr 3m remaining"}, nil
}

func hotSpotHomePageHandler(w http.ResponseWriter, r *http.Request, title string) {
	p, err := loadHotSpotHomePage()
	print(p)
	if err != nil {
		http.Redirect(w, r, "/hotspot/home/", http.StatusFound)
		return
	}
	renderHotSpotTemplate(w, "hotspot-home", p)
}

func hotSpotAcceptPageHandler(w http.ResponseWriter, r *http.Request, title string) {
	p, err := loadHotSpotAcceptPage()
	print(p)
	if err != nil {
		http.Redirect(w, r, "/hotspot/accept/", http.StatusFound)
		return
	}
	renderHotSpotTemplate(w, "hotspot-accept", p)
}

func hotSpotDeclinePageHandler(w http.ResponseWriter, r *http.Request, title string) {
	p, err := loadHotSpotDeclinePage()
	print(p)
	if err != nil {
		http.Redirect(w, r, "/hotspot/decline/", http.StatusFound)
		return
	}
	renderHotSpotTemplate(w, "hotspot-decline", p)
}

func registratoinHomePageHandler(w http.ResponseWriter, r *http.Request, title string) {
	p, err := loadRegistrationHomePage()
	print(p)
	if err != nil {
		http.Redirect(w, r, "/registration/home/", http.StatusFound)
		return
	}
	renderRegistrationTemplate(w, "registration-home", p)
}

func registratoinInfoPageHandler(w http.ResponseWriter, r *http.Request, title string) {
	p, err := loadRegistrationInfoPage()
	print(p)
	if err != nil {
		http.Redirect(w, r, "/registration/info/", http.StatusFound)
		return
	}
	renderRegistrationTemplate(w, "registration-information", p)
}

func registratoinLoginPageHandler(w http.ResponseWriter, r *http.Request, title string) {
	p, err := loadRegistrationLoginPage()
	print(p)
	if err != nil {
		http.Redirect(w, r, "/registration/login/", http.StatusFound)
		return
	}
	renderRegistrationTemplate(w, "registration-login", p)
}

func registratoinAcceptPageHandler(w http.ResponseWriter, r *http.Request, title string) {
	p, err := loadRegistrationLoginPage()
	print(p)
	if err != nil {
		http.Redirect(w, r, "/registration/accept/", http.StatusFound)
		return
	}
	renderRegistrationTemplate(w, "registration-accept", p)
}

func sponsorLoginPageHandler(w http.ResponseWriter, r *http.Request, title string) {
	p, err := loadSponsorLoginPage()
	print(p)
	if err != nil {
		http.Redirect(w, r, "/sponsor/login/", http.StatusFound)
		return
	}
	renderSponsorTemplate(w, "sponsor-login", p)
}

func sponsorUsersPageHandler(w http.ResponseWriter, r *http.Request, title string) {
	p, err := loadSponsorUsersPage()
	print(p)
	if err != nil {
		http.Redirect(w, r, "/sponsor/users/", http.StatusFound)
		return
	}
	renderSponsorTemplate(w, "sponsor-users", p)
}

var templates = template.Must(template.ParseFiles("templates/hotspot/hotspot-accept.html", "templates/hotspot/hotspot-decline.html", "templates/hotspot/hotspot-home.html", "templates/registration/registration-home.html", "templates/registration/registration-information.html", "templates/registration/registration-login.html", "templates/registration/registration-accept.html", "templates/sponsor/sponsor-login.html", "templates/sponsor/sponsor-users.html"))

func renderHotSpotTemplate(w http.ResponseWriter, tmpl string, p *HotSpotPage) {
	err := templates.ExecuteTemplate(w, tmpl+".html", p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func renderRegistrationTemplate(w http.ResponseWriter, tmpl string, p *RegistrationPage) {
	err := templates.ExecuteTemplate(w, tmpl+".html", p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func renderSponsorTemplate(w http.ResponseWriter, tmpl string, p *SponsorPage) {
	err := templates.ExecuteTemplate(w, tmpl+".html", p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

var validPath = regexp.MustCompile("^/([a-zA-Z0-9]+)/(home|accept|decline|info|login|users)/$")

func makeHandler(fn func(http.ResponseWriter, *http.Request, string)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		m := validPath.FindStringSubmatch(r.URL.Path)
		if m == nil {
			http.NotFound(w, r)
			return
		}
		fn(w, r, m[2])
	}
}

func main() {
	http.HandleFunc("/hotspot/home/", makeHandler(hotSpotHomePageHandler))
	http.HandleFunc("/hotspot/accept/", makeHandler(hotSpotAcceptPageHandler))
	http.HandleFunc("/hotspot/decline/", makeHandler(hotSpotDeclinePageHandler))

	http.HandleFunc("/registration/home/", makeHandler(registratoinHomePageHandler))
	http.HandleFunc("/registration/info/", makeHandler(registratoinInfoPageHandler))
	http.HandleFunc("/registration/login/", makeHandler(registratoinLoginPageHandler))
	http.HandleFunc("/registration/accept/", makeHandler(registratoinAcceptPageHandler))

	http.HandleFunc("/sponsor/login/", makeHandler(sponsorLoginPageHandler))
	http.HandleFunc("/sponsor/users/", makeHandler(sponsorUsersPageHandler))

	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	log.Fatal(http.ListenAndServe(":8080", nil))
}
