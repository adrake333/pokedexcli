package pokeapi




import(
	"encoding/json"
	"fmt"
	"io"
)




func (c *Client) PokemonData(name string) (PokemonDataResp, error) {
	endpoint := "https://pokeapi.co/api/v2/pokemon/" + name
	if val, ok := c.cache.Get(endpoint); ok {
		data := PokemonDataResp{}
		err := json.Unmarshal(val, &data)
		if err != nil {
			return PokemonDataResp{}, err
		}
		return data, nil
	}
	resp, err := c.httpClient.Get(endpoint)
	if err != nil {
		return PokemonDataResp{}, err
	}
	defer resp.Body.Close()
	if resp.StatusCode > 399 {
		return PokemonDataResp{}, fmt.Errorf("bad status code: %v", resp.StatusCode)
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return PokemonDataResp{}, err
	}
	var data PokemonDataResp
	err = json.Unmarshal(body, &data)
	if err != nil {
		return PokemonDataResp{}, err
	}
	c.cache.Add(endpoint, body)

	return data, nil
}
