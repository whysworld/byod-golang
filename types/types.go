package types

type HotSpotPage struct {
	PortalID	   string
	Query		   string
	Title          string
	WelcomeTitle   string
	WelcomeMessage string
	Content        string
}

type GuestInfo struct {
	PortalID	   string
	Query		   string
	Name         string
	Email        string
	Company      string
	SponsorEmail string
	Option1      string
	Option2      string
}

type RegistrationPage struct {
	PortalID	   string
	Query		   string
	Title          string
	WelcomeTitle   string
	WelcomeMessage string
	Content        string
	Information    []GuestInfo
}

type SponsorPage struct {
	PortalID	   string
	Query		   string
	Title          string
	WelcomeTitle   string
	WelcomeMessage string
	Content        string
	Information    []GuestInfo
	Status         string
	TimeLeft       string
	ErrorMessage   string
}
