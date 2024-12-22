package pokeapi


import (
	"net/http"
	"time"
)
type Client struct {
	baseURL string
	httpClient *http.Client
}

func NewClient(url string, timeout time.Duration) *Client {
	return &Client{
		baseURL: url,
		httpClient: &http.Client{
			Timeout: timeout,
		},

	}
}
