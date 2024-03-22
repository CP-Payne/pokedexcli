package pokeapi

import (
	"encoding/json"
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
