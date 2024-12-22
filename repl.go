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
	callback    func() error
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
	}
	return supportedCommands
}

func startRepl() {
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
		err := cliCommand.callback()
		if err != nil {
			os.Exit(1)
		}
	}
}
func cleanInput(text string) []string {
	return strings.Fields(strings.ToLower(text))
}
