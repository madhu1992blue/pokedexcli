package main

import (
	"fmt"
)

func view_pokedex(cfg *config, args ...string) error {
	fmt.Println("Your Pokedex:")
	for pokemonName := range cfg.Pokeclient.Pokedex {
		fmt.Println(" - "+pokemonName)
	}
	return nil
}

