package go_perform_football

type GetSquad struct {
	Squad []Squad `json:"squad"`
}

type Squad struct {
	ContestantId        string   `json:"contestantId"`
	ContestantName      string   `json:"contestantName"`
	ContestantShortName string   `json:"contestantShortName"`
	ContestantClubName  string   `json:"contestantClubName"`
	ContestantCode      string   `json:"contestantCode"`
	Type                string   `json:"type"`
	TeamType            string   `json:"teamType"`
	VenueName           string   `json:"venueName"`
	VenueId             string   `json:"venueId"`
	Person              []Person `json:"person"`
}

type Person struct {
	Id               string `json:"id"`
	FirstName        string `json:"firstName"`
	LastName         string `json:"lastName"`
	MatchName        string `json:"matchName"`
	Nationality      string `json:"nationality"`
	NationalityId    string `json:"nationalityId"`
	Position         string `json:"position"`
	Type             string `json:"type"`
	DateOfBirth      string `json:"dateOfBirth"`
	PlaceOfBirth     string `json:"placeOfBirth"`
	CountryOfBirth   string `json:"countryOfBirth"`
	CountryOfBirthId string `json:"countryOfBirthId"`
	Height           int    `json:"height"`
	Weight           int    `json:"weight"`
	Foot             string `json:"foot"`
	Status           string `json:"status"`
	Active           string `json:"active"`
}
