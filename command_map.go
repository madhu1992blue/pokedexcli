package main

import (
	"fmt"
)

func map_area(cfg *config, args ...string) error {
	locationAreas, err := cfg.pokeclient.GetLocationAreas(cfg.NextLocationAreaURL)
	if err != nil {
		return err
	}
	cfg.NextLocationAreaURL = locationAreas.Next
	cfg.PreviousLocationAreaURL = locationAreas.Previous
	if err != nil {
		return err
	}
	for _, area := range locationAreas.Results {
		fmt.Println(area.Name)
	}

	return nil
}

func map_back_area(cfg *config, args ...string) error {
	if cfg.PreviousLocationAreaURL == nil {
		fmt.Println("you're on the first page")
	}
	locationAreas, err := cfg.pokeclient.GetLocationAreas(cfg.PreviousLocationAreaURL)
	cfg.NextLocationAreaURL = locationAreas.Next
	cfg.PreviousLocationAreaURL = locationAreas.Previous
	if err != nil {
		return err
	}
	for _, area := range locationAreas.Results {
		fmt.Println(area.Name)
	}
	return nil
}
