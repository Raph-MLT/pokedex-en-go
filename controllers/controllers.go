package controllers

import (
	"errors"
	"net/http"

	"pokedex/helpers"
	"pokedex/models"
	"pokedex/services"
)

func ListPokemon(w http.ResponseWriter, r *http.Request) {
	pokemons, err := services.GetAllPokemon()
	if err != nil {
		showError(w, err)
		return
	}
	helpers.Render(w, "list.html", pokemons)
}

func DetailPokemon(w http.ResponseWriter, r *http.Request) {
	pokemon, err := services.GetPokemon(r.PathValue("id"))
	if err != nil {
		showError(w, err)
		return
	}
	helpers.Render(w, "pokemon.html", pokemon)
}

func showError(w http.ResponseWriter, err error) {
	status := http.StatusServiceUnavailable
	switch {
	case errors.Is(err, models.ErrNotFound):
		status = http.StatusNotFound
	case errors.Is(err, models.ErrBadRequest):
		status = http.StatusBadRequest
	case errors.Is(err, models.ErrTimeout):
		status = http.StatusGatewayTimeout
	}

	w.WriteHeader(status)
	helpers.Render(w, "error.html", err.Error())
}
