package main

import (
	"fmt"
	"errors"
	"math/rand"
)

func catch_pokemon(cfg *config, args ...string) error {
	if len(args) == 0 {
		return errors.New("name required as argument for catching a pokemon")
	}
	pokemonName := args[0]
	fmt.Println("Throwing a Pokeball at " + pokemonName+"...")
	pokemonResp, err := cfg.Pokeclient.GetPokemon(nil, pokemonName)
	if err != nil {
		return err
	}
	if err != nil {
		return err
	}
	prob := float32(rand.Intn(pokemonResp.BaseExperience)) /float32(pokemonResp.BaseExperience)
	if prob > 0.5 {
		fmt.Println(pokemonName + " was caught!")
		fmt.Println("You may now inspect it with the inspect command.")
		cfg.Pokeclient.Pokedex[pokemonName] = *pokemonResp
	} else {
		fmt.Println(pokemonName + " escaped!")
	}
	return nil
}

