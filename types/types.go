package types

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
	ErrorMessage   string
}
