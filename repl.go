package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func startRepl() {
	locationApiURL := "https://pokeapi.co/api/v2/location-area/"
	reader := bufio.NewScanner(os.Stdin)
	cfg := &Config{
		NextLocationAreaURL: stringPtr(locationApiURL),
	}

	for {
		fmt.Print("Pokedex > ")
		reader.Scan()
		words := cleanInput(reader.Text())
		if len(words) == 0 {
			continue
		}
		cmdName := words[0]

		cmd, ok := getCommands()[cmdName]
		if ok {
			err := cmd.callback(cfg)
			if err != nil {
				fmt.Println(err)
			}
			continue
		} else {
			fmt.Println("Unknown command")
			continue
		}
	}
}

func stringPtr(s string) *string {
	return &s
}

func cleanInput(text string) []string {
	output := strings.ToLower(text)
	words := strings.Fields(output)
	return words
}

type cliCommand struct {
	name        string
	description string
	callback    func(*Config) error
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"map": {
			name:        "map",
			description: "Display next 20 location areas",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Display previous 20 location areas",
			callback:    commandMapBack,
		},
	}
}
