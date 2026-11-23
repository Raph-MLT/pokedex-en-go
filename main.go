package main

import (
	"log"
	"net/http"

	"pokedex/controllers"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /{$}", controllers.ListPokemon)
	mux.HandleFunc("GET /pokemon/{id}", controllers.DetailPokemon)

	log.Println("Serveur lancé sur http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", mux))
}
