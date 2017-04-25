package api

import (
	"encoding/json"
	"fmt"
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
	ResultCount int64   `json:"results"`
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



// LookupEmail returns the definitions from Urban Dictionary for a given word.
func LookupEmail(word string) (response *Response, err error) {
	endpoint := fmt.Sprintf("%s/api?q=%s", hackedemailsAPIURI, url.QueryEscape(word))
	client := &http.Client{
		Timeout: 15 * time.Second,
	}
	resp, err := client.Get(endpoint)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	dec := json.NewDecoder(resp.Body)
	if err = dec.Decode(&response); err != nil {
		return nil, err
	}

	return response, nil
}
