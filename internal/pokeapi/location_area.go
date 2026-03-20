package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) GetLocationArea(areaName string) (RespLocationArea, error) {
	url := fmt.Sprintf("%s/location-area/%s/", baseURL, areaName)

	if val, ok := c.cache.Get(url); ok {
		areaResp := RespLocationArea{}
		err := json.Unmarshal(val, &areaResp)
		if err != nil {
			return RespLocationArea{}, err
		}
		return areaResp, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return RespLocationArea{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return RespLocationArea{}, err
	}
	defer resp.Body.Close()

	dat, err := io.ReadAll(resp.Body)
	if err != nil {
		return RespLocationArea{}, err
	}

	if resp.StatusCode > 299 {
		return RespLocationArea{}, fmt.Errorf("response failed with a status code: %d", resp.StatusCode)
	}

	areaResp := RespLocationArea{}
	err = json.Unmarshal(dat, &areaResp)
	if err != nil {
		return RespLocationArea{}, err
	}

	c.cache.Add(url, dat)
	return areaResp, nil
}
