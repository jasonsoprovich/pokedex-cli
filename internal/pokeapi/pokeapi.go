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
	ID             int           `json:"id"`
	Name           string        `json:"name"`
	Height         int           `json:"height"`
	Weight         int           `json:"weight"`
	BaseExperience int           `json:"base_experience"`
	Stats          []PokemonStat `json:"stats"`
	Types          []PokemonType `json:"types"`
}

type PokemonStat struct {
	BaseStat int      `json:"base_stat"`
	Stat     StatInfo `json:"stat"`
}

type StatInfo struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

type PokemonType struct {
	Slot int      `json:"slot"`
	Type TypeInfo `json:"type"`
}

type TypeInfo struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}
