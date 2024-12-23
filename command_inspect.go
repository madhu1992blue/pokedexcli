package main

import (
	"fmt"
	"errors"
)

func inspect_pokemon(cfg *config, args ...string) error {
	if len(args) == 0 {
		return errors.New("name required as argument for catching a pokemon")
	}
	pokemonName := args[0]
	pokemon, ok := cfg.Pokeclient.Pokedex[pokemonName]
	if !ok {
		fmt.Println("you have not caught that pokemon")
		return nil
	}
	fmt.Printf("Name: %s\n",pokemon.Name)
	fmt.Printf("Height: %d\n", pokemon.Height)
	fmt.Printf("Weight: %d\n",pokemon.Weight)
	fmt.Println("Stats:")
	for _, statInfo := range pokemon.Stats {
		fmt.Printf("\t-%s: %d\n",statInfo.Stat.Name,statInfo.BaseStat)
	}
	fmt.Println("Types:")
	for _, t := range pokemon.Types {
		fmt.Printf("\t- %s\n",t.Type.Name)
	}
	return nil
}

