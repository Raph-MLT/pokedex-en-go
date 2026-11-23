package api

import (
	"encoding/json"
	"errors"
	"fmt"
	"net"
	"net/http"
	"time"

	"pokedex/models"
)

const baseURL = "https://tyradex.app/api/v1"

var client = &http.Client{Timeout: 5 * time.Second}

func GetAllPokemon() ([]models.Pokemon, error) {
	var pokemons []models.Pokemon
	if err := fetch(baseURL+"/pokemon", &pokemons); err != nil {
		return nil, err
	}
	return pokemons, nil
}

func GetPokemon(id int) (models.Pokemon, error) {
	var pokemon models.Pokemon
	if err := fetch(fmt.Sprintf("%s/pokemon/%d", baseURL, id), &pokemon); err != nil {
		return models.Pokemon{}, err
	}
	if pokemon.PokedexID == 0 {
		return models.Pokemon{}, models.ErrNotFound
	}
	return pokemon, nil
}

func fetch(url string, target any) error {
	request, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return models.ErrUnavailable
	}
	request.Header.Set("User-Agent", "pokedex-ynov/1.0")

	response, err := client.Do(request)
	if err != nil {
		var netErr net.Error
		if errors.As(err, &netErr) && netErr.Timeout() {
			return models.ErrTimeout
		}
		return models.ErrUnavailable
	}
	defer response.Body.Close()

	switch response.StatusCode {
	case http.StatusOK:
	case http.StatusNotFound:
		return models.ErrNotFound
	default:
		return models.ErrUnavailable
	}

	if err := json.NewDecoder(response.Body).Decode(target); err != nil {
		return models.ErrUnavailable
	}
	return nil
}
