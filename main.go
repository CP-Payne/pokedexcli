package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type cliCommand struct {
	name        string
	description string
	callback    func() error
}

var commands map[string]cliCommand

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	commands = getCommands()

	for {
		fmt.Print("\nPokedex > ")
		scanner.Scan()

		err := scanner.Err()
		if err != nil {
			log.Fatal(err)
		}
		input := scanner.Text()
		command, ok := commands[input]
		if !ok {
			fmt.Println("Invalid command. See 'help' for valid commands.")
			continue
		}

		err = command.callback()
		if err != nil {
			fmt.Printf("Error executing %s command: %s", input, err)
		}

		// fmt.Printf("Your input is: %s\n", scanner.Text())
	}
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
	fmt.Println("Usage: \n")

	for key, value := range commands {
		fmt.Printf("%s: %s\n", key, value.description)
	}

	return nil
}

func commandExit() error {
	fmt.Println("Exiting Pokedex...\n")
	os.Exit(0)
	return nil
}
