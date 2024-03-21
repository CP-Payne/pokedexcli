package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
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

func GetLocations(url string) (previousUrl, nextUrl string) {
	res, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	body, err := io.ReadAll(res.Body)
	res.Body.Close()
	if res.StatusCode > 299 {
		log.Fatalf("Response failed with status code: %d and\nbody: %s\n", res.StatusCode, body)
	}
	if err != nil {
		log.Fatal(err)
	}

	var locationsResponse LocationApiResponse
	if err := json.Unmarshal(body, &locationsResponse); err != nil {
		log.Fatalf("Error unmarshaling JSON: %v", err)
	}

	for _, location := range locationsResponse.Results {
		fmt.Println(location.Name)
	}

	return locationsResponse.Previous, locationsResponse.Next
}
