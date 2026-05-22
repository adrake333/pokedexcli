package pokeapi




import (
	"net/http"
	"fmt"
	"io"
)




func ListLocations(url *string) (LocationAreaResp, error) {
	endpoint := "https://pokeapi.co/api/v2/location-area"
	if url != nil {
		endpoint = *url
	}
	resp, err := http.Get(endpoint)
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
	return locations, nil
}

