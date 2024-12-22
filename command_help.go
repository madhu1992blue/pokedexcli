package main

import (
	"fmt"
)

func commandHelp() error {
	fmt.Println(`Welcome to the Pokedex!
Usage:
`)
	for _, cmd := range getSupportedCommands() {
		fmt.Printf("%s: %s\n", cmd.name, cmd.description)
	}
	return nil
}
