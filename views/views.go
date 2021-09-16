package views

import (
	"net/http"
	"regexp"
	"whysworld.net/byod/types"
)
var DefaultInfo = types.GuestInfo{Name: "", Email: "", Company: "", SponsorEmail: "", Option1: "Other", Option2: "Other"}

var validPath = regexp.MustCompile("^/guestportal/([a-zA-Z0-9-_]+)/(home|accept|decline|info|login|users|logout)/$")

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
