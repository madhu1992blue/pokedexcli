package main

import (
	"fmt"
	"errors"
)

func explore_area(cfg *config, args ...string) error {
	if len(args) == 0 {
		return errors.New("name required as argument for exploring an area")
	}
	areaName := args[0]
	locationAreasExplored, err := cfg.Pokeclient.ExploreLocationArea(nil, areaName)
	if err != nil {
		return err
	}
	if err != nil {
		return err
	}
	for _, encounter := range locationAreasExplored.PokemonEncounters {
		fmt.Println(encounter.Pokemon.Name)
	}

	return nil
}

