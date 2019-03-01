package go_perform_football

type BaseSeason struct {
	Id          string `json:"id"`
	Name        string `json:"name"`
	StartDate   string `json:"startDate"`
	EndDate     string `json:"endDate"`
	LastUpdated string `json:"lastUpdated"`
}

type Season struct {
	BaseSeason
	IncludesVenues    string `json:"includesVenues"`
	OcId              string `json:"ocId"`
	Active            string `json:"active"`
	IncludesStandings string `json:"includesStandings"`
}

type Country struct {
	Country     string `json:"country"`
	CountryId   string `json:"countryId"`
	CountryCode string `json:"countryCode"`
}

type Competition struct {
	Id                 string   `json:"id"`
	OcId               string   `json:"ocId"`
	Name               string   `json:"name"`
	CompetitionCode    string   `json:"competitionCode"`
	DisplayOrder       int      `json:"displayOrder"`
	IsFriendly         string   `json:"isFriendly"`
	CompetitionFormat  string   `json:"competitionFormat"`
	CompetitionType    string   `json:"competitionType"`
	Type               string   `json:"type"`
	TournamentCalendar []Season `json:"tournamentCalendar"`
	Country
}

type GetTournaments struct {
	Competition []Competition `json:"competition"`
}
