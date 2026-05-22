package pokeapi




import (
	"encoding/json"
	"io"
)




func (c *Client) LocationDetails(area string) (LocationDetailsResp, error) {
	endpoint := "https://pokeapi.co/api/v2/location-area/" + area
	if val, ok := c.cache.Get(endpoint); ok {
		details := LocationDetailsResp{}
		err := json.Unmarshal(val, &details)
		if err != nil {
			return LocationDetailsResp{}, err
		}
		return details, nil
	}
	resp, err := c.httpClient.Get(endpoint)
	if err != nil {
		return LocationDetailsResp{}, err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return LocationDetailsResp{}, err
	}
	var details LocationDetailsResp
	err = json.Unmarshal(body, &details)
	if err != nil {
		return LocationDetailsResp{}, err
	}
	c.cache.Add(endpoint, body)
	return details, nil
}
