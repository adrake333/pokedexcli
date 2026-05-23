package pokeapi

import (
	"encoding/json"
	"io"
)

func (c *Client) ListLocations(url *string) (LocationAreaResp, error) {
	endpoint := "https://pokeapi.co/api/v2/location-area"
	if url != nil {
		endpoint = *url
	}
	if val, ok := c.cache.Get(endpoint); ok {
		locations := LocationAreaResp{}
		err := json.Unmarshal(val, &locations)
		if err != nil {
			return LocationAreaResp{}, err
		}
		return locations, nil
	}
	resp, err := c.httpClient.Get(endpoint)
	if err != nil {
		return LocationAreaResp{}, err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return LocationAreaResp{}, err
	}
	var locations LocationAreaResp
	err = json.Unmarshal(body, &locations)
	if err != nil {
		return LocationAreaResp{}, err
	}
	c.cache.Add(endpoint, body)
	return locations, nil
}
