package dreamhost

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

const (
	defaultURL = "https://api.dreamhost.com/"
)

func setup() (client *Client, mux *http.ServeMux, serverURL string, teardown func()) {
	mux = http.NewServeMux()

	apiHandler := http.NewServeMux()
	apiHandler.Handle("/", http.StripPrefix("/", mux))

	server := httptest.NewServer(apiHandler)

	client = NewClient("some-api-key", nil)
	url, _ := url.Parse(server.URL + "/")

	client.URL = url

	return client, mux, server.URL, server.Close
}

func testMethod(t *testing.T, r *http.Request, want string) {
	if got := r.Method; got != want {
		t.Errorf("Request method: %v, want %v", got, want)
	}
}

func TestNewClient(t *testing.T) {
	c := NewClient("some-api-key", nil)

	if got, want := c.URL.String(), defaultURL; got != want {
		t.Errorf("NewClient URL is %v, want %v", got, want)
	}
	if got, want := c.UserAgent, userAgent; got != want {
		t.Errorf("NewClient UserAgent is %v, want %v", got, want)
	}

	c2 := NewClient("some-api-key", nil)
	if c.client == c2.client {
		t.Error("NewClient returned same http.Clients, but they should differ")
	}
}
