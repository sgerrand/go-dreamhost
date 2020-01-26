package dreamhost

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"

	"github.com/google/go-querystring/query"
	"github.com/google/uuid"
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

// Params specifies the optional parameters to API requests
type Params struct {
	apiKey   string `url:"key"`
	Command  string `url:"cmd"`
	UniqueID string `url:"unique_id"`

	// Optional parameters
	Format  string `url:"format,omitempty"`
	Account string `url:"account,omitempty"`
}

func addParams(s string, params interface{}) (string, error) {
	u, err := url.Parse(s)
	if err != nil {
		return s, err
	}

	qs, err := query.Values(params)
	if err != nil {
		return s, err
	}

	u.RawQuery = qs.Encode()

	return u.String(), nil
}

func newUniqueID() string {
	return uuid.New().String()
}

// NewRequest creates an API request. An API command can be passed in cmd.
func (c *Client) NewRequest(method, cmd string, body interface{}) (*http.Request, error) {
	if !strings.HasSuffix(c.URL.Path, "/") {
		return nil, fmt.Errorf("URL must have a trailing slash, but %q does not", c.URL)
	}

	if c.apiKey == "" {
		return nil, fmt.Errorf("An API key must be set to make API requests")
	}

	params := Params{apiKey: c.apiKey, Command: cmd, Format: "json", UniqueID: newUniqueID()}
	u, err := addParams(c.URL.String(), params)
	if err != nil {
		fmt.Println(fmt.Errorf("Error: %v", err))
		return nil, err
	}

	var buf io.ReadWriter
	if body != nil {
		buf = new(bytes.Buffer)
		enc := json.NewEncoder(buf)
		enc.SetEscapeHTML(false)
		err := enc.Encode(body)
		if err != nil {
			return nil, err
		}
	}

	req, err := http.NewRequest(method, u, buf)
	if err != nil {
		fmt.Println(fmt.Errorf("Error: %v", err))
		return nil, err
	}

	if body != nil {
		req.Header.Set("Content-Type", "application/json")
	}
	if c.UserAgent != "" {
		req.Header.Set("User-Agent", c.UserAgent)
	}

	return req, nil
}
