package controllers

import (
	"errors"
	"html/template"
	"net/http"
	"strconv"

	"pokedex/models"
	"pokedex/services"
)

var (
	listTemplate   = template.Must(template.ParseFiles("templates/list.html"))
	detailTemplate = template.Must(template.ParseFiles("templates/detail.html"))
	errorTemplate  = template.Must(template.ParseFiles("templates/error.html"))
)

func ListPokemon(w http.ResponseWriter, r *http.Request) {
	pokemons, err := services.GetAllPokemon()
	if err != nil {
		showError(w, err)
		return
	}
	listTemplate.Execute(w, pokemons)
}

func DetailPokemon(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		showError(w, models.ErrBadRequest)
		return
	}

	pokemon, err := services.GetPokemon(id)
	if err != nil {
		showError(w, err)
		return
	}
	detailTemplate.Execute(w, pokemon)
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
	errorTemplate.Execute(w, err.Error())
}
