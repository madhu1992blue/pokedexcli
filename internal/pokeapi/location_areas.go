package pokeapi

import (
	"net/http"
	"encoding/json"
)

type locationAreas struct {
	Count int `json:"count"`
	Next *string `json:"next"`
	Previous *string `json:"previous"`
	Results []struct{
		Name string `json:"name"`
		Url string `json:"url"`
	} `json:"results"`

}
func (c *Client) GetLocationAreas(pageURL *string) (*locationAreas, error) {
	url := c.baseURL+"/location-area"
	if pageURL != nil {
		url = *pageURL
	}
	req, err:= http.NewRequest("GET",url,nil)
	if err != nil {
		return nil,err
	}
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	decoder := json.NewDecoder(resp.Body)
	var locationAreaResp locationAreas
	if err := decoder.Decode(&locationAreaResp); err != nil {
		return nil,err
	}
	return &locationAreaResp, nil
}
