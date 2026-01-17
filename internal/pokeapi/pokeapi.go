package pokeapi

import (
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
		URL  string `json:"url"`
	} `json:"results"`
}

func (c *Client) GetAreas(pageURL *string) (locationAreas, error) {
	url := baseURL + "/location-area"
	if pageURL != nil {
		url = *pageURL
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return locationAreas{}, fmt.Errorf("Unable to connect to PokeAPI: %w", err)
	}

	res, err := c.httpClient.Do(req)
	if err != nil {
		return locationAreas{}, fmt.Errorf("Error obtaining PokeAPI response: %w", err)
	}

	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)
	if res.StatusCode > 299 {
		return locationAreas{}, fmt.Errorf("Response failed with status code: %d and\nbody: %s\n", res.StatusCode, data)
	}

	var areas locationAreas
	err = json.Unmarshal(data, &areas)
	if err != nil {
		return locationAreas{}, fmt.Errorf("Error unmarshalling JSON from API call: %w", err)
	}

	return areas, nil
}
