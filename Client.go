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
)

type Client struct {
	token string
}

func (c *Client) Init(token string) {
	c.token = token
}

func (c *Client) createUrl(method string) string {
	return fmt.Sprintf("%ssoccerdata/%s/%s?_fmt=json", baseUrl, method, c.token)
}

func (c *Client) GetTournaments() (GetTournaments, error) {
	var (
		url  string
		res  GetTournaments
		err  error
		body []byte
	)

	url = c.createUrl(urlTournaments)

	fmt.Println(url)
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
