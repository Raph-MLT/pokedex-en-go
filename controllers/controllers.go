package controllers

import (
	"errors"
	"log"
	"net/http"

	"pokedex/helpers"
	"pokedex/models"
	"pokedex/services"
)

func ListPokemon(w http.ResponseWriter, r *http.Request) {
	pokemons, err := services.GetAllPokemon()
	if err != nil {
		showError(w, r, err)
		return
	}
	helpers.Render(w, "list.html", pokemons)
}

func DetailPokemon(w http.ResponseWriter, r *http.Request) {
	pokemon, err := services.GetPokemon(r.PathValue("id"))
	if err != nil {
		showError(w, r, err)
		return
	}
	helpers.Render(w, "pokemon.html", pokemon)
}

func NotFound(w http.ResponseWriter, r *http.Request) {
	log.Printf("%s : page inconnue", r.URL.Path)
	showPage(w, http.StatusNotFound, "Cette page n'existe pas.")
}

func showError(w http.ResponseWriter, r *http.Request, err error) {
	log.Printf("%s : %v", r.URL.Path, err)

	switch {
	case errors.Is(err, models.ErrBadRequest):
		showPage(w, http.StatusBadRequest, "Numéro de Pokémon invalide.")
	case errors.Is(err, models.ErrNotFound):
		showPage(w, http.StatusNotFound, "Ce Pokémon n'existe pas.")
	case errors.Is(err, models.ErrTimeout):
		showPage(w, http.StatusGatewayTimeout, "Le service met trop de temps à répondre.")
	default:
		showPage(w, http.StatusServiceUnavailable, "Service momentanément indisponible.")
	}
}

func showPage(w http.ResponseWriter, status int, message string) {
	w.WriteHeader(status)
	helpers.Render(w, "error.html", message)
}
