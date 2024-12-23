package pokeapi

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type locationAreas struct {
	Count    int     `json:"count"`
	Next     *string `json:"next"`
	Previous *string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		Url  string `json:"url"`
	} `json:"results"`
}

type locationAreaExplored struct {
	PokemonEncounters []struct{
		Pokemon struct {
			Name string `json:"name"`
		} `json:"pokemon"`

	} `json:"pokemon_encounters"`
}
func (c *Client) GetLocationAreas(pageURL *string) (*locationAreas, error) {
	url := c.baseURL + "/location-area"
	if pageURL != nil {
		url = *pageURL
	}
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
	var locationAreaResp locationAreas
	if err := decoder.Decode(&locationAreaResp); err != nil {
		fmt.Println("Here")
		return nil, err
	}
	return &locationAreaResp, nil
}

func (c *Client) ExploreLocationArea(pageURL *string, areaName string) (*locationAreaExplored, error) {
	url := c.baseURL + "/location-area/"
	if pageURL != nil {
		url = *pageURL
	}
	url += areaName
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
	var locationAreaExploredResp locationAreaExplored
	if err := decoder.Decode(&locationAreaExploredResp); err != nil {
		return nil, err
	}
	return &locationAreaExploredResp, nil
}
