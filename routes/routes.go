package routes

import (
	"net/http"

	"pokedex/controllers"
)

func New() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /{$}", controllers.ListPokemon)
	mux.HandleFunc("GET /pokemon/{id}", controllers.DetailPokemon)
	return mux
}
