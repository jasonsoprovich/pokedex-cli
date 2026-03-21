package pokeapi

const (
	baseURL = "https://pokeapi.co/api/v2"
)

type RespLocationArea struct {
	ID                int                `json:"id"`
	Name              string             `json:"name"`
	PokemonEncounters []PokemonEncounter `json:"pokemon_encounters"`
}

type PokemonEncounter struct {
	Pokemon PokemonRef `json:"pokemon"`
}

type PokemonRef struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

type RespPokemon struct {
	ID             int    `json:"id"`
	Name           string `json:"name"`
	BaseExperience int    `json:"base_experience"`
}
