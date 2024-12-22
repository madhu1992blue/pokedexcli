package main

import (
	"fmt"
	"bufio"
	"strings"
	"os"
)

func startRepl() {
	scanner := bufio.NewScanner(os.Stdin)
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
		fmt.Printf("Your command was: %s\n", cleanedWords[0])
	}
}
func cleanInput(text string) []string {
	return strings.Fields(strings.ToLower(text))
}
