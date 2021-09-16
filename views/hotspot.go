package views

import (
	"html/template"
	"net/http"
	"whysworld.net/byod/types"
	"github.com/gorilla/mux"
	"log"
	"fmt"
)

func LoadHotSpotHomePage() (*types.HotSpotPage, error) {
	title := "Acceptable Use Policy"
	welcomeTitle := "BYOD Welcome"
	welcomeMessage := "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Aliquam sed enim sollicitudin, venenatis massa sit amet, placerat neque. Proin blandit arcu."
	content := "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Aliquam sed enim sollicitudin, venenatis massa sit amet, placerat neque. Proin blandit arcu at metus bibendum, a fringilla purus feugiat. Mauris nec nulla auctor, porta enim et, eleifend magna. Phasellus blandit quam ut ante posuere tempus. Sed condimentum elementum libero, sit amet semper sapien tempus eget. Duis egestas vehicula semper. Nulla eleifend, felis tincidunt porttitor porta, quam lacus rhoncus ipsum.Lorem ipsum dolor sit amet, consectetur adipiscing elit. Aliquam sed enim sollicitudin, venenatis massa sit amet, placerat neque. Proin blandit arcu at metus bibendum, a fringilla purus feugiat. Mauris nec nulla auctor, porta enim et, eleifend magna. Phasellus blandit quam ut ante posuere tempus. Sed condimentum elementum libero, sit amet semper sapien tempus eget. Duis egestas vehicula semper. Nulla eleifend, felis tincidunt porttitor porta, quam lacus rhoncus ipsum.Lorem ipsum dolor sit amet, consectetur adipiscing elit. Aliquam sed enim sollicitudin, venenatis massa sit amet, placerat neque. Proin blandit arcu at metus bibendum, a fringilla purus feugiat. Mauris nec nulla auctor, porta enim et, eleifend magna. Phasellus blandit quam ut ante posuere tempus. Sed condimentum elementum libero, sit amet semper sapien tempus eget. Duis egestas vehicula semper. Nulla eleifend, felis tincidunt porttitor porta, quam lacus rhoncus ipsum"

	return &types.HotSpotPage{Title: title, WelcomeTitle: welcomeTitle, WelcomeMessage: welcomeMessage, Content: content}, nil
}

func LoadHotSpotAcceptPage() (*types.HotSpotPage, error) {
	title := "Success"
	welcomeTitle := "Install"
	welcomeMessage := "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Aliquam sed enim sollicitudin, venenatis massa sit amet, placerat neque. Proin blandit arcu."
	content := ""
	return &types.HotSpotPage{Title: title, WelcomeTitle: welcomeTitle, WelcomeMessage: welcomeMessage, Content: content}, nil
}

func LoadHotSpotDeclinePage() (*types.HotSpotPage, error) {
	title := "DECLINED"
	welcomeTitle := "Install"
	welcomeMessage := "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Aliquam sed enim sollicitudin, venenatis massa sit amet, placerat neque. Proin blandit arcu."
	content := ""
	return &types.HotSpotPage{Title: title, WelcomeTitle: welcomeTitle, WelcomeMessage: welcomeMessage, Content: content}, nil
}

func HotSpotHomePageHandler(w http.ResponseWriter, r *http.Request) {
	//get params (portal_id)
	vars := mux.Vars(r)
	portal_id, ok := vars["portal_id"]
	if !ok {
		log.Print("portal_id is missing in parameters")
	}
	log.Print(`portal_id := `, portal_id)

	//get query
	log.Print("query: ", r.URL.Query())
	p, err := LoadHotSpotHomePage()
	print(p)
	if err != nil {
		http.Redirect(w, r, "/hotspot/home/", http.StatusFound)
		return
	}
	switch r.Method {
		case "GET":
			templateName := fmt.Sprintf("%s-home", portal_id)
			RenderHotSpotTemplate(w, templateName, p)
		case "POST":
			r.ParseForm()
			accept := r.Form.Get("accept")
			if accept == "accept" {
				redirectURI := fmt.Sprintf("/guestportal/%s/accept?accept=true", portal_id)
				http.Redirect(w, r, redirectURI, http.StatusFound)
				return
			} else {
				redirectURI := fmt.Sprintf("/guestportal/%s/decline?accept=false", portal_id)
				http.Redirect(w, r, redirectURI, http.StatusFound)
				return
			}
		default: 
			RenderHotSpotTemplate(w, "hotspot-home", p)
	}
	
}

func HotSpotAcceptPageHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	portal_id, ok := vars["portal_id"]
	if !ok {
		log.Print("portal_id is missing in parameters")
	}
	p, err := LoadHotSpotAcceptPage()
	if err != nil {
		return
	}
	templateName := fmt.Sprintf("%s-accept", portal_id)
	RenderHotSpotTemplate(w, templateName, p)
}

func HotSpotDeclinePageHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	portal_id, ok := vars["portal_id"]
	if !ok {
		log.Print("portal_id is missing in parameters")
	}
	p, err := LoadHotSpotDeclinePage()
	print(p)
	if err != nil {
		return
	}
	switch r.Method {
		case "GET":
			templateName := fmt.Sprintf("%s-decline", portal_id)
			RenderHotSpotTemplate(w, templateName, p)
		case "POST":
			r.ParseForm()
			action := r.Form.Get("action")
			var redirectURI string
			if action == "back" {
				redirectURI = fmt.Sprintf("/guestportal/%s/home?refresh=true", portal_id)
			} else {
				redirectURI = fmt.Sprintf("/guestportal/%s/decline?accept=false", portal_id)
			}
			http.Redirect(w, r, redirectURI, http.StatusFound)
		default: 
			RenderHotSpotTemplate(w, "hotspot-decline", p)
	}

}

var hotSpotTemplate = template.Must(template.ParseFiles("templates/hotspot/hotspot-accept.html", "templates/hotspot/hotspot-decline.html", "templates/hotspot/hotspot-home.html"))

func RenderHotSpotTemplate(w http.ResponseWriter, tmpl string, p *types.HotSpotPage) {
	log.Print("template: ", tmpl)
	err := hotSpotTemplate.ExecuteTemplate(w, tmpl+".html", p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
