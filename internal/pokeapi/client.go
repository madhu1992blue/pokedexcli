package pokeapi


import (
	"net/http"
	"time"
	"github.com/madhu1992blue/pokedexcli/internal/pokecache"
)
type Client struct {
	baseURL string
	httpClient *http.Client
	cache *pokecache.Cache
}

func NewClient(url string, timeout time.Duration) *Client {
	return &Client{
		baseURL: url,
		httpClient: &http.Client{
			Timeout: timeout,
		},
		cache: pokecache.NewCache(timeout),
	}
}