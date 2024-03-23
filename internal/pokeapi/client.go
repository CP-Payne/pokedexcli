package pokeapi

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/CP-Payne/pokedexcli/internal/pokecache"
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
