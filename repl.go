package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type cliCommand struct {
	name        string
	description string
	callback    func(cfg *config, args ...string) error
}

func getSupportedCommands() map[string]cliCommand {

	var supportedCommands map[string]cliCommand = map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"help": {
			name:        "help",
			description: "Displays the help message",
			callback:    commandHelp,
		},
		"map": {
			name:        "map",
			description: "Map Location area",
			callback:    map_area,
		},
		"mapb": {
			name:        "mapb",
			description: "Map back location area",
			callback:    map_back_area,
		},
		"explore": {
			name: "explore",
			description: "Explore with: explore <area_name>",
			callback: explore_area,
		},
		"catch": {
			name: "catch",
			description: "Try to catch a pokemon with: catch <pokemon>",
			callback: catch_pokemon,
		},
		"inspect": {
			name: "inspect",
			description: "Inspect a caught pokemon: inspect <pokemon>",
			callback: inspect_pokemon,
		},
		"pokedex": {
			name: "pokedex",
			description: "View your pokedex",
			callback: view_pokedex,
		},
	}
	return supportedCommands
}

func startRepl(cfg *config) {
	scanner := bufio.NewScanner(os.Stdin)
	supportedCommands := getSupportedCommands()
	for {
		fmt.Print("Pokedex > ")
		moreTokens := scanner.Scan()
		if !moreTokens {
			break
		}
		input := scanner.Text()
		cleanedWords := cleanInput(input)
		if len(cleanedWords) == 0 {
			continue
		}
		commandName := cleanedWords[0]

		cliCommand, ok := supportedCommands[commandName]

		if !ok {
			fmt.Println("Unknown command")
			continue
		}
		err := cliCommand.callback(cfg, cleanedWords[1:]...)
		if err != nil {
			os.Exit(1)
		}
	}
}
func cleanInput(text string) []string {
	return strings.Fields(strings.ToLower(text))
}
