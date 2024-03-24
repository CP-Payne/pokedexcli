package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/CP-Payne/pokedexcli/internal/pokeapi"
	"github.com/CP-Payne/pokedexcli/internal/pokecache"
)

type Config struct {
	NextUrl  string
	PrevUrl  string
	Cache    *pokecache.Cache
	Location *pokeapi.LocationPokemonApiResponse
	Pokedex  *pokeapi.Pokedex
}

var commands map[string]cliCommand

func main() {
	cache := pokecache.NewCache(5 * time.Minute)
	c := &Config{
		NextUrl:  "https://pokeapi.co/api/v2/location-area?offset=0&limit=20",
		PrevUrl:  "",
		Cache:    cache,
		Location: &pokeapi.LocationPokemonApiResponse{},
		Pokedex: &pokeapi.Pokedex{
			Pokedex: make(map[string]pokeapi.PokemonDetails),
		},
	}

	scanner := bufio.NewScanner(os.Stdin)
	commands = c.getCommands()

	for {
		fmt.Print("\nPokedex > ")
		scanner.Scan()

		err := scanner.Err()
		if err != nil {
			log.Fatal(err)
		}
		input := scanner.Text()
		inputSlice := strings.Split(input, " ")
		commandInput := inputSlice[0]
		params := inputSlice[1:]

		command, ok := commands[commandInput]
		if !ok {
			fmt.Println("Unknown command. See 'help' for valid commands.")
			continue
		}

		err = command.callback(params...)
		if err != nil {
			fmt.Printf("Error executing %s command: %s", commandInput, err)
		}

		// fmt.Printf("Your input is: %s\n", scanner.Text())
	}
}
