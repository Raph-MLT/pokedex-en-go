package main

import (
	"fmt"
	"log"

	"pokedex/services"
)

func main() {
	pikachu, err := services.GetPokemon(25)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Pokémon n°25 :", pikachu.Name.Fr)

	pokemons, err := services.GetAllPokemon()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Nombre de Pokémon :", len(pokemons))

	_, err = services.GetPokemon(0)
	fmt.Println("Test id 0 :", err)
}
