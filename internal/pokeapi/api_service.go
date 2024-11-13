package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/imigrance/pokedexcli/pokedex"
)

func (c *Client) GetLocations(pageURL *string) (RespLocation, error) {
	url := baseURL + "/location-area"
	if pageURL != nil {
		url = *pageURL
	}

	data, ok := c.cache.Get(url)
	if ok {
		locations := RespLocation{}
		if err := json.Unmarshal(data, &locations); err != nil {
			return RespLocation{}, err
		}

		return locations, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return RespLocation{}, nil
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return RespLocation{}, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return RespLocation{}, err
	}

	c.cache.Add(url, body)

	locations := RespLocation{}
	if err := json.Unmarshal(body, &locations); err != nil {
		return RespLocation{}, err
	}

	return locations, nil
}

func (c *Client) GetInArea(area string) (RespLocationArea, error) {
	url := baseURL + "/location-area/" + area

	data, ok := c.cache.Get(url)
	if ok {
		areaResp := RespLocationArea{}
		if err := json.Unmarshal(data, &areaResp); err != nil {
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

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return RespLocationArea{}, err
	}

	c.cache.Add(url, body)

	areaResp := RespLocationArea{}
	if err := json.Unmarshal(body, &areaResp); err != nil {
		return RespLocationArea{}, err
	}

	return areaResp, nil
}

func (c *Client) GetPokemon(pokemonName string) (pokedex.Pokemon, error) {
	url := baseURL + "/pokemon/" + pokemonName

	data, ok := c.cache.Get(url)
	if ok {
		respPokemon := pokedex.Pokemon{}
		if err := json.Unmarshal(data, &respPokemon); err != nil {
			return pokedex.Pokemon{}, err
		}
		return respPokemon, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return pokedex.Pokemon{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return pokedex.Pokemon{}, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return pokedex.Pokemon{}, err
	}

	c.cache.Add(url, body)

	respPokemon := pokedex.Pokemon{}
	if err := json.Unmarshal(body, &respPokemon); err != nil {
		return pokedex.Pokemon{}, err
	}

	return respPokemon, nil
}
