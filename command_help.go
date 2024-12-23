package main

import (
	"fmt"
)

func commandHelp(cfg *config) error {
	fmt.Print(`Welcome to the Pokedex!
Usage:

`)
	for _, cmd := range getSupportedCommands() {
		fmt.Printf("%s: %s\n", cmd.name, cmd.description)
	}
	return nil
}
