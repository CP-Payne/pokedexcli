package main

import (
	"errors"
	"fmt"
	"os"

	"github.com/CP-Payne/pokedexcli/internal/pokeapi"
)

func (c *config) mapn() error {
	if c.nextUrl == "" {
		fmt.Println("There are no more locations. Use 'mapb' command the view previous locations")
		return errors.New("there are no more locations to view")
	}
	c.prevUrl, c.nextUrl = pokeapi.GetLocations(c.nextUrl)
	return nil
}

func (c *config) mapb() error {
	if c.prevUrl == "" {
		fmt.Println("There are no previous locations. Use 'map' command to view next locations")
		return errors.New("no locations to navigate back to")
	}
	c.prevUrl, c.nextUrl = pokeapi.GetLocations(c.prevUrl)
	return nil
}

func (c *config) getCommands() map[string]cliCommand {
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
	}
}

func commandHelp() error {
	fmt.Println("\nWelcome to the Pokedex!")
	fmt.Print("Usage: \n\n")

	for key, value := range commands {
		fmt.Printf("%s: %s\n", key, value.description)
	}

	return nil
}

func commandExit() error {
	fmt.Print("Exiting Pokedex...\n\n")
	os.Exit(0)
	return nil
}
