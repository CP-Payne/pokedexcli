package pokeapi

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/CP-Payne/pokedexcli/internal/pokecache"
	"github.com/CP-Payne/pokedexcli/internal/utils"
)

type Location struct {
	Name string
	Url  string
}

type LocationApiResponse struct {
	Next     string
	Previous string
	Results  []Location
}

type LocationPokemonApiResponse struct {
	Location          Location            `json:"location"`
	PokemonEncounters []PokemonEncounters `json:"pokemon_encounters"`
}

type PokemonEncounters struct {
	Pokemon Pokemon `json:"pokemon"`
}

type Pokemon struct {
	Name string `json:"name"`
	Url  string `json:"url"`
}

type Pokedex struct {
	Pokedex map[string]PokemonDetails
}

type PokemonDetails struct {
	BaseExperience int     `json:"base_experience"`
	Name           string  `json:"name"`
	Height         int     `json:"height"`
	Weight         int     `json:"weight"`
	Stats          []Stats `json:"stats"`
	Types          []Types `json:"types"`
}

type Types struct {
	Slot int  `json:"slot"`
	Type Type `json:"type"`
}

type Type struct {
	Name string `json:"name"`
	Url  string `json:"url"`
}

type Stats struct {
	BaseStat int  `json:"base_stat"`
	Effort   int  `json:"effort"`
	Stat     Stat `json:"stat"`
}

type Stat struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

func GetLocations(url string, cache *pokecache.Cache) (previousUrl, nextUrl string) {
	data, inCache := cache.Get(url)
	if !inCache {
		res, err := http.Get(url)
		if err != nil {
			log.Fatal(err)
		}
		data, err = io.ReadAll(res.Body)
		res.Body.Close()
		// data is the body of the response (json)
		if res.StatusCode > 299 {
			log.Fatalf("Response failed with status code: %d and\nbody: %s\n", res.StatusCode, data)
		}
		if err != nil {
			log.Fatal(err)
		}

		cache.Add(url, data)

	}

	var locationsResponse LocationApiResponse
	if err := json.Unmarshal(data, &locationsResponse); err != nil {
		log.Fatalf("Error unmarshaling JSON: %v", err)
	}

	for _, location := range locationsResponse.Results {
		fmt.Println(location.Name)
	}

	return locationsResponse.Previous, locationsResponse.Next
}

func LocationPokemons(locationName string, cache *pokecache.Cache) (locationPokemons *LocationPokemonApiResponse, err error) {
	locationDetailsUrl := fmt.Sprintf("https://pokeapi.co/api/v2/location-area/%s/", locationName)
	// fmt.Println(locationDetailsUrl)
	fmt.Printf("Exploring %s...\n", locationName)
	data, inCache := cache.Get(locationDetailsUrl)
	if !inCache {
		res, err := http.Get(locationDetailsUrl)
		if err != nil {
			log.Fatal(err)
		}
		data, err = io.ReadAll(res.Body)
		// fmt.Println(string(data))
		res.Body.Close()
		// data is the body of the response (json)
		if res.StatusCode > 299 {
			// log.Fatalf("Response failed with status code: %d and\nbody: %s\n", res.StatusCode, data)
			return &LocationPokemonApiResponse{}, errors.New("location not found")
		}
		if err != nil {
			log.Fatal(err)
		}

		cache.Add(locationDetailsUrl, data)

	}

	var locationPokemon LocationPokemonApiResponse
	if err := json.Unmarshal(data, &locationPokemon); err != nil {
		log.Fatalf("Error unmarshaling JSON: %v", err)
	}

	// fmt.Printf("Current location %s", locationPokemonResponse.Location.Name)

	fmt.Printf("Found %d Pokemon:\n", len(locationPokemon.PokemonEncounters))

	for _, encounter := range locationPokemon.PokemonEncounters {
		fmt.Printf("- %s\n", encounter.Pokemon.Name)
	}
	return &locationPokemon, nil
}

func CatchPokemon(pokemonName string, cache *pokecache.Cache, pokedex *Pokedex, location *LocationPokemonApiResponse) (err error) {
	pokemonInArea := false
	for _, encounter := range location.PokemonEncounters {
		if pokemonName == encounter.Pokemon.Name {
			pokemonInArea = true
			break
		}
	}

	if !pokemonInArea {
		return errors.New("pokemon not found in area")
	}

	// fmt.Printf("Catching pokemon %s....", pokemonName)
	pokemonDetailsUrl := fmt.Sprintf("https://pokeapi.co/api/v2/pokemon/%s/", pokemonName)
	fmt.Printf("Throwing a Pokeball at %s...\n", pokemonName)

	data, inCache := cache.Get(pokemonDetailsUrl)
	if !inCache {
		res, err := http.Get(pokemonDetailsUrl)
		if err != nil {
			log.Fatal(err)
		}
		data, err = io.ReadAll(res.Body)
		// fmt.Println(string(data))
		res.Body.Close()
		// data is the body of the response (json)
		if res.StatusCode > 299 {
			// log.Fatalf("Response failed with status code: %d and\nbody: %s\n", res.StatusCode, data)
			return errors.New("location not found")
		}
		if err != nil {
			log.Fatal(err)
		}

		cache.Add(pokemonDetailsUrl, data)

	}

	var pokemonDetails PokemonDetails
	if err := json.Unmarshal(data, &pokemonDetails); err != nil {
		log.Fatalf("Error unmarshaling JSON: %v", err)
	}

	// Attempting to catch pokemon
	catchStatus := utils.CatchStatus(pokemonDetails.BaseExperience)

	if catchStatus {
		pokedex.Pokedex[pokemonDetails.Name] = pokemonDetails
		fmt.Printf("%s was caught!\n", pokemonDetails.Name)
	} else {
		fmt.Printf("%s escaped!\n", pokemonDetails.Name)
	}

	return nil
}
