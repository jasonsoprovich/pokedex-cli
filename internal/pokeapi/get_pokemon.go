package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) GetPokemon(name string) (RespPokemon, error) {
	url := fmt.Sprintf("%s/pokemon/%s/", baseURL, name)
	if val, ok := c.cache.Get(url); ok {
		pokemonResp := RespPokemon{}
		if err := json.Unmarshal(val, &pokemonResp); err != nil {
			return RespPokemon{}, err
		}
		return pokemonResp, nil
	}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return RespPokemon{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return RespPokemon{}, err
	}
	defer resp.Body.Close()

	dat, err := io.ReadAll(resp.Body)
	if err != nil {
		return RespPokemon{}, err
	}

	if resp.StatusCode > 299 {
		return RespPokemon{}, fmt.Errorf("response failed with status code: %d", resp.StatusCode)
	}

	pokemonResp := RespPokemon{}
	if err := json.Unmarshal(dat, &pokemonResp); err != nil {
		return RespPokemon{}, err
	}

	c.cache.Add(url, dat)
	return pokemonResp, nil
}
