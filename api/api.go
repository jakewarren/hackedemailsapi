package api

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"
)

const (
	hackedemailsAPIURI = "https://hacked-emails.com"
)

// Response describes the data structure that comes back from the
// hacked-emails API.
type Response struct {
	Status      string   `json:"status"`
	Query       string   `json:"query"`
	ResultCount int64    `json:"results"`
	Breaches    []Breach `json:"data"`
}

// Breach contains the details for each breach
type Breach struct {
	Author         string `json:"author"`
	DateCreated    string `json:"date_created"`
	DateLeaked     string `json:"date_leaked"`
	Details        string `json:"details"`
	EmailsCount    int64  `json:"emails_count"`
	SourceLines    int64  `json:"source_lines"`
	SourceNetwork  string `json:"source_network"`
	SourceProvider string `json:"source_provider"`
	SourceSize     int64  `json:"source_size"`
	SourceURL      string `json:"source_url"`
	Title          string `json:"title"`
	Verified       bool   `json:"verified"`
}

// LookupEmail returns the breach data from hacked-emails.com for a given email
func LookupEmail(email string) (response *Response, err error) {
	endpoint := fmt.Sprintf("%s/api?q=%s", hackedemailsAPIURI, url.QueryEscape(email))
	client := &http.Client{
		Timeout: 15 * time.Second,
	}
	resp, err := client.Get(endpoint)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)

	dec := json.NewDecoder(resp.Body)
	if err = dec.Decode(&response); err != nil {
		//attempt to decode status and data field to check for error message
		var objmap map[string]*json.RawMessage

		json.Unmarshal(body, &objmap)

		var status, data string
		jsonerr := json.Unmarshal(*objmap["status"], &status)
		jsonerr = json.Unmarshal(*objmap["data"], &data)

		if jsonerr == nil {
			return nil, fmt.Errorf("status: %s - data: %s", status, data)
		}

		return nil, err
	}

	return response, nil
}
