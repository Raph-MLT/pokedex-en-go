package routes

import (
	"net/http"

	"pokedex/controllers"
)

func New() *http.ServeMux {
	mux := http.NewServeMux()
	mux.Handle("GET /{$}", http.RedirectHandler("/pokemon", http.StatusMovedPermanently))
	mux.HandleFunc("GET /pokemon", controllers.ListPokemon)
	mux.HandleFunc("GET /pokemon/{id}", controllers.DetailPokemon)
	mux.Handle("GET /static/", http.StripPrefix("/static/", http.FileServer(http.Dir("assets"))))
	mux.HandleFunc("/", controllers.NotFound)
	return mux
}
