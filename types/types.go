package types

type NullString struct {
    String string
    Valid  bool // Valid is true if String is not NULL
}

type HotSpotPage struct {
	PortalID	   string
	Query		   string
	Title		   string
	SubTitle       string
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
	Status		   string
	CreatedAt      string
}

type RegistrationPage struct {
	PortalID	   string
	Query		   string
	Title          string
	SubTitle       string
	WelcomeTitle   string
	WelcomeMessage string
	Content        string
	Information    []GuestInfo
}

type SponsorPage struct {
	PortalID	   string
	Query		   string
	Title          string
	SubTitle       string
	WelcomeTitle   string
	WelcomeMessage string
	Content        string
	Information    []GuestInfo
	Status         string
	TimeLeft       string
	ErrorMessage   string
}

type User struct {
	Name 	       string
	Email	       string
	Password       string
	Company		   string
	Option1		   string
	Option2        string
	Role           string
	CreatedBy      string
	Status		   string
	CreatedAt      string
}
