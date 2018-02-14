package go_perform_football

type GetMatch struct {
	Match []DetailMatch `json:"match"`
}

type GetMatchStat struct {
	DetailMatch
}

type DetailMatch struct {
	MatchInfo MatchInfo `json:"matchInfo"`
	LiveData  LiveData  `json:"liveData"`
}

type Event struct {
	ContestantId string `json:"contestantId"`
	PeriodId     int    `json:"periodId"`
	TimeMin      int    `json:"timeMin"`
	LastUpdated  string `json:"lastUpdated"`
}

type Goal struct {
	Event
	Type             string `json:"type"`
	ScorerId         string `json:"scorerId"`
	ScorerName       string `json:"scorerName"`
	AssistPlayerId   string `json:"assistNameId"`
	AssistPlayerName string `json:"assistNameName"`
	OptaEventId      string `json:"optaEventId"`
	HomeScore        int    `json:"homeScore"`
	AwayScore        int    `json:"awayScore"`
}

type Card struct {
	Event
	Type        string `json:"type"`
	PlayerId    string `json:"playerId"`
	PlayerName  string `json:"playerName"`
	OptaEventId string `json:"optaEventId"`
}

type Substitute struct {
	Event
	PlayerOnId    string `json:"playerOnId"`
	PlayerOnName  string `json:"playerOnName"`
	PlayerOffId   string `json:"playerOffId"`
	PlayerOffName string `json:"playerOffName"`
}

type LiveData struct {
	MatchDetails struct {
		PeriodId       int    `json:"periodId"`
		MatchStatus    string `json:"matchStatus"`
		Winner         string `json:"winner"`
		MatchLengthMin int    `json:"matchLengthMin"`
		MatchLengthSec int    `json:"matchLengthSec"`
		Period         []struct {
			Id        int    `json:"id"`
			Start     string `json:"start"`
			End       string `json:"end"`
			LengthMin int    `json:"lengthMin"`
			LengthSec int    `json:"lengthSec"`
		} `json:"period"`
		Scores struct {
			Ht struct {
				Home int `json:"home"`
				Away int `json:"away"`
			} `json:"ht"`
			Ft struct {
				Home int `json:"home"`
				Away int `json:"away"`
			} `json:"ft"`
			Total struct {
				Home int `json:"home"`
				Away int `json:"away"`
			} `json:"total"`
		} `json:"scores"`
	} `json:"matchDetails"`
	Goal       []Goal       `json:"goal"`
	Card       []Card       `json:"card"`
	Substitute []Substitute `json:"substitute"`
	LineUp     []LineUp     `json:"lineUp"`
}

type LineUp struct {
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
	Position string `json:"position"`
	Country  struct {
		Id   string `json:"id"`
		Name string `json:""`
	} `json:"country"`
}

type Venue struct {
	Id        string `json:"id"`
	LongName  string `json:"longName"`
	ShortName string `json:"shortName"`
}

type MatchInfo struct {
	Id          string `json:"id"`
	Date        string `json:"date"`
	Time        string `json:"time"`
	Week        string `json:"week"`
	LastUpdated string `json:"lastUpdated"`
	Description string `json:"description"`
	//Competition Competition `json:"competition"`
	TournamentCalendar struct {
		Id        string `json:"id"`
		StartDate string `json:"startDate"`
		EndDate   string `json:"endDate"`
		Name      string `json:"name"`
	} `json:"tournamentCalendar"`
	Stage struct {
		Id        string `json:"id"`
		FormatId  string `json:"formatId"`
		StartDate string `json:"startDate"`
		EndDate   string `json:"endDate"`
		Name      string `json:"name"`
	} `json:"stage"`
	Contestant []Contestant `json:"contestant"`
	Venue      Venue        `json:"venue"`
}
