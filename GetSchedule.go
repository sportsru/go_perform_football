package go_perform_football

type GetSchedule struct {
	Competition        Competition `json:"competition"`
	TournamentCalendar Season      `json:"tournamentCalendar"`
	MatchDate          []MatchDate `json:"matchDate"`
}

type MatchDate struct {
	Date          string  `json:"date"`
	NumberOfGames string  `json:"numberOfGames"`
	Match         []Match `json:"match"`
}

type Match struct {
	Id                 string `json:"id"`
	CoverageLevel      string `json:"coverageLevel"`
	Date               string `json:"date"`
	Time               string `json:"time"`
	HomeContestantId   string `json:"homeContestantId"`
	AwayContestantId   string `json:"awayContestantId"`
	HomeContestantName string `json:"homeContestantName"`
	AwayContestantName string `json:"awayContestantName"`
}
