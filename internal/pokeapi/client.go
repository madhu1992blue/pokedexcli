package pokeapi

import (
	"github.com/madhu1992blue/pokedexcli/internal/pokecache"
	"net/http"
	"time"
)

type Pokemon struct {
	Name string
}
type Client struct {
	baseURL    string
	httpClient *http.Client
	cache      *pokecache.Cache
	pokedex    map[string]Pokemon
}

func NewClient(url string, timeout, interval time.Duration) *Client {
	return &Client{
		baseURL: url,
		httpClient: &http.Client{
			Timeout: timeout,
		},
		cache: pokecache.NewCache(interval),
		pokedex: make(map[string]Pokemon),
	}
}
