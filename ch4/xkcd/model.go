package xkcd

import (
	"encoding/json"
	"fmt"
	"net/http"
)

const (
	linkTemplate = "https://xkcd.com/%d"
	urlTemplate  = "https://xkcd.com/%d/info.0.json"
	currentUrl   = "https://xkcd.com/info.0.json"
)

type Comic struct {
	Num int

	Day   string
	Month string
	Year  string

	Title      string
	SafeTitle  string `json:"safe_title"`
	News       string
	Transcript string
	Alt        string

	Img string
}

// Link returns link to the comic
func (c *Comic) Link() string {
	return fmt.Sprintf(linkTemplate, c.Num)
}

type Client http.Client

var DefaultClient = &Client{}

// Current returns the current xkcd comic
func (c *Client) Current() (*Comic, error) {
	return c.getByUrl(currentUrl)
}

// Get returns the xkcd comic with the num number
func (c *Client) Get(num int) (*Comic, error) {
	url := fmt.Sprintf(urlTemplate, num)
	return c.getByUrl(url)
}

func (c *Client) getByUrl(url string) (*Comic, error) {
	httpClient := http.Client(*c)

	resp, err := httpClient.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("can't get comic: %s", resp.Status)
	}

	var comic Comic
	if err := json.NewDecoder(resp.Body).Decode(&comic); err != nil {
		return nil, err
	}

	return &comic, nil
}
