package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func cleanInput(text string) []string {
	output := strings.ToLower(text)
	words := strings.Fields(output)
	return words
}

func startRepl() {
	scanner := bufio.NewScanner(os.Stdin)
	prompt := "Pokedex > "

	for {
		fmt.Printf("%v", prompt)
		scanner.Scan()
		input := scanner.Text()
		words := cleanInput(input)
		if len(words) == 0 {
			continue
		}
		fmt.Printf("Your command was: %s\n", words[0])
	}
}
