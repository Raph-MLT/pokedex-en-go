package services

import (
	"strconv"
	"strings"

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

func SearchPokemon(nom string) ([]models.Pokemon, error) {
	pokemons, err := GetAllPokemon()
	if err != nil {
		return nil, err
	}

	nom = strings.ToLower(strings.TrimSpace(nom))
	var result []models.Pokemon
	for _, pokemon := range pokemons {
		if strings.Contains(strings.ToLower(pokemon.Name.Fr), nom) {
			result = append(result, pokemon)
		}
	}
	return result, nil
}

func GetPokemon(idText string) (models.Pokemon, error) {
	id, err := strconv.Atoi(idText)
	if err != nil || id < 1 {
		return models.Pokemon{}, models.ErrBadRequest
	}
	return api.GetPokemon(id)
}
