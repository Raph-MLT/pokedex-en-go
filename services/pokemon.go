package services

import (
	"pokedex/api"
	"pokedex/models"
)

func GetAllPokemon() ([]models.Pokemon, error) {
	pokemons, err := api.GetAllPokemon()
	if err != nil {
		return nil, err
	}

	var result []models.Pokemon
	for _, pokemon := range pokemons {
		if pokemon.PokedexID != 0 {
			result = append(result, pokemon)
		}
	}
	return result, nil
}

func GetPokemon(id int) (models.Pokemon, error) {
	if id < 1 {
		return models.Pokemon{}, models.ErrBadRequest
	}
	return api.GetPokemon(id)
}
