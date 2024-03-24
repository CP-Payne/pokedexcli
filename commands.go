package main

import (
	"errors"
	"fmt"
	"os"

	"github.com/CP-Payne/pokedexcli/internal/pokeapi"
)

type cliCommand struct {
	name        string
	description string
	callback    func(params ...string) error
}

func (c *Config) getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"map": {
			name:        "map",
			description: "Obtain list of locations on the map. Repetitive calling will retrieve the next set of locations",
			callback:    c.mapn,
		},
		"mapb": {
			name:        "mapb",
			description: "Obtain previous list of locations on the map.",
			callback:    c.mapb,
		},
		"explore": {
			name:        "explore",
			description: "Obtain a list of pokemons in an area.",
			callback:    c.explore,
		},
		"catch": {
			name:        "catch",
			description: "Catch a pokemon.",
			callback:    c.catch,
		},
		"inspect": {
			name:        "inspect",
			description: "Inspect a pokemon.",
			callback:    c.inspect,
		},
	}
}

func commandHelp(params ...string) error {
	fmt.Println("\nWelcome to the Pokedex!")
	fmt.Print("Usage: \n\n")

	for key, value := range commands {
		fmt.Printf("%s: %s\n", key, value.description)
	}

	return nil
}

func commandExit(params ...string) error {
	fmt.Print("Exiting Pokedex...\n\n")
	os.Exit(0)
	return nil
}

func (c *Config) mapn(params ...string) error {
	if c.NextUrl == "" {
		fmt.Println("There are no more locations. Use 'mapb' command the view previous locations")
		return errors.New("there are no more locations to view")
	}
	c.PrevUrl, c.NextUrl = pokeapi.GetLocations(c.NextUrl, c.Cache)
	return nil
}

func (c *Config) mapb(params ...string) error {
	if c.PrevUrl == "" {
		fmt.Println("There are no previous locations. Use 'map' command to view next locations")
		return errors.New("no locations to navigate back to")
	}
	c.PrevUrl, c.NextUrl = pokeapi.GetLocations(c.PrevUrl, c.Cache)
	return nil
}

func (c *Config) explore(params ...string) error {
	if len(params) == 0 {
		return errors.New("please provide a location to explore (list pokemon)")
	}
	if len(params) > 1 {
		return errors.New("can only explore one location at a time")
	}

	var err error
	c.Location, err = pokeapi.LocationPokemons(params[0], c.Cache)
	if err != nil {
		return err
	}

	return nil
}

func (c *Config) catch(params ...string) error {
	if len(params) == 0 {
		return errors.New("Provide pokemon name to catch")
	}
	if len(params) > 1 {
		return errors.New("can only catch one pokemon at a time")
	}

	var err error
	err = pokeapi.CatchPokemon(params[0], c.Cache, c.Pokedex, c.Location)
	if err != nil {
		return err
	}

	return nil
}

func (c *Config) inspect(params ...string) error {
	pokemonName := params[0]
	pokemonInfo, found := c.Pokedex.Pokedex[pokemonName]
	if !found {
		return errors.New("you have not caught that pokemon")
	}

	fmt.Printf("Name: %v\n", pokemonInfo.Name)
	fmt.Printf("Height: %d\n", pokemonInfo.Height)
	fmt.Printf("Height: %d\n", pokemonInfo.Height)
	fmt.Println("Stats:")
	for _, stat := range pokemonInfo.Stats {
		fmt.Printf("  -%s: %d\n", stat.Stat.Name, stat.BaseStat)
	}
	fmt.Println("Types:")
	for _, pokemonType := range pokemonInfo.Types {
		fmt.Printf("  -%s\n", pokemonType.Type.Name)
	}
	return nil
}
