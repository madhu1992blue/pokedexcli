package pokeapi

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
)


type PokemonResp struct {
	BaseExperience int `json:"base_experience"`
}

func (c *Client) GetPokemon(pageURL *string, pokemonName string) (*PokemonResp, error) {
	url := c.baseURL + "/pokemon/"
	if pageURL != nil {
		url = *pageURL
	}
	url += pokemonName
	var dataReader io.Reader
	if data, ok := c.cache.Get(url); ok {
		dataReader = bytes.NewReader(data)
	} else {
		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			return nil, err
		}
		resp, err := c.httpClient.Do(req)
		if err != nil {
			return nil, err
		}
		defer resp.Body.Close()
		dataBytes, err := io.ReadAll(resp.Body)
		if err != nil && err != io.EOF {
			return nil, err
		}
		dataReader = bytes.NewReader(dataBytes)
		c.cache.Add(url, dataBytes)

	}
	decoder := json.NewDecoder(dataReader)
	var pokemonResp PokemonResp
	if err := decoder.Decode(&pokemonResp); err != nil {
		return nil, err
	}
	return &pokemonResp, nil
}
