package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

func commandExit(cfg *Config) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp(cfg *Config) error {
	fmt.Println()
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println()

	for _, cmd := range getCommands() {
		fmt.Printf("%s: %s\n", cmd.name, cmd.description)
	}
	fmt.Println()
	return nil
}

func commandMap(cfg *Config) error {
	if cfg.NextLocationAreaURL == nil {
		return fmt.Errorf("no more location areas")
	}

	res, err := http.Get(*cfg.NextLocationAreaURL)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return err
	}

	if res.StatusCode > 299 {
		return fmt.Errorf("Response failed with status code: %d", res.StatusCode)
	}

	locationResponse := LocationAreaResponse{}
	err = json.Unmarshal(body, &locationResponse)
	if err != nil {
		return err
	}

	cfg.NextLocationAreaURL = locationResponse.Next
	cfg.PreviousLocationAreaURL = locationResponse.Previous

	for _, location := range locationResponse.Results {
		fmt.Println(location.Name)
	}

	return nil
}

func commandMapBack(cfg *Config) error {
	if cfg.PreviousLocationAreaURL == nil {
		return fmt.Errorf("you're on the first page")
	}

	res, err := http.Get(*cfg.PreviousLocationAreaURL)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return err
	}

	if res.StatusCode > 299 {
		return fmt.Errorf("Response failed with status code: %d", res.StatusCode)
	}

	locationResponse := LocationAreaResponse{}
	err = json.Unmarshal(body, &locationResponse)
	if err != nil {
		return err
	}

	cfg.NextLocationAreaURL = locationResponse.Next
	cfg.PreviousLocationAreaURL = locationResponse.Previous

	for _, location := range locationResponse.Results {
		fmt.Println(location.Name)
	}

	return nil
}
