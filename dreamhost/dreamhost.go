package dreamhost

import (
	"net/http"
	"net/url"
)

const (
	defaultURL = "https://api.dreamhost.com/"
	userAgent  = "go-dreamhost"
)

// A Client manages communication with the Dreamhost API.
type Client struct {
	apiKey		string // API key used for communication with the API
	client		*http.Client // HTTP client used for communication with the API
	URL		*url.URL // URL of the API server
	UserAgent	string // User agent used for communication with the API
}

// NewClient returns a new Dreamhost API client. No checks for a nil apiKey are
// made. If a nil httpClient is provided, then a new http.Client will be used.
func NewClient(apiKey string, httpClient *http.Client) *Client {
	if httpClient == nil {
		httpClient = &http.Client{}
	}

	url, _ := url.Parse(defaultURL)

	c := &Client{apiKey: apiKey, client: httpClient, URL: url, UserAgent: userAgent}

	return c
}
