package dreamhost

import (
	"net/http"
	"net/url"
)

const (
	defaultUrl = "https://api.dreamhost.com/"
	userAgent  = "go-dreamhost"
)

type Client struct {
	apiKey		string // API key used for communication with the API
	client		*http.Client // HTTP client used for communication with the API
	URL		*url.URL // URL of the API server
	UserAgent	string // User agent used for communication with the API
}

func NewClient(k string, httpClient *http.Client) *Client {
	if httpClient == nil {
		httpClient = &http.Client{}
	}

	url, _ := url.Parse(defaultUrl)

	c := &Client{apiKey: k, client: httpClient, URL: url, UserAgent: userAgent}

	return c
}
