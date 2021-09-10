package views

import (
	"log"
	"net/http"
	"regexp"
	"whysworld.net/byod/types"
)
var DefaultInfo = types.GuestInfo{Name: "", Email: "", Company: "", SponsorEmail: "", Option1: "Other", Option2: "Other"}

var validPath = regexp.MustCompile("^/([a-zA-Z0-9]+)/(home|accept|decline|info|login|users|logout)/$")

func MakeHandler(fn func(http.ResponseWriter, *http.Request)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		m := validPath.FindStringSubmatch(r.URL.Path)
		if m == nil {
			http.NotFound(w, r)
			return
		}
		fn(w, r)
	}
}

var validPathForSponsorLogin = regexp.MustCompile("^/([a-zA-Z0-9]+)/sponsor/login/$")

func MakeSponsorLoginHandler(fn func(http.ResponseWriter, *http.Request)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Print("Paht: ", r.URL.Path)
		m := validPathForSponsorLogin.FindStringSubmatch(r.URL.Path)
		if m == nil {
			http.NotFound(w, r)
			return
		}
		fn(w, r)
	}
}