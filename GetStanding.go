package go_perform_football

type GetStanding struct {
	Stage []Stage `json:"stage"`
}

type Stage struct {
	Id        string     `json:"id"`
	FormatId  string     `json:"formatId"`
	Name      string     `json:"name"`
	StartDate string     `json:"startDate"`
	EndDate   string     `json:"endDate"`
	Division  []Division `json:"division"`
}

type Division struct {
	Type    string    `json:"type"`
	Ranking []Ranking `json:"ranking"`
}

type Ranking struct {
	Rank                int    `json:"rank"`
	RankStatus          string `json:"rankStatus"`
	RankId              string `json:"rankId"`
	LastRank            int    `json:"lastRank"`
	ContestantId        string `json:"contestantId"`
	ContestantName      string `json:"contestantName"`
	ContestantShortName string `json:"contestantShortName"`
	ContestantClubName  string `json:"contestantClubName"`
	ContestantCode      string `json:"contestantCode"`
	Points              int    `json:"points"`
	MatchesPlayed       int    `json:"matchesPlayed"`
	MatchesWon          int    `json:"matchesWon"`
	MatchesLost         int    `json:"matchesLost"`
	MatchesDrawn        int    `json:"matchesDrawn"`
	GoalsFor            int    `json:"goalsFor"`
	GoalsAgainst        int    `json:"goalsAgainst"`
	Goaldifference      string `json:"goaldifference"`
}
