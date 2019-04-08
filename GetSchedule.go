package go_perform_football

type GetSchedule struct {
	Competition        Competition `json:"competition"`
	TournamentCalendar BaseSeason  `json:"tournamentCalendar"`
	MatchDate          []MatchDate `json:"matchDate"`
}

type MatchDate struct {
	Date          string  `json:"date"`
	NumberOfGames string  `json:"numberOfGames"`
	Match         []Match `json:"match"`
}

type Match struct {
	BaseMatch

	CoverageLevel string `json:"coverageLevel"`

	HomeContestantId string `json:"homeContestantId"`
	AwayContestantId string `json:"awayContestantId"`

	HomeContestantName string `json:"homeContestantName"`
	AwayContestantName string `json:"awayContestantName"`

	HomeContestantOfficialName string `json:"homeContestantOfficialName"`
	AwayContestantOfficialName string `json:"awayContestantOfficialName"`

	HomeContestantShortName string `json:"homeContestantShortName"`
	AwayContestantShortName string `json:"awayContestantShortName"`

	HomeContestantCode string `json:"homeContestantCode"`
	AwayContestantCode string `json:"awayContestantCode"`
}
