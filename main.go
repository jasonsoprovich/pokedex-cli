package main

import (
	"math/rand"
	"time"

	"github.com/jasonsoprovich/pokedex-cli/internal/pokeapi"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	pokeClient := pokeapi.NewClient(5*time.Second, time.Minute*5)
	cfg := &config{
		pokeapiClient: pokeClient,
		pokedex:       make(map[string]pokeapi.RespPokemon),
	}

	startRepl(cfg)
}
