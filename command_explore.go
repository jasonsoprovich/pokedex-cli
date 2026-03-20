package main

import (
	"errors"
	"fmt"
)

func commandExplore(cfg *config, args []string) error {
	if len(args) != 1 {
		return errors.New("usage: explore <location-area-name>")
	}

	areaName := args[0]
	fmt.Printf("Exploring %s...\n", areaName)

	areaResp, err := cfg.pokeapiClient.GetLocationArea(areaName)
	if err != nil {
		return err
	}

	fmt.Println("Found Pokemon:")
	for _, encounter := range areaResp.PokemonEncounters {
		fmt.Printf(" - %s\n", encounter.Pokemon.Name)
	}

	return nil
}
