package go_perform_football

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)

const baseUrl = "http://api.performfeeds.com/"

const (
	urlTournaments = "tournamentcalendar"
	urlTeam        = "squads"
	urlSchedule    = "tournamentschedule"
	urlMatch       = "match"
	urlStandings   = "standings"
	urlMatchStats  = "matchstats"
)

type Client struct {
	token string
}

func (c *Client) Init(token string) {
	c.token = token
}

func (c *Client) GetSquad(teamId string) (GetSquad, error) {

	var (
		url  string
		res  GetSquad
		err  error
		body []byte
	)

	url = fmt.Sprintf("%ssoccerdata/%s/%s?ctst=%s&_fmt=json&_rt=b&detailed=yes", baseUrl, urlTeam, "1vmmaetzoxkgg1qf6pkpfmku0k", teamId)
	fmt.Println(url)

	if body, err = c.getUrl(url); err != nil {
		return GetSquad{}, err
	}

	if err = json.Unmarshal(body, &res); err != nil {
		return GetSquad{}, err
	}

	return res, nil
}

func (c *Client) GetStanding(seasonId string) (GetStanding, error) {

	var (
		url  string
		res  GetStanding
		err  error
		body []byte
	)

	url = fmt.Sprintf("%ssoccerdata/%s/%s?tmcl=%s&_fmt=json&_rt=b&type=total", baseUrl, urlStandings, "1vmmaetzoxkgg1qf6pkpfmku0k", seasonId)

	fmt.Println(url)
	if body, err = c.getUrl(url); err != nil {
		return GetStanding{}, err
	}

	if err = json.Unmarshal(body, &res); err != nil {
		return GetStanding{}, err
	}

	return res, nil

}

func (c *Client) GetMatchStat(matchId string) (MatchStat, error) {

	var (
		url  string
		res  MatchStat
		err  error
		body []byte
	)

	url = fmt.Sprintf("%ssoccerdata/%s/%s/%s?_fmt=json&_rt=b&detailed=yes", baseUrl, urlMatchStats, "1vmmaetzoxkgg1qf6pkpfmku0k", matchId)

	fmt.Println(url)
	if body, err = c.getUrl(url); err != nil {
		return MatchStat{}, err
	}

	if err = json.Unmarshal(body, &res); err != nil {
		return MatchStat{}, err
	}

	return res, nil
}

func (c *Client) GetMatch(matchId string) (GetMatch, error) {

	var (
		url  string
		res  GetMatch
		err  error
		body []byte
	)

	url = fmt.Sprintf("%ssoccerdata/%s/%s?tmcl=%s&_fmt=json&live=yes", baseUrl, urlMatch, "1vmmaetzoxkgg1qf6pkpfmku0k", matchId)

	fmt.Println(url)
	if body, err = c.getUrl(url); err != nil {
		return GetMatch{}, err
	}

	if err = json.Unmarshal(body, &res); err != nil {
		return GetMatch{}, err
	}

	return res, nil
}

func (c *Client) GetSchedule(tournamentId string) (GetSchedule, error) {
	var (
		url  string
		res  GetSchedule
		err  error
		body []byte
	)

	url = fmt.Sprintf("%ssoccerdata/%s/%s/%s?_fmt=json", baseUrl, urlSchedule, c.token, tournamentId)

	fmt.Println(url)

	if body, err = c.getUrl(url); err != nil {
		return GetSchedule{}, err
	}

	if err = json.Unmarshal(body, &res); err != nil {
		return GetSchedule{}, err
	}

	return res, nil
}

func (c *Client) GetTournaments() (GetTournaments, error) {
	var (
		url  string
		res  GetTournaments
		err  error
		body []byte
	)

	url = fmt.Sprintf("%ssoccerdata/%s/%s?_fmt=json", baseUrl, urlTournaments, c.token)

	if body, err = c.getUrl(url); err != nil {
		return GetTournaments{}, err
	}

	if err = json.Unmarshal(body, &res); err != nil {
		return GetTournaments{}, err
	}

	return res, nil
}

func (c *Client) getUrl(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != 200 {
		return nil, errors.New(fmt.Sprintf("API refused with code %v", resp.Status))
	}

	return body, nil
}
