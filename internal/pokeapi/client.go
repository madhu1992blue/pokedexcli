package pokeapi

import (
	"github.com/madhu1992blue/pokedexcli/internal/pokecache"
	"net/http"
	"time"
)

type Client struct {
	baseURL    string
	httpClient *http.Client
	cache      *pokecache.Cache
	Pokedex    map[string]PokemonResp
}

func NewClient(url string, timeout, interval time.Duration) *Client {
	return &Client{
		baseURL: url,
		httpClient: &http.Client{
			Timeout: timeout,
		},
		cache: pokecache.NewCache(interval),
		Pokedex: make(map[string]PokemonResp),
	}
}
