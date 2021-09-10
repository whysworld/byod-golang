package views

import (
	"html/template"
	"net/http"

	"whysworld.net/byod/types"
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
	p, err := LoadHotSpotHomePage()
	print(p)
	if err != nil {
		http.Redirect(w, r, "/hotspot/home/", http.StatusFound)
		return
	}
	RenderHotSpotTemplate(w, "hotspot-home", p)
}

func HotSpotAcceptPageHandler(w http.ResponseWriter, r *http.Request) {
	p, err := LoadHotSpotAcceptPage()
	print(p)
	if err != nil {
		http.Redirect(w, r, "/hotspot/accept/", http.StatusFound)
		return
	}
	RenderHotSpotTemplate(w, "hotspot-accept", p)
}

func HotSpotDeclinePageHandler(w http.ResponseWriter, r *http.Request) {
	p, err := LoadHotSpotDeclinePage()
	print(p)
	if err != nil {
		http.Redirect(w, r, "/hotspot/decline/", http.StatusFound)
		return
	}
	RenderHotSpotTemplate(w, "hotspot-decline", p)
}

var hotSpotTemplate = template.Must(template.ParseFiles("templates/hotspot/hotspot-accept.html", "templates/hotspot/hotspot-decline.html", "templates/hotspot/hotspot-home.html"))

func RenderHotSpotTemplate(w http.ResponseWriter, tmpl string, p *types.HotSpotPage) {
	err := hotSpotTemplate.ExecuteTemplate(w, tmpl+".html", p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
