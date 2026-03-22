package main

import (
	"fmt"
	"sort"
)

func commandPokedex(cfg *config, args []string) error {
	if len(args) != 0 {
		return fmt.Errorf("pokedex does not take any arguments")
	}

	fmt.Println("Your Pokedex:")

	caughtPokemon := make([]string, 0, len(cfg.pokedex))

	for p := range cfg.pokedex {
		caughtPokemon = append(caughtPokemon, p)
	}
	sort.Strings(caughtPokemon)
	for _, p := range caughtPokemon {
		fmt.Printf(" - %s\n", p)
	}

	return nil
}
