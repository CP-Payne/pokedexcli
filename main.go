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

type config struct {
	nextUrl string
	prevUrl string
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
			fmt.Println("Unknown command. See 'help' for valid commands.")
			continue
		}

		err = command.callback()
		if err != nil {
			fmt.Printf("Error executing %s command: %s", input, err)
		}

		// fmt.Printf("Your input is: %s\n", scanner.Text())
	}
}
