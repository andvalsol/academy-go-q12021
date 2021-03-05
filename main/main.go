package main

import "fmt"

func main() {
	pokemonRepository := Repository{}

	pokemonUseCase := GetPokemonUseCase{pokemonRepository}

	pokemon, err := pokemonUseCase.GetPokemonById(1) // TODO accept user input as an ID instead of hardcoding it

	if err != nil {
		// Show the user an error
		fmt.Print(err.Error())
	} else {
		// Show the user the pokemon information
		fmt.Print(pokemon.Name)
	}
}
