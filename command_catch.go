package main

import (
	"errors"
	"fmt"
	"math/rand"
)

func commandCatch(cfg *config, args []string) error {
	if len(args) != 1 {
		return errors.New("usage: catch <pokemon-name>")
	}

	name := args[0]
	fmt.Printf("Throwing a Pokeball at %s...\n", name)

	pokemon, err := cfg.pokeapiClient.GetPokemon(name)
	if err != nil {
		return err
	}

	// rand.Seed(time.Now().UnixNano())

	chance := rand.Intn(pokemon.BaseExperience + 1)

	if chance > pokemon.BaseExperience/2 {
		fmt.Printf("%s escaped!\n", pokemon.Name)
		return nil
	}

	cfg.pokedex[pokemon.Name] = pokemon
	fmt.Printf("%s was caught!\n", pokemon.Name)
	return nil
}
