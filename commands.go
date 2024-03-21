package main

import (
	"fmt"
	"os"
)

func (c *config) mapn() error {
	return nil
}

func (c *config) mapb() error {
	return nil
}

func getCommands() map[string]cliCommand {
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
