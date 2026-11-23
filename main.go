package main

import (
	"fmt"
	"log"

	"pokedex/api"
)

func main() {
	pikachu, err := api.GetPokemon(25)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Pokémon n°25 :", pikachu.Name.Fr)

	pokemons, err := api.GetAllPokemon()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Nombre de Pokémon :", len(pokemons))
}
