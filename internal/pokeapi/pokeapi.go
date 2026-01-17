package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type locationAreas struct {
	Count    int `json:"count"`
	Next     any `json:"next"`
	Previous any `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

func GetAreas(apiURL string) (locationAreas, error) {
	res, err := http.Get(apiURL)
	if err != nil {
		return locationAreas{}, fmt.Errorf("Unable to connect to PokeAPI: %w", err)
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
