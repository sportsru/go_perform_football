package go_perform_football

type GetMatch struct {
	Match []DetailMatch `json:"match"`
}

type MatchStat struct {
	DetailMatch
}

type DetailMatch struct {
	MatchInfo MatchInfo `json:"matchInfo"`
	LiveData  LiveData  `json:"liveData"`
}

type BaseEvent struct {
	ContestantId string  `json:"contestantId"`
	PeriodId     int     `json:"periodId"`
	TimeMin      int     `json:"timeMin"`
	TimeMinSec   *string `json:"timeMinSec"`
	LastUpdated  string  `json:"lastUpdated"`
	OptaEventId  string  `json:"optaEventId"`
}

type Goal struct {
	BaseEvent
	Type             string  `json:"type"`
	ScorerId         *string `json:"scorerId"`
	ScorerName       *string `json:"scorerName"`
	AssistPlayerId   *string `json:"assistPlayerId"`
	AssistPlayerName *string `json:"assistPlayerName"`
	HomeScore        *int    `json:"homeScore"`
	AwayScore        *int    `json:"awayScore"`
}

type Card struct {
	BaseEvent
	Type        string  `json:"type"`
	PlayerId    *string `json:"playerId"`
	PlayerName  *string `json:"playerName"`
	OptaEventId string  `json:"optaEventId"`
}

type Substitute struct {
	BaseEvent
	PlayerOnId    string `json:"playerOnId"`
	PlayerOnName  string `json:"playerOnName"`
	PlayerOffId   string `json:"playerOffId"`
	PlayerOffName string `json:"playerOffName"`
}

type BaseMatch struct {
	Id   string `json:"id"`
	Date string `json:"date"`
	Time string `json:"time"`
}

type LiveData struct {
	MatchDetails struct {
		PeriodId       int     `json:"periodId"`
		MatchStatus    string  `json:"matchStatus"`
		Winner         *string `json:"winner"`
		MatchLengthMin *int    `json:"matchLengthMin"`
		MatchLengthSec *int    `json:"matchLengthSec"`
		Period         []struct {
			Id        int    `json:"id"`
			Start     string `json:"start"`
			End       string `json:"end"`
			LengthMin int    `json:"lengthMin"`
			LengthSec int    `json:"lengthSec"`
		} `json:"period"`
		Scores map[string]struct {
			Home int `json:"home"`
			Away int `json:"away"`
		} `json:"scores"`
	} `json:"matchDetails"`
	Goal       []Goal       `json:"goal"`
	Card       []Card       `json:"card"`
	Substitute []Substitute `json:"substitute"`
	LineUp     []Lineup     `json:"lineUp"`
}

type Lineup struct {
	ContestantId string `json:"contestantId"`
	Player       []struct {
		PlayerId     string `json:"playerId"`
		FirstName    string `json:"firstName"`
		LastName     string `json:"lastName"`
		MatchName    string `json:"matchName"`
		ShirtNumber  int    `json:"shirtNumber"`
		Position     string `json:"position"`
		PositionSide string `json:"positionSide"`
		Stat         []struct {
			Type  string `json:"type"`
			Value string `json:"value"`
		} `json:"stat"`
	}
}

type Team struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

type Contestant struct {
	Team

	ShortName    string `json:"shortName"`
	OfficialName string `json:"officialName"`
	Code         string `json:"code"`
	Position     string `json:"position"`
	Country      struct {
		Id   string `json:"id"`
		Name string `json:"name"`
	} `json:"country"`
}

type Venue struct {
	Id        string `json:"id"`
	LongName  string `json:"longName"`
	ShortName string `json:"shortName"`
}

type MatchInfo struct {
	BaseMatch
	Week               string          `json:"week"`
	LastUpdated        string          `json:"lastUpdated"`
	Description        string          `json:"description"`
	Competition        BaseCompetition `json:"competition"`
	TournamentCalendar BaseSeason      `json:"tournamentCalendar"`
	Stage              struct {
		Id        string `json:"id"`
		FormatId  string `json:"formatId"`
		StartDate string `json:"startDate"`
		EndDate   string `json:"endDate"`
		Name      string `json:"name"`
	} `json:"stage"`
	Contestant []Contestant `json:"contestant"`
	Venue      Venue        `json:"venue"`
}
